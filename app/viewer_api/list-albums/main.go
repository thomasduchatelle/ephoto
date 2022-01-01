package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/thomasduchatelle/dphoto/app/viewer_api/common"
	"github.com/thomasduchatelle/dphoto/domain/catalog"
)

func Handler(ctx context.Context) (common.Response, error) {
	if err := common.ConnectCatalog("tomdush@gmail.com"); err != nil {
		return common.InternalError(err)
	}

	albums, err := catalog.FindAllAlbumsWithStats()
	if err != nil {
		return common.InternalError(errors.Wrapf(err, "failed to fetch albums"))
	}

	return common.Ok(albums)
}

func main() {
	lambda.Start(Handler)
}
