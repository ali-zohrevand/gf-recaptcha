package routers

import (
	"github.com/ali-zohrevand/gf-recapcha/release/Handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", AddContext(Handler.Index))

	router.HandleFunc("/ver3", AddContext(Handler.Ver3))
	router.HandleFunc("/puzzle", AddContext(Handler.Puzzle))
	router.HandleFunc("/v2", AddContext(Handler.V2))

	return router

}
