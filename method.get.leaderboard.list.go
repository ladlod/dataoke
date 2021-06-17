package dataoke

import "fmt"

type GetLeaderboardListResp struct {
	CommonRespHeader
	Data *GetLeaderboardListRespBody `json:"data"`
}

type GetLeaderboardListRespBody struct {
	List     []*TaobaoGoodsSt `json:"list"`
	TotalNum int64            `json:"totalNum"`
	PageId   string           `json:"pageId"`
}

func (a *DaTaoKeApp) GetNineOpList(pageId string, pageSize int64, nineCid string, params Params) (res *GetLeaderboardListRespBody, err error) {
	params.Set("version", "v2.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	params.Set("nineCid", nineCid)
	bResp, err := a.postQuery(params, GetNineOpListURI)
	if err != nil {
		return
	}
	resp := &GetLeaderboardListResp{}
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

func (a *DaTaoKeApp) GetHighCommissionList(pageId string, pageSize int64, params Params) (res *GetLeaderboardListRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	bResp, err := a.postQuery(params, GetHighCommissionListURI)
	if err != nil {
		return
	}
	resp := &GetLeaderboardListResp{}
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

func (a *DaTaoKeApp) GetExplosiveList(pageId string, pageSize int64, params Params) (res *GetLeaderboardListRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	bResp, err := a.postQuery(params, GetExplosiveListURI)
	if err != nil {
		return
	}
	resp := &GetLeaderboardListResp{}
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

func (a *DaTaoKeApp) GetHistoryLowPriceList(pageId string, pageSize int64, params Params) (res *GetLeaderboardListRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	bResp, err := a.postQuery(params, GetHistoryLowPriceListURI)
	if err != nil {
		return
	}
	resp := &GetLeaderboardListResp{}
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
