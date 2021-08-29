package utils

import (
	"math/rand"
	"time"
)

func GenerateKey(length uint) string {
	chars := "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	shortUrl := ""
	for i := 0; i < int(length); i++ {
		r := rand.Intn(len(chars))
		shortUrl += string(chars[r])
	}
	return shortUrl
}
