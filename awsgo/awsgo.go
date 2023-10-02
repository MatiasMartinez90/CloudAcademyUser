package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context //
var Cfg aws.Config      // Tipo de dato de aws config, una estructura para manejar las configuraciones de inicio de sesion
var err error

func InicializoAWS() {
	Ctx = context.TODO() //TODO retorna un no nullo, que lo haga sin ningun tipo de configuracion (un contexto vacio)

	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1")) // Trae todas las configuraciones que tenga AWS en nuestra cuenta, el Ctx(contexto se lo manda vacio), Y le especifica la region

	if err != nil {
		panic("Error al cargar la config .aws/config " + err.Error())
	}
}
