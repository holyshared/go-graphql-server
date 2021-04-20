//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/holyshared/go-graphql-server/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DataStore *gorm.DB
	todos     []*model.Todo
}
