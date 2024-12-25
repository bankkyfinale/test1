package router

// import (
// 	"fmt"
// 	"test1/internal/handler"

// 	"github.com/gorilla/mux"
// )
import (
	"test1/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes application routes
// func InitRoutes() *mux.Router {
// 	r := mux.NewRouter()
// 	//Define routes
// 	r.HandleFunc("/", handler.HelloHandler).Methods("GET")
// 	r.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")

// 	return r
// }

// ฟังก์ชัน printRoutes สำหรับพิมพ์ routes ทั้งหมดออกบน console
// func PrintRoutes(r *mux.Router) {
// 	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 		pathTemplate, _ := route.GetPathTemplate()
// 		methods, _ := route.GetMethods()
// 		fmt.Printf("Route: %s | Methods: %v\n", pathTemplate, methods)
// 		return nil
// 	})
// }

// RegisterRoutes ลงทะเบียน Routes ทั้งหมด
func RegisterRoutes(r *gin.Engine) {
	// ใช้ ForceConsoleColor เพื่อบังคับให้มีสีใน console output
	gin.ForceConsoleColor()

	// Health Check Route
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Welcome to Gin!"})
	// })

	// User Routes
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", handlers.GetUser)  // ดึงข้อมูลผู้ใช้
		userGroup.POST("/", handlers.CreateUser) // สร้างผู้ใช้ใหม่
	}

}

// func PrintRoutes(r *gin.Engine) {
// 	for _, route := range r.Routes() {
// 		fmt.Printf("Method: %s | Path: %s\n", route.Method, route.Path)
// 	}
// }
