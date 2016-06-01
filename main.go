package main

import (
	"harrysbaraini/monoedit/modules/product"
	"harrysbaraini/monoedit/system"
)

func main() {
	// Register router
	router := system.RegisterRouter()

	// Create version 1 routing
	v1 := router.Group("/v1")
	{
		product.Service.Register(v1)
	}

	router.Run(":8888")
}
