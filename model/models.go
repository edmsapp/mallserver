package model

type Point_log struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Balance         int    `gorm:"balance" json:"balance"`                     // 当前积分
	Credit          int    `gorm:"credit" json:"credit"`                       // 获取积分
	Debit           int    `gorm:"debit" json:"debit"`                         // 扣除积分
	Memo            string `gorm:"memo" json:"memo"`                           // 备注
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Operator        string `gorm:"operator" json:"operator"`                   // 操作员
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdateBy    string `gorm:"last_update_by" json:"last_update_by"`       // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Gift_item struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Gift            int    `gorm:"gift" json:"gift"`                           // 赠品
	Quantity        int    `gorm:"quantity" json:"quantity"`                   // 数量
	PromotionId     int    `gorm:"promotion_id" json:"promotion_id"`           // 促销
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Navigation struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Url             string `gorm:"url" json:"url"`                             // 链接地址
	IsBlankTarget   int    `gorm:"is_blank_target" json:"is_blank_target"`     // 是否新窗口打开
	Position        int    `gorm:"position" json:"position"`                   // 位置
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Payment_method struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Method          int    `gorm:"method" json:"method"`                       // 方式
	Timeout         int    `gorm:"timeout" json:"timeout"`                     // 超时时间
	Icon            string `gorm:"icon" json:"icon"`                           // 图标
	Description     string `gorm:"description" json:"description"`             // 介绍
	Content         string `gorm:"content" json:"content"`                     // 内容
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	Type            int    `gorm:"type" json:"type"`                           // 类型
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Specification struct {
	Id                int    `gorm:"id" json:"id"`                                   // 主键_id
	Name              string `gorm:"name" json:"name"`                               // 名称
	ProductCategoryId int    `gorm:"product_category_id" json:"product_category_id"` // 绑定分类
	Orders            int    `gorm:"orders" json:"orders"`                           // 排序
	CreateBy          string `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int    `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Member_attribute struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Type            int    `gorm:"type" json:"type"`                           // 类型
	PropertyIndex   int    `gorm:"property_index" json:"property_index"`       // 属性序号
	IsEnabled       int    `gorm:"is_enabled" json:"is_enabled"`               // 是否启用
	IsRequired      int    `gorm:"is_required" json:"is_required"`             // 是否必填
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Order_item struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Sn              string  `gorm:"sn" json:"sn"`                               // 商品编号
	Name            string  `gorm:"name" json:"name"`                           // 商品名称
	Price           float64 `gorm:"price" json:"price"`                         // 商品价格
	Weight          int     `gorm:"weight" json:"weight"`                       // 商品重量
	Thumbnail       string  `gorm:"thumbnail" json:"thumbnail"`                 // 商品缩略图
	Specifications  string  `gorm:"specifications" json:"specifications"`       // 规格
	Quantity        int     `gorm:"quantity" json:"quantity"`                   // 数量
	ShippedQuantity int     `gorm:"shipped_quantity" json:"shipped_quantity"`   // 已发货数量
	ReturnQuantity  int     `gorm:"return_quantity" json:"return_quantity"`     // 已退货数量
	IsGift          int     `gorm:"is_gift" json:"is_gift"`                     // 是否为赠品
	IsDelivery      int     `gorm:"is_delivery" json:"is_delivery"`             // 是否需要物流
	OrderId         int     `gorm:"order_id" json:"order_id"`                   // 订单
	SkuId           int     `gorm:"sku_id" json:"sku_id"`                       // SKU
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_image struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Title           string `gorm:"title" json:"title"`                         // 标题
	ProductId       int    `gorm:"product_id" json:"product_id"`               // 商品
	Source          string `gorm:"source" json:"source"`                       // 原图片
	Large           string `gorm:"large" json:"large"`                         // 大图片
	Medium          string `gorm:"medium" json:"medium"`                       // 中图片
	Thumbnail       string `gorm:"thumbnail" json:"thumbnail"`                 // 缩略图
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Delivery_center struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Contact         string `gorm:"contact" json:"contact"`                     // 联系人
	Mobile          string `gorm:"mobile" json:"mobile"`                       // 手机
	Phone           string `gorm:"phone" json:"phone"`                         // 电话
	AreaId          int    `gorm:"area_id" json:"area_id"`                     // 地区
	AreaName        string `gorm:"area_name" json:"area_name"`                 // 地区名称
	ZipCode         string `gorm:"zip_code" json:"zip_code"`                   // 邮编
	Address         string `gorm:"address" json:"address"`                     // 地址
	IsDefault       int    `gorm:"is_default" json:"is_default"`               // 是否默认
	Memo            string `gorm:"memo" json:"memo"`                           // 备注
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Message_config struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_ID
	IsMailEnabled   int    `gorm:"is_mail_enabled" json:"is_mail_enabled"`     // 是否启用邮件
	IsSmsEnabled    int    `gorm:"is_sms_enabled" json:"is_sms_enabled"`       // 是否启用短信
	Type            int    `gorm:"type" json:"type"`                           // 类型
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Payment struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Sn              string  `gorm:"sn" json:"sn"`                               // 编号
	Method          int     `gorm:"method" json:"method"`                       // 方式
	PaymentMethod   string  `gorm:"payment_method" json:"payment_method"`       // 支付方式
	Bank            string  `gorm:"bank" json:"bank"`                           // 收款银行
	Account         string  `gorm:"account" json:"account"`                     // 收款账号
	Fee             float64 `gorm:"fee" json:"fee"`                             // 支付手续费
	Amount          float64 `gorm:"amount" json:"amount"`                       // 付款金额
	Payer           string  `gorm:"payer" json:"payer"`                         // 付款人
	Operator        string  `gorm:"operator" json:"operator"`                   // 操作员
	PaymentDate     string  `gorm:"payment_date" json:"payment_date"`           // 付款日期
	Memo            string  `gorm:"memo" json:"memo"`                           // 备注
	PaymentPluginId string  `gorm:"payment_plugin_id" json:"payment_plugin_id"` // 支付插件ID
	Expire          string  `gorm:"expire" json:"expire"`                       // 到期时间
	MemberId        int     `gorm:"member_id" json:"member_id"`                 // 会员
	OrderId         int     `gorm:"order_id" json:"order_id"`                   // 订单
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Brand struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Introduction    string `gorm:"introduction" json:"introduction"`           // 介绍
	Url             string `gorm:"url" json:"url"`                             // 网址
	Logo            string `gorm:"logo" json:"logo"`                           // logo
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Member struct {
	Id                int     `gorm:"id" json:"id"`                                   // 主键_id
	Username          string  `gorm:"username" json:"username"`                       // 用户名
	Password          string  `gorm:"password" json:"password"`                       // 密码
	Name              string  `gorm:"name" json:"name"`                               // 姓名
	Avatar            string  `gorm:"avatar" json:"avatar"`                           // 头像
	Gender            int     `gorm:"gender" json:"gender"`                           // 性别,0未知,1男,2女性别
	Mobile            string  `gorm:"mobile" json:"mobile"`                           // 手机
	Phone             string  `gorm:"phone" json:"phone"`                             // 电话
	Birth             string  `gorm:"birth" json:"birth"`                             // 出生日期
	ZipCode           string  `gorm:"zip_code" json:"zip_code"`                       // 邮编
	OpenId            string  `gorm:"open_id" json:"open_id"`                         // openID
	AreaId            int     `gorm:"area_id" json:"area_id"`                         // 地区
	Address           string  `gorm:"address" json:"address"`                         // 地址
	Email             string  `gorm:"email" json:"email"`                             // E-mail
	Amount            float64 `gorm:"amount" json:"amount"`                           // 消费金额
	Balance           float64 `gorm:"balance" json:"balance"`                         // 余额
	Point             int     `gorm:"point" json:"point"`                             // 积分
	MemberRankId      int     `gorm:"member_rank_id" json:"member_rank_id"`           // 会员等级
	RegisterIp        string  `gorm:"register_ip" json:"register_ip"`                 // 注册IP
	AttributeValue0   string  `gorm:"attribute_value0" json:"attribute_value0"`       // 会员注册项值0
	AttributeValue1   string  `gorm:"attribute_value1" json:"attribute_value1"`       // 会员注册项值1
	AttributeValue2   string  `gorm:"attribute_value2" json:"attribute_value2"`       // 会员注册项值2
	AttributeValue3   string  `gorm:"attribute_value3" json:"attribute_value3"`       // 会员注册项值3
	AttributeValue4   string  `gorm:"attribute_value4" json:"attribute_value4"`       // 会员注册项值4
	AttributeValue5   string  `gorm:"attribute_value5" json:"attribute_value5"`       // 会员注册项值5
	AttributeValue6   string  `gorm:"attribute_value6" json:"attribute_value6"`       // 会员注册项值6
	AttributeValue7   string  `gorm:"attribute_value7" json:"attribute_value7"`       // 会员注册项值7
	AttributeValue8   string  `gorm:"attribute_value8" json:"attribute_value8"`       // 会员注册项值8
	AttributeValue9   string  `gorm:"attribute_value9" json:"attribute_value9"`       // 会员注册项值9
	IsEnabled         int     `gorm:"is_enabled" json:"is_enabled"`                   // 是否启用
	IsLocked          int     `gorm:"is_locked" json:"is_locked"`                     // 是否锁定
	LockedDate        string  `gorm:"locked_date" json:"locked_date"`                 // 锁定日期
	LoginDate         string  `gorm:"login_date" json:"login_date"`                   // 最后登录日期
	LoginFailureCount int     `gorm:"login_failure_count" json:"login_failure_count"` // 连续登录失败次数
	LoginIp           string  `gorm:"login_ip" json:"login_ip"`                       // 最后登录IP
	SafeKeyExpire     string  `gorm:"safe_key_expire" json:"safe_key_expire"`         // 安全密匙有效期
	SafeKeyValue      string  `gorm:"safe_key_value" json:"safe_key_value"`           // 安全密匙值
	CreateBy          string  `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string  `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string  `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string  `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int     `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Parameter_group struct {
	Id                int    `gorm:"id" json:"id"`                                   // 主键_id
	Name              string `gorm:"name" json:"name"`                               // 名称
	ProductCategoryId int    `gorm:"product_category_id" json:"product_category_id"` // 绑定分类
	Orders            int    `gorm:"orders" json:"orders"`                           // 排序
	CreateBy          string `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int    `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Promotion_product struct {
	Promotions int `gorm:"promotions" json:"promotions"`
	Products   int `gorm:"products" json:"products"`
}

type Promotion_product_category struct {
	Promotions        int `gorm:"promotions" json:"promotions"`                 // 促销Id
	ProductCategories int `gorm:"product_categories" json:"product_categories"` // 商品分类
}

type Stock_log struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	InQuantity      int    `gorm:"in_quantity" json:"in_quantity"`             // 入库数量
	OutQuantity     int    `gorm:"out_quantity" json:"out_quantity"`           // 出库数量
	Stock           int    `gorm:"stock" json:"stock"`                         // 当前库存
	SkuId           int    `gorm:"sku_id" json:"sku_id"`                       // SKU
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Memo            string `gorm:"memo" json:"memo"`                           // 备注
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Tag struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Icon            string `gorm:"icon" json:"icon"`                           // 图标
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	Memo            string `gorm:"memo" json:"memo"`                           // 备注
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Article struct {
	Id                int    `gorm:"id" json:"id"`                                   // 主键_id
	Title             string `gorm:"title" json:"title"`                             // 标题
	Author            string `gorm:"author" json:"author"`                           // 作者
	Content           string `gorm:"content" json:"content"`                         // 内容
	SeoTitle          string `gorm:"seo_title" json:"seo_title"`                     // 页面标题
	SeoDescription    string `gorm:"seo_description" json:"seo_description"`         // 页面描述
	SeoKeywords       string `gorm:"seo_keywords" json:"seo_keywords"`               // 页面关键词
	IsPublication     int    `gorm:"is_publication" json:"is_publication"`           // 是否发布
	IsTop             int    `gorm:"is_top" json:"is_top"`                           // 是否置顶
	Hits              int    `gorm:"hits" json:"hits"`                               // 点击数
	ArticleCategoryId int    `gorm:"article_category_id" json:"article_category_id"` // 文章分类
	CreateBy          string `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int    `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Permission struct {
	Id              int    `gorm:"id" json:"id"`
	Name            string `gorm:"name" json:"name"`                           // 名称
	Value           string `gorm:"value" json:"value"`                         // 值
	Module          string `gorm:"module" json:"module"`                       // 所属模块
	Description     string `gorm:"description" json:"description"`             // 简介
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_notify struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Email           string `gorm:"email" json:"email"`                         // E-mail
	HasSent         int    `gorm:"has_sent" json:"has_sent"`                   // 是否已发送
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	ProductId       int    `gorm:"product_id" json:"product_id"`               // 商品
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Order_log struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Content         string `gorm:"content" json:"content"`                     // 内容
	Operator        string `gorm:"operator" json:"operator"`                   // 操作员
	OrderId         int    `gorm:"order_id" json:"order_id"`                   // 订单
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Payment_shipping_method struct {
	PaymentMethods  int `gorm:"payment_methods" json:"payment_methods"`
	ShippingMethods int `gorm:"shipping_methods" json:"shipping_methods"`
}

type Permission_role struct {
	Permissions int `gorm:"permissions" json:"permissions"`
	Roles       int `gorm:"roles" json:"roles"` // 角色
}

type Role struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Value           string `gorm:"value" json:"value"`                         // 值
	Description     string `gorm:"description" json:"description"`             // 简介
	IsSystem        int    `gorm:"is_system" json:"is_system"`                 // 是否系统角色
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Ad struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Title           string `gorm:"title" json:"title"`                         // 标题
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Content         string `gorm:"content" json:"content"`                     // 内容
	AdPositionId    int    `gorm:"ad_position_id" json:"ad_position_id"`       // 广告位
	Path            string `gorm:"path" json:"path"`                           // 路径
	Url             string `gorm:"url" json:"url"`                             // 链接地址
	BeginDate       string `gorm:"begin_date" json:"begin_date"`               // 起始日期
	EndDate         string `gorm:"end_date" json:"end_date"`                   // 结束日期
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Area_freight_config struct {
	Id               int     `gorm:"id" json:"id"`                                 // 主键_ID
	FirstPrice       float64 `gorm:"first_price" json:"first_price"`               // 首重价格
	ContinuePrice    float64 `gorm:"continue_price" json:"continue_price"`         // 续重价格
	ShippingMethodId int     `gorm:"shipping_method_id" json:"shipping_method_id"` // 配送方式
	AreaId           int     `gorm:"area_id" json:"area_id"`                       // 地区
	CreateBy         string  `gorm:"create_by" json:"create_by"`                   // 创建人
	CreationDate     string  `gorm:"creation_date" json:"creation_date"`           // 创建日期
	LastUpdatedBy    string  `gorm:"last_updated_by" json:"last_updated_by"`       // 最后修改人
	LastUpdatedDate  string  `gorm:"last_updated_date" json:"last_updated_date"`   // 最后修改日期
	DeleteFlag       int     `gorm:"delete_flag" json:"delete_flag"`               // 删除标记
}

type Cart_item struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	CartId          int    `gorm:"cart_id" json:"cart_id"`                     // 购物车
	SkuId           int    `gorm:"sku_id" json:"sku_id"`                       // SKU
	Quantity        int    `gorm:"quantity" json:"quantity"`                   // 数量
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product struct {
	Id                int     `gorm:"id" json:"id"`                                   // 主键_id
	Sn                string  `gorm:"sn" json:"sn"`                                   // 编号
	Name              string  `gorm:"name" json:"name"`                               // 名称
	Caption           string  `gorm:"caption" json:"caption"`                         // 副标题
	Image             string  `gorm:"image" json:"image"`                             // 展示图片
	Unit              string  `gorm:"unit" json:"unit"`                               // 单位
	Weight            int     `gorm:"weight" json:"weight"`                           // 重量
	Type              int     `gorm:"type" json:"type"`                               // 类型
	IsMarketable      int     `gorm:"is_marketable" json:"is_marketable"`             // 是否上架
	IsList            int     `gorm:"is_list" json:"is_list"`                         // 是否列出
	IsTop             int     `gorm:"is_top" json:"is_top"`                           // 是否置顶
	IsDelivery        int     `gorm:"is_delivery" json:"is_delivery"`                 // 是否需要物流
	IsGift            int     `gorm:"is_gift" json:"is_gift"`                         // 是否为赠品
	Introduction      string  `gorm:"introduction" json:"introduction"`               // 介绍
	Memo              string  `gorm:"memo" json:"memo"`                               // 备注
	Keyword           string  `gorm:"keyword" json:"keyword"`                         // 搜索关键词
	SeoTitle          string  `gorm:"seo_title" json:"seo_title"`                     // 页面标题
	SeoKeywords       string  `gorm:"seo_keywords" json:"seo_keywords"`               // 页面关键词
	SeoDescription    string  `gorm:"seo_description" json:"seo_description"`         // 页面描述
	Score             float64 `gorm:"score" json:"score"`                             // 评分
	TotalScore        int     `gorm:"total_score" json:"total_score"`                 // 总评分
	ScoreCount        int     `gorm:"score_count" json:"score_count"`                 // 评分数
	Hits              int     `gorm:"hits" json:"hits"`                               // 点击数
	WeekHits          int     `gorm:"week_hits" json:"week_hits"`                     // 周点击数
	MonthHits         int     `gorm:"month_hits" json:"month_hits"`                   // 月点击数
	Sales             int     `gorm:"sales" json:"sales"`                             // 销量
	WeekSales         int     `gorm:"week_sales" json:"week_sales"`                   // 周销量
	MonthSales        int     `gorm:"month_sales" json:"month_sales"`                 // 月销量
	WeekHitsDate      string  `gorm:"week_hits_date" json:"week_hits_date"`           // 周点击数更新日期
	MonthHitsDate     string  `gorm:"month_hits_date" json:"month_hits_date"`         // 月点击数更新日期
	WeekSalesDate     string  `gorm:"week_sales_date" json:"week_sales_date"`         // 周销量更新日期
	MonthSalesDate    string  `gorm:"month_sales_date" json:"month_sales_date"`       // 月销量更新日期
	AttributeValue0   string  `gorm:"attribute_value0" json:"attribute_value0"`       // 商品属性值0
	AttributeValue1   string  `gorm:"attribute_value1" json:"attribute_value1"`       // 商品属性值1
	AttributeValue2   string  `gorm:"attribute_value2" json:"attribute_value2"`       // 商品属性值2
	AttributeValue3   string  `gorm:"attribute_value3" json:"attribute_value3"`       // 商品属性值3
	AttributeValue4   string  `gorm:"attribute_value4" json:"attribute_value4"`       // 商品属性值4
	AttributeValue5   string  `gorm:"attribute_value5" json:"attribute_value5"`       // 商品属性值5
	AttributeValue6   string  `gorm:"attribute_value6" json:"attribute_value6"`       // 商品属性值6
	AttributeValue7   string  `gorm:"attribute_value7" json:"attribute_value7"`       // 商品属性值7
	AttributeValue8   string  `gorm:"attribute_value8" json:"attribute_value8"`       // 商品属性值8
	AttributeValue9   string  `gorm:"attribute_value9" json:"attribute_value9"`       // 商品属性值9
	AttributeValue10  string  `gorm:"attribute_value10" json:"attribute_value10"`     // 商品属性值10
	AttributeValue11  string  `gorm:"attribute_value11" json:"attribute_value11"`     // 商品属性值11
	AttributeValue12  string  `gorm:"attribute_value12" json:"attribute_value12"`     // 商品属性值12
	AttributeValue13  string  `gorm:"attribute_value13" json:"attribute_value13"`     // 商品属性值13
	AttributeValue14  string  `gorm:"attribute_value14" json:"attribute_value14"`     // 商品属性值14
	AttributeValue15  string  `gorm:"attribute_value15" json:"attribute_value15"`     // 商品属性值15
	AttributeValue16  string  `gorm:"attribute_value16" json:"attribute_value16"`     // 商品属性值16
	AttributeValue17  string  `gorm:"attribute_value17" json:"attribute_value17"`     // 商品属性值17
	AttributeValue18  string  `gorm:"attribute_value18" json:"attribute_value18"`     // 商品属性值18
	AttributeValue19  string  `gorm:"attribute_value19" json:"attribute_value19"`     // 商品属性19
	BrandId           int     `gorm:"brand_id" json:"brand_id"`                       // 品牌
	ProductCategoryId int     `gorm:"product_category_id" json:"product_category_id"` // 商品分类
	CreateBy          string  `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string  `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string  `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string  `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int     `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Promotion_member_rank struct {
	Promotions  int `gorm:"promotions" json:"promotions"`
	MemberRanks int `gorm:"member_ranks" json:"member_ranks"`
}

type Register_code struct {
	Mobile         string `gorm:"mobile" json:"mobile"`                   // 接收短信的手机号码
	SmsCode        string `gorm:"sms_code" json:"sms_code"`               // 短信内容
	Op             string `gorm:"op" json:"op"`                           // 短信类型
	ExpirationDate string `gorm:"expiration_date" json:"expiration_date"` // 结束时间
}

type Seo struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Title           string `gorm:"title" json:"title"`                         // 页面标题
	Keywords        string `gorm:"keywords" json:"keywords"`                   // 页面关键词
	Description     string `gorm:"description" json:"description"`             // 页面描述
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Shipping_method struct {
	Id                    int     `gorm:"id" json:"id"`
	Name                  string  `gorm:"name" json:"name"`                                         // 名称
	FirstWeight           int     `gorm:"first_weight" json:"first_weight"`                         // 首重量
	ContinueWeight        int     `gorm:"continue_weight" json:"continue_weight"`                   // 续重量
	FirstPrice            float64 `gorm:"first_price" json:"first_price"`                           // 首重价格
	ContinuePrice         float64 `gorm:"continue_price" json:"continue_price"`                     // 续重价格
	Orders                int     `gorm:"orders" json:"orders"`                                     // 订单
	Icon                  string  `gorm:"icon" json:"icon"`                                         // 图标
	Description           string  `gorm:"description" json:"description"`                           // description
	DefaultDeliveryCorpId int     `gorm:"default_delivery_corp_id" json:"default_delivery_corp_id"` // 默认物流公司
	CreateBy              string  `gorm:"create_by" json:"create_by"`                               // 创建人
	CreationDate          string  `gorm:"creation_date" json:"creation_date"`                       // 创建日期
	LastUpdatedBy         string  `gorm:"last_updated_by" json:"last_updated_by"`                   // 最后修改人
	LastUpdatedDate       string  `gorm:"last_updated_date" json:"last_updated_date"`               // 最后修改日期
	DeleteFlag            int     `gorm:"delete_flag" json:"delete_flag"`                           // 删除标记
}

type Delivery_corp struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Url             string `gorm:"url" json:"url"`                             // 网址
	Code            string `gorm:"code" json:"code"`                           // 代码
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Feedback struct {
	Id           int    `gorm:"id" json:"id"`
	MemberId     int    `gorm:"member_id" json:"member_id"`         // 会员ID
	CreationDate string `gorm:"creation_date" json:"creation_date"` // 反馈的时间
	Suggestion   string `gorm:"suggestion" json:"suggestion"`       // 反馈内容
}

type Friend_link struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Type            int    `gorm:"type" json:"type"`                           // 类型
	Logo            string `gorm:"logo" json:"logo"`                           // logo
	Url             string `gorm:"url" json:"url"`                             // 链接地址
	Orders          int    `gorm:"orders" json:"orders"`                       // 订单
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Order_coupon struct {
	Orders  int `gorm:"orders" json:"orders"`   // 订单
	Coupons int `gorm:"coupons" json:"coupons"` // 优惠券
}

type Plugin_config struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	PluginId        string `gorm:"plugin_id" json:"plugin_id"`                 // 插件ID
	Attributes      string `gorm:"attributes" json:"attributes"`               // 属性
	IsEnabled       int    `gorm:"is_enabled" json:"is_enabled"`               // 是否启用
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	Type            int    `gorm:"type" json:"type"`                           // 类型
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Promotion_gift struct {
	Promotions int `gorm:"promotions" json:"promotions"`
	Gifts      int `gorm:"gifts" json:"gifts"`
}

type Member_attribute_option struct {
	MemberAttribute int    `gorm:"member_attribute" json:"member_attribute"`
	Options         string `gorm:"options" json:"options"`
}

type Member_favorite_product struct {
	FavoriteMembers  int `gorm:"favorite_members" json:"favorite_members"`
	FavoriteProducts int `gorm:"favorite_products" json:"favorite_products"`
}

type Order struct {
	Id                 int     `gorm:"id" json:"id"`                                     // 主键_id
	Sn                 string  `gorm:"sn" json:"sn"`                                     // 订单编号
	Status             int     `gorm:"status" json:"status"`                             // 订单状态
	Amount             float64 `gorm:"amount" json:"amount"`                             // 订单金额
	AmountPaid         float64 `gorm:"amount_paid" json:"amount_paid"`                   // 已付金额
	CouponDiscount     float64 `gorm:"coupon_discount" json:"coupon_discount"`           // 优惠券折扣
	OffsetAmount       float64 `gorm:"offset_amount" json:"offset_amount"`               // 调整金额
	Price              float64 `gorm:"price" json:"price"`                               // 商品价格
	Fee                float64 `gorm:"fee" json:"fee"`                                   // 支付手续费
	Freight            float64 `gorm:"freight" json:"freight"`                           // 运费
	RewardPoint        int     `gorm:"reward_point" json:"reward_point"`                 // 赠送积分
	Consignee          string  `gorm:"consignee" json:"consignee"`                       // 收货人
	AreaName           string  `gorm:"area_name" json:"area_name"`                       // 地区名称
	Address            string  `gorm:"address" json:"address"`                           // 地址
	ZipCode            string  `gorm:"zip_code" json:"zip_code"`                         // 邮编
	Phone              string  `gorm:"phone" json:"phone"`                               // 电话
	ExchangePoint      int     `gorm:"exchange_point" json:"exchange_point"`             // 兑换积分
	IsAllocatedStock   int     `gorm:"is_allocated_stock" json:"is_allocated_stock"`     // 是否已分配库存
	IsExchangePoint    int     `gorm:"is_exchange_point" json:"is_exchange_point"`       // 是否已兑换积分
	IsUseCouponCode    int     `gorm:"is_use_coupon_code" json:"is_use_coupon_code"`     // 是否已使用优惠码
	IsNotify           int     `gorm:"is_notify" json:"is_notify"`                       // 是否通知
	InvoiceContent     string  `gorm:"invoice_content" json:"invoice_content"`           // 发票内容
	InvoiceTitle       string  `gorm:"invoice_title" json:"invoice_title"`               // 发票标题
	Tax                float64 `gorm:"tax" json:"tax"`                                   // 税金
	Memo               string  `gorm:"memo" json:"memo"`                                 // 附言
	PaymentMethodName  string  `gorm:"payment_method_name" json:"payment_method_name"`   // 支付方式名称
	PaymentMethodType  int     `gorm:"payment_method_type" json:"payment_method_type"`   // 支付方式类型
	PromotionName      string  `gorm:"promotion_name" json:"promotion_name"`             // 促销名称
	PromotionDiscount  float64 `gorm:"promotion_discount" json:"promotion_discount"`     // 促销折扣
	Expire             string  `gorm:"expire" json:"expire"`                             // 到期时间
	LockKey            string  `gorm:"lock_key" json:"lock_key"`                         // 锁定KEY
	LockExpire         string  `gorm:"lock_expire" json:"lock_expire"`                   // 锁定到期时间
	Weight             int     `gorm:"weight" json:"weight"`                             // 重量
	Source             int     `gorm:"source" json:"source"`                             // 订单来源
	Type               int     `gorm:"type" json:"type"`                                 // 类型
	Quantity           int     `gorm:"quantity" json:"quantity"`                         // 商品数量
	RefundAmount       float64 `gorm:"refund_amount" json:"refund_amount"`               // 退款金额
	ReturnedQuantity   int     `gorm:"returned_quantity" json:"returned_quantity"`       // 退款商品数量
	ShippedQuantity    int     `gorm:"shipped_quantity" json:"shipped_quantity"`         // 已发货数量
	ShippingMethodName string  `gorm:"shipping_method_name" json:"shipping_method_name"` // 配送方式名称
	ShippingDate       string  `gorm:"shipping_date" json:"shipping_date"`               // 配送时间
	PaymentMethodId    int     `gorm:"payment_method_id" json:"payment_method_id"`       // 支付方式
	ShippingMethodId   int     `gorm:"shipping_method_id" json:"shipping_method_id"`     // 配送方式
	DeliveryCorpId     int     `gorm:"delivery_corp_id" json:"delivery_corp_id"`         // 物流公司
	DeliveryCorpName   string  `gorm:"delivery_corp_name" json:"delivery_corp_name"`     // 物流公司名称
	OperatorId         int     `gorm:"operator_id" json:"operator_id"`                   // 操作员
	AreaId             int     `gorm:"area_id" json:"area_id"`                           // 地区
	MemberId           int     `gorm:"member_id" json:"member_id"`                       // 会员
	CouponCodeId       int     `gorm:"coupon_code_id" json:"coupon_code_id"`             // 优惠码
	CompleteDate       string  `gorm:"complete_date" json:"complete_date"`               // 完成日期
	CreateBy           string  `gorm:"create_by" json:"create_by"`                       // 创建人
	CreationDate       string  `gorm:"creation_date" json:"creation_date"`               // 创建日期
	LastUpdatedBy      string  `gorm:"last_updated_by" json:"last_updated_by"`           // 最后修改人
	LastUpdatedDate    string  `gorm:"last_updated_date" json:"last_updated_date"`       // 最后修改日期
	DeleteFlag         int     `gorm:"delete_flag" json:"delete_flag"`                   // 删除标记
}

type Delivery_template struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Content         string `gorm:"content" json:"content"`                     // 内容
	Background      string `gorm:"background" json:"background"`               // 背景图
	Height          int    `gorm:"height" json:"height"`                       // 高度
	Width           int    `gorm:"width" json:"width"`                         // 宽度
	Offsetx         int    `gorm:"offsetx" json:"offsetx"`                     // 偏移量X
	Offsety         int    `gorm:"offsety" json:"offsety"`                     // 偏移量Y
	IsDefault       int    `gorm:"is_default" json:"is_default"`               // 是否默认
	Memo            string `gorm:"memo" json:"memo"`                           // 备注
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_tag struct {
	Products int `gorm:"products" json:"products"`
	Tags     int `gorm:"tags" json:"tags"`
}

type Returns_item struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Sn              string `gorm:"sn" json:"sn"`                               // 商品编号
	Name            string `gorm:"name" json:"name"`                           // 商品名称
	Quantity        int    `gorm:"quantity" json:"quantity"`                   // 数量
	ReturnsId       int    `gorm:"returns_id" json:"returns_id"`               // 退货单-关联returns
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Specification_value struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Name            string `gorm:"name" json:"name"`                           // 名称
	Image           string `gorm:"image" json:"image"`                         // 图片
	SpecificationId int    `gorm:"specification_id" json:"specification_id"`   // 规格
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Attribute_option struct {
	Attribute int    `gorm:"attribute" json:"attribute"`
	Options   string `gorm:"options" json:"options"` // 可选项
}

type Cart struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	CartKey         string `gorm:"cart_key" json:"cart_key"`                   // 密钥
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Coupon struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Name            string  `gorm:"name" json:"name"`                           // 名称
	Condition       string  `gorm:"condition" json:"condition"`                 // 条件
	Prefix          string  `gorm:"prefix" json:"prefix"`                       // 前缀
	Introduction    string  `gorm:"introduction" json:"introduction"`           // 介绍
	BeginDate       string  `gorm:"begin_date" json:"begin_date"`               // 使用起始日期
	EndDate         string  `gorm:"end_date" json:"end_date"`                   // 使用结束日期
	IsEnabled       int     `gorm:"is_enabled" json:"is_enabled"`               // 是否启用
	IsExchange      int     `gorm:"is_exchange" json:"is_exchange"`             // 是否允许积分兑换
	MaximumPrice    float64 `gorm:"maximum_price" json:"maximum_price"`         // 最大商品价格
	MaximumQuantity int     `gorm:"maximum_quantity" json:"maximum_quantity"`   // 最大商品数量
	MinimumPrice    float64 `gorm:"minimum_price" json:"minimum_price"`         // 最小商品价格
	MinimumQuantity int     `gorm:"minimum_quantity" json:"minimum_quantity"`   // 最小商品数量
	Point           int     `gorm:"point" json:"point"`                         // 积分兑换数
	PriceExpression string  `gorm:"price_expression" json:"price_expression"`   // 价格运算表达式
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Receiver struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Consignee       string `gorm:"consignee" json:"consignee"`                 // 收货人
	AreaName        string `gorm:"area_name" json:"area_name"`                 // 地区名称
	Address         string `gorm:"address" json:"address"`                     // 地址
	ZipCode         string `gorm:"zip_code" json:"zip_code"`                   // 邮编
	Phone           string `gorm:"phone" json:"phone"`                         // 电话
	IsDefault       int    `gorm:"is_default" json:"is_default"`               // 是否默认
	AreaId          int    `gorm:"area_id" json:"area_id"`                     // 地区
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Shipping_item struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Sn              string `gorm:"sn" json:"sn"`                               // 商品编号
	Name            string `gorm:"name" json:"name"`                           // 商品名称
	Quantity        int    `gorm:"quantity" json:"quantity"`                   // 数量
	ShippingId      int    `gorm:"shipping_id" json:"shipping_id"`             // 发货单
	IsDelivery      int    `gorm:"is_delivery" json:"is_delivery"`             // 是否需要物流
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Member_rank struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Name            string  `gorm:"name" json:"name"`                           // 名称
	Amount          float64 `gorm:"amount" json:"amount"`                       // 消费金额
	IsDefault       int     `gorm:"is_default" json:"is_default"`               // 是否默认
	IsSpecial       int     `gorm:"is_special" json:"is_special"`               // 是否特殊
	Scale           float64 `gorm:"scale" json:"scale"`                         // 优惠比例
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_member_price struct {
	Product        int     `gorm:"product" json:"product"`
	MemberPrice    float64 `gorm:"member_price" json:"member_price"`
	MemberPriceKey int     `gorm:"member_price_key" json:"member_price_key"`
}

type Product_specification_value struct {
	Products            int `gorm:"products" json:"products"`
	SpecificationValues int `gorm:"specification_values" json:"specification_values"`
}

type Promotion struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Name            string  `gorm:"name" json:"name"`                           // 名称
	Title           string  `gorm:"title" json:"title"`                         // 标题
	BeginDate       string  `gorm:"begin_date" json:"begin_date"`               // 起始日期
	EndDate         string  `gorm:"end_date" json:"end_date"`                   // 结束日期
	MinimumQuantity int     `gorm:"minimum_quantity" json:"minimum_quantity"`   // 最小商品数量
	MaximumQuantity int     `gorm:"maximum_quantity" json:"maximum_quantity"`   // 最大商品数量
	MinimumPrice    float64 `gorm:"minimum_price" json:"minimum_price"`         // 最小商品价格
	MaximumPrice    float64 `gorm:"maximum_price" json:"maximum_price"`         // 最大商品价格
	PriceExpression string  `gorm:"price_expression" json:"price_expression"`   // 价格运算表达式
	PointExpression string  `gorm:"point_expression" json:"point_expression"`   // 积分运算表达式
	IsFreeShipping  int     `gorm:"is_free_shipping" json:"is_free_shipping"`   // 是否免运费
	IsCouponAllowed int     `gorm:"is_coupon_allowed" json:"is_coupon_allowed"` // 是否允许使用优惠券
	Introduction    string  `gorm:"introduction" json:"introduction"`           // 介绍
	Image           string  `gorm:"image" json:"image"`                         // 图片
	Orders          int     `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Promotion_coupon struct {
	Promotions int `gorm:"promotions" json:"promotions"`
	Coupons    int `gorm:"coupons" json:"coupons"`
}

type Refunds struct {
	Id              int     `gorm:"id" json:"id"`                               // id
	Sn              string  `gorm:"sn" json:"sn"`                               // 编号
	IsReview        int     `gorm:"is_review" json:"is_review"`                 // 是否审核通过
	Method          int     `gorm:"method" json:"method"`                       // 方式
	PaymentMethod   string  `gorm:"payment_method" json:"payment_method"`       // 支付方式
	Bank            string  `gorm:"bank" json:"bank"`                           // 退款银行
	Account         string  `gorm:"account" json:"account"`                     // 退款账号
	Amount          float64 `gorm:"amount" json:"amount"`                       // 退款金额
	Payee           string  `gorm:"payee" json:"payee"`                         // 收款人
	Operator        string  `gorm:"operator" json:"operator"`                   // 操作员
	Memo            string  `gorm:"memo" json:"memo"`                           // 备注
	OrderId         int     `gorm:"order_id" json:"order_id"`                   // 订单
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Sms struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Mobile          string `gorm:"mobile" json:"mobile"`                       // 手机
	Code            string `gorm:"code" json:"code"`                           // 短信值
	Type            int    `gorm:"type" json:"type"`                           // 短信类型
	ExpireDate      string `gorm:"expire_date" json:"expire_date"`             // 过期日期
	UsedDate        string `gorm:"used_date" json:"used_date"`                 // 使用日期
	IsUsed          int    `gorm:"is_used" json:"is_used"`                     // 是否已使用
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
}

type Ad_position struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Description     string `gorm:"description" json:"description"`             // 描述
	Height          int    `gorm:"height" json:"height"`                       // 高度
	Name            string `gorm:"name" json:"name"`                           // 名称
	Template        string `gorm:"template" json:"template"`                   // 模板
	Width           int    `gorm:"width" json:"width"`                         // 宽度
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Log struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Content         string `gorm:"content" json:"content"`                     // 内容
	Operation       string `gorm:"operation" json:"operation"`                 // 操作
	Operator        string `gorm:"operator" json:"operator"`                   // 操作员
	Parameter       string `gorm:"parameter" json:"parameter"`                 // 请求参数
	Ip              string `gorm:"ip" json:"ip"`                               // IP
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Payment_log struct {
	Id                int     `gorm:"id" json:"id"`                                   // 主键_id
	Amount            float64 `gorm:"amount" json:"amount"`                           // 支付金额
	Fee               float64 `gorm:"fee" json:"fee"`                                 // 支付手续费
	PaymentPluginId   string  `gorm:"payment_plugin_id" json:"payment_plugin_id"`     // 支付插件ID
	PaymentPluginName string  `gorm:"payment_plugin_name" json:"payment_plugin_name"` // 支付插件名称
	Sn                string  `gorm:"sn" json:"sn"`                                   // 编号
	Status            int     `gorm:"status" json:"status"`                           // 状态
	Type              int     `gorm:"type" json:"type"`                               // 类型
	Memo              string  `gorm:"memo" json:"memo"`                               // 备注
	MemberId          int     `gorm:"member_id" json:"member_id"`                     // 会员
	OrderId           int     `gorm:"order_id" json:"order_id"`                       // 订单
	CreateBy          string  `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string  `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string  `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string  `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int     `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Product_specification struct {
	Products       int `gorm:"products" json:"products"`
	Specifications int `gorm:"specifications" json:"specifications"`
}

type Sku struct {
	Id                  int     `gorm:"id" json:"id"`                                     // 主键_id
	Sn                  string  `gorm:"sn" json:"sn"`                                     // 编号
	Stock               int     `gorm:"stock" json:"stock"`                               // 库存
	AllocatedStock      int     `gorm:"allocated_stock" json:"allocated_stock"`           // 已分配库存
	Cost                float64 `gorm:"cost" json:"cost"`                                 // 成本价
	MarketPrice         float64 `gorm:"market_price" json:"market_price"`                 // 市场价
	Price               float64 `gorm:"price" json:"price"`                               // 售价
	ExchangePoint       int     `gorm:"exchange_point" json:"exchange_point"`             // 兑换积分
	RewardPoint         int     `gorm:"reward_point" json:"reward_point"`                 // 赠送积分
	SpecificationValues string  `gorm:"specification_values" json:"specification_values"` // 规格值
	ProductId           int     `gorm:"product_id" json:"product_id"`                     // 商品
	IsDefault           int     `gorm:"is_default" json:"is_default"`                     // 是否默认
	IsEnable            int     `gorm:"is_enable" json:"is_enable"`                       // 是否启用
	CreateBy            string  `gorm:"create_by" json:"create_by"`                       // 创建人
	CreationDate        string  `gorm:"creation_date" json:"creation_date"`               // 创建日期
	LastUpdatedBy       string  `gorm:"last_updated_by" json:"last_updated_by"`           // 最后修改人
	LastUpdatedDate     string  `gorm:"last_updated_date" json:"last_updated_date"`       // 最后修改日期
	DeleteFlag          int     `gorm:"delete_flag" json:"delete_flag"`                   // 删除标记
}

type Admin struct {
	Id                int    `gorm:"id" json:"id"`                                   // 主键_id
	Name              string `gorm:"name" json:"name"`                               // 姓名
	Username          string `gorm:"username" json:"username"`                       // 用户名
	Password          string `gorm:"password" json:"password"`                       // 密码
	Hasher            string `gorm:"hasher" json:"hasher"`                           // 加密类型
	Salt              string `gorm:"salt" json:"salt"`                               // 加密盐
	Department        string `gorm:"department" json:"department"`                   // 部门
	Email             string `gorm:"email" json:"email"`                             // E-mail
	IsEnabled         int    `gorm:"is_enabled" json:"is_enabled"`                   // 是否启用
	IsLocked          int    `gorm:"is_locked" json:"is_locked"`                     // 是否锁定
	LockedDate        string `gorm:"locked_date" json:"locked_date"`                 // 锁定日期
	LoginDate         string `gorm:"login_date" json:"login_date"`                   // 最后登录日期
	LoginIp           string `gorm:"login_ip" json:"login_ip"`                       // 最后登录IP
	LoginFailureCount int    `gorm:"login_failure_count" json:"login_failure_count"` // 连续登录失败次数
	CreateBy          string `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int    `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Admin_role struct {
	Admins int `gorm:"admins" json:"admins"`
	Roles  int `gorm:"roles" json:"roles"` // 角色
}

type Article_tag struct {
	Articles int `gorm:"articles" json:"articles"`
	Tags     int `gorm:"tags" json:"tags"`
}

type Parameter struct {
	Id               int    `gorm:"id" json:"id"`                                 // 主键_id
	Name             string `gorm:"name" json:"name"`                             // 名称
	ParameterGroupId int    `gorm:"parameter_group_id" json:"parameter_group_id"` // 参数组
	Orders           int    `gorm:"orders" json:"orders"`                         // 排序
	CreateBy         string `gorm:"create_by" json:"create_by"`                   // 创建人
	CreationDate     string `gorm:"creation_date" json:"creation_date"`           // 创建日期
	LastUpdatedBy    string `gorm:"last_updated_by" json:"last_updated_by"`       // 最后修改人
	LastUpdatedDate  string `gorm:"last_updated_date" json:"last_updated_date"`   // 最后修改日期
	DeleteFlag       int    `gorm:"delete_flag" json:"delete_flag"`               // 删除标记
}

type Coupon_code struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	Code            string `gorm:"code" json:"code"`                           // 号码
	CouponId        int    `gorm:"coupon_id" json:"coupon_id"`                 // 优惠券
	IsUsed          int    `gorm:"is_used" json:"is_used"`                     // 是否已使用
	UsedDate        string `gorm:"used_date" json:"used_date"`                 // 使用日期
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_category struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	SeoTitle        string `gorm:"seo_title" json:"seo_title"`                 // 页面标题
	SeoKeywords     string `gorm:"seo_keywords" json:"seo_keywords"`           // 页面关键词
	SeoDescription  string `gorm:"seo_description" json:"seo_description"`     // 页面描述
	TreePath        string `gorm:"tree_path" json:"tree_path"`                 // 树路径
	Grade           int    `gorm:"grade" json:"grade"`                         // 层级
	Image           string `gorm:"image" json:"image"`                         // 图片
	ParentId        int    `gorm:"parent_id" json:"parent_id"`                 // 上级分类
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	IsMarketable    int    `gorm:"is_marketable" json:"is_marketable"`         // 是否上架
	IsTop           int    `gorm:"is_top" json:"is_top"`                       // 是否置顶
	IsShow          int    `gorm:"is_show" json:"is_show"`                     // 是否显示
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Promotion_brand struct {
	Promotions int `gorm:"promotions" json:"promotions"`
	Brands     int `gorm:"brands" json:"brands"`
}

type Returns struct {
	Id              int     `gorm:"id" json:"id"`                               // id
	Sn              string  `gorm:"sn" json:"sn"`                               // 编号
	ShippingMethod  string  `gorm:"shipping_method" json:"shipping_method"`     // 配送方式
	DeliveryCorp    string  `gorm:"delivery_corp" json:"delivery_corp"`         // 物流公司
	TrackingNo      string  `gorm:"tracking_no" json:"tracking_no"`             // 运单号
	Freight         float64 `gorm:"freight" json:"freight"`                     // 物流费用
	Shipper         string  `gorm:"shipper" json:"shipper"`                     // 发货人
	Area            string  `gorm:"area" json:"area"`                           // 地区
	Address         string  `gorm:"address" json:"address"`                     // 地址
	ZipCode         string  `gorm:"zip_code" json:"zip_code"`                   // 邮编
	Phone           string  `gorm:"phone" json:"phone"`                         // 电话
	Operator        string  `gorm:"operator" json:"operator"`                   // 操作员
	Memo            string  `gorm:"memo" json:"memo"`                           // 备注
	OrderId         int     `gorm:"order_id" json:"order_id"`                   // 订单
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Shipping struct {
	Id               int     `gorm:"id" json:"id"`                                 // id
	Sn               string  `gorm:"sn" json:"sn"`                                 // 编号
	ShippingMethod   string  `gorm:"shipping_method" json:"shipping_method"`       // 配送方式
	DeliveryCorp     string  `gorm:"delivery_corp" json:"delivery_corp"`           // 物流公司
	DeliveryCorpUrl  string  `gorm:"delivery_corp_url" json:"delivery_corp_url"`   // 物流公司网址
	DeliveryCorpCode string  `gorm:"delivery_corp_code" json:"delivery_corp_code"` // 物流公司代码
	TrackingNo       string  `gorm:"tracking_no" json:"tracking_no"`               // 运单号
	Freight          float64 `gorm:"freight" json:"freight"`                       // 物流费用
	Consignee        string  `gorm:"consignee" json:"consignee"`                   // 收货人
	Area             string  `gorm:"area" json:"area"`                             // 地区
	Address          string  `gorm:"address" json:"address"`                       // 地址
	ZipCode          string  `gorm:"zip_code" json:"zip_code"`                     // 邮编
	Phone            string  `gorm:"phone" json:"phone"`                           // 电话
	Operator         string  `gorm:"operator" json:"operator"`                     // 操作员
	Memo             string  `gorm:"memo" json:"memo"`                             // 备注
	OrderId          int     `gorm:"order_id" json:"order_id"`                     // 订单
	CreateBy         string  `gorm:"create_by" json:"create_by"`                   // 创建人
	CreationDate     string  `gorm:"creation_date" json:"creation_date"`           // 创建日期
	LastUpdatedBy    string  `gorm:"last_updated_by" json:"last_updated_by"`       // 最后修改人
	LastUpdatedDate  string  `gorm:"last_updated_date" json:"last_updated_date"`   // 最后修改日期
	DeleteFlag       int     `gorm:"delete_flag" json:"delete_flag"`               // 删除标记
}

type Area struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	FullName        string `gorm:"full_name" json:"full_name"`                 // 全称
	Name            string `gorm:"name" json:"name"`                           // 名称
	TreePath        string `gorm:"tree_path" json:"tree_path"`                 // 树路径
	ParentId        int    `gorm:"parent_id" json:"parent_id"`                 // 上级地区
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Article_category struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Name            string `gorm:"name" json:"name"`                           // 名称
	SeoTitle        string `gorm:"seo_title" json:"seo_title"`                 // 页面标题
	SeoDescription  string `gorm:"seo_description" json:"seo_description"`     // 页面描述
	SeoKeywords     string `gorm:"seo_keywords" json:"seo_keywords"`           // 页面关键词
	ParentId        int    `gorm:"parent_id" json:"parent_id"`                 // 上级分类
	Grade           int    `gorm:"grade" json:"grade"`                         // 层级
	TreePath        string `gorm:"tree_path" json:"tree_path"`                 // 树路径
	Orders          int    `gorm:"orders" json:"orders"`                       // 排序
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Attribute struct {
	Id                int    `gorm:"id" json:"id"`                                   // 主键_id
	Name              string `gorm:"name" json:"name"`                               // 名称
	PropertyIndex     int    `gorm:"property_index" json:"property_index"`           // 属性序号
	ProductCategoryId int    `gorm:"product_category_id" json:"product_category_id"` // 绑定分类
	Orders            int    `gorm:"orders" json:"orders"`                           // 排序
	CreateBy          string `gorm:"create_by" json:"create_by"`                     // 创建人
	CreationDate      string `gorm:"creation_date" json:"creation_date"`             // 创建日期
	LastUpdatedBy     string `gorm:"last_updated_by" json:"last_updated_by"`         // 最后修改人
	LastUpdatedDate   string `gorm:"last_updated_date" json:"last_updated_date"`     // 最后修改日期
	DeleteFlag        int    `gorm:"delete_flag" json:"delete_flag"`                 // 删除标记
}

type Sn struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Type            int    `gorm:"type" json:"type"`                           // 类型
	LastValue       int    `gorm:"last_value" json:"last_value"`               // 末值
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Message struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	Title           string `gorm:"title" json:"title"`                         // 标题
	Content         string `gorm:"content" json:"content"`                     // 内容
	SenderId        int    `gorm:"sender_id" json:"sender_id"`                 // 发件人
	ReceiverId      int    `gorm:"receiver_id" json:"receiver_id"`             // 收件人
	ForMessage      int    `gorm:"for_message" json:"for_message"`             // 原消息
	ReceiverRead    int    `gorm:"receiver_read" json:"receiver_read"`         // 收件人已读
	ReceiverDelete  int    `gorm:"receiver_delete" json:"receiver_delete"`     // 收件人删除
	SenderRead      int    `gorm:"sender_read" json:"sender_read"`             // 发件人已读
	SenderDelete    int    `gorm:"sender_delete" json:"sender_delete"`         // 发件人删除
	IsDraft         int    `gorm:"is_draft" json:"is_draft"`                   // 是否为草稿
	Ip              string `gorm:"ip" json:"ip"`                               // ip
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_parameter_value struct {
	Product           int    `gorm:"product" json:"product"`
	ParameterValue    string `gorm:"parameter_value" json:"parameter_value"`
	ParameterValueKey int    `gorm:"parameter_value_key" json:"parameter_value_key"`
}

type Review struct {
	Id              int    `gorm:"id" json:"id"`                               // id
	Score           int    `gorm:"score" json:"score"`                         // 评分
	Content         string `gorm:"content" json:"content"`                     // 内容
	IsShow          int    `gorm:"is_show" json:"is_show"`                     // 是否显示
	Ip              string `gorm:"ip" json:"ip"`                               // IP
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	ProductId       int    `gorm:"product_id" json:"product_id"`               // 商品
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Consultation struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键_id
	ForConsultation int    `gorm:"for_consultation" json:"for_consultation"`   // 咨询
	ProductId       int    `gorm:"product_id" json:"product_id"`               // 商品
	Content         string `gorm:"content" json:"content"`                     // 内容
	MemberId        int    `gorm:"member_id" json:"member_id"`                 // 会员
	IsShow          int    `gorm:"is_show" json:"is_show"`                     // 是否显示
	Ip              string `gorm:"ip" json:"ip"`                               // IP
	CreateBy        string `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int    `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Deposit_log struct {
	Id              int     `gorm:"id" json:"id"`                               // 主键_id
	Type            int     `gorm:"type" json:"type"`                           // 类型
	MemberId        int     `gorm:"member_id" json:"member_id"`                 // 会员
	Operator        string  `gorm:"operator" json:"operator"`                   // 操作员
	Credit          float64 `gorm:"credit" json:"credit"`                       // 收入金额
	Detinyint       float64 `gorm:"detinyint" json:"detinyint"`                 // 支出金额
	Balance         float64 `gorm:"balance" json:"balance"`                     // 当前余额
	Memo            string  `gorm:"memo" json:"memo"`                           // 备注
	CreateBy        string  `gorm:"create_by" json:"create_by"`                 // 创建人
	CreationDate    string  `gorm:"creation_date" json:"creation_date"`         // 创建日期
	LastUpdatedBy   string  `gorm:"last_updated_by" json:"last_updated_by"`     // 最后修改人
	LastUpdatedDate string  `gorm:"last_updated_date" json:"last_updated_date"` // 最后修改日期
	DeleteFlag      int     `gorm:"delete_flag" json:"delete_flag"`             // 删除标记
}

type Product_category_brand struct {
	ProductCategories int `gorm:"product_categories" json:"product_categories"`
	Brands            int `gorm:"brands" json:"brands"`
}
