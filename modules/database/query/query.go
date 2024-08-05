package query

import (
	"github.com/LocTr/reverse_uber_be/modules/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppDB struct {
	App config.AppTools
	DB  mongo.Client
}

func NewAppDB(app config.AppTools, db mongo.Client) database.DBRepo {
	return &AppDB{
		App: config.AppTools{},
		DB:  mongo.Client{},
	}
}

func (g *GoAppDB) InsertMenu() {

}
