package db

import (
	"os"

	"github.com/ddessilvestri/gambit-user/models"
	"github.com/ddessilvestri/gambit-user/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
