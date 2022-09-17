package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,
    "getUserDetails":getUserDetails,
    "updateUserDetails":updateUserDetails,
    "deleteUser":deleteUser,
    "getUsersList":getUsersList,
    "createUser":createUser,
}