package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/marcosvieirajr/payment/configs"
)

func NewDB(envs configs.Config) (*sql.DB, error) {

	host := envs.DBHost
	port := envs.DBPort
	user := envs.DBUser
	password := envs.DBPassword
	dbname := envs.DBName

	// starting db
	dataSourceInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err := sql.Open("postgres", dataSourceInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
