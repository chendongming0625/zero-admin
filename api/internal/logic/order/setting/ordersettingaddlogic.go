package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingAddLogic {
	return OrderSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingAddLogic) OrderSettingAdd(req types.AddOrderSettingReq) (*types.AddOrderSettingResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddOrderSettingResp{}, nil
}
