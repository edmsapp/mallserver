package apps

import (
	. "newsServ/handler"
	"newsServ/pkg/errno"
	"newsServ/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
)

// @Summary List the app in the database
// @Description List app product
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body app.ListRequest true "List app"
// @Success 200 {object} user.SwaggerProductListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /productlists [get]
func ProductLists(c *gin.Context) {
	log.Info("ProductLists function called.")
	var r ProductListResponse
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ProductListRequest{
		TotalCount: count,
		UserList:   infos,
	})
}

// @Summary List the app in the database
// @Description List app product
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body app.ListRequest true "List app"
// @Success 200 {object} user.SwaggerProductListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /productlists [get]
func AreaLists(c *gin.Context) {
	var r AreaListsResponse
	parentID := c.Query("id")
	id, err := strconv.Atoi(parentID)
	if err != nil {
		id = 0
	}
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.AreaLists(id)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, AreaListsRequest{
		TotalCount: count,
		AreaLists:  infos,
	})
}
