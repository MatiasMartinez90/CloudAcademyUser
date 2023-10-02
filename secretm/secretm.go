package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/CloudAcademyUser/awsgo"
	"github.com/CloudAcademyUser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Recibe el nombre del secreto, y devuelve un modelo (de models) y un error (x si hubo)
func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {

	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	// la variable usa el servicio de secret manager, y el newfromconfig es inicializar el servicio
	// Nos pide el awsconfig de cfg
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	// Nos traemos el secreto. Pasamos el contexto (que rige la ejecucion de toda la lambda), le pasamos el puntero de secret manager-->
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		//Tranformamos nos pide el secret id, y lo trasformamos con una funcion de aws.
		SecretId: aws.String(nombreSecret), // Recibe un string, y la funcion del paquete aws lo transforma en un puntero de string
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// Parsea el json codificado que me llego y lo va a enviar a la estructura de models.
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret OK " + nombreSecret)
	return datosSecret, nil
}
