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

func (a *DaTaoKeApp) SuperDiscountGoods(pageId string, pageSize int64, params Params) (res *GetLeaderboardListRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	bResp, err := a.postQuery(params, SuperDiscountGoodsURI)
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

type GetHalfPriceDayResp struct {
	CommonRespHeader
	Data *GetHalfPriceDayRespBody `json:"data"`
}

type GetHalfPriceDayRespBody struct {
	HalfPriceInfo *HalfPriceInfoSt `json:"halfPriceInfo"`
	SessionsList  []*HalfSessionSt `json:"sessionsList"`
}

type HalfPriceInfoSt struct {
	Banner string             `json:"banner"`
	List   []*HalfPriceItemSt `json:"list"`
}

type HalfPriceItemSt struct {
	Id           int64   `json:"id"`
	ItemId       string  `json:"itemId"`
	Top          int64   `json:"top"`
	HdLeixing    int64   `json:"hdLeixing"`
	QiangNum     int64   `json:"qiangNum"`
	UpdateTime   string  `json:"updateTime"`
	ServerTime   string  `json:"serverTime"`
	StartTime    string  `json:"startTime"`
	ItemSoldNum  int64   `json:"itemSoldNum"`
	TodaySellNum string  `json:"todaySellNum"`
	Name         string  `json:"name"`
	PicUrl       string  `json:"picUrl"`
	Price        float64 `json:"price"`
	Yijuhua      string  `json:"yijuhua"`
	Preferential string  `json:"preferential"`
	CouponAmount float64 `json:"couponAmount"`
	ActivityId   int64   `json:"activityId"`
	RestCount    int64   `json:"restCount"`
	Tmall        int64   `json:"tmall"`
	ActivityType int64   `json:"activityType"`
	UseQuan      int64   `json:"useQuan"`
	IsMamaQuan   int64   `json:"isMamaQuan"`
}

type HalfSessionSt struct {
	HpdTime string `json:"hpdTime"`
	Status  string `json:"status"`
}

func (a *DaTaoKeApp) GetHalfPriceDay(pageId string, pageSize int64, params Params) (res *GetHalfPriceDayRespBody, err error) {
	params.Set("version", "v1.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	bResp, err := a.postQuery(params, GetHalfPriceDayURI)
	if err != nil {
		return
	}
	resp := &GetHalfPriceDayResp{}
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

type ListTipOffResp struct {
	CommonRespHeader
	Data *ListTipOffRespBody `json:"data"`
}

type ListTipOffRespBody struct {
	PageId   string              `json:"pageId"`
	TotalNum int64               `json:"totalNum"`
	GoScroll bool                `json:"goScroll"`
	List     []*ListTipOffItemSt `json:"list"`
}

type ListTipOffItemSt struct {
	ItemId     string `json:"itemId"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Remark     string `json:"remark"`
	Url        string `json:"url"`
	UpdateTime string `json:"updateTime"`
}

// 仅实现了topic为1的形式
func (a *DaTaoKeApp) ListTipOff(pageId string, pageSize int64, params Params) (res *ListTipOffRespBody, err error) {
	params.Set("version", "v3.0.0")
	params.Set("pageId", pageId)
	params.Set("pageSize", pageSize)
	params.Set("topic", 1)
	bResp, err := a.postQuery(params, ListTipOffURI)
	if err != nil {
		return
	}
	resp := &ListTipOffResp{}
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
