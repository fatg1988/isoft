package controllers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_learning_web/models"
	"time"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) FilterTopicTheme() {
	// 获取课程 id
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	topic_theme, err := models.FilterTopicTheme(topic_id, topic_type)

	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "topic_theme": topic_theme}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CommentController) FilterTopicReply() {
	// 获取 topic_id 和 topic_type
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id")

	topic_replys, err := models.FilterTopicReply(topic_id, topic_type, parent_id)

	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "topic_replys": topic_replys}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err.Error()}
	}
	this.ServeJSON()
}

func (this *CommentController) AddTopicReply() {
	user_name := this.Ctx.Input.Session("UserName").(string)

	// 获取 topic_id 和 topic_type
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	// 查询 topicTheme
	topicTheme, _ := models.FilterTopicTheme(topic_id, topic_type)

	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id")
	// 获取评论内容
	reply_content := this.GetString("reply_content")
	// 获取被评论人员
	refer_user_name := this.GetString("refer_user_name")

	// 构造 TopicReply 实例
	var topic_reply models.TopicReply
	topic_reply.ParentId = parent_id
	topic_reply.ReplyType = "comment"
	topic_reply.ReplyContent = reply_content
	topic_reply.TopicTheme = &topicTheme
	topic_reply.ReferUserName = refer_user_name
	topic_reply.SubReplyAmount = 0
	topic_reply.CreatedBy = user_name
	topic_reply.CreatedTime = time.Now()
	topic_reply.LastUpdatedBy = user_name
	topic_reply.LastUpdatedTime = time.Now()

	_, err := models.AddTopicReply(&topic_reply)
	models.ModifySubReplyAmount(parent_id)

	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
