package utils

import "github.com/google/uuid"

func RandomUniqueStringGenerator() string {
	return uuid.New().String()
}
