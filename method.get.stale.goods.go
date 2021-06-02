package dataoke

import (
	"fmt"
	"time"
)

type GetStaleGoodsResp struct {
	CommonRespHeader
	Data *GetStaleGoodsRespBody `json:"data"`
}

type GetStaleGoodsRespBody struct {
	List     []*StaleGoodsSt `json:"list"`
	PageId   string          `json:"pageId"`
	TotalNum int64           `json:"totalNum"`
}

type StaleGoodsSt struct {
	Id      int64  `json:"id"`
	GoodsId string `json:"goodsId"`
}

func (a *DaTaoKeApp) GetStaleGoods(pageId int64, startTime int64, endTime int64, params Params) (res *GetStaleGoodsRespBody, err error) {
	params.Set("version", "v1.0.1")
	params.Set("pageId", pageId)
	params.Set("startTime", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"))
	if endTime != 0 {
		params.Set("endTime", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	}
	bResp, err := a.postQuery(params, GetStaleGoodsURI)
	if err != nil {
		return
	}
	resp := &GetStaleGoodsResp{}
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
