package routers

import (
	"FybBackend/routers/v1/backend/comment/modifyComment"
	"FybBackend/routers/v1/backend/comment/selectComment"
	"FybBackend/routers/v1/backend/dashboard"
	"FybBackend/routers/v1/backend/enterprise/enterpriseLogin"
	modifyEnterprise "FybBackend/routers/v1/backend/enterprise/modifyJob"
	selectEnterprise "FybBackend/routers/v1/backend/enterprise/selectJob"
	"FybBackend/routers/v1/backend/feedback/modifyFeedback"
	"FybBackend/routers/v1/backend/feedback/selectFeedback"
	"FybBackend/routers/v1/backend/job/modifyJob"
	"FybBackend/routers/v1/backend/job/selectJob"
	"FybBackend/routers/v1/backend/news/modifyNews"
	"FybBackend/routers/v1/backend/news/selectNews"
	"FybBackend/routers/v1/backend/post/modifyPost"
	"FybBackend/routers/v1/backend/post/selectPost"
	"FybBackend/routers/v1/backend/user/admin"
	"FybBackend/routers/v1/backend/user/modifyUser"
	"FybBackend/routers/v1/backend/user/selectUser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBackend(r *gin.Engine, db *gorm.DB) {
	//user
	admin.Login(r, db)
	admin.GetCode(r)
	admin.RsetPwd(r, db)
	modifyUser.AddUser(r, db)
	modifyUser.UpdateUser(r, db)
	modifyUser.DeleteUser(r, db)
	selectUser.SelectUserByAccount(r, db)
	selectUser.SelectUsersByPage(r, db)

	//post
	modifyPost.DeletePost(r, db)
	modifyPost.AddPost(r, db)
	modifyPost.CheckPost(r, db)
	modifyPost.UpdatePost(r, db)
	selectPost.SelectPostByPage(r, db)
	selectPost.SelectPostByAccount(r, db)

	//job
	modifyJob.DeleteJob(r, db)
	modifyJob.AddJob(r, db)
	modifyJob.UpdateJob(r, db)
	selectJob.SelectJobByPage(r, db)
	selectJob.SelectJobByID(r, db)

	//enterprise
	enterpriseLogin.Login(r, db)
	modifyEnterprise.RegisterEnterprise(r, db)
	modifyEnterprise.DeleteEnterprise(r, db)
	modifyEnterprise.AddEnterprise(r, db)
	modifyEnterprise.UpdateEnterprise(r, db)
	selectEnterprise.SelectEnterpriseByPage(r, db)
	selectEnterprise.SelectEnterpriseByID(r, db)

	//news
	modifyNews.AddNews(r, db)
	modifyNews.DeleteNews(r, db)
	modifyNews.UpdateNews(r, db)
	selectNews.SelectNewsById(r, db)
	selectNews.SelectNewsByPage(r, db)

	//comment
	selectComment.SelectCommentByPage(r, db)
	selectComment.SelectCommentById(r, db)
	modifyComment.DeleteComment(r, db)
	modifyComment.UpdateComment(r, db)
	modifyComment.AddComment(r, db)

	//feedback
	selectFeedback.SelectFeedbackByPage(r, db)
	modifyFeedback.UpdateFeedback(r, db)

	//dashboard
	dashboard.GetPostData(r, db)
	dashboard.GetHomeData(r, db)

}
