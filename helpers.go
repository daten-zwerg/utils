package helpers

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SetProxy() {
	host := GetEnv("PROXY_HOST", "")
	port := GetEnv("PROXY_PORT", "")
	credentials := GetEnv("PROXY_CREDENTIALS", "")

	service := fmt.Sprintf("http://%s@%s:%s", credentials, host, port)

	if err := os.Setenv("HTTP_PROXY", service); err != nil {
		log.Println("error on setting HTTP Proxy", err)
	}

	if err := os.Setenv("HTTPS_PROXY", service); err != nil {
		log.Println("error on setting HTTPS Proxy", err)
	}
}

func NewMongodbUri() string {
	host := GetEnv("MONGODB_HOST", "localhost")
	port := GetEnv("MONGODB_PORT", "27017")
	credentials := GetEnv("MONGODB_CREDENTIALS", "")

	uri := fmt.Sprintf("mongodb://%s@%s:%s", credentials, host, port)
	if credentials == "" {
		uri = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	return uri
}
