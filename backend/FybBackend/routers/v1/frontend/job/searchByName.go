package job

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SearchByName(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/job/searchByName/:Name", func(context *gin.Context) {
		var result *multierror.Error
		Name := context.Param("Name")

		err, jobs, count := fybDatabase.SearchJobByName(db, Name)
		result = multierror.Append(result, err)
		if count > 0 && result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "岗位搜索成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": jobs,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data": map[string]interface{}{
					"num":  count,
					"list": jobs,
				},
			})
		}
	})
}
