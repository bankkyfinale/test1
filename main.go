package main

import (
	"log"
	"net/http"
	"test1/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server on :8080...")

	// Initialize routes
	//r := router.InitRoutes()
	// สร้าง Gin Engine
	r := gin.Default()

	// เรียกใช้ฟังก์ชัน RegisterRoutes
	router.RegisterRoutes(r)

	// พิมพ์ route ทั้งหมดออกใน console
	//router.PrintRoutes(r)

	// Start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
