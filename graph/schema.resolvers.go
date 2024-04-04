package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
)

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context) (*model.PostMutation, error) {
	return &model.PostMutation{}, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context) (*model.PostQuery, error) {
	return &model.PostQuery{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
