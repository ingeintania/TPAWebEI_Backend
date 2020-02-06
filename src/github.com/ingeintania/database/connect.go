package database


import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

const Dbname = "postgres"
const Dbhost = "127.0.0.1"
const Dbport = "5432" //default is 5432
const Dbuser = "postgres" //default is postgres
const Dbpassword = "prk"

func Connect()(*gorm.DB, error){

	return gorm.Open("postgres","host="+Dbhost+" port="+Dbport+" user="+Dbuser+" dbname="+Dbname+" password="+Dbpassword+" sslmode=disable")
}