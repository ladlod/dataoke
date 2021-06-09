package dataoke

type CommonRespHeader struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

type ErrorResponse struct {
	Msg     string `json:"msg"`
	Code    int64  `json:"code"`
	SubMsg  string `json:"sub_msg"`
	SubCode int64  `json:"sub_code"`
}

type TaobaoGoodsSt struct {
	// 商品相关信息
	Id                     int64    `json:"id"`                     // 大淘客商品id
	GoodsId                string   `json:"goodsId"`                // 淘宝商品id
	Title                  string   `json:"title"`                  // 淘宝标题
	Dtitle                 string   `json:"dtitle"`                 // 大淘客短标题
	ItemLink               string   `json:"itemLink"`               // 商品链接
	OriginalPrice          float64  `json:"originalPrice"`          // 原价
	ActualPrice            float64  `json:"actualPrice"`            // 现价
	MonthSales             int64    `json:"monthSales"`             // 近30天销量
	TwoHoursSales          int64    `json:"twoHoursSales"`          // 近两小时销量
	DailySales             int64    `json:"dailySales"`             // 当日销量
	Desc                   string   `json:"desc"`                   // 推广文案
	CreateTime             string   `json:"createTime"`             // 商品创建时间
	MainPic                string   `json:"mainPic"`                // 商品主图
	MarketingMainPic       string   `json:"marketingMainPic"`       // 营销主图
	SpecialText            []string `json:"specialText"`            // 特色文案
	Video                  string   `json:"video"`                  // 商品视频
	Cid                    int64    `json:"cid"`                    // 大淘客分类id
	Subcid                 []int64  `json:"subcid"`                 // 大淘客二级分类id
	Tbcid                  int64    `json:"tbcid"`                  // 淘宝品类id
	Haitao                 int64    `json:"haitao"`                 // 是否为海淘
	Tchaoshi               int64    `json:"tchaoshi"`               // 是否为天猫超市
	Yunfeixian             int64    `json:"yunfeixian"`             // 是否有运费险
	FreeshipRemoteDistrict int64    `json:"freeshipRemoteDistrict"` // 偏远地区是否包邮

	// 店铺相关信息
	ShopType    TaoBaoShopType `json:"shopType"`    // 店铺类型
	GoldSellers int64          `json:"goldSellers"` // 是否为金牌卖家
	SellerId    string         `json:"sellerId"`    // 淘宝卖家id
	ShopName    string         `json:"shopName"`    // 店铺名称
	ShopLogo    string         `json:"shopLogo"`    // 店铺logo
	ShopLevel   interface{}    `json:"shopLevel"`   // 店铺等级, 因字段类型不明确，暂时使用interface类型

	// 优惠券相关信息
	CouponLink       string  `json:"couponLink"`       // 领券链接
	CouponStartTime  string  `json:"couponStartTime"`  // 优惠券开始时间
	CouponEndTime    string  `json:"couponEndTime"`    // 优惠券结束时间
	CouponPrice      float64 `json:"couponPrice"`      // 优惠券金额
	CouponConditions string  `json:"couponConditions"` // 优惠券最低消费
	CouponReceiveNum int64   `json:"couponReceiveNum"` // 领券量
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

type TaobaoGoodsDetailSt struct {
	TaobaoGoodsSt
	DetailPics string `json:"detailPics"`
	Imgs       string `json:"imgs"`
	Reimgs     string `json:"reimgs"`
}

type TaoBaoOrderSt struct {
	// 订单基本信息
	TradeId          string           `json:"trade_id"`           // 订单编号
	TradeParentId    string           `json:"trade_parent_id"`    // 买家在淘宝后台显示的订单编号
	TkStatus         TaoBaoOrderState `json:"tk_status"`          // 订单状态
	PayPrice         string           `json:"pay_price"`          // 买家确认收货的付款金额（不包含运费金额）
	AlipayTotalPrice string           `json:"alipay_total_price"` // 买家拍下付款的金额（不包含运费金额）
	DepositPrice     string           `json:"deposit_price"`      // 预售时期，用户对预售商品支付的定金金额
	RefundTag        int64            `json:"refund_tag"`         // 维权标签 0.非维权 1.维权订单
	OrderType        string           `json:"order_type"`         // 订单所属平台类型，包括天猫、淘宝、聚划算等
	TerminalType     string           `json:"terminal_type"`      // 成交设备类型

	// 时间信息
	ClickTime     string `json:"click_time"`      // 点击时间
	ModifiedTime  string `json:"modified_time"`   // 订单更新时间
	TkCreateTime  string `json:"tk_create_time"`  // 订单创建的时间
	TbPaidTime    string `json:"tb_paid_time"`    // 订单在淘宝拍下付款的时间
	TbDepositTime string `json:"tb_deposit_time"` // 预售时期，用户对预售商品支付定金的付款时间
	TkEarningTime string `json:"tk_earning_time"` // 订单确认收货后且商家完成佣金支付的时间

	// 商品信息
	ItemId           int64  `json:"item_id"`            // 商品id
	ItemPrice        string `json:"item_price"`         // 商品单价
	ItemImg          string `json:"item_img"`           // 商品主图
	ItemTitle        string `json:"item_title"`         // 商品标题
	ItemLink         string `json:"item_link"`          // 商品链接
	ItemNum          int64  `json:"item_num"`           // 商品数量
	ItemCategoryName string `json:"item_category_name"` // 商品所属的根类目，即一级类目的名称
	FlowSource       string `json:"flow_source"`        // 产品类型
	SellerNick       string `json:"seller_nick"`        // 掌柜旺旺
	SellerShopTitle  string `json:"seller_shop_title"`  // 店铺名称

	// 结算信息
	TkOrderRole         int64  `json:"tk_order_role"`         // 2.佣金收益的第一归属者 3.从其他淘宝客佣金中进行分成的推广者
	PubShareFee         string `json:"pub_share_fee"`         // 结算实际收入
	PubSharePreFee      string `json:"pub_share_pre_fee"`     // 结算预估收入
	PubShareRate        string `json:"pub_share_rate"`        // 从结算佣金中分得的收益比率
	SubsidyRate         string `json:"subsidy_rate"`          // 平台给与的补贴比率
	TkTotalRate         string `json:"tk_total_rate"`         // 提成=收入比率×分成比率。指实际获得收益的比率
	AlimamaRate         string `json:"alimama_rate"`          // 推广者赚取佣金后支付给阿里妈妈的技术服务费用的比率
	AlimamaShareFee     string `json:"alimama_share_fee"`     // 技术服务费=结算金额×收入比率×技术服务费率
	SubsidyType         string `json:"subsidy_type"`          // 平台出资方，如天猫、淘宝、或聚划算等
	SubsidyFee          string `json:"subsidy_fee"`           // 补贴金额=结算金额×补贴比率
	TotalCommissionRate string `json:"total_commission_rate"` // 佣金比率
	TotalCommissionFee  string `json:"total_commission_fee"`  // 佣金金额=结算金额＊佣金比率
	IncomeRate          string `json:"income_rate"`           // 订单结算的佣金比率+平台的补贴比率
	ShareRelativeRate   string `json:"share_relative_rate"`   // 专项服务费率
	ShareFee            string `json:"share_fee"`             // 结算专项服务费

	// 归因信息
	AdzoneId   int64  `json:"adzone_id"`   // 推广位管理下的推广位名称对应的ID
	PubId      int64  `json:"pub_id"`      // 推广者的会员id
	SiteName   string `json:"site_name"`   // 媒体管理下的对应ID的自定义名称
	AdzoneName string `json:"adzone_name"` // 推广位名称
	SiteId     int64  `json:"site_id"`     // 媒体管理下的ID
	SpecialId  int64  `json:"special_id"`  // 会员运营id
	RelationId int64  `json:"relation_id"` // 渠道关系id
	AppKey     string `json:"app_key"`     // 开发者调用api的appkey
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

const (
	QueryTypeCt = 1 // 订单创建时间
	QueryTypePt = 2 // 订单支付时间
	QueryTypeFt = 3 // 订单结算时间
	QueryTypeUt = 4 // 订单更新时间
)

type TaoBaoOrderState int64

const (
	OrderStateTimeout = 1
	OrderStateCancel  = 2
	OrderStateRefund  = 3
)
