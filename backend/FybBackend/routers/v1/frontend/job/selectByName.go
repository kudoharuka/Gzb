package job

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SelectByCode(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/job/detail/:name", func(context *gin.Context) {
		var result *multierror.Error
		name := context.Param("name")

		err, job, _ := fybDatabase.SearchSingleJobByName(db, name)
		result = multierror.Append(result, err)
		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "用户查看岗位成功",
				"data":    job,
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
