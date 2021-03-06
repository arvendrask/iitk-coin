package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var (
	router = gin.Default()
)

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		panic(err)
	}
	router.POST("/signup", Signup)
	router.POST("/login", Login)
	router.GET("/secretpage", Secretpage)
	router.POST("/awardcoins", Awardcoins)
	router.GET("/getcoins", Getcoins)
	router.GET("/redeemadminlist", Redeemadminlist)
	router.POST("/redeemuser", Redeemuser)
	router.POST("/transfercoins", Transfercoins)
	router.POST("/redeemadminaction", Redeemadminaction)
	log.Fatal(router.Run(":8080"))

}
