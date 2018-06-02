package models

import "time"

type Catalog struct {
	Id              int64     `json:"id"`
	Author          string    `json:"author"`       // 分类作者
	CatalogName     string    `json:"catalog_name"` // 分类名称
	CatalogDesc     string    `json:"catalog_desc"` // 分类简介
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	Blogs           []*Blog   `json:"blogs" orm:"reverse(many)"`
}

type Blog struct {
	Id              int64     `json:"id"`
	Author          string    `json:"author"` // 博客作者
	BlogTitle       string    `json:"blog_title"`
	KeyWords        string    `json:"key_words"` // 搜索关键词
	Catalog         *Catalog  `json:"catalog" orm:"rel(fk)"`
	Content         string    `json:"content" orm:"type(text)"`
	BlogType        int8      `json:"blog_type"`   // 0:original, 1:translate, 2:reprint 分别表示原创、翻译、转载
	BlogStatus      int8      `json:"blog_status"` // 0:未发布 1:已发布
	Views           int64     `json:"views"`       // 观看次数
	Edits           int64     `json:"edits"`       // 编辑次数
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}
