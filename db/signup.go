package db

import (
	"fmt"

	"github.com/ddessilvestri/gambit-user/models"
	"github.com/ddessilvestri/gambit-user/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(signupModel models.SignUp) error {
	fmt.Println("Start Registring")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "INSERT INTO users(User_Email, User_UUID, User_DateAdd) VALUES (')" + signupModel.UserEmail + "','" + signupModel.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully Signup ")

	return nil
}
