package datasources

import (
	"github.com/avinashb98/blaast/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoORM(mongoConfig *config.Mongo) error {
	err := mgm.SetDefaultConfig(nil, mongoConfig.BlaastDB, options.Client().ApplyURI(mongoConfig.HostURI))
	return err
}
