package handlers

import (
	"log"
	"net/http"
	"test1/connection"

	"github.com/gin-gonic/gin"
)

type ColumnType struct {
	Name             string `json:"nam"`
	DatabaseTypeName string `json:"dtn"`
}

// GetUser ดึงข้อมูลผู้ใช้
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "message": "User found"})
}

// CreateUser สร้างผู้ใช้ใหม่
func CreateUser(c *gin.Context) {
	var user struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ทดสอบ connect db
	db, err := connection.ConnectDBGorm("dh_vendor_portal")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Custom Query แบบ Raw SQL
	query := `SELECT id, calendar_id, calendar_name, effective_date 
FROM appointment_calendar 
WHERE 0=0`

	var calendars []Calendar
	result := db.Raw(query).Scan(&calendars) // ส่งค่า "My Calendar Name" เป็นตัวแปร
	if result.Error != nil {
		log.Fatalf("Error querying data: %v\n", result.Error)
	}
	// แสดงผลข้อมูล
	// for _, cal := range calendars {
	// 	fmt.Printf("ID: %d, CalendarID: %s, CalendarName: %s, EffectiveDate: %s\n",
	// 		cal.ID, cal.CalendarID, cal.CalendarName, cal.EffectiveDate.Format("2006-01-02"))
	// }

	//ss
	// query := `SELECT id, calendar_id, calendar_name, effective_date
	// 		  FROM appointment_calendar
	// 		  WHERE 0=0`
	// rows, err := dx.Query(query)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// ดึงข้อมูลจากตาราง
	// var calendars []Calendar
	// if err := dx.Table("your_table_name").Find(&calendars).Error; err != nil {
	// 	log.Fatalf("Error querying calendars: %v\n", err)
	// }
	//defer dx.Close()

	// var calendars []Calendar
	// for rows.Next() {
	// 	var cal Calendar
	// 	err := rows.Scan(&cal.ID, &cal.CalendarID, &cal.CalendarName, &cal.EffectiveDate)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// 	calendars = append(calendars, cal)
	// }

	// if err := rows.Err(); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }

	//c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
	c.JSON(http.StatusCreated, calendars)
}

// Struct สำหรับเก็บข้อมูล
type Calendar struct {
	ID            string `json:"id"`
	CalendarID    string `json:"calendar_id"`
	CalendarName  string `json:"calendar_name"`
	EffectiveDate string `json:"effective_date"`
}
