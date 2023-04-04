package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *sql.DB

// func MysqlConnection() {
// 	// Capture connection properties.
// 	cfg := mysql.Config{
// 		User:   os.Getenv("DBUSER"),
// 		Passwd: os.Getenv("DBPASS"),
// 		// Net:    "tcp",
// 		Addr:   "127.0.0.1:3306",
// 		DBName: "recordings",
// 	}
// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }

func DBconnection() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Mani*123@tcp(127.0.0.1:3306)/record?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connected", db)
	}
}
