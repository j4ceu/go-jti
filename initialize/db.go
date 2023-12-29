package initialize

import (
	"fmt"
	"go-jti/config"
	"go-jti/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() *gorm.DB {
	postgresConfig := config.LoadPostgresConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Name,
		postgresConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	tx := db.Session(&gorm.Session{PrepareStmt: true})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	migrateDatabase(tx)

	return tx

}

func migrateDatabase(db *gorm.DB) {

	errMigrate := db.AutoMigrate(
		&models.PhoneNumber{},
	)

	if errMigrate != nil {
		log.Fatal("Gagal Migrate")
	}

	log.Println("Migrate Berhasil!")

}
