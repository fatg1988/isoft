package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/isoft/isoft_learning_web/models"
	"strings"
	"time"
)

type NoteController struct {
	beego.Controller
}

func (this *NoteController) ViewNote() {
	personal := this.GetString("personal")
	if personal == "personal" {
		this.Data["Personal"] = "personal"
	}

	user_name := this.Ctx.Input.Session("UserName").(string)
	note_id, err := this.GetInt("note_id")

	if err == nil {
		note, _ := models.QueryNoteById(note_id)
		this.Data["Note"] = &note
		this.Data["NoteCollect"] = models.IsFavorite(user_name, note_id, "note_collect")
		this.TplName = "note/note_view.html"
	}
}

func (this *NoteController) QueryNoteExist() {
	//初始化
	note_name := this.GetString("note_name")
	user_name := this.Ctx.Input.Session("UserName").(string)

	if strings.TrimSpace(note_name) != "" {
		count, err := models.QueryNoteExist(note_name, user_name)
		if err == nil && count == 0 {
			this.Data["json"] = &map[string]interface{}{"flag": false}
		} else {
			this.Data["json"] = &map[string]interface{}{"flag": true}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"flag": false}
	}
	this.ServeJSON()
}

func (this *NoteController) QueryNoteHtmlById() {
	note_id, err := this.GetInt("note_id")
	if err == nil {
		note, _ := models.QueryNoteById(note_id)
		this.Data["NoteContent"] = note.NoteContent.String()
	}
	this.TplName = "note/note_iframe.html"
}

func (this *NoteController) QueryNoteById() {
	note_id, err := this.GetInt("note_id")
	if err == nil {
		note, _ := models.QueryNoteById(note_id)
		this.Data["json"] = &map[string]interface{}{"note": note}
	}
	this.ServeJSON()
}

func (this *NoteController) CreateOrUpdateNote() {
	method := this.Ctx.Request.Method
	if method == "GET" {
		note_id, err := this.GetInt("note_id")
		if err == nil {
			// 表单回显需要的笔记 id
			this.Data["NoteId"] = note_id
		}
		this.Data["NoteOwnerTotalAmount"] = models.GetNoteOwnerTotalAmount()
		this.Data["NoteTotalAmount"] = models.GetNoteTotalAmount()
		this.TplName = "note/note_edit.html"
	} else {
		user_name := this.Ctx.Input.Session("UserName").(string)
		var note models.Note
		note.NoteName = this.GetString("note_name")
		note.NoteOwner = user_name
		note.NoteKeyWords = this.GetString("note_key_words")
		note.NoteContent = orm.TextField(this.GetString("note_content"))
		note.CreatedBy = user_name
		note.CreatedTime = time.Now()
		note.LastUpdatedBy = user_name
		note.LastUpdatedTime = time.Now()
		note.EditTime = 1 // 新增时编辑次数为 1

		note_id, err := this.GetInt("note_id")
		if err == nil {
			// 表单回显需要的笔记 id
			note.Id = note_id
			err := models.UpdateNoteById(&note)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"flag": true}
			} else {
				this.Data["json"] = &map[string]interface{}{"flag": false, "msg": err.Error()}
			}
		} else {
			// 新增操作
			_, err := models.AddNote(&note)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"flag": true}
			} else {
				this.Data["json"] = &map[string]interface{}{"flag": false, "msg": err.Error()}
			}
		}
		this.ServeJSON()
	}
}

func (this *NoteController) ListNote() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	condArr := make(map[string]string)
	filter_type := this.GetString("filter_type")
	page, _ := this.GetInt("page", 1)      // 页数
	offset, _ := this.GetInt("offset", 10) // 每页记录数
	// 查询私人笔记
	if filter_type == "personal" {
		condArr["NoteOwner"] = user_name
		this.Data["Personal"] = true
	}
	notes, _, _ := models.QueryNote(condArr, page, offset)
	this.Data["Notes"] = &notes
	// 课程管理界面
	this.Layout = "course/home_manage.html"
	this.TplName = "note/note_list.html"
}

func (this *NoteController) CollectList() {
	// 课程管理界面
	this.Layout = "course/home_manage.html"
	this.TplName = "note/note_collect.html"
}
