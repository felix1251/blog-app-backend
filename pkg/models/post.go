package models

import(
	"github.com/jinzhu/gorm"
	"github.com/felix1251/go-rest-api/pkg/config"
)

var db *gorm.DB

type Post struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Description string `json:"author"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

