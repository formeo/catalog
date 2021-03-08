package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	View *View
}

func NewApp(view *View) *App {
	return &App{view}
}

func (app *App) NewRouter() *mux.Router {

	rootRouter := mux.NewRouter()
	apiRouter := rootRouter.PathPrefix("/api/v1/").Subrouter()
	apiRouter.Use(UserHandler)

	apiRouter.HandleFunc("/categories", app.View.GetCategories()).Methods("GET")
	apiRouter.HandleFunc("/categories/tree", app.View.GetCategoriesTree()).Methods("GET")
	apiRouter.HandleFunc("/products", app.View.GetProducts()).Methods("GET")
	apiRouter.HandleFunc("/products/{sku}/view", app.View.GetProduct()).Methods("GET")
	apiRouter.HandleFunc("/products/{sku}/for_cart", app.View.GetProductCart()).Methods("GET")
	apiRouter.HandleFunc("/hierarchy/products/", app.View.Hierarchy()).Methods("GET")
	apiRouter.HandleFunc("/products/breadcrumbs/", app.View.GetProductsBreadcrumbs()).Methods("GET")

	return rootRouter
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}
