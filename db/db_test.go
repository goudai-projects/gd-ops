package db

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

type TestUser struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
}

func TestInit(t *testing.T) {
	Init("root:5c15faccacf44346b725c39d79e02b73@(192.168.0.201:3306)/gd_ops?charset=utf8mb4&parseTime=True&loc=Local")
	db.AutoMigrate(&TestUser{})
	user := TestUser{
		Username: "gd",
	}
	db.Table("test_users").FirstOrCreate(&user)
	u, _ := json.Marshal(user)
	fmt.Println(string(u))
	var users []TestUser
	db.Find(&users)
	for i := range users {
		u, _ = json.Marshal(users[i])
		fmt.Println(string(u))
	}
}
