package services

import (
	"context"
	"myapp/graph/model"
	"time"
)

func (s *Service) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post

	if err := s.DB.Model(&posts).Find(&posts).Error; err != nil {
		panic(err)
	}

	return posts, nil
}

func (s *Service) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := &model.Post{
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.DB.Model(&post).Create(&post).Omit("updated_at").Error; err != nil {
		panic(err)
	}

	return post, nil
}
