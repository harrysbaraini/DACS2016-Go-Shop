package string

import (
	"harrysbaraini/monoedit/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// Service settings
var Service = utils.Service{
	Name:   "String",
	Prefix: "/string",
	Routes: []utils.Route{
		utils.Route{Method: "GET", Url: "/uppercase", Handler: uppercase},
		utils.Route{Method: "GET", Url: "/count", Handler: uppercase},
	},
}

func uppercase(ctx *gin.Context) {
	str := ctx.Query("s")
	if str == "" {
		ctx.JSON(422, gin.H{
			"result": gin.H{
				"message": "Empty value",
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"result": strings.ToUpper(str),
		})
	}
}

func count(ctx *gin.Context) {
	str := ctx.Query("s")
	ctx.JSON(200, gin.H{
		"result": len(str),
	})
}
