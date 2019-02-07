package dao

import (
	"errors"
	"project/model"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

var _ DataDao = (*DataDaoMock)(nil)

var servers = make([]string, 0)

// DataDaoMock stores all the datas in a mock
type DataDaoMock struct {
	storage map[string]*model.Data
}

// NewDataDAOMock get the storage
func NewDataDAOMock() DataDao {
	daoMock := &DataDaoMock{
		storage: make(map[string]*model.Data),
	}
	return daoMock
}

// Get a data in the storage
func (dao *DataDaoMock) Get(id string) (*model.Data, error) {
	data, ok := dao.storage[id]

	if !ok {
		return nil, errors.New("Data not found with id -> " + id)
	}

	return data, nil
}

// Upsert modify or create a data
func (dao *DataDaoMock) Upsert(data *model.Data) (*model.Data, error) {
	if data.ID == "" {
		tmpID, _ := uuid.NewV4()
		data.ID = tmpID.String()
	}
	data.LastUse = time.Now().Add(5 * time.Second)
	dao.storage[data.ID] = data
	return data, nil
}

// GetAll tasks in the storage
func (dao *DataDaoMock) GetAll() ([]model.Data, error) {
	log.Info("Server.GetAll() = Mock")
	var datas []model.Data
	for dataID := range dao.storage {
		data := dao.storage[dataID]
		datas = append(datas, *data)
	}
	return datas, nil
}

// AddServer add a server
func (dao *DataDaoMock) AddServer(server *model.Server) ([]model.Server, error) {
	servers = append(servers, server.IP)
	log.Info("New server IP list : ")
	for _, ip := range servers {
		log.Info(ip)
	}

	return nil, nil
}
