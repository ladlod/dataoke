package dataoke

import "fmt"

type GetPrivilegeLinkResp struct {
	CommonRespHeader
	Data *GetPrivilegeLinkRespBody `json:"data"`
}

type GetPrivilegeLinkRespBody struct {
	ItemId            int64   `json:"itemId"`            // 商品id
	OriginalPrice     float64 `json:"originalPrice"`     // 商品原价
	ActualPrice       float64 `json:"actualPrice"`       // 商品券后价
	ItemUrl           string  `json:"itemUrl"`           // 商品淘客链接
	ShortUrl          string  `json:"shortUrl"`          // 短连接
	Tpwd              string  `json:"tpwd"`              // 淘口令
	LongTpwd          string  `json:"longTpwd"`          // 针对ios14版本的长口令
	CouponClickUrl    string  `json:"couponClickUrl"`    // 优惠券推广链接
	CouponEndTime     string  `json:"couponEndTime"`     // 优惠券结束时间
	CouponStartTime   string  `json:"couponStartTime"`   // 优惠券开始时间
	CouponInfo        string  `json:"couponInfo"`        // 优惠券面额描述
	CouponTotalCount  int64   `json:"couponTotalCount"`  // 优惠券总量
	CouponRemainCount int64   `json:"couponRemainCount"` // 优惠券余量
	MaxCommissionRate string  `json:"maxCommissionRate"` // 佣金比例
	MinCommissionRate string  `json:"minCommissionRate"` // 最低佣金比
}

func (a *DaTaoKeApp) GetPrivilegeLink(pid string, relationId string, goodsId string, params Params) (res *GetPrivilegeLinkRespBody, err error) {
	params.Set("version", "v1.3.1")
	params.Set("pid", pid)
	params.Set("channelId", relationId)
	params.Set("goodsId", goodsId)
	bResp, err := a.postQuery(params, GetPrivilegeLinkURI)
	if err != nil {
		return
	}
	resp := &GetPrivilegeLinkResp{}
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
