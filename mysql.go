package mysqladapter

import (
	"io/fs"

	"github.com/debyten/database/dbconf"
	gormdb "github.com/debyten/gorm-adapter"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	msq "gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func NewConfiguration[V string | uint | uint32 | uint64 | int | int32 | int64](defaultDB string, idGenerator gormdb.IDGeneratorRegistry[V], additionalParams ...map[string]string) *gormdb.Configuration {
	ds := dbconf.NewMysql(defaultDB, additionalParams...)
	prov := NewProvider(ds)
	return gormdb.NewConfiguration(ds, prov, idGenerator)
}

func WithMigrations[V string | uint | uint32 | uint64 | int | int32 | int64](defaultDB string, idGenerator gormdb.IDGeneratorRegistry[V], migrations fs.FS, execMigrations bool, additionalParams ...map[string]string) *gormdb.Configuration {
	cfg := NewConfiguration(defaultDB, idGenerator, additionalParams...)
	return cfg.MustSetMigrations(migrations, execMigrations)
}

// NewProvider returns a function that can be used to create a database connection for MySQL.
// The function takes a Datasource configuration as an argument and returns a database connection of
// type *gorm.DB and an error if any.
func NewProvider(cfg dbconf.Datasource) gormdb.ConnProvider {
	return func() (*gorm.DB, error) {
		return gorm.Open(msq.Open(cfg.ConnURL()), &gorm.Config{})
	}
}
