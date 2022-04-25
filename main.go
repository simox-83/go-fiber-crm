package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/simox-83/go-fiber-crm-basic/database"
	"github.com/simox-83/go-fiber-crm-basic/model"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:@/crm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection opened")
	//defer database.DBConn.Close()
	database.DBConn.AutoMigrate(&model.Lead{})

}

func main() {
	initDatabase()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.DB().Close()

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", model.GetLeads)
	app.Get("/api/v1/lead/:id", model.GetLead)
	app.Post("/api/v1/lead", model.NewLead)
	app.Delete("/api/v1/lead/:id", model.DeleteLead)
}
