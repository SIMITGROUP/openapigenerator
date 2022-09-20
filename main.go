package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	// r := gin.Default()
	file := "spec.yaml"
	file = "jwt.yaml"
	_ = NewServer(file)

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	// errInit := authMiddleware.MiddlewareInit()

	// if errInit != nil {
	// 	log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	// }

	// r.POST("/login", authMiddleware.LoginHandler)

	// auth := r.Group("/auth")
	// // Refresh time can be longer than token timeout
	// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// auth.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	auth.GET("/hello", helloHandler)
	// }

	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
	// r.Run(port)
}
