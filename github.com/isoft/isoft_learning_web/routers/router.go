package routers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_learning_web/controllers"
	//"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin",
	//		"Access-Control-Allow-Headers", "Content-Type", "X-Requested-With"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "get,post:Index")

	beego.Router("/course/index", &controllers.CourseController{}, "get,post:Index")
	beego.Router("/course/queryCourse", &controllers.CourseController{}, "get,post:QueryCourse")
	beego.Router("/course/play", &controllers.CourseController{}, "get,post:Play")
	beego.Router("/course/homemanage", &controllers.CourseController{}, "get,post:HomeManage")
	beego.Router("/course/newcourse", &controllers.CourseController{}, "get,post:NewCourse")
	beego.Router("/course/courselist", &controllers.CourseController{}, "get,post:CourseList")
	beego.Router("/course/showCourseDetail", &controllers.CourseController{}, "get,post:ShowCourseDetail")
	beego.Router("/course/newcourse/add", &controllers.CourseController{}, "get,post:AddNewCourse")
	beego.Router("/course/newcourse/changeImage", &controllers.CourseController{}, "get,post:ChangeImage")
	beego.Router("/course/newcourse/uploadvedio", &controllers.CourseController{}, "get,post:UploadVedio")
	beego.Router("/course/newcourse/endUpdate", &controllers.CourseController{}, "get,post:EndUpdate")
	beego.Router("/course/queryCourseExist", &controllers.CourseController{}, "get,post:QueryCourseExist")
	beego.Router("/course/search", &controllers.CourseController{}, "get,post:SearchCourse")
	beego.Router("/course/coursetype/list", &controllers.CourseController{}, "get,post:CourseTypeList")
	beego.Router("/course/coursesubtype/list", &controllers.CourseController{}, "get,post:CourseSubTypeList")

	beego.Router("/common/toggle_favorite", &controllers.CommonController{}, "get,post:ToggleFavorite")
	beego.Router("/common/checkLoginUser", &controllers.CommonController{}, "get,post:CheckLoginUser")
	beego.Router("/common/login", &controllers.CommonController{}, "get,post:Login")
	beego.Router("/common/logout", &controllers.CommonController{}, "get,post:Logout")
	beego.Router("/common/query_configuration", &controllers.CommonController{}, "get,post:QueryConfiguration")

	beego.Router("/comment/topicTheme/filter", &controllers.CommentController{}, "get,post:FilterTopicTheme")
	beego.Router("/comment/topicReply/add", &controllers.CommentController{}, "get,post:AddTopicReply")
	beego.Router("/comment/topicReply/filter", &controllers.CommentController{}, "get,post:FilterTopicReply")

	beego.Router("/note/list", &controllers.NoteController{}, "get,post:ListNote")
	beego.Router("/note/queryNoteExist", &controllers.NoteController{}, "get,post:QueryNoteExist")
	beego.Router("/note/queryNoteById", &controllers.NoteController{}, "get,post:QueryNoteById")
	beego.Router("/note/queryNoteHtmlById", &controllers.NoteController{}, "get,post:QueryNoteHtmlById")
	beego.Router("/note/edit", &controllers.NoteController{}, "get,post:CreateOrUpdateNote")
	beego.Router("/note/view", &controllers.NoteController{}, "get,post:ViewNote")
	beego.Router("/note/collect_list", &controllers.NoteController{}, "get,post:CollectList")
}
