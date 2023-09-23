package controller

import (
	"encoding/json"
	// "errors"
	"net/http"

	"github.com/Kyohei-takiyama/GoRestApi/entity"
	"github.com/Kyohei-takiyama/GoRestApi/errors"
	"github.com/Kyohei-takiyama/GoRestApi/service"
)

type controller struct {}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter , req *http.Request)
	AddPost(resp http.ResponseWriter , req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter , req *http.Request) {
	resp.Header().Set("Content-Type" , "application/json")
	posts , err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		// return error message
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter , req *http.Request) {
	resp.Header().Set("Content-Type" , "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		// return error message

		return
	}
	errValidate := postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		// return error message
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errValidate.Error()})
		return
	}

	result , errCreate := postService.Create(&post)
	if errCreate != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		// return error message
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}