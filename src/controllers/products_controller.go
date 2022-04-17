package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

type ProductsController struct {
	Controller
	UserServices *services.UserService
	ProductServices *services.ProductServices
}

func (ctx *ProductsController) CreateRoutes(router *mux.Router) {
	router.Handle("/products", utils.UseAdminOnly(ctx.Tokens.JwtKey, ctx.UserServices, http.HandlerFunc(ctx.create_product))).Methods("POST")
}

func (ctx *ProductsController) create_product(w http.ResponseWriter, r *http.Request) {
	var create_product types.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&create_product); err != nil {
		utils.FailResponse(w, "could not parse body data")
		return
	}
	new_product := ctx.ProductServices.Create(&create_product)
	utils.DataResponse(w, &new_product)
}