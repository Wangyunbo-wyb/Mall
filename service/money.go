package service

import (
	"Mall/pkg/utils/ctl"
	"Mall/pkg/utils/log"
	"Mall/repository/db/dao"
	"Mall/types"
	"context"
	"github.com/spf13/cast"
	"sync"
)

var MoneySrvIns *MoneySrv
var MoneySrvOnce sync.Once

type MoneySrv struct {
}

func GetMoneySrv() *MoneySrv {
	MoneySrvOnce.Do(func() {
		MoneySrvIns = &MoneySrv{}
	})
	return MoneySrvIns
}

// MoneyShow 展示用户的金额
func (s *MoneySrv) MoneyShow(ctx context.Context, req *types.MoneyShowReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	user, err := dao.NewUserDao(ctx).GetUserById(u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	money, err := user.DecryptMoney(req.Key)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	resp = &types.MoneyShowResp{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserMoney: cast.ToString(money),
	}

	return
}
