package clients

import (
	"fmt"

	"github.com/maxabella/appgym/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql interface {
	Migrate()
	Connect()
}

type Mysql_Client struct { // Cliente de mysql database que se dedica a hacer la conexion, migracion y proveer un punto de acceso a la base de datos
	db *gorm.DB
}

func (s *Mysql_Client) Connect() { // Funcion para Inicializar
	dsn := "root:12345678@tcp(127.0.0.1:3307)/gymdb"
	var err error
	s.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}
	fmt.Println("Mysql Client Connected to database")
}

func (s *Mysql_Client) Migrate() {
	s.db.AutoMigrate(dao.ActivityType{}, dao.UserType{}, dao.User{}, dao.Activity{}, dao.ActivityHour{}, dao.Inscription{})
}
