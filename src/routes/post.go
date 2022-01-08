package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (app *AppBase) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := new([]types.Post)
	app.DB.Table("posts").Find(&posts)
	utils.JsonResponse(w, types.Result{Data: &posts})
}

func (app *AppBase) CreatePost(w http.ResponseWriter, r *http.Request) {
	data := types.CreatePost{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	
	// Create post
	slug := strings.ReplaceAll(data.Title, " ", "-")
	slug = strings.ToLower(slug)
	new_post := types.Post{
		Title: data.Title,
		Slug: slug,
		Summary: data.Summary,
		Content: data.Content,
	}
	app.DB.Create(&new_post)
	utils.JsonResponse(w, new_post)
}

func (app *AppBase) GetPostFromId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	post := types.Post{}
	app.DB.Table("posts").Find(&post, "id = ?", id)
	if post.ID == uuid.Nil {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find post"})
		return
	}
	utils.JsonResponse(w, &post)
}