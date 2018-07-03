package user

import (
    `learnlanguage/go/apiserver_demos/demo01/pkg/errno`
    `learnlanguage/go/apiserver_demos/demo01/service`

    `github.com/gin-gonic/gin`
    ."learnlanguage/go/apiserver_demos/demo01/handler"
)


// List list the users in the database.
func List(c *gin.Context) {
    var r ListRequest
    if err := c.Bind(&r); err != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
    if err != nil {
        SendResponse(c, err, nil)
        return
    }

    SendResponse(c, nil, ListResponse{
        TotalCount: count,
        UserList:   infos,
    })
}