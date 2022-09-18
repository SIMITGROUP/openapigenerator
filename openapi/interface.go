package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,
    "refreshAccessToken":refreshAccessToken,
    "updateTask":updateTask,
    "checkAccessToken":checkAccessToken,
    "getTask":getTask,
    "getTaskByUsername":getTaskByUsername,
    "login":login,
    "registerUser":registerUser,
    "addTask":addTask,
    "deleteTask":deleteTask,
    "getTaskById":getTaskById,
    "getUsers":getUsers,
}