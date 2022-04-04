package app

import (
	"net/http"
	"strconv"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

func (app *AppBase) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.FailResponse(w, "could not parse POST data")
		return
	}
	
	body := r.PostForm
	title := body.Get("title")
	product_key := body.Get("product_key")
	rating, rating_err := strconv.ParseFloat(body.Get("rating"), 32)
	if rating_err != nil {
		utils.FailResponse(w, "could not parse product rating")
		return
	}
	price, price_err := strconv.ParseFloat(body.Get("price"), 32)
	if price_err != nil {
		utils.FailResponse(w, "could not parse product price")
		return
	}
	original_url := body.Get("original_url")
	website_origin := body.Get("website_origin")
	total_reviews, total_reviews_err := strconv.Atoi(body.Get("total_reviews"))
	if total_reviews_err != nil {
		utils.FailResponse(w, "could not parse total_reviews")
		return
	}
	new_product := types.Product{
		Title: title,
		ProductKey: product_key,
		Rating: float32(rating),
		Price: float32(price),
		OriginalUrl: original_url,
		WebsiteOrigin: website_origin,
		TotalReviews: total_reviews,
	}
	app.DB.Table("products").Create(&new_product)
	utils.DataResponse(w, &new_product)
}