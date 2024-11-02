package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
}

func (p *Product) TableName() string {
	return "products"
}

var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "mysql_root", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("database auto migrate failed")
	}

	// insert
	if err := db.Create(&Product{Code: "D1", Price: 100}).Error; err != nil {
		log.Fatalf("create product failed: %v", err)
	}
	// query
	PrintProducts(db)

	// find first record
	product := &Product{}
	if err := db.Where("code = ?", "D2").First(&product).Error; err != nil {
		log.Fatalf("get product failed: %v", err)
	}

	// update
	product.Price = 300
	if err := db.Save(product).Error; err != nil {
		log.Fatalf("update product failed: %v", err)
	}

	// delete
	if err := db.Where("code = ?", "D1").Delete(&Product{}).Error; err != nil {
		log.Fatalf("delete product failed: %v", err)
	}

	PrintProducts(db)
}

func PrintProducts(db *gorm.DB) {
	products := make([]*Product, 0)
	var count int64
	// 可以一次查询出来吗？
	d := db.Where("code LIKE ?", "%D%").Offset(0).Limit(10).Order("id desc").Find(&products).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("list products failed: %v", d.Error)
	}
	log.Printf("total count: %d", count)
	for _, product := range products {
		log.Printf("code: %s, price: %d \n", product.Code, product.Price)
	}
}
