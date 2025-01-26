package database

var connection string

func init() { // This will executed when imported
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
