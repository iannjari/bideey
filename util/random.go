package util

import (
	"math/rand"
	"time"
)

func GetRandomStrCode() string {
	length := 5
	ran_str := make([]byte, length)
	for i := 0; i < length; i++ {

		rand.New(rand.NewSource(time.Now().UnixNano()))

		// String
		charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

		// Getting random character
		c := charset[rand.Intn(len(charset))]

		ran_str[i] = c

	}

	return string(ran_str)
}
