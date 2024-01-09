package enterprise

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SearchJobsByEnterprise(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/enterprise/job", func(context *gin.Context) {
		var result *multierror.Error
		data, err1 := context.GetRawData()
		var mp map[string]interface{}
		err2 := json.Unmarshal(data, &mp)
		if mp["region"] == "岗位地点" {
			delete(mp, "region")
		} else {
			mp["job.region"] = mp["region"]
			delete(mp, "region")
		}
		if mp["type"] == "工作性质" {
			delete(mp, "type")
		}
		if mp["wage"] == "薪资水平" {
			delete(mp, "wage")
		}

		if mp["category"] == "工作类型" {
			delete(mp, "category")
		}
		where := make(map[string]interface{})
		where["name"] = mp["name"]
		enterprise, _, _ := fybDatabase.SelectSingleEnterpriseByCondition(db, where)
		delete(mp, "name")
		mp["enterpriseID"] = enterprise.ID
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
