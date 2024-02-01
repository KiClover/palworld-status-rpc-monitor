package monitor

import (
	"context"
	"github.com/KiClover/palworld-status-rpc-monitor/internal/svc"
	"github.com/KiClover/palworld-status-rpc-monitor/types/palworldmonitor"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
	"math"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMonitorInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMonitorInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMonitorInfoLogic {
	return &GetMonitorInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	netInSpeed, netOutSpeed, netInTransfer, netOutTransfer, lastUpdateNetStats uint64
)

var (
	//Version           string
	//expectDiskFsTypes = []string{
	//	"apfs", "ext4", "ext3", "ext2", "f2fs", "reiserfs", "jfs", "btrfs",
	//	"fuseblk", "zfs", "simfs", "ntfs", "fat32", "exfat", "xfs", "fuse.rclone",
	//}
	excludeNetInterfaces = []string{
		"lo", "tun", "docker", "veth", "br-", "vmbr", "vnet", "kube",
	}
	//getMacDiskNo = regexp.MustCompile(`\/dev\/disk(\d)s.*`)
)

func (l *GetMonitorInfoLogic) GetMonitorInfo(in *palworldmonitor.Empty) (*palworldmonitor.MonitorInfo, error) {
	memory, _ := mem.VirtualMemory()
	percent, _ := cpu.Percent(0, false)
	bootTime, _ := host.BootTime()
	var cachedBootTime time.Time
	cachedBootTime = time.Unix(int64(bootTime), 0)
	TrackNetworkSpeed()
	return &palworldmonitor.MonitorInfo{
		MemUsed:    memory.Used / MB,
		MemTotal:   memory.Total / MB,
		MemPercent: float32(Round(memory.UsedPercent, 2)),
		CpuPercent: float32(Round(percent[0], 0)),
		BootTime:   GetHourDiffer(cachedBootTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
		NetIn:      float32(Round(float64(netInSpeed/KB), 2)),
		NetOut:     float32(Round(float64(netOutSpeed/KB), 2)),
	}, nil
}

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n // TODO +0.5 是为了四舍五入，如果不希望这样去掉这个
}

// GetHourDiffer 获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}
func TrackNetworkSpeed() {
	var innerNetInTransfer, innerNetOutTransfer uint64
	nc, err := net.IOCounters(true)
	if err == nil {
		for _, v := range nc {
			if isListContainsStr(excludeNetInterfaces, v.Name) {
				continue
			}
			innerNetInTransfer += v.BytesRecv
			innerNetOutTransfer += v.BytesSent
		}
		now := uint64(time.Now().Unix())
		diff := now - lastUpdateNetStats
		if diff > 0 {
			netInSpeed = (innerNetInTransfer - netInTransfer) / diff
			netOutSpeed = (innerNetOutTransfer - netOutTransfer) / diff
		}
		netInTransfer = innerNetInTransfer
		netOutTransfer = innerNetOutTransfer
		lastUpdateNetStats = now
	}
}

func isListContainsStr(list []string, str string) bool {
	for i := 0; i < len(list); i++ {
		if strings.Contains(str, list[i]) {
			return true
		}
	}
	return false
}
