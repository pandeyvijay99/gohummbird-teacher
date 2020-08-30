package handlers

import (
	"github.com/pandeyvijay99/gohummbird-teacher/db"
	"github.com/pandeyvijay99/gohummbird-teacher/models"

	"github.com/gofiber/fiber"
)

// AddNewStudent new Student
func AddNewTeacher(c *fiber.Ctx) {
	db := db.DB
	teacher := new(models.Teacher)
	if err := c.BodyParser(teacher); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't add teacher", "data": err})
		return
	}
	db.Create(&teacher)
	
	c.JSON(fiber.Map{"status": "success", "message": "Created Teacher", "data": teacher})
}

// UpdateStudent Student
func UpdateTeacher(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DB
	teacher := new(models.Teacher)
	// var teacher models.Teacher
	db.First(&teacher, id)
	if teacher.EmpID == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No teacher found with ID", "data": nil})
		return
	}
	if err := c.BodyParser(teacher); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher", "data": err})
		return
	}
	// db.Delete(&teacher)
	db.Save(&teacher)
	
	c.JSON(fiber.Map{"status": "success", "message": "Teacher successfully updated", "data": teacher})
}

// GetAllStudents Function
func GetAllTeachers(c *fiber.Ctx) {
	db := db.DB
	type Teacher struct {
		EmpID         string
		Name          string
		Designation   string
		MobileNo      string
		//InChargeOf    string 
		Email         string
		Photo         string
	}
	var teacher []Teacher

	var total int64
	db.Model(&teacher).Count(&total)
	db.Select("teachers.name,teachers.teacher_profile_pic,teachers.phone_number, teachers.emp_id,teachers.personal_email").Find(&teacher)

	c.JSON(fiber.Map{"status": "success", "message": "All teachers", "data": teacher, "total": total})
}

// GetStudent Function
func GetTeacher(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DB
	var teacher models.Teacher
	db.Find(&teacher, id)
	if teacher.id == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No teacher found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Teacher found", "data": teacher})
}
