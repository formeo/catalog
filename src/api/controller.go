package api

type Controller struct {
	storage *Storage
}

func NewController(storage *Storage) *Controller {
	return &Controller{storage: storage}
}
