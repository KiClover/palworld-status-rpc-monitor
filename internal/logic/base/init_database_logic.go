package base

import (
	"context"

	"github.com/KiClover/palworld-status-rpc-monitor/internal/svc"
	"github.com/KiClover/palworld-status-rpc-monitor/types/palworldmonitor"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *palworldmonitor.Empty) (*palworldmonitor.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &palworldmonitor.BaseResp{}, nil
}
