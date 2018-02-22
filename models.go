package main

import "time"

// Country is
type Country struct {
	ID        uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:id"`
	SID       string `sql:"size:50"`
	Name      string `sql:"size:500"`
	Count     int
	UID       string
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
}
