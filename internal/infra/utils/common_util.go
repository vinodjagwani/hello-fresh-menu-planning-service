package utils

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func ValidUuid(ud string) (uuid.UUID, error) {
	uuidFrmStr, uuidErr := uuid.FromString(ud)
	if uuidErr != nil {
		log.Printf("An invalid uuid: %s with error %s", uuidFrmStr, uuidErr)
		return uuidFrmStr, uuidErr
	}
	return uuidFrmStr, uuidErr
}
