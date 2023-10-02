package models

// La estructura que usa secretmanager, que nos devuelve 6 valores en formato json
type SecretRDSJson struct {
	Username            string `json:"username"` // Alt izq + 96 // Se usa para pasar a minuscula
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

// Para el registro de usuarios --> Los eventos que recibimos de cognito vamos a recibir parametros, y solo nos interesan estos 2:
type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
