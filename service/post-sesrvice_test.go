package service

import (
	"testing"

	"github.com/Kyohei-takiyama/GoRestApi/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct{
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post , error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *mockRepository) FindAll() ([]entity.Post , error){
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	post := entity.Post{
		ID: 1,
		Title: "A",
		Text: "B",
	}

	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post} , nil)

	testService := NewPostService(mockRepo)

	result , _ :=testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t , result[0].ID , post.ID)
	assert.Equal(t , result[0].Title , post.Title)
	assert.Equal(t , result[0].Text , post.Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)

	post := entity.Post{
		ID: 1,
		Title: "A",
		Text: "B",
	}

	// Setup expectations
	mockRepo.On("Save").Return(&post , nil)

	testService := NewPostService(mockRepo)

	result , _ :=testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Equal(t , result.ID , post.ID)
	assert.Equal(t , result.Title , post.Title)
	assert.Equal(t , result.Text , post.Text)
}

func TestValidateEmptyPost(t *testing.T){
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t , err)

	assert.Equal(t, "The post is empty" , err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T){
	post := entity.Post{ID: 1,Title: "",Text: "test"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t , err)

	assert.Equal(t , "The post title is empty" , err.Error())
}

func TestValidateEmptyPostText(t *testing.T){
	post := entity.Post{ID: 1,Title: "test",Text: ""}
	testService := NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t , err)

	assert.Equal(t , "The post text is empty" , err.Error())
}