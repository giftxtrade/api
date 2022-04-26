package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductsController struct {
	Controller
	UserServices *services.UserService
	ProductServices *services.ProductServices
}

func (ctx *ProductsController) CreateRoutes(router *mux.Router, path string) {
	router.Handle(path, ctx.Controller.UseJwtAuth(http.HandlerFunc(ctx.find_all_products))).Methods("GET")
	router.Handle(path, ctx.Controller.UseAdminOnly(http.HandlerFunc(ctx.create_product))).Methods("POST")
	router.Handle(path + "/{id}", ctx.Controller.UseJwtAuth(http.HandlerFunc(ctx.find_product))).Methods("GET")
}

func (ctx *ProductsController) find_all_products(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, types.Response{Message: "all products"})
}

func (ctx *ProductsController) create_product(w http.ResponseWriter, r *http.Request) {
	var create_product types.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&create_product); err != nil {
		utils.FailResponse(w, "could not parse body data")
		return
	}

	// validation
	var errors []string
	if create_product.ProductKey == "" {
		errors = append(errors, "product key is required")
	}
	if create_product.Category == "" {
		errors = append(errors, "category name is required")
	}
	if len(errors) > 0 {
		utils.FailResponse(w, &errors)
		return
	}

	new_product := ctx.ProductServices.CreateOrUpdate(&create_product)
	utils.DataResponse(w, new_product)
}

func (ctx *ProductsController) find_product(w http.ResponseWriter, r *http.Request) {
	query_params := mux.Vars(r)
	id := query_params["id"]
	product := ctx.ProductServices.Find(id)
	if product.ID == uuid.Nil {
		utils.FailResponse(w, "product not found")
		return
	}
	utils.DataResponse(w, product)
}