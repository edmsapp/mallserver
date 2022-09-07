package webapis

import (
	"fmt"
	. "newsServ/handler"
	"newsServ/model"
	"newsServ/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.ListRequest true "List users"
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /user [get]
func GetIndexListsHandler(c *gin.Context) {
	var r GetIndexListsRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := model.GetListsNewsModel()
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetIndexListsRespones{
		TotalCount: count,
		Lists:      infos,
	})
}

func GetNewsHandler(c *gin.Context) {
	newsId := c.Param("id")
	id, err := strconv.Atoi(newsId)
	if err != nil {
		id = 1
	}
	fmt.Println("GetNewsHandler", id)

	infos, err := model.GetNewsOneModel(id)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, GetNewsOneRespones{
		Lists: infos,
	})
}
