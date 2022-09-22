package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"os"

	"github.com/ShinsakuYagi/gql-upload-sample/graphql/generated"
	"github.com/ShinsakuYagi/gql-upload-sample/model"
)

// Example is the resolver for the example field.
func (r *queryResolver) Example(ctx context.Context) (*model.Example, error) {
	return &model.Example{
		ProjectName: "gql-upload-sample",
		Envcode:     os.Getenv("ENVCODE"),
		Version:     model.Version,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }