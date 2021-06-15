package dataoke

import (
	"fmt"
	"time"
)

type GetOrderDetailResp struct {
	CommonRespHeader
	Data *GetOrderDetailRespBody `json:"data"`
}

type GetOrderDetailRespBody struct {
	HasPre        bool                    `json:"has_pre"`
	HasNext       bool                    `json:"has_next"`
	PageNo        int64                   `json:"page_no"`
	PageSize      int64                   `json:"page_size"`
	PositionIndex string                  `json:"position_index"`
	Results       *GetOrderDetailResultSt `json:"results"`
	ErrorResponse *ErrorResponse          `json:"error_response,omitempty"`
}

type GetOrderDetailResultSt struct {
	PublisherOrderDto []*TaoBaoOrderSt `json:"publisher_order_dto"`
}

func (a *DaTaoKeApp) GetOrderDetails(queryType TaoBaoOrderQueryType, startTime int64, endTime int64, params Params) (res *GetOrderDetailRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("queryType", queryType)
	params.Set("startTime", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"))
	params.Set("endTime", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	bResp, err := a.postQuery(params, GetOrderDetailsURI)
	if err != nil {
		return
	}
	resp := &GetOrderDetailResp{}
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
	if resp.Data.ErrorResponse != nil {
		err = fmt.Errorf("%v", resp.Data.ErrorResponse.SubMsg)
		return
	}
	res = resp.Data

	return
}
