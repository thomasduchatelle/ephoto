package bootstrap

import (
	cmd2 "github.com/thomasduchatelle/ephoto/cmd/fphoto/cmd"
	config2 "github.com/thomasduchatelle/ephoto/cmd/fphoto/config"
	"github.com/thomasduchatelle/ephoto/pkg/acl/aclcore"
	"github.com/thomasduchatelle/ephoto/pkg/acl/acldynamodb"
	"github.com/thomasduchatelle/ephoto/pkg/acl/catalogacl"
	"github.com/thomasduchatelle/ephoto/pkg/catalog"
)

type catalogPort struct {
}

func (c catalogPort) FindAlbum(owner, folderName string) (*catalog.Album, error) {
	return catalog.FindAlbum(owner, folderName)
}

func init() {
	config2.Listen(func(cfg config2.Config) {
		repository := acldynamodb.Must(acldynamodb.New(cfg.GetAWSSession(), cfg.GetString(config2.CatalogDynamodbTable), false))
		createUser := &aclcore.CreateUser{
			ScopesReader: repository,
			ScopeWriter:  repository,
		}
		cmd2.CreateUserCase = createUser.CreateUser

		sharedAlbum := &catalogacl.ShareAlbumCase{
			ScopeWriter: repository,
			CatalogPort: new(catalogPort),
		}
		cmd2.ShareAlbumCase = sharedAlbum.ShareAlbumWith

		unSharedAlbum := &catalogacl.UnShareAlbumCase{
			RevokeScopeRepository: repository,
		}
		cmd2.UnShareAlbumCase = unSharedAlbum.StopSharingAlbum
	})
}
