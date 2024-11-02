package user

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-im/apps/user/rpc/user"
	"zero-im/pkg/ctxdata"

	"zero-im/apps/user/api/internal/svc"
	"zero-im/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDetailLogic 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Detail 用户详情
func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	uid := ctxdata.GetUid(l.ctx)

	userInfoResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	var res types.User
	copier.Copy(&res, userInfoResp.User)

	return &types.UserInfoResp{Info: res}, nil

}
