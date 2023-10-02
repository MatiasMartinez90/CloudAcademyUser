package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/CloudAcademyUser/models"
	"github.com/CloudAcademyUser/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error

// creo la variable Db con un puntero del tipo DB (por un tema de velocidad se usa un puntero para cosas de DB, en vez de ser un tipo de dato)
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName")) // Se trae del env del SO el nombre del secreto y consulta la funcion GetSecret
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel)) // myqsql es el data source, llama a la funcion ConnStr que le pasa los datos de conexion
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName) // Paso un parametro para que no de error la conexion a la db, y paso las credenciales de la db
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza UserIsAdmin")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecución exitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "User is not Admin"

}

func UserExists(UserUUID string) (error, bool) {
	fmt.Println("Comienza UserExists")

	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + UserUUID + "'"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserExists > Ejecución exitosa - valor devuelto " + valor)

	if valor == "1" {
		return nil, true
	}
	return nil, false
}
