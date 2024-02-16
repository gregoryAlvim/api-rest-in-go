package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gregoryAlvim/api-rest-in-go/models"
	"github.com/gregoryAlvim/api-rest-in-go/services"
)

type BoxController struct{}

type CreateBody struct {
	models.Box
}

type CreateResponse struct {
	Message string `json:"message"`
}

func (BoxController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var body CreateBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil || len(body.Box.Name) == 0 || len(body.Box.Content) == 0 {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	boxService := services.BoxService{}

	boxService.Create(&body.Box)

	res := CreateResponse{Message: "Caixa criada com sucesso!"}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}

func (BoxController) Find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	idParam := chi.URLParam(r, "id")

	boxService := services.BoxService{}
	box, err := boxService.Find(idParam)

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusNotExtended)
	} else {
		json.NewEncoder(w).Encode(box)
	}
}

func (BoxController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	boxService := services.BoxService{}
	boxes, err := boxService.FindAll()

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusNotExtended)
	} else {
		json.NewEncoder(w).Encode(boxes)
	}
}

func (BoxController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	idParam := chi.URLParam(r, "id")

	var body CreateBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil || len(body.Box.Name) == 0 || len(body.Box.Content) == 0 {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	boxService := services.BoxService{}
	err = boxService.Update(idParam, &body.Box)

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	res := CreateResponse{Message: "Caixa atualizada com sucesso!"}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func (BoxController) Remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	idParam := chi.URLParam(r, "id")

	boxService := services.BoxService{}
	err := boxService.Remove(idParam)

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	res := CreateResponse{Message: "Caixa deletada com sucesso!"}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
