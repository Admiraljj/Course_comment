package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	drive := "mysql"
	username := "root"
	password := "1029admiral"
	host := "127.0.0.1"
	port := "3306"
	database := "course_comment"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(drive, args)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil
	}
	return db
}
