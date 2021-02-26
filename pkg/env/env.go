package env

import "os"

// Getenv is
func Getenv(key string, defaultvalue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultvalue
	}
	return value
}
