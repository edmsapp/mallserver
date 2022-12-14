package apps

import (
	"newsServ/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ProductListResponse struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ProductListRequest struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}

type AreaListsResponse struct {
	Id int `json:"id"`
}

type AreaListsRequest struct {
	TotalCount uint64             `json:"totalCount"`
	AreaLists  []*model.AreaModel `json:"areaLists"`
}
