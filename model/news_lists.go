package model

import "fmt"

type Details struct {
	Id         int    `gorm:"id" json:"id"`
	NewsId     int    `gorm:"news_id" json:"news_id"`
	Figure     string `gorm:"figure" json:"figure"`
	PList      string `gorm:"p_list" json:"p_list"`
	Alt        string `gorm:"alt" json:"alt"`
	Src        string `gorm:"src" json:"src"`
	LocalSrc   string `gorm:"local_src" json:"local_src"`
	CreateTime string `gorm:"create_time" json:"create_time"`
	Content    string `gorm:"content" json:"content"`
}

func (t *Details) TableName() string {
	return "details"
}

type News struct {
	Id          int    `gorm:"id" json:"id"`           // ID
	Href        string `gorm:"href" json:"href"`       // 新闻源页面URL
	ImgUrl      string `gorm:"img_url" json:"img_url"` // 图片URL
	LoadImg     string `gorm:"load_img" json:"load_img"`
	Title       string `gorm:"title" json:"title"`             // 新闻标题
	Country     string `gorm:"country" json:"country"`         // 国家缩写
	Source      string `gorm:"source" json:"source"`           // 来源的网站名称
	Description string `gorm:"description" json:"description"` // 新闻摘要
	Status      int    `gorm:"status" json:"status"`           // 状态, 0=未处理获取内容, 1=处理内容, 9=处理失败
	CreateTime  string `gorm:"create_time" json:"create_time"` // 获取时间
	Menu        string `gorm:"menu" json:"menu"`               // 新分类
	HrefHash    string `gorm:"href_hash" json:"href_hash"`     // hash值
}

func (t *News) TableName() string {
	return "news"
}

func GetListsNewsModel() ([]*News, uint64, error) {

	var limit uint64 = 10

	lists := make([]*News, 0)
	var count uint64
	status := 1
	offset := 0

	where := fmt.Sprintf("status =  %d", status)
	if err := DB.Self.Model(&News{}).Where(where).Count(&count).Error; err != nil {
		return lists, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&lists).Error; err != nil {
		return lists, limit, err
	}
	return lists, count, nil
}

func GetNewsOneModel(id int) (*Details, error) {
	u := &Details{}
	d := DB.Self.Where("news_id = ?", id).First(&u)
	return u, d.Error
}

func GetNewesModel(id int) ([]*Details, error) {

	lists := make([]*Details, 0)
	fmt.Println("model", id)
	where := fmt.Sprintf("news_id =  %d", id)
	err := DB.Self.Model(&Details{}).Where(where).Find(&lists).Error
	if err != nil {
		return lists, err
	}

	return lists, nil
}
