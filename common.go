package dataoke

type CommonRespHeader struct {
	Status int64  `json:"status"`
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Time   int64  `json:"time"`
}
