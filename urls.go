package dataoke

const (
	ReqDomain = "https://openapi.dataoke.com"

	// goods
	GetGoodsListURI    = "api/goods/get-goods-list"
	PullGoodsByTimeURI = "api/goods/pull-goods-by-time"
	GetStaleGoodsURI   = "api/goods/get-stale-goods-by-time"
	GetNewestGoodsURI  = "api/goods/get-newest-goods"
	GetGoodsDetailsURI = "api/goods/get-goods-details"

	// order
	GetOrderDetailsURI = "api/tb-service/get-order-details"

	// trans
	ParseTaokoulingURI  = "api/tb-service/parse-taokouling"
	GetPrivilegeLinkURI = "api/tb-service/get-privilege-link"

	// leaderboard
	GetNineOpListURI         = "api/goods/nine/op-goods-list"
	GetHighCommissionListURI = "api/goods/singlePage/list-height-commission "
	GetExplosiveListURI      = "api/goods/explosive-goods-list"
)
