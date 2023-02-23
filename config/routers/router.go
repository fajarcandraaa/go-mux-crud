package routers

import (
	"github.com/fajarcandraaa/go-mux-crud/handler"
	"github.com/fajarcandraaa/go-mux-crud/repositories"
	"github.com/fajarcandraaa/go-mux-crud/src/user"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== USER ===
	s := user.NewService(r)
	h := handler.NewUserHandler(s)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== USER ===
	se.Router.HandleFunc("/user", h.RegisterNewUser).Methods("POST")
	//==========================================================

}
