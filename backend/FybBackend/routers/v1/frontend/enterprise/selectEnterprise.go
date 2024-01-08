package enterprise

import (
	fybDatabase "FybBackend/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SelectEnterpriseByName(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/enterprise/detail/:name", func(context *gin.Context) {
		var result *multierror.Error
		name := context.Param("name")
		mp := make(map[string]interface{})
		conditions := make(map[string]interface{})
		mp["name"] = name
		enterprise, count, err := fybDatabase.SelectSingleEnterpriseByCondition(db, mp)
		conditions["enterpriseID"] = enterprise.ID
		fmt.Println(conditions)
		jobs, _, err1 := fybDatabase.SelectJobByCondition(db, conditions)
		result = multierror.Append(result, err, err1)
		if count > 0 && result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "用户查看企业成功",
				"data": map[string]interface{}{
					"desc":        enterprise.Desc,
					"recruitment": enterprise.Recruitment,
					"name":        enterprise.Name,
					"type":        enterprise.Type,
					"region":      enterprise.Region,
					"city":        enterprise.City,
					"logo":        enterprise.Logo,
					"foundDate":   enterprise.FoundDate,
					"founder":     enterprise.Founder,
					"list":        jobs,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
