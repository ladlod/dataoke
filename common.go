package dataoke

type CommonRespHeader struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

type TaobaoGoodsSt struct {
	// 商品相关信息
	Id               int64    `json:"id"`               // 大淘客商品id
	GoodsId          string   `json:"goodsId"`          // 淘宝商品id
	Title            string   `json:"title"`            // 淘宝标题
	Dtitle           string   `json:"dtitle"`           // 大淘客短标题
	ItemLink         string   `json:"itemLink"`         // 商品链接
	OriginalPrice    float64  `json:"originalPrice"`    // 原价
	ActualPrice      float64  `json:"actualPrice"`      // 现价
	MonthSales       int64    `json:"monthSales"`       // 近30天销量
	TwoHoursSales    int64    `json:"twoHoursSales"`    // 近两小时销量
	DailySales       int64    `json:"dailySales"`       // 当日销量
	Desc             string   `json:"desc"`             // 推广文案
	CreateTime       string   `json:"createTime"`       // 商品创建时间
	MainPic          string   `json:"mainPic"`          // 商品主图
	MarketingMainPic string   `json:"marketingMainPic"` // 营销主图
	SpecialText      []string `json:"specialText"`      // 特色文案
	Video            string   `json:"video"`            // 商品视频
	Cid              int64    `json:"cid"`              // 大淘客分类id
	Subcid           []int64  `json:"subcid"`           // 大淘客二级分类id
	Tbcid            int64    `json:"tbcid"`            // 淘宝品类id
	Haitao           int64    `json:"haitao"`           // 是否为海淘
	Tchaoshi         int64    `json:"tchaoshi"`         // 是否为天猫超市
	Yunfeixian       int64    `json:"yunfeixian"`       // 是否有运费险

	// 店铺相关信息
	ShopType    TaoBaoShopType `json:"shopType"`    // 店铺类型
	GoldSellers int64          `json:"goldSellers"` // 是否为金牌卖家
	SellerId    string         `json:"sellerId"`    // 淘宝卖家id
	ShopName    string         `json:"shopName"`    // 店铺名称
	ShopLogo    string         `json:"shopLogo"`    // 店铺logo
	ShopLevel   int64          `json:"shopLevel"`   // 店铺等级

	// 优惠券相关信息
	CouponReceiveNum int64   `json:"couponReceiveNum"` // 领券量
	CouponLink       string  `json:"couponLink"`       // 领券链接
	CouponStartTime  string  `json:"couponStartTime"`  // 优惠券开始时间
	CouponEndTime    string  `json:"couponEndTime"`    // 优惠券结束时间
	CouponPrice      float64 `json:"couponPrice"`      // 优惠券金额
	CouponConditions string  `json:"couponConditions"` // 优惠券最低消费
	CouponTotalNum   int64   `json:"couponTotalNum"`   // 优惠券总数

	// 活动相关信息
	ActivityType      TaoBaoActivityType `json:"activityType"`      // 活动类型
	ActivityStartTime string             `json:"activityStartTime"` // 活动开始时间
	ActivityEndTime   string             `json:"activityEndTime"`   // 活动结束时间

	// 评价相关信息
	DescScore      float64 `json:"descScore"`      // 描述分
	DsrScore       float64 `json:"dsrScore"`       // 描述相符
	DsrPercent     float64 `json:"dsrPercent"`     // 描述同行比
	ShipScore      float64 `json:"shipScore"`      // 服务评分
	ShipPercent    float64 `json:"shipPercent"`    // 服务同行比
	ServiceScore   float64 `json:"serviceScore"`   // 物流评分
	ServicePercent float64 `json:"servicePercent"` // 物流同行比

	// 佣金相关信息
	CommissionType TaoBaoCommissionType `json:"commissionType"` // 佣金类型
	CommissionRate float64              `json:"commissionRate"` // 佣金比例
}

type TaoBaoOrderSt struct {
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

// TaoBaoOrderQueryType
type TaoBaoOrderQueryType int64
