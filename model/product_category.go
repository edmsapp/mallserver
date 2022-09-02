package model

func (t *Product_category) TableName() string {
	return "product_category"
}

// GetAllProductCategory retrieves all ProductCategory matches certain condition. Returns empty list if
// no records exist
func GetAllProductCategory() {
}
