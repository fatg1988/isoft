package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Note struct {
	Id              int           `json:"id"`
	NoteName        string        `json:"note_name"`         // 笔记名称
	NoteOwner       string        `json:"note_owner"`        // 笔记作者
	NoteKeyWords    string        `json:"note_key_words"`    // 笔记关键字,用于模糊搜索使用
	NoteContent     orm.TextField `json:"note_content"`      // 笔记内容
	CreatedBy       string        `json:"created_by"`        // 创建人
	CreatedTime     time.Time     `json:"created_time"`      // 创建时间
	LastUpdatedBy   string        `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time     `json:"last_updated_time"` // 修改时间
	EditTime        int           `json:"edit_time"`         // 编辑次数
	ViewTime        int           `json:"view_time"`         // 阅读次数
}

func GetNoteOwnerTotalAmount() (count int64) {
	o := orm.NewOrm()
	var list orm.ParamsList
	num, err := o.Raw("SELECT COUNT(DISTINCT note_owner) FROM note").ValuesFlat(&list)
	if err == nil && num > 0 {
		// reflect.TypeOf(list[0]).Name() == string
		count, _ = strconv.ParseInt(list[0].(string), 10, 64)
	}
	return
}

func QueryNoteById(note_id int) (note Note, err error) {
	o := orm.NewOrm()
	o.QueryTable("note").Filter("id", note_id).One(&note)
	return
}

func UpdateNoteById(note *Note) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("note").Filter("id", note.Id).Update(orm.Params{
		"note_name": note.NoteName, "note_key_words": note.NoteKeyWords,
		"note_content": note.NoteContent, "edit_time": orm.ColValue(orm.ColAdd, 1),
	})
	return
}

func GetNoteTotalAmount() (count int64) {
	o := orm.NewOrm()
	count, _ = o.QueryTable("note").Count()
	return
}

func QueryNoteExist(note_name, user_name string) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("note").Filter("note_name", note_name).Filter("note_owner", user_name).Count()
	return
}

func AddNote(note *Note) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(note)
	return id, err
}

func QueryNote(condArr map[string]string, page int, offset int) (notes []Note, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("note")
	cond := orm.NewCondition()

	if _, ok := condArr["NoteOwner"]; ok {
		cond = cond.And("NoteOwner", condArr["NoteOwner"])
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&notes)
	return
}
