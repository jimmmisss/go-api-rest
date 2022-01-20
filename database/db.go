package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conexao() {
	dns := "host=localhost user=lezzr_go password=lezzr_go dbname=lezzr_go port=5431 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
