package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/go_crud_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)

	type User struct {
		gorm.Model
		Name string
		Age  int64
	}
	var user User
	app := fiber.New()

	app.Post("/adduser", func(c *fiber.Ctx) error {
		username := c.Query("username")
		age, err := strconv.ParseInt(c.Query("age"), 64, 64)
		fmt.Println(err)

		user := User{Name: username, Age: age}
		db.Create(&user)
		return c.JSON(&user)
	})

	app.Get("/username", func(c *fiber.Ctx) error {
		username := c.Query("username")
		db.First(&user, "Name = ?", username)
		return c.JSON(user)
	})

	app.Listen(":3001")
}
