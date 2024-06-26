package service

import (
	util "Mall/pkg/utils/log"
	"Mall/repository/db/dao"
	"Mall/types"
	"context"
	"sync"
)

var CarouselSrvIns *CarouselSrv
var CarouselSrvOnce sync.Once

type CarouselSrv struct {
}

func GetCarouselSrv() *CarouselSrv {
	CarouselSrvOnce.Do(func() {
		CarouselSrvIns = &CarouselSrv{}
	})
	return CarouselSrvIns
}

// ListCarousel 列表
func (s *CarouselSrv) ListCarousel(ctx context.Context, req *types.ListCarouselReq) (resp interface{}, err error) {
	carousels, err := dao.NewCarouselDao(ctx).ListCarousel()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  carousels,
		Total: int64(len(carousels)),
	}

	return
}
