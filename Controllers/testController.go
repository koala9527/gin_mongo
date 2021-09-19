package Controllers

import (
	"gin_mongodb/Services"
	"gin_mongodb/global"

	"github.com/gin-gonic/gin"
)

func TestInsert(c *gin.Context) {
	Services.TestInsert()
}

func Log(c *gin.Context) {

	live_id := c.PostForm("live_id")
	insert_data := c.PostForm("data")
	msg_type := c.PostForm("msg_type")
	result := global.NewResult(c)
	data, err := Services.Log(live_id, insert_data, msg_type)
	if err != nil {
		result.Error(5201, err.Error(), data)
		return
	}
	result.Success(data)
}
