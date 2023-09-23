package main

import (
	"fmt"
	"net/http"

	"github.com/Kyohei-takiyama/GoRestApi/controller"
	router "github.com/Kyohei-takiyama/GoRestApi/http"
	"github.com/Kyohei-takiyama/GoRestApi/repository"
	"github.com/Kyohei-takiyama/GoRestApi/service"
)

var (
	postRepo repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewChaiRouter()
)

const port string = "8000"


func main(){
	httpRouter.GET("/" , func (resp http.ResponseWriter , req *http.Request)  {
		fmt.Fprintln(resp , "Up and running...")
	})
	httpRouter.GET("/posts" , postController.GetPosts)
	httpRouter.POST("/posts" , postController.AddPost)
	httpRouter.SERVE("8000")
}