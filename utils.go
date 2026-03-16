package utils

import (
	"regexp"
	"math/rand"
	"time"
)

// GenerateID generates a random string of the given length
func GenerateID(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// ValidateURL checks if the URL is valid
func ValidateURL(url string) bool {
	regex := "^(http|https)://[^\\s/$.?#].[^\\s]*$"
	re := regexp.MustCompile(regex)
	return re.MatchString(url)
}

// ValidatePackageName checks if the package name is valid
func ValidatePackageName(name string) bool {
	regex := "^[a-zA-Z][a-zA-Z0-9._]*$"
	re := regexp.MustCompile(regex)
	return re.MatchString(name)
}