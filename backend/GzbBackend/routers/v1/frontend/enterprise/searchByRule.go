package enterprise

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SearchByRule(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/enterprise/searchByRule", func(context *gin.Context) {
		var result *multierror.Error
		data, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			result = multierror.Append(result, err)
		}
		region := m["region"]
		belong := m["belong"]
		_type := m["type"]

		requestBody := map[string]interface{}{
			"region": region,
			"belong": belong,
			"type":   _type,
		}

		if region == "企业地点" {
			delete(requestBody, "region")
		}
		if belong == "所属行业" {
			delete(requestBody, "belong")
		}
		if _type == "企业类型" {
			delete(requestBody, "type")
		}
		var responseBody []fybDatabase.Enterprise
		var count int64
		err, responseBody, count = fybDatabase.SearchEnterpriseByRule(db, requestBody)
		if err != nil {
			result = multierror.Append(result, err)
		}
		if count >= 0 && result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "企业筛选成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": responseBody,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data": map[string]interface{}{
					"num":  count,
					"list": responseBody,
				},
			})
		}

	})
}
