package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,
    "addTask":addTask,
    "checkAccessToken":checkAccessToken,
    "login":login,
    "updateTask":updateTask,
    "getUsers":getUsers,
    "refreshAccessToken":refreshAccessToken,
    "registerUser":registerUser,
    "deleteTask":deleteTask,
    "getTask":getTask,
    "getTaskById":getTaskById,
    "getTaskByUsername":getTaskByUsername,
}