package dataoke

import (
	"fmt"
	"time"
)

type PullGoodsByTimeResp struct {
	CommonRespHeader
	Data *PullGoodsByTimeRespBody `json:"data"`
}

type PullGoodsByTimeRespBody struct {
	List     []*TaobaoGoodsSt `json:"list"`
	PageId   string           `json:"pageId"`
	TotalNum int64            `json:"totalNum"`
}

func (a *DaTaoKeApp) PullGoodsByTime(pageId int64, startTime int64, endTime int64, params Params) (res *PullGoodsByTimeRespBody, err error) {
	params.Set("version", "v1.2.3")
	params.Set("pageId", pageId)
	params.Set("startTime", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"))
	if endTime != 0 {
		params.Set("endTime", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	}
	bResp, err := a.postQuery(params, PullGoodsByTimeURI)
	if err != nil {
		return
	}
	resp := &PullGoodsByTimeResp{}
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
