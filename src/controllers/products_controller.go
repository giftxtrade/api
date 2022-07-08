package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	var errors []string
	q := r.URL.Query()
	search := strings.TrimSpace(q.Get("search"))
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	minPrice, err := strconv.ParseFloat(q.Get("minPrice"), 32)
	if err != nil || minPrice < 0 {
		minPrice = 0
	}
	maxPrice, err := strconv.ParseFloat(q.Get("maxPrice"), 32)
	if err != nil || maxPrice < minPrice {
		maxPrice = 5000
	}
	sort := strings.TrimSpace(q.Get("sort"))

	products, err := ctx.
		ProductService.
		Search(search, limit, page, float32(minPrice), float32(maxPrice), sort)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
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