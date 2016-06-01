package utils

import "github.com/gin-gonic/gin"

// Route ... Route configuration struct
type Route struct {
	Method  string
	Url     string
	Handler gin.HandlerFunc
}

// Service ... Service configuration struct
type Service struct {
	Name   string
	Prefix string
	Routes []Route
}

// Register ...
// Register the service
func (s Service) Register(engine *gin.RouterGroup) {
	for i := 0; i < len(s.Routes); i++ {
		engine.Handle(s.Routes[i].Method, s.Prefix+s.Routes[i].Url, s.Routes[i].Handler)
	}
}
