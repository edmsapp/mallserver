package web

import (
	"mallServer/model"
)

type GetIndexBannerRequest struct {
	TotalCount uint64      `json:"totalCount"`
	Lists      []*model.Ad `json:"lists"`
}

type GetIndexBannerResponse struct {
	TotalCount uint64      `json:"totalCount"`
	Lists      []*model.Ad `json:"lists"`
}

type GetIndexGategoryResponse struct {
	TotalCount uint64                    `json:"totalCount"`
	Lists      []*model.Product_category `json:"lists"`
}
