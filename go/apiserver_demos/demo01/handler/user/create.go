package user

import (
    . "learnlanguage/go/apiserver_demos/demo01/handler"
    "learnlanguage/go/apiserver_demos/demo01/model"
    "learnlanguage/go/apiserver_demos/demo01/pkg/errno"
    "learnlanguage/go/apiserver_demos/demo01/util"

    "github.com/gin-gonic/gin"
    "github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
    s:= make(map[string]interface{})
    s["X-Request-Id"] =  util.GetReqID(c)

    log.Info("User Create function called.", s)
    var r CreateRequest
    if err := c.Bind(&r); err != nil {
        SendResponse(c, errno.ErrBind, nil)
        return
    }

    u := model.UserModel{
        Username: r.Username,
        Password: r.Password,
    }

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
    // Insert the user to the database.
    if err := u.Create(); err != nil {
        SendResponse(c, errno.ErrDatabase, nil)
        return
    }

    rsp := CreateResponse{
        Username: r.Username,
    }

    // Show the user information.
    SendResponse(c, nil, rsp)
}
