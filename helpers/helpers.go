package helpers

import (
	"UserRestApi/validators"
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Get environment parameter using a key, if not exists it returns the default value provide it
func GetEnvParamDefault(key string, defaultValue string) string {
	if defaultValue == "" {
		panic("Default value must not be an empty value")
	}
	value, exists := os.LookupEnv(key)
	if (exists && value == "") || (!exists) {
		value = defaultValue
	}
	return value
}

// Get environment parameter, if you need set a default value consider use GetEnvParamDefault
func GetEnvParam(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Env variable does not exist")
	}
	return value
}

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Bcrypt(value string) (string, error) {
	fmt.Println(value)
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost) // Second parameter is number of rounds
	return string(bytes), err
}

func CheckBcrypt(hash, compareValue string) bool {
	hashedValue := []byte(hash)
	plainValue := []byte(compareValue)
	err := bcrypt.CompareHashAndPassword(hashedValue, plainValue)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}

func ValidateRequestError(err error, typ int) (interface{}, error) {
	var verr validator.ValidationErrors
	var failed error = nil
	if errors.As(err, &verr) {
		if typ == validators.SIMPLE {
			return validators.NewJSONFormatter().Simple(verr), failed
		}
		if typ == validators.DESCRIPTIVE {
			return validators.NewJSONFormatter().Descriptive(verr), failed
		}
	}
	failed = errors.New("error por defecto")
	return []validators.ValidationError{}, failed
}
