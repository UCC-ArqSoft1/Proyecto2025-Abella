package clients

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/maxabella/appgym/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql interface {
	Migrate()
	Connect()
	waitForDB(string) (*gorm.DB, error)
}

type Mysql_Client struct { // Cliente de mysql database que se dedica a hacer la conexion, migracion y proveer un punto de acceso a la base de datos
	db *gorm.DB
}

func (s *Mysql_Client) Connect() { // Funcion para Inicializar

	err := godotenv.Load("../../ENV.env")
	if err != nil {

	}

	fmt.Println(os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//dsn := "root:12345678@tcp(127.0.0.1:3307)/gymdb"
	dsn := fmt.Sprintf("%s:%v@tcp(%s:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	fmt.Println("DSN: ", dsn)
	s.db, err = s.waitForDB(dsn)
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}
	fmt.Println("Mysql Client Connected to database")
}

func (s *Mysql_Client) waitForDB(dsn string) (*gorm.DB, error) { // Funcion de la profe para timeout de la base de datos
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil && sqlDB.Ping() == nil {
				return db, nil
			}
		}

		log.Printf("⏳ Esperando base de datos... intento %d/10\n", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("❌ no se pudo conectar a la BD: %v", err)
}

func (s *Mysql_Client) Migrate() {
	s.db.AutoMigrate(dao.ActivityType{}, dao.UserType{}, dao.User{}, dao.Activity{}, dao.ActivityHour{}, dao.Inscription{})
}
