package system

import "github.com/gin-gonic/gin"

// RegisterRouter ...
// Registers an instance of the Gin Framework
func RegisterRouter() *gin.Engine {
	return gin.Default() // default middleware: logger and recovery (crash-free)
}
