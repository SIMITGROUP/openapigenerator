package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DummyFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "DummyFunction"})
}

func getUserDetails(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "age": "age",
        "firstName": "firstName",
        "id": "id",
        "lastName": "lastName",
        "password": "password",
        "username": "username",})
}
func updateUserDetails(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{})
}
func deleteUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{})
}
func getUsersList(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "list_of_users": "list_of_users",
        "number_of_users": "number_of_users",})
}
func createUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{})
}
