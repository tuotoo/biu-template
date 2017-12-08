package main

import (
	"os"

	"github.com/emicklei/go-restful"
	"github.com/rs/zerolog"
	"github.com/tuotoo/biu"
	_ "github.com/tuotoo/biu-template/routes"
)

func main() {
	biu.SetLoggerOutput(zerolog.ConsoleWriter{Out: os.Stderr})
	restful.Filter(biu.LogFilter())
	restful.Filter(restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}.Filter)
	restful.Filter(restful.DefaultContainer.OPTIONSFilter)

	swaggerService := biu.NewSwaggerService(biu.SwaggerInfo{
		Title:        "UserService",
		Description:  "Resource for managing Users",
		ContactName:  "Jqs7",
		ContactEmail: "7@jqs7.com",
		ContactURL:   "https://jqs7.com",
		LicenseName:  "MIT",
		LicenseURL:   "http://mit.org",
		Version:      "1.0.0",
	})
	restful.Add(swaggerService)

	biu.Run(":"+os.Getenv("PORT"), nil)
}
