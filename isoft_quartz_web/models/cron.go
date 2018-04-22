package models

import "time"

type CronMeta struct {
	CronType        string    `json:"cron_type"`
	CronName        string    `json:"cron_name"`
	Id              int       `json:"id"`
	Second          string    `json:"second" orm:"default(*)"`
	Minute          string    `json:"minute" orm:"default(*)"`
	Hour            string    `json:"hour" orm:"default(*)"`
	Day             string    `json:"day" orm:"default(*)"`
	DayOfWeek       string    `json:"day_of_week" orm:"default(*)"`
	Year            string    `json:"year" orm:"default(*)"`
	CronExpression  string    `json:"cron_expression"`
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}
