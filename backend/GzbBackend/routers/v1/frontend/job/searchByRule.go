package job

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SearchByRule(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/job/searchByRule", func(context *gin.Context) {
		var result *multierror.Error
		data, err1 := context.GetRawData()
		var mp map[string]interface{}
		err2 := json.Unmarshal(data, &mp)
		mp["job.region"] = mp["region"]
		mp["job.type"] = mp["type"]
		delete(mp, "region")
		delete(mp, "type")
		if mp["job.region"] == "岗位地点" {
			delete(mp, "job.region")
		}
		if mp["job.type"] == "工作性质" {
			delete(mp, "job.type")
		}
		if mp["wage"] == "薪资水平" {
			delete(mp, "wage")
		}

		if mp["category"] == "工作类型" {
			delete(mp, "category")
		}
		jobs, count, err3 := fybDatabase.SelectJobByCondition(db, mp)

		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "岗位筛选成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": jobs,
				},
			})
		} else {
			var code int
			if err3 != nil || err2 != nil {
				code = 500
			} else {
				code = 404
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})
		}

	})
}
