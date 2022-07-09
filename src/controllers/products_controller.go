package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

type ProductsController struct {
	Controller
	UserService *services.UserService
	ProductService *services.ProductService
}

func (ctx *ProductsController) CreateRoutes(router *mux.Router, path string) {
	router.Handle(path, ctx.Controller.UseJwtAuth(http.HandlerFunc(ctx.find_all_products))).Methods("GET")
	router.Handle(path, ctx.Controller.UseAdminOnly(http.HandlerFunc(ctx.create_product))).Methods("POST")
	router.Handle(path + "/{id}", ctx.Controller.UseJwtAuth(http.HandlerFunc(ctx.find_product))).Methods("GET")
}

func (ctx *ProductsController) find_all_products(w http.ResponseWriter, r *http.Request) {
	var filter types.ProductFilter
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		utils.FailResponse(w, "could not parse body data")
		return
	}
	
	products, err := ctx.
		ProductService.
		Search(filter)
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		utils.FailResponse(w, errors)
		return
	}

	utils.JsonResponse(w, types.Result{
		Data: products,
	})
}

func (ctx *ProductsController) create_product(w http.ResponseWriter, r *http.Request) {
	var create_product types.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&create_product); err != nil {
		utils.FailResponse(w, "could not parse body data")
		return
	}

	var new_product types.Product
	_, err := ctx.ProductService.CreateOrUpdate(&create_product, &new_product)
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		utils.FailResponse(w, errors)
		return
	}
	utils.DataResponse(w, new_product)
}

func (ctx *ProductsController) find_product(w http.ResponseWriter, r *http.Request) {
	query_params := mux.Vars(r)
	id := query_params["id"]
	var product types.Product
	if ctx.ProductService.Find(id, &product) != nil {
		utils.FailResponse(w, "product not found")
		return
	}
	utils.DataResponse(w, product)
}