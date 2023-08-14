package exception

import (
	"errors"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func ValidationError(err error) error {
	return errors.New("validation error")
}
