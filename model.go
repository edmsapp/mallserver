package model

type P_comment_reply struct {
	Id         int    `gorm:"id" json:"id"`                   // 回复ID
	CommentId  int    `gorm:"comment_id" json:"comment_id"`   // 评论ID
	UserId     int    `gorm:"user_id" json:"user_id"`         // 用户ID
	AtUserId   int    `gorm:"at_user_id" json:"at_user_id"`   // @用户ID
	Content    string `gorm:"content" json:"content"`         // 内容
	Ip         string `gorm:"ip" json:"ip"`                   // IP地址
	IpLoc      string `gorm:"ip_loc" json:"ip_loc"`           // IP城市地址
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_post_collection struct {
	Id         int `gorm:"id" json:"id"`                   // 收藏ID
	PostId     int `gorm:"post_id" json:"post_id"`         // POST ID
	UserId     int `gorm:"user_id" json:"user_id"`         // 用户ID
	CreatedOn  int `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_wallet_statement struct {
	Id              int    `gorm:"id" json:"id"`                             // 账单ID
	UserId          int    `gorm:"user_id" json:"user_id"`                   // 用户ID
	ChangeAmount    int    `gorm:"change_amount" json:"change_amount"`       // 变动金额
	BalanceSnapshot int    `gorm:"balance_snapshot" json:"balance_snapshot"` // 资金快照
	Reason          string `gorm:"reason" json:"reason"`                     // 变动原因
	PostId          int    `gorm:"post_id" json:"post_id"`                   // 关联动态
	CreatedOn       int    `gorm:"created_on" json:"created_on"`             // 创建时间
	ModifiedOn      int    `gorm:"modified_on" json:"modified_on"`           // 修改时间
	DeletedOn       int    `gorm:"deleted_on" json:"deleted_on"`             // 删除时间
	IsDel           int    `gorm:"is_del" json:"is_del"`                     // 是否删除 0 为未删除、1 为已删除
}

type P_post_content struct {
	Id         int    `gorm:"id" json:"id"`                   // 内容ID
	PostId     int    `gorm:"post_id" json:"post_id"`         // POST ID
	UserId     int    `gorm:"user_id" json:"user_id"`         // 用户ID
	Content    string `gorm:"content" json:"content"`         // 内容
	Type       int    `gorm:"type" json:"type"`               // 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源
	Sort       int    `gorm:"sort" json:"sort"`               // 排序，越小越靠前
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_post_star struct {
	Id         int `gorm:"id" json:"id"`                   // 收藏ID
	PostId     int `gorm:"post_id" json:"post_id"`         // POST ID
	UserId     int `gorm:"user_id" json:"user_id"`         // 用户ID
	CreatedOn  int `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_tag struct {
	Id         int    `gorm:"id" json:"id"`                   // 标签ID
	UserId     int    `gorm:"user_id" json:"user_id"`         // 创建者ID
	Tag        string `gorm:"tag" json:"tag"`                 // 标签名
	QuoteNum   int    `gorm:"quote_num" json:"quote_num"`     // 引用数
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_captcha struct {
	Id         int    `gorm:"id" json:"id"`                   // 验证码ID
	Phone      string `gorm:"phone" json:"phone"`             // 手机号
	Captcha    string `gorm:"captcha" json:"captcha"`         // 验证码
	UseTimes   int    `gorm:"use_times" json:"use_times"`     // 使用次数
	ExpiredOn  int    `gorm:"expired_on" json:"expired_on"`   // 过期时间
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_comment struct {
	Id         int    `gorm:"id" json:"id"`                   // 评论ID
	PostId     int    `gorm:"post_id" json:"post_id"`         // POST ID
	UserId     int    `gorm:"user_id" json:"user_id"`         // 用户ID
	Ip         string `gorm:"ip" json:"ip"`                   // IP地址
	IpLoc      string `gorm:"ip_loc" json:"ip_loc"`           // IP城市地址
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_post struct {
	Id              int    `gorm:"id" json:"id"`                               // 主题ID
	UserId          int    `gorm:"user_id" json:"user_id"`                     // 用户ID
	CommentCount    int    `gorm:"comment_count" json:"comment_count"`         // 评论数
	CollectionCount int    `gorm:"collection_count" json:"collection_count"`   // 收藏数
	UpvoteCount     int    `gorm:"upvote_count" json:"upvote_count"`           // 点赞数
	Visibility      int    `gorm:"visibility" json:"visibility"`               // 可见性 0公开 1私密 2好友可见
	IsTop           int    `gorm:"is_top" json:"is_top"`                       // 是否置顶
	IsEssence       int    `gorm:"is_essence" json:"is_essence"`               // 是否精华
	IsLock          int    `gorm:"is_lock" json:"is_lock"`                     // 是否锁定
	LatestRepliedOn int    `gorm:"latest_replied_on" json:"latest_replied_on"` // 最新回复时间
	Tags            string `gorm:"tags" json:"tags"`                           // 标签
	AttachmentPrice int    `gorm:"attachment_price" json:"attachment_price"`   // 附件价格(分)
	Ip              string `gorm:"ip" json:"ip"`                               // IP地址
	IpLoc           string `gorm:"ip_loc" json:"ip_loc"`                       // IP城市地址
	CreatedOn       int    `gorm:"created_on" json:"created_on"`               // 创建时间
	ModifiedOn      int    `gorm:"modified_on" json:"modified_on"`             // 修改时间
	DeletedOn       int    `gorm:"deleted_on" json:"deleted_on"`               // 删除时间
	IsDel           int    `gorm:"is_del" json:"is_del"`                       // 是否删除 0 为未删除、1 为已删除
}

type P_post_attachment_bill struct {
	Id         int `gorm:"id" json:"id"`                   // 购买记录ID
	PostId     int `gorm:"post_id" json:"post_id"`         // POST ID
	UserId     int `gorm:"user_id" json:"user_id"`         // 用户ID
	PaidAmount int `gorm:"paid_amount" json:"paid_amount"` // 支付金额
	CreatedOn  int `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_user struct {
	Id         int    `gorm:"id" json:"id"`                   // 用户ID
	Nickname   string `gorm:"nickname" json:"nickname"`       // 昵称
	Username   string `gorm:"username" json:"username"`       // 用户名
	Phone      string `gorm:"phone" json:"phone"`             // 手机号
	Password   string `gorm:"password" json:"password"`       // MD5密码
	Salt       string `gorm:"salt" json:"salt"`               // 盐值
	Status     int    `gorm:"status" json:"status"`           // 状态，1正常，2停用
	Avatar     string `gorm:"avatar" json:"avatar"`           // 用户头像
	Balance    int    `gorm:"balance" json:"balance"`         // 用户余额（分）
	IsAdmin    int    `gorm:"is_admin" json:"is_admin"`       // 是否管理员
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_wallet_recharge struct {
	Id          int    `gorm:"id" json:"id"`                     // 充值ID
	UserId      int    `gorm:"user_id" json:"user_id"`           // 用户ID
	Amount      int    `gorm:"amount" json:"amount"`             // 充值金额
	TradeNo     string `gorm:"trade_no" json:"trade_no"`         // 支付宝订单号
	TradeStatus string `gorm:"trade_status" json:"trade_status"` // 交易状态
	CreatedOn   int    `gorm:"created_on" json:"created_on"`     // 创建时间
	ModifiedOn  int    `gorm:"modified_on" json:"modified_on"`   // 修改时间
	DeletedOn   int    `gorm:"deleted_on" json:"deleted_on"`     // 删除时间
	IsDel       int    `gorm:"is_del" json:"is_del"`             // 是否删除 0 为未删除、1 为已删除
}

type P_attachment struct {
	Id         int    `gorm:"id" json:"id"`
	UserId     int    `gorm:"user_id" json:"user_id"`
	FileSize   int    `gorm:"file_size" json:"file_size"`
	ImgWidth   int    `gorm:"img_width" json:"img_width"`
	ImgHeight  int    `gorm:"img_height" json:"img_height"`
	Type       int    `gorm:"type" json:"type"` // 1图片，2视频，3其他附件
	Content    string `gorm:"content" json:"content"`
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_comment_content struct {
	Id         int    `gorm:"id" json:"id"`                   // 内容ID
	CommentId  int    `gorm:"comment_id" json:"comment_id"`   // 评论ID
	UserId     int    `gorm:"user_id" json:"user_id"`         // 用户ID
	Content    string `gorm:"content" json:"content"`         // 内容
	Type       int    `gorm:"type" json:"type"`               // 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址
	Sort       int    `gorm:"sort" json:"sort"`               // 排序，越小越靠前
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type P_message struct {
	Id             int    `gorm:"id" json:"id"`                             // 消息通知ID
	SenderUserId   int    `gorm:"sender_user_id" json:"sender_user_id"`     // 发送方用户ID
	ReceiverUserId int    `gorm:"receiver_user_id" json:"receiver_user_id"` // 接收方用户ID
	Type           int    `gorm:"type" json:"type"`                         // 通知类型，1动态，2评论，3回复，4私信，99系统通知
	Brief          string `gorm:"brief" json:"brief"`                       // 摘要说明
	Content        string `gorm:"content" json:"content"`                   // 详细内容
	PostId         int    `gorm:"post_id" json:"post_id"`                   // 动态ID
	CommentId      int    `gorm:"comment_id" json:"comment_id"`             // 评论ID
	ReplyId        int    `gorm:"reply_id" json:"reply_id"`                 // 回复ID
	IsRead         int    `gorm:"is_read" json:"is_read"`                   // 是否已读
	CreatedOn      int    `gorm:"created_on" json:"created_on"`             // 创建时间
	ModifiedOn     int    `gorm:"modified_on" json:"modified_on"`           // 修改时间
	DeletedOn      int    `gorm:"deleted_on" json:"deleted_on"`             // 删除时间
	IsDel          int    `gorm:"is_del" json:"is_del"`                     // 是否删除 0 为未删除、1 为已删除
}
