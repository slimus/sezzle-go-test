package main

import (
	"fmt"
	"os"

	"github.com/slimus/calc/src/calc"
	"github.com/slimus/calc/src/storage"
	"github.com/slimus/calc/src/ws"
)

func main() {
	storage := storage.NewMemoryStorage(10)

	wsServer := ws.NewWS()
	c := calc.NewCalc()
	calcHandler := ws.New(storage, c)
	wsServer.RegisterHandlers(calcHandler.Handle())

	port := os.Getenv("CALC_WS_PORT")
	if port == "" {
		port = "4321" // default port
	}

	fmt.Println(wsServer.Run(port))
}
