package main

import (
	"gofr.dev/pkg/gofr"
	"Gofr_Zopsmart/datastore"
	"Gofr_Zopsmart/handler"
)

func main() {
	app := gofr.New()
	s := datastore.New()
	h := handler.New(*s)
	app.GET("/Passbook/{id}", h.GetByID)
	app.POST("/Passbook", h.Create)
	app.PUT("/Passbook/{id}", h.Update)
	app.DELETE("/Passbook/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
	app.Start()
}
