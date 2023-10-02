package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/CloudAcademyUser/awsgo"
	"github.com/CloudAcademyUser/bd"
	"github.com/CloudAcademyUser/models"

	//"github.com/CloudAcademyUser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

// En lambda a dif del desarrollo tradicional no se puede poner el codigo en la funcion main, sino que esta tiene que llamar a otra donde si este el codigo.
//func main() {
//	lambda.Start(EjecutoLambda)
//}

// ctx context.Context --> Contexto --> Es como un director de orquesta para multihilo, en este proyecto no se usa pero hay que ponerlo
// events.CognitoEvent ---> Nos traemos de cognito, el evento de post confirmacion. Hay muchos eventos por las cosas que pasan en el userpool de cognito
// La funcion devuelve dos objetos, el mismo evento que recibio con valores, y el error para grabar en cloudwatch y que la lambda falle por error
//func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
//	awsgo.InicializoAWS()
//
//	if !ValidoParametros() {
//		fmt.Println("Error en los parametros. Debe enviar SecretName")
//		err := errors.New("error en los parametros debe enviar secretname")
//		return event, err
//	}

//	var datos models.SignUp

//	for row, att := range event.Request.UserAttributes {
//		switch row {
//		case "email":
//			datos.UserEmail = att
//			fmt.Println("Email = " + datos.UserEmail)
//		case "sub":
//			datos.UserUUID = att
//			fmt.Println("Sub = " + datos.UserUUID)
//		}
//	}

//	err := bd.ReadSecret()

//	if err != nil {
//		fmt.Println("error al leer el Secret " + err.Error())
//		return event, err
//	}

//	err = bd.SignUp(datos)

//	return event, err

//}

// Uso el paquete OS para manejar las variables de entorno de la lambda
// Para validar que devuelve un boolen de true o false (recibi bien los parametros o no)
//func ValidoParametros() bool {
// LookupEnv ---> Que se fije en las variables de entorno
//	_, traeParametro := os.LookupEnv("SecretName")
//	if !traeParametro {
//		return traeParametro
//	}
//
//	_, traeParametro = os.LookupEnv("UrlPrefix")
//	if !traeParametro {
//		return traeParametro
//	}

//	return traeParametro

//}

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. Debe enviar SecretName")
		err := errors.New("error en los parametros debe enviar secretname")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("error al leer el Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)

	return event, err
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
