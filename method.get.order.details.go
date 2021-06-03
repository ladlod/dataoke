package dataoke

import "time"

func (a *DaTaoKeApp) GetOrderDetails(queryType TaoBaoOrderQueryType, startTime int64, endTime int64, params Params) (err error) {
	params.Set("version", "v1.0.0")
	params.Set("queryType", queryType)
	params.Set("startTime", time.Unix(startTime, 0).Format("2006-01-02 15:04:05"))
	params.Set("endTime", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))
	_, err = a.postQuery(params, GetOrderDetailsURI)
	if err != nil {
		return
	}

	return
}
