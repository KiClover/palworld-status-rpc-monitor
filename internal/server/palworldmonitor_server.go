// Code generated by goctl. DO NOT EDIT.
// Source: palworldmonitor.proto

package server

import (
	"context"

	"github.com/KiClover/palworld-status-rpc-monitor/internal/logic/base"
	"github.com/KiClover/palworld-status-rpc-monitor/internal/logic/monitor"
	"github.com/KiClover/palworld-status-rpc-monitor/internal/svc"
	"github.com/KiClover/palworld-status-rpc-monitor/types/palworldmonitor"
)

type PalworldmonitorServer struct {
	svcCtx *svc.ServiceContext
	palworldmonitor.UnimplementedPalworldmonitorServer
}

func NewPalworldmonitorServer(svcCtx *svc.ServiceContext) *PalworldmonitorServer {
	return &PalworldmonitorServer{
		svcCtx: svcCtx,
	}
}

func (s *PalworldmonitorServer) InitDatabase(ctx context.Context, in *palworldmonitor.Empty) (*palworldmonitor.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

func (s *PalworldmonitorServer) GetMonitorInfo(ctx context.Context, in *palworldmonitor.Empty) (*palworldmonitor.MonitorInfo, error) {
	l := monitor.NewGetMonitorInfoLogic(ctx, s.svcCtx)
	return l.GetMonitorInfo(in)
}
