package model

import "github.com/elton/project-layout/app/myapp/global"

// User represents a user.
type User struct {
	global.COMMODEL
	Name   string `gorm:"index:idx_name" json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
