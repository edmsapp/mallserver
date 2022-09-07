package webapis

import (
	"newsServ/model"
)

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type GetIndexListsRequest struct {
	Offset int `json:"offset"`
	Page   int `json:"page"`
}

type GetIndexListsRespones struct {
	TotalCount uint64        `json:"totalCount"`
	Lists      []*model.News `json:"lists"`
}

type GetNewsOneRespones struct {
	Lists *model.Details `json:"lists"`
}
