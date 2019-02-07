package dao

import "central/model"

// DataDao interfaces is used to know all the allowed methods
type DataDao interface {
	Get(id string) (*model.Data, error)

	Upsert(data *model.Data) (*model.Data, error)

	GetAll() ([]model.Data, error)

	AddServer(server *model.Server) ([]model.Server, error)

	DeleteAll()
}
