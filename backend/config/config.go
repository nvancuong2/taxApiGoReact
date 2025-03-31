package config

import "os"

func GetMongoURI() string {
	println("MONGO_URI", os.Getenv("MONGO_URI"))
	return os.Getenv("MONGO_URI")
}
