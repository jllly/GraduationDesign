package main

import (
	"Homework_system/Middleware"
	"Homework_system/class"
	"Homework_system/homework"
	"Homework_system/title"
	"Homework_system/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerinit(r *gin.Engine) {
	r.GET("class_management_content", func(context *gin.Context) {
		context.HTML(http.StatusOK, "class_management_content.html", gin.H{
			"title": "Main website",
		})
	})
	//登录
	r.GET("/", func(context *gin.Context) {
		//清除浏览器缓存
		context.Header("Pragma", "no-cache")
		context.Header("Cache-Control", "no-cache")
		context.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Main website",
		})
	})
	//登录成功后跳转到学生首页
	r.GET("/Student", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "student_index.html", gin.H{
			"page": "1",
		})
	})
	//登录成功后跳转到老师首页
	r.GET("/Teacher", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "teacher.html", gin.H{
			"page": "1",
		})
	})
	r.GET("/Get_information", Middleware.UserAuth(), user.Obtain_name) //获取显示名字

	r.GET("/Verificationcode", user.Verificationcode) //验证码生成请求
	r.GET("/captcha/:captchaId", user.Captchaview)
	r.POST("/submitregistered", user.Submitregistered) //提交注册申请表单
	r.POST("/login", user.Login)                       //登录
	r.GET("/login", user.Login)
	r.POST("/forget", user.Submitphone) //忘记密码,提交手机号
	r.POST("forgetform", user.Forget)   //忘记密码请求
	r.POST("/loginOut", Middleware.UserAuth(), user.LoginOut) //退出
	//老师班级作业管理
	r.GET("/Correct_students_papers", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "Correct_students_papers.html", gin.H{
			"page": "2",
		})
	})
	//布置作业
	r.GET("/decoratework", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "decoratework.html", gin.H{
			"page": "2",
		})
	})
	r.POST("/Class_Add", Middleware.UserAuth(), class.Claa_Add)                    //创建班级
	r.GET("/Class_List", Middleware.UserAuth(), class.Claa_List)                   //获取班级列表
	r.GET("/Class_List_students", Middleware.UserAuth(), class.Claa_List_students) //获取班级列表

	r.POST("//Class_nameModify", Middleware.UserAuth(), class.ClassName_editor) //修改班级ID
	r.POST("/Delete_class", Middleware.UserAuth(), class.Class_delete)          //删除班级

	r.POST("/Class_join", Middleware.UserAuth(), class.Student_class)         //学生通过邀请码加入班级
	r.POST("/Teacher_student_join", Middleware.UserAuth(), class.Add_student) //老师输入ID或手机号加入班级

	r.GET("/T_homework_is_in_progress", Middleware.UserAuth(), class.Teacher_homework_progress_List) //获取班级列表

	//选择班级
	r.GET("/Selectproblem", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "Selectproblem.html", gin.H{
			"page": "3",
		})
	})
	//选择班级
	r.GET("/Selecttheclass", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "Selecttheclass.html", gin.H{
			"page": "3",
		})
	})
	//学生错题本
	r.GET("/WrongT", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "WrongT.html", gin.H{
			"page": "2",
		})
	})

	//学生收藏
	r.GET("/mycollection", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "mycollection.html", gin.H{
			"page": "3",
		})
	})
	//题目创建页面解答题
	r.GET("/Answerquestions", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "Answerquestions.html", gin.H{
			"title": "Main website",
		})
	})
	//学生班级
	r.GET("/myclass", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "myclass.html", gin.H{
			"page": "4",
		})
	})

	//学生做作业
	r.GET("/dowork", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "dowork.html", gin.H{})
	})

	//创建一个 填空题
	//CreateAFilliInQuestion
	r.GET("/CreateAFilliInQuestion", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "CreateAFilliInQuestion.html", gin.H{
			"title": "Main website",
		})
	})
	//解答题作答区域
	r.GET("/jSolutionarea", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "jSolutionarea.html", gin.H{
			"title": "Main website",
		})
	})

	//题目创建页面选择题
	r.GET("/Createproblem", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "Createproblem.html", gin.H{
			"title": "Main website",
		})
	})

	//添加题目
	r.POST("/Add_Question", Middleware.UserAuth(), title.Add_question)
	//显示题目列表
	r.GET("/Question_List", Middleware.UserAuth(), title.Question_List)

	//老师班级管理
	r.GET("/class_management", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "class_management_list.html", gin.H{
			"page": "2",
		})
	})



	//老师题库管理
	r.GET("/questionbank_management", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "questionbank_management.html", gin.H{
			"page": "3",
		})
	})


	//跳到老师管理学生界面
	r.GET("/Student_management", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "class_management_content.html", gin.H{
			"page": "1",
		})
	})

	//跳到老师管理学生界面
	r.GET("/Myclass_main", Middleware.UserAuth(), func(context *gin.Context) {
		context.HTML(http.StatusOK, "my_class_content.html", gin.H{
			"page": "1",
		})
	})

	r.GET("/Student_list", Middleware.UserAuth(), class.Student_list)      //获取班级内学生列表
	r.GET("/Search_student", Middleware.UserAuth(), class.Search_student)  //搜索获取班级内学生列表
	r.GET("/get_student_list", Middleware.UserAuth(), class.Download_list) //搜索获取班级内学生列表
	r.POST("/Batch_delete", Middleware.UserAuth(), class.Batch_delete)     //删除班级内学生

	r.GET("/Choose_class", Middleware.UserAuth(), homework.Choose_class)                //班级id获取班级信息
	r.GET("/Search_Class_list", Middleware.UserAuth(), homework.Search_Class_list)      //所有班级列表
	r.GET("/Search_Class", Middleware.UserAuth(), homework.Search_Class)                //搜索班级
	r.POST("/Complete_assignment", Middleware.UserAuth(), homework.Complete_assignment) //搜索班级
	r.GET("/fornow_Student_worklist", Middleware.UserAuth(), homework.Fornow_Student)   //学生往期作业名单
	//r.GET("/Before_Student_worklist", Middleware.UserAuth(), homework.Before_Student)   //学生往期作业名单
	//r.POST("/workbegin", user.Workbegin) //开始做作业

}
