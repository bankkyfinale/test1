package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับเชื่อมต่อกับฐานข้อมูล PostgreSQL โดยรับชื่อฐานข้อมูลเป็นพารามิเตอร์
func ConnectDB(dbName string) (*sql.DB, error) {
	// กำหนดข้อมูลการเชื่อมต่อ (Connection String)

	PG_HOST := `internal-rds-uat.cluster-cthn4dylzmeq.ap-southeast-1.rds.amazonaws.com`
	PG_USER := `postgres`
	PG_PASS := `D4IlqPv0Yy`
	PG_PORT := `5432`

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", PG_USER, dbName, PG_PASS, PG_HOST, PG_PORT)

	// เชื่อมต่อกับ PostgreSQL
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	dbConn.SetConnMaxLifetime(0)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetMaxOpenConns(100)

	// ตรวจสอบการเชื่อมต่อ
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database:", dbName)
	return dbConn, nil
}

func ConnectDBGorm(dbName string) (*gorm.DB, error) {
	// กำหนดข้อมูลการเชื่อมต่อ (Connection String)

	PG_HOST := `internal-rds-uat.cluster-cthn4dylzmeq.ap-southeast-1.rds.amazonaws.com`
	PG_USER := `postgres`
	PG_PASS := `D4IlqPv0Yy`
	PG_PORT := `5432`

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", PG_USER, dbName, PG_PASS, PG_HOST, PG_PORT)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
