package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type TopicTheme struct {
	Id              int       `json:"id"`
	TopicId         int       `json:"topic_id"`                       // 评论主题 id
	TopicType       string    `json:"topic_type"`                     // 评论主题类型
	TopicContent    string    `json:"topic_content" orm:"size(4000)"` // 评论主题内容
	CreatedBy       string    `json:"created_by"`                     // 评论主题创建人
	CreatedTime     time.Time `json:"created_time"`                   // 评论主题创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`                // 评论主题修改人
	LastUpdatedTime time.Time `json:"last_updated_time"`              // 评论主题修改时间
}

type TopicReply struct {
	Id              int         `json:"id"`
	ParentId        int         `json:"parent_id" orm:"default(0)` // 父级评论回复 id
	TopicTheme      *TopicTheme `orm:"rel(fk)" json:"topic_theme"`
	ReplyType       string      `json:"reply_type"`                     // 评论类型
	ReplyContent    string      `json:"reply_content" orm:"size(4000)"` // 评论内容
	ReferUserName   string      `json:"refer_user_name"`                // 被评论人
	SubReplyAmount  int         `json:"sub_reply_amount"`               // 子评论数
	CreatedBy       string      `json:"created_by"`                     // 评论回复创建人
	CreatedTime     time.Time   `json:"created_time"`                   // 评论回复创建时间
	LastUpdatedBy   string      `json:"last_updated_by"`                // 评论回复修改人
	LastUpdatedTime time.Time   `json:"last_updated_time"`              // 评论回复修改时间
}

func AddTopicTheme(topic_theme *TopicTheme) (id int64, err error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("topic_theme").Filter("topic_id", topic_theme.TopicId).Filter("topic_type", topic_theme.TopicType).Count()
	if count == 0 {
		id, err = o.Insert(topic_theme)
	}
	return
}

func FilterTopicTheme(topic_id int, topic_type string) (topic_theme TopicTheme, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("topic_theme").Filter("topic_id", topic_id).Filter("topic_type", topic_type).One(&topic_theme)
	return
}

func AddTopicReply(topic_reply *TopicReply) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(topic_reply)
	return
}

func FilterTopicReply(topic_id int, topic_type string, parent_id int) (topic_replys []TopicReply, err error) {
	o := orm.NewOrm()
	// 查询 topicTheme
	topicTheme, _ := FilterTopicTheme(topic_id, topic_type)
	_, err = o.QueryTable("topic_reply").Filter("topic_theme_id", topicTheme.Id).Filter("parent_id", parent_id).
		OrderBy("-created_time").All(&topic_replys)
	return
}

func ModifySubReplyAmount(id int) {
	o := orm.NewOrm()
	o.QueryTable("topic_reply").Filter("id", id).Update(orm.Params{
		"sub_reply_amount": orm.ColValue(orm.ColAdd, 1),
	})
}
