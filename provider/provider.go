package provider

import (
	"context"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/database"
	"github.com/christiansoetanto/better-servant-of-servus-dei/provider/dbms"
	"sync"
)

type provider struct {
	Dbms      *dbms.Obj
	AppConfig config.AppConfig
}

type Resource struct {
	AppConfig config.AppConfig
	Database  *database.Obj
}

type Provider interface {
	HelloWorld(ctx context.Context) error
}

var obj Provider
var once sync.Once

//GetProvider get provider client
func GetProvider(resource *Resource) Provider {
	once.Do(func() {
		obj = &provider{
			//Dbms:      dbms.GetDbmsClient(resource.Database),
			AppConfig: resource.AppConfig,
		}
	})
	return obj
}
