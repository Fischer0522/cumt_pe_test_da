package v1

import (
	"SportAnalyse/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInfos(c *gin.Context) {

	groupType := c.Query("groupType")
	fmt.Println(groupType)
	infos := models.SelectByGroup(groupType)
	c.JSON(200, gin.H{
		"code":200,
		"data":infos,
		"message": "OK",
	})

}
