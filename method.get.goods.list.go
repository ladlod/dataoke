package dataoke

import "fmt"

type GetGoodsListResp struct {
	CommonRespHeader
	Data *GetGoodsListRespBody `json:"data"`
}

type GetGoodsListRespBody struct {
	List     []*TaobaoGoodsSt `json:"list"`
	PageId   string           `json:"pageId"`
	TotalNum int64            `json:"totalNum"`
	GoScroll bool             `json:"goScroll"`
}

func (a *DaTaoKeApp) GetGoodsList(pageId string, params Params) (res *GetGoodsListRespBody, err error) {
	params.Set("version", "v1.2.4")
	params.Set("pageId", pageId)
	bResp, err := a.postQuery(params, GetGoodsListURI)
	if err != nil {
		return
	}
	resp := &GetGoodsListResp{}
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
