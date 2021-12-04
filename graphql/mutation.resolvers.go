package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/ShinsakuYagi/gql-upload-sample/graphql/generated"
	"github.com/ShinsakuYagi/gql-upload-sample/model"
)

func (r *mutationResolver) UploadFile(ctx context.Context, input model.UploadInput) (*model.UploadPayload, error) {
	const (
		fileDirName string      = "file"
		filePerm    fs.FileMode = 0o666
	)

	// io.Reader provided in input.File.File, so it is easy to throw it to S3 or something.
	file, err := io.ReadAll(input.Data.File)
	if err != nil {
		return nil, err
	}

	filePath := path.Join(fileDirName, input.Data.Filename)
	if _, err := os.Create(filePath); err != nil {
		return nil, err
	}
	if err := os.WriteFile(filePath, file, filePerm); err != nil {
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &model.UploadPayload{
		Path: path.Join(pwd, filePath),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
