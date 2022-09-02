package model

func (t *Ad) TableName() string {
	return "ad"
}

// ListUser List all area
func GetIndexBanner() ([]*Ad, uint64, error) {

	banner := make([]*Ad, 0)
	var count uint64

	// where := fmt.Sprintf("parent_id =  %d", id)
	if err := DB.Self.Model(&Ad{}).Count(&count).Error; err != nil {
		return banner, count, err
	}

	if err := DB.Self.Find(&banner).Error; err != nil {
		return banner, count, err
	}
	return banner, count, nil
}
