package user

import (
    `strconv`

    . `learnlanguage/go/apiserver_demos/demo01/handler`
    `learnlanguage/go/apiserver_demos/demo01/model`
    `learnlanguage/go/apiserver_demos/demo01/pkg/errno`
    `learnlanguage/go/apiserver_demos/demo01/util`

    `github.com/gin-gonic/gin`
    "github.com/lexkong/log"

)

// Update update a exist user account info.
func Update(c *gin.Context) {
    s:= make(map[string]interface{})
    s["X-Request-Id"] =  util.GetReqID(c)

    log.Info("User Create function called.", s)
    // log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
 // Get the user id from the url parameter.
    userId, _ := strconv.Atoi(c.Param("id"))

    // Binding the user data.
    var u model.UserModel
    if err := c.Bind(&u); err != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    // We update the record based on the user id.
    u.Id = uint64(userId)

    // Validate the data.
    if err := u.Validate(); err != nil {
        SendResponse(c, errno.ErrValidation, nil)
        return
    }

    // Encrypt the user password.
    if err := u.Encrypt(); err != nil {
        SendResponse(c, errno.ErrEncrypt, nil)
        return
    }

    // Save changed fields.
    if err := u.Update(); err != nil {
        SendResponse(c, errno.ErrDatabase, nil)
        return
    }

    SendResponse(c, nil, nil)
}
