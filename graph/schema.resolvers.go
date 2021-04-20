package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	domain "github.com/holyshared/go-graphql-server/model"

	"github.com/holyshared/go-graphql-server/graph/generated"
	"github.com/holyshared/go-graphql-server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	ormTodo := &domain.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	r.DataStore.Create(ormTodo)

	todo := &model.Todo{
		Text:   ormTodo.Text,
		ID:     ormTodo.ID,
		UserID: ormTodo.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []domain.Todo
	dtoTodos := []*model.Todo{}

	r.DataStore.Find(&todos)

	for _, t := range todos {
		dtoTodos = append(dtoTodos, &model.Todo{
			ID:     t.ID,
			Text:   t.Text,
			Done:   t.Done,
			UserID: t.UserID,
		})
	}
	r.todos = dtoTodos

	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
