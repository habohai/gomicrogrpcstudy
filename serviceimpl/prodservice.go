package serviceimpl

import (
	"context"
	"gomicrogrpcstudy/services"
	"strconv"
	"time"
)

type ProdService struct{}

// 测试方法
func newProd(pid int32, pname string) *services.ProdModel {
	return &services.ProdModel{
		ProdID:   pid,
		ProdName: pname,
	}
}

func (p *ProdService) GetProdsList(ctx context.Context, in *services.ProdsRequest, res *services.ProdListResponse) error {
	time.Sleep(time.Second * 3)
	var models []*services.ProdModel
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}

func (p *ProdService) GetProdsDetail(ctx context.Context, in *services.ProdsRequest, res *services.ProdDetailResponse) error {
	time.Sleep(time.Second * 3)
	res.Data = newProd(in.ProdId, "测试商品详情")
	return nil
}
