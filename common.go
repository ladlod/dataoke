package dataoke

type CommonRespHeader struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

// TaoBaoShopType 店铺类型
type TaoBaoShopType int64

const (
	ShopTaobao  TaoBaoShopType = 0 // 淘宝店铺
	ShopTianMao TaoBaoShopType = 1 // 天猫店铺
)

// TaoBaoCommissionType 佣金类型
type TaoBaoCommissionType int64

const (
	CommissionDefault = 0 // 通用
	CommissionFixed   = 1 // 定向
	CommissionHigh    = 2 // 高佣
	CommissionMarket  = 3 // 营销
)

// TaoBaoActivityType 淘宝活动类型
type TaoBaoActivityType int64

const (
	ActivityNone      = 0 // 无活动
	ActivityRush      = 1 // 淘抢购
	ActivityJuhuasuan = 2 // 聚划算
)
