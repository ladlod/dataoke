package dataoke

import "fmt"

type GetSuperCategoryResp struct {
	CommonRespHeader
	Data []*TaobaoCategorySt `json:"data"`
}

type TaobaoCategorySt struct {
	Cid   int64  `json:"cid"`
	Cname string `json:"cname"`
	Cpic  string `json:"cpic"`
}

type TaobaoSubCategorySt struct {
	SubCid   int64  `json:"subcid"`
	SubCname string `json:"subcname"`
	SCpic    string `json:"scpic"`
}

func (a *DaTaoKeApp) GetSuperCategory() (res []*TaobaoCategorySt, err error) {
	params := NewParams()
	params.Set("version", "v1.1.0")
	bResp, err := a.postQuery(params, GetSuperCategoryURI)
	if err != nil {
		return
	}
	resp := &GetSuperCategoryResp{}
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
