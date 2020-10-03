package main

import (
	"doc-server/db/conf"
	"doc-server/interactor"
	serverMiddleware "doc-server/presenter/middleware"
	"doc-server/presenter/router"
	"fmt"
	"net/http"
)

var initMessage = `
	starting doc-server
	*******************
	http://localhost:8080/api/v1/docs
	*******************
`

func main() {
	conn, err := conf.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println(initMessage)
	i := interactor.NewInteractor(conn)
	r := i.NewRepository()
	u := i.NewUsecase(r)
	// NewHandlerにNewRepositoryを渡す
	h := i.NewHandler(u)
	middleware := serverMiddleware.NewServerMiddleware()
	server := router.NewRouter()
	server.Router(h, middleware)
	http.ListenAndServe(":8080", server.Route)
}


