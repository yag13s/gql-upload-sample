package main

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ShinsakuYagi/gql-upload-sample/model"
	"github.com/gin-gonic/gin"
	"github.com/k1LoW/runn"
)

func Test_runn(t *testing.T) {
	t.Run("test runn with gin and GraphQL", func(t *testing.T) {
		if err := os.Setenv("ENVCODE", "test"); err != nil {
			panic(err)
		}
		model.Version = "testing"

		r := gin.Default()
		defineRoutes(r)

		ts := httptest.NewServer(r.Handler())
		t.Cleanup(func() {
			ts.Close()
		})

		opts := []runn.Option{
			runn.T(t),
			runn.Runner("req", ts.URL),
		}

		o, err := runn.Load("../testutil/books/**/*.yml", opts...)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()
		if err := o.RunN(ctx); err != nil {
			t.Fatal(err)
		}
	})
}
