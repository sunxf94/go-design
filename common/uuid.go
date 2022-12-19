package common

import (
	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	return uuid.NewV1().String()
}
