package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-im/apps/user/rpc/internal/svc"
	"zero-im/apps/user/rpc/user"
	"zero-im/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		err   error
		users []*models.User
	)

	if in.Phone != "" {
		userEntity, err := l.svcCtx.UserDao.FindByPhone(in.Phone)
		if err == nil {
			users = append(users, userEntity)
		}
	} else if in.Name != "" {
		users, err = l.svcCtx.UserDao.ListByName(in.Name)
	} else if len(in.Ids) > 0 {
		users, err = l.svcCtx.UserDao.ListByIds(in.Ids)
	}

	if err != nil {
		return nil, err
	}

	var resp []*user.UserEntity
	copier.Copy(&resp, users)
	return &user.FindUserResp{
		User: resp,
	}, err
}
