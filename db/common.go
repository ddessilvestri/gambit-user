package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ddessilvestri/gambit-user/models"
	"github.com/ddessilvestri/gambit-user/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error{
	Db, err = sql.Open("mysql",ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err 
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err	
	}

	fmt.Println("Successfully database connection")
	return nil
}

func ConnStr(json models.SecretRDSJson) string {
	var dbUser, authToken,dbEndpoint,dbName string
	dbUser = json.Username
	authToken = json.Password
	dbEndpoint = json.Host
	dbName = "gambit"
	dsn :=fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPassword=true",dbUser,authToken,dbEndpoint,dbName)
	fmt.Println(dsn)
	return dsn
}