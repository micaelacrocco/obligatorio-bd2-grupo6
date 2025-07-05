package configuration

import "os"

const DbHostKey = "DB_HOST"
const DbPortKey = "DB_PORT"
const DbUserKey = "DB_USER"
const DbPasswordKey = "DB_PASSWORD"
const DbNameKey = "DB_NAME"

var JWT_SECRET = "mi-clave"

func GetDSN() string {
	dbHost := os.Getenv(DbHostKey)
	dbPort := os.Getenv(DbPortKey)
	dbUser := os.Getenv(DbUserKey)
	dbPassword := os.Getenv(DbPasswordKey)
	dbName := os.Getenv(DbNameKey)

	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}
