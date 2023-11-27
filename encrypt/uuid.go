package encrypt

import "github.com/google/uuid"

func GetUuid() string {
	v4 := uuid.New()
	return v4.String()
}
