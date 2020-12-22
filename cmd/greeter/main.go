package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/timdrysdale/swagger-example-server/gen/restapi"
	"github.com/timdrysdale/swagger-example-server/gen/restapi/operations"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

//https://shashankvivek-7.medium.com/go-swagger-user-authentication-securing-api-using-jwt-part-2-c80fdc1a020a
/*
func validateToken(name string, authenticate security.ScopedTokenAuthentication) runtime.Authenticator {

	return

}
*/

// context values - requestScopedValues
// https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte("jwtsecret"), nil
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if token.Valid {
		return claims["user"].(string), nil
	}
	return nil, errors.New("invalid token")
}

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	//create new service API
	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	//parse flags
	flag.Parse()

	// set the port this service will run on
	server.Port = *portFlag

	// set the Authorizer
	api.BearerAuth = ValidateHeader

	// set the Handler
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(
		func(params operations.GetGreetingParams, principal interface{}) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}

			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})

	//serve API

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
