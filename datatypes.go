package main

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
