package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	// Set the environment variable
	os.Setenv(key, "postgres://tn:vnhcmmnmn@10.0.4.200:5432/tn?sslmode=verify-full")
	val := GetEnvDefault(key, "postgres://tn:vnhcmmnmn@10.0.4.200:5432/tn?sslmode=verify-full")
	log.Println("The value is :" + val)

	os.Unsetenv("The value is :" + val)

	os.Unsetenv(key)
	val = GetEnvDefault(key, "postgres://tn:vnhcmmnmn@10.0.4.200:5432/tn?/sslmode=verify-full")
	log.Println("The default value is :" + val)
}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
