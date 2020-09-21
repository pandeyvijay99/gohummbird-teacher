package handlers

import (
	"github.com/pandeyvijay99/gohummbird-teacher/db"
	"github.com/pandeyvijay99/gohummbird-teacher/models"
	//"strconv"
	//"fmt"
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
	id := c.Params("emp_id")
	keyname := c.Params("type")
	db := db.DB
	teacher := new(models.Teacher)
	
	db.Where("emp_id = ?", id).First(&teacher)
	
	if(keyname == "bank-details"){
	
		bankdetails := new(models.Bankdetails)
		db.Where("user_id = ?", teacher.UserId).First(&bankdetails)
		
		if err := c.BodyParser(bankdetails); err != nil {
			c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher bank details", "data": err, "st": bankdetails})
			return
		}
		db.Save(&bankdetails)
		c.JSON(fiber.Map{"status": "success", "message": "Teacher successfully updated", "data": bankdetails})
		return
	}else if (keyname == "additional-info") {

		bankdetails1 := new(models.Bankdetails)
		db.Where("user_id = ?", teacher.UserId).First(&bankdetails1)
		//c.JSON(fiber.Map{"status": "success", "message": "Teacher Additional successfully updated", "data": bankdetails1})	
		//return

		if err := c.BodyParser(bankdetails1); err != nil {
			c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher bank details", "data": err, "st": bankdetails1})
			return
		}
		db.Save(&bankdetails1)
		
		// db.Model(&bankdetails).Updates(bankdetails{Name: "hello", Age: 18})
		c.JSON(fiber.Map{"status": "success", "message": "Teacher Additional successfully updated", "data": bankdetails1})	
		return
	}else if (keyname == "academic-info") {

		academicinfo := new(models.TeacherTimeTable)
		db.Where("user_id = ?", teacher.UserId).First(&academicinfo)

		//c.JSON(fiber.Map{"status": "success", "message": "Teacher Additional successfully updated", "data":academicinfo })	
		//return

		if err := c.BodyParser(academicinfo); err != nil {
			c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher academic details", "data": err, "st": academicinfo})
			return
		}
		db.Save(&academicinfo)
		
		// db.Model(&bankdetails).Updates(bankdetails{Name: "hello", Age: 18})
		c.JSON(fiber.Map{"status": "success", "message": "Teacher Academic details  successfully updated", "data": academicinfo})	
		return
	}
	if (keyname == "default") {
		teacher := new(models.Teacher)	
	db.Where("emp_id = ?", id).First(&teacher)
	if err := c.BodyParser(teacher); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher  details", "data": err, "st": teacher})
		return
	}
	db.Save(&teacher)
	
	
	c.JSON(fiber.Map{"status": "success", "message": "Teacher Academic details  successfully updated", "data": teacher})	
	return	
	 
	// db.Save(&teacher)
	
	// c.JSON(fiber.Map{"status": "success", "message": "Teacher successfully updated", "data": teacher})
}
}

// GetAllStudents Function
func GetAllTeachers(c *fiber.Ctx) {
	db := db.DB
	//type Teacher struct {
	//	EmpID         string
	//	Name          string
	//	Designation   string
	//	PhoneNumber      string
		//InChargeOf    string 
	//	Email         string
	//	Photo         string
	//}
	var teacher []models.Teacher

	var total int64
	db.Model(&teacher).Count(&total)
	db.Find(&teacher)

	c.JSON(fiber.Map{"status": "success", "message": "All teachers", "data": teacher, "total": total})
}

// GetStudent Function
func GetTeacher(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DB
	var teacher models.Teacher
	db.Find(&teacher, id)
	if teacher.Name == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No teacher found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Teacher found", "data": teacher})
}
//Store Bank Details 
func StoreBankDetails(c *fiber.Ctx) {
	id := c.Params("id")
	keyname := c.Params("keyname")
	db := db.DB
	teacher := new(models.Teacher)
	db.First(&teacher,id)
	if err := c.BodyParser(teacher); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't add teacher", "data": err})
		return
	}
	

	if keyname== "bank-details" {
		bankdetails := new(models.Bankdetails)
		db.Where("user_id = ?", id).First(&bankdetails)
		if err := c.BodyParser(bankdetails); err != nil {
			c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update teacher bank details", "data": err, "st": bankdetails})
			return
		}
		db.Save(&bankdetails)
		c.JSON(fiber.Map{"status": "success", "message": "Teacher successfully updated", "data": bankdetails})
		return
} else if keyname == "additional-info"{
	
}

}

