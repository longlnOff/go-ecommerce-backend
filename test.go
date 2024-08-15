package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    sqlDB, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
    }
 
    _, err = gorm.Open(mysql.New(mysql.Config{
        Conn: sqlDB,
        }), &gorm.Config{})
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
    // defer the close till after the main function has finished
    // executing
    defer sqlDB.Close()

}

