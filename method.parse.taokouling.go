package dataoke

import "fmt"

type ParseTaokoulingResp struct {
	CommonRespHeader
	Data *ParseTaokoulingRespBody `json:"data"`
}

type ParseTaokoulingRespBody struct {
	GoodsId    string                 `json:"goodsId"`
	OriginUrl  string                 `json:"originUrl"`
	OriginType string                 `json:"originType"`
	OriginInfo *TaokoulingOrginInfoSt `json:"originInfo"`
}

type TaokoulingOrginInfoSt struct {
	// 商品信息
	Pid      string  `json:"pid"`
	Image    string  `json:"image"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ShopLogo string  `json:"shop_logo"`
	ShopName string  `json:"shopName"`

	// 优惠券信息
	ActivityId string  `json:"activityId"`
	Amount     float64 `json:"amount"`
	StartFee   float64 `json:"startFee"`
	StartTime  string  `json:"startTime"`
	EndTime    string  `json:"endTime"`
	Status     int64   `json:"status"`
}

func (a *DaTaoKeApp) ParseTaokouling(taokouling string) (res *ParseTaokoulingRespBody, err error) {
	params := NewParams()
	params.Set("version", "v1.0.0")
	params.Set("content", taokouling)
	bResp, err := a.postQuery(params, ParseTaokoulingURI)
	if err != nil {
		return
	}
	resp := &ParseTaokoulingResp{}
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
