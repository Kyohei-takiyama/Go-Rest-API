package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Kyohei-takiyama/GoRestApi/entity"
)


type repo struct {}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId = "pragmatic-reviews-84315"
	collectionName = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post , error){
	ctx := context.Background()
	client , err := firestore.NewClient(ctx , projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client : %v" , err)
		return nil , err
	}

	defer client.Close()
	_ , _ , err = client.Collection(collectionName).Add(ctx , map[string]interface{}{
		"ID": post.ID,
		"Title": post.Title,
		"Text": post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post : %v" , err)
		return nil , err
	}

	return post , nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if doc == nil {
			break
		}
		log.Println(doc.Data())

		if err != nil {
			log.Printf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}