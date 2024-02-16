package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gregoryAlvim/api-rest-in-go/controllers"
)

const routeBase = "/api/box"

func createBoxRoutes(router *chi.Mux) {
	boxController := controllers.BoxController{}

	addMethodPost(router, &boxController)
	addMethodsGet(router, &boxController)
	addMethodUpdate(router, &boxController)
	addMethodDelete(router, &boxController)
}

func addMethodPost(router *chi.Mux, boxController *controllers.BoxController) {
	router.Post(routeBase+"/", func(w http.ResponseWriter, r *http.Request) {
		boxController.Create(w, r)
	})
}

func addMethodsGet(router *chi.Mux, boxController *controllers.BoxController) {
	router.Get(routeBase+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		boxController.Find(w, r)
	})

	router.Get(routeBase+"/", func(w http.ResponseWriter, r *http.Request) {
		boxController.FindAll(w, r)
	})
}

func addMethodUpdate(router *chi.Mux, boxController *controllers.BoxController) {
	router.Put(routeBase+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		boxController.Update(w, r)
	})
}

func addMethodDelete(router *chi.Mux, boxController *controllers.BoxController) {
	router.Delete(routeBase+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		boxController.Remove(w, r)
	})
}
