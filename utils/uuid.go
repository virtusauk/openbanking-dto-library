package utils

import (
	"fmt"
	"github.com/gofrs/uuid"
)

func GenarateUUID() string {

	var u2 uuid.UUID
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return ""
	}

	return u2.String()
}
