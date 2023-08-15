package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
    host := "localhost"
    port := 5432
    user := "user"
    password := "password"
    dbname := "database"

    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("fail to connect to database")
    }

    return db
}

// TODO CRIAR VARIAVEIS DE AMBIENTE