package web

import (
	. "mallServer/handler"
	"mallServer/model"
	"mallServer/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary List the app in the database
// @Description List app product
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body app.ListRequest true "List app"
// @Success 200 {object} user.SwaggerProductListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /index [get]
func GetIndexLists(c *gin.Context) {
	var r GetIndexGategoryResponse
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := model.GetIndexBanner()
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexGategoryResponse{
		TotalCount: count,
		Lists:      infos,
	})
}

// @Summary List the app in the database
// @Description List app product
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body app.ListRequest true "List app"
// @Success 200 {object} user.SwaggerProductListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /banner [get]
func GetIndexBanner(c *gin.Context) {
	var r GetIndexBannerResponse
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := model.GetIndexBanner()
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexBannerRequest{
		TotalCount: count,
		Lists:      infos,
	})
}
