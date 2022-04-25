package model

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/simox-83/go-fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string `gorm:"" json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)

}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var l Lead
	db.Find(&l, id)
	if l.Name == "" {
		c.Status(404).Send("No ID found")
		return
	}
	c.JSON(l)

}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	l := new(Lead)
	err := c.BodyParser(l)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&l)
	c.JSON(l)

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var l Lead
	db.Delete(&l, id)
	if l.Name == "" {
		c.Status(404).Send("No ID found")
		return
	}

	c.Send("Lead Successfully Deleted")

}
