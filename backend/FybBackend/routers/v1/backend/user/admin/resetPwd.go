package admin

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"time"
)

func RsetPwd(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/resetPwd", func(context *gin.Context) {
		var result error
		flag := false

		mp := make(map[string]interface{})
		b, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to read request body: %w", err))
		}
		if err := json.Unmarshal(b, &mp); err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to unmarshal JSON: %w", err))
		}

		vCode := mp["code"]
		account := mp["account"]
		password := mp["password"]

		c, err := redis.Dial("tcp", ":6379", redis.DialConnectTimeout(time.Second))
		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "无法连接到 Redis",
				"error":   err.Error(),
			})
			return
		}
		defer c.Close()

		vCodeRaw, err := redis.Bytes(c.Do("GET", account))
		if err != nil {
			if err != redis.ErrNil {
				result = multierror.Append(result, fmt.Errorf("failed to retrieve verification code from Redis: %w", err))
			}
		} else {
			if vCodeRaw != nil && string(vCodeRaw) != "" && vCode == string(vCodeRaw) {
				flag = true
			}
		}

		if result == nil && flag {
			// 发送注册成功邮件
			content := fmt.Sprintf(`
亲爱的%s，重置密码成功

工作帮_Official
	`, account)
			err := SendEmail(account.(string), "【工作帮】系统消息，请勿回复", content)
			if err != nil {
				// 处理邮件发送失败的错误
				context.JSON(500, gin.H{
					"code":    500,
					"message": "邮件发送失败",
					"error":   err.Error(),
				})
			}
			where := make(map[string]interface{})
			values := make(map[string]interface{})
			where["account"] = account
			values["password"] = password
			count, err1 := fybDatabase.UpdateSingleEnterpriseByCondition(db, where, values)
			if err1 != nil || count == 0 {
				// 处理邮件发送失败的错误
				context.JSON(500, gin.H{
					"code":    500,
					"message": "密码重置",
					"error":   err1.Error(),
				})
			}
			context.JSON(200, gin.H{
				"code":    200,
				"message": "验证成功",
			})

		} else {
			errMsg := ""
			if result != nil {
				errMsg = result.Error()
			}
			context.JSON(404, gin.H{
				"code":    404,
				"message": "验证码错误或已失效",
				"error":   errMsg,
			})
		}
	})
}
