package model

import (
	"fmt"
)

// `id` int(11) NOT NULL AUTO_INCREMENT,
// `parent_id` int(11) NOT NULL DEFAULT '0',
// `level` tinyint(4) NOT NULL DEFAULT '1',
// `code` bigint(20) NOT NULL,
// `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
type AreaModel struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	ParentId int    `gorm:"parent_id" json:"parentId"`
	Level    uint16 `json:"level"`
	Code     uint64 `json:"code"`
	Name     string `json:"name"`
}

func (c *AreaModel) TableName() string {
	return "area"
}

// Create creates a new user account.
func (u *AreaModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteArea(id uint64) error {
	user := AreaModel{}
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *AreaModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetArea(id int) (*AreaModel, error) {
	u := &AreaModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

// ListUser List all area
func AreaLists(id int) ([]*AreaModel, uint64, error) {

	area := make([]*AreaModel, 0)
	var count uint64
	where := fmt.Sprintf("parent_id =  %d", id)
	if err := DB.Self.Model(&AreaModel{}).Where(where).Count(&count).Error; err != nil {
		return area, count, err
	}

	if err := DB.Self.Where(where).Find(&area).Error; err != nil {
		return area, count, err
	}
	return area, count, nil
}
