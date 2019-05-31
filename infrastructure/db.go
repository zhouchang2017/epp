package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zhouchang2017/epp/app/modules/inventories/models"
	models2 "github.com/zhouchang2017/epp/app/modules/supplies/models"
	"github.com/zhouchang2017/epp/common"
	"github.com/zhouchang2017/epp/config"
	"log"
)

var db *gorm.DB

func dbInit() {
	conf := config.Config.DB
	entity, err := gorm.Open(
		conf.Driver,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
			conf.Charset,
			conf.ParseTime,
			conf.Local,
		),
	)

	if err != nil {
		log.Fatalln("mysql conn failed,", err)
	}

	log.Println("连接数据库")

	entity = entity.Debug()

	db = entity

	db.AutoMigrate(models.Inventory{})
	db.AutoMigrate(models2.SupplyOrder{})
	db.AutoMigrate(models2.SupplyOrderItem{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}

// 分页
func Page(db *gorm.DB, out interface{}, page, perPage uint, where interface{}, args ...interface{}) (res *common.Paginate, err error) {
	if perPage == 0 {
		perPage = 15
	}
	if page == 0 {
		page = 1
	}

	var count uint

	db.Find(out).Count(&count)

	err = db.Where(where, args...).Limit(perPage).Offset((page - 1) * perPage).Find(out).Error

	return common.NewPaginate(out, page, perPage, count), err
}


// http://api.test/hubs/1?include=hot_users:limit(3).posts:fields(id|title):limit(3)
func ResloveInclude(params string)  {
	
}
