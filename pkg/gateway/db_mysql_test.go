package gateway

import (
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const testDBURL = "user:password@tcp(127.0.0.1:3307)/testdb?charset=utf8&parseTime=True&loc=Asia%2FTokyo"

func openMySQLForTest() (*gorm.DB, error) {
	return gorm.Open("mysql", testDBURL)
}

func setup() {
	db, err := gorm.Open("mysql", testDBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pos := strings.Index(wd, "pkg")
	dir := wd[0:pos] + "/sqls"
	driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+dir, "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
func withSetup(fn func()) {
	setup()
	fn()
}
