package modifyEnterprise

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/exceptionHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func RegisterEnterprise(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/enterprise/register", func(context *gin.Context) {
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		_, err3 := fybDatabase.AddEnterprise(db, mp1)
		result = multierror.Append(result, err1, err2, err3)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "添加成功",
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
