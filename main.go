package main

import "github.com/gin-gonic/gin"
import "harrysbaraini/monoedit/string"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		string.Service.Register(v1)
	}

	router.Run(":8888")
}
