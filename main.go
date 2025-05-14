package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/ddessilvestri/gambit-user/awsgo"
	"github.com/ddessilvestri/gambit-user/db"
	"github.com/ddessilvestri/gambit-user/models"
)

func main() {
	lambda.Start(LambdaExec)
}

// test
func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSInit()

	if !ParameterValidation() {
		const errorMsg string = "parameter Error. 'SecretName' must be sent"
		fmt.Println(errorMsg)
		err := errors.New(errorMsg)
		return event, err
	}

	var data models.SignUp
	for row, attr := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = attr
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = attr
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error when retrieving the Secret" + err.Error())
		return event, err
	}

	err = db.SignUp(data)

	return event, err
}

func ParameterValidation() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
