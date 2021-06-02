package dataoke

import "fmt"

type GetGoodsDetailsResp struct {
	CommonRespHeader
	Data *TaobaoGoodsSt `json:"data"`
}

func (a *DaTaoKeApp) GetGoodsDetails(pageId int64, id int64, goodsId string) (res *TaobaoGoodsSt, err error) {
	params := NewParams()
	params.Set("version", "v1.2.3")
	if id != 0 {
		params.Set("id", id)
	} else if goodsId != "" {
		params.Set("goodsId", goodsId)
	}
	bResp, err := a.postQuery(params, GetGoodsDetailsURI)
	if err != nil {
		return
	}
	resp := &GetGoodsDetailsResp{}
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
