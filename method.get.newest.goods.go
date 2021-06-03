package dataoke

import (
	"fmt"
	"time"
)

type GetNewestGoodsResp struct {
	CommonRespHeader
	Data *GetNewestGoodsRespBody `json:"data"`
}

type GetNewestGoodsRespBody struct {
	List []*GoodsUpdateInfoSt `json:"list"`
}

type GoodsUpdateInfoSt struct {
	// 商品相关信息
	Id            int64    `json:"id"`            // 大淘客商品id
	GoodsId       string   `json:"goodsId"`       // 淘宝商品id
	OriginalPrice float64  `json:"originalPrice"` // 原价
	ActualPrice   float64  `json:"actualPrice"`   // 券后价
	MonthSales    int64    `json:"monthSales"`    // 近30天销量
	TwoHoursSales int64    `json:"twoHoursSales"` // 近两小时销量
	DailySales    int64    `json:"dailySales"`    // 当日销量
	Subcid        []int64  `json:"subcid"`        // 大淘客二级分类id
	SpecialText   []string `json:"specialText"`   // 特色文案

	// 优惠券相关信息
	CouponPrice       float64 `json:"couponPrice"`       // 优惠券金额
	CouponReceiveNum  int64   `json:"couponReceiveNum"`  // 领券量
	CouponRemainCount int64   `json:"couponRemainCount"` // 券剩余量
	CouponLink        string  `json:"couponLink"`        // 领券链接

	// 佣金相关信息
	CommissionType TaoBaoCommissionType `json:"commissionType"` // 佣金类型
	CommissionRate float64              `json:"commissionRate"` // 佣金比例
}

func (a *DaTaoKeApp) GetNewestGoods(pageId string, startTime int64, endTime int64, params Params) (res *GetNewestGoodsRespBody, err error) {
	params.Set("version", "v1.2.0")
	params.Set("pageId", pageId)
	params.Set("startTime", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"))
	if endTime != 0 {
		params.Set("endTime", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	}
	bResp, err := a.postQuery(params, GetNewestGoodsURI)
	if err != nil {
		return
	}
	resp := &GetNewestGoodsResp{}
	err = dJson.Unmarshal(bResp, resp)
	if err != nil {
		return
	}
	if resp.Code != 0 {
		err = fmt.Errorf("%v", resp.Msg)
		return
	}
	if resp.Data == nil {
		err = fmt.Errorf("nil resp")
		return
	}
	res = resp.Data

	return
}
