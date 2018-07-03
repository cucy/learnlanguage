package user

import (
    `strconv`

    `learnlanguage/go/apiserver_demos/demo01/model`
    `learnlanguage/go/apiserver_demos/demo01/pkg/errno`

    `github.com/gin-gonic/gin`
    . "learnlanguage/go/apiserver_demos/demo01/handler"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
    userId, _ := strconv.Atoi(c.Param("id"))
    if err := model.DeleteUser(uint64(userId)); err != nil {
        SendResponse(c, errno.ErrDatabase, nil)
        return
    }

    SendResponse(c, nil, nil)
}

