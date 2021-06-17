package dataoke

const (
	ReqDomain = "https://openapi.dataoke.com"

	// goods
	GetGoodsListURI     = "api/goods/get-goods-list"
	PullGoodsByTimeURI  = "api/goods/pull-goods-by-time"
	GetStaleGoodsURI    = "api/goods/get-stale-goods-by-time"
	GetNewestGoodsURI   = "api/goods/get-newest-goods"
	GetGoodsDetailsURI  = "api/goods/get-goods-details"
	GetSuperCategoryURI = "api/category/get-super-category"

	// order
	GetOrderDetailsURI = "api/tb-service/get-order-details"

	// trans
	ParseTaokoulingURI  = "api/tb-service/parse-taokouling"
	GetPrivilegeLinkURI = "api/tb-service/get-privilege-link"

	// leaderboard
	GetNineOpListURI          = "api/goods/nine/op-goods-list"
	GetHighCommissionListURI  = "api/goods/singlePage/list-height-commission"
	GetExplosiveListURI       = "api/goods/explosive-goods-list"
	GetHistoryLowPriceListURI = "api/goods/get-history-low-price-list"
	GetHalfPriceDayURI        = "api/goods/get-half-price-day"
	ListTipOffURI             = "api/dels/spider/list-tip-off"
	SuperDiscountGoodsURI     = "api/goods/super-discount-goods"
)
