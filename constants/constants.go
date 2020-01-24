package constants

import "os"

// MongoURI : MongoDB URI link
var MongoURI string = os.Getenv("mongo")

// JWTKey : Secret key
var JWTKey = []byte(os.Getenv("jwt_key"))
