package dao

import (
	"central/model"
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
	redis "gopkg.in/redis.v5"
)

var _ DataDao = (*DataDAORedis)(nil)

type DataDAORedis struct {
	redisCli *redis.Client
}

func NewDataRedis(redisCli *redis.Client) DataDao {
	return &DataDAORedis{
		redisCli: redisCli,
	}
}

// Get a data in the storage
func (dao *DataDAORedis) Get(id string) (*model.Data, error) {
	resData, err := dao.redisCli.Get(id).Result()

	if err != nil {
		return nil, err
	} else if len(resData) == 0 {
		return nil, errors.New("REDIS -> " + id + " does not exist")
	}

	data := model.Data{}

	err = json.Unmarshal([]byte(resData), &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Upsert modify or create a data
func (dao *DataDAORedis) Upsert(data *model.Data) (*model.Data, error) {
	if data.ID == "" {
		tmp, _ := uuid.NewV4()
		data.ID = tmp.String()
	}

	resData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	status := dao.redisCli.Set(data.ID, string(resData), 0)

	if status.Err() != nil {
		return nil, status.Err()
	}

	return data, nil
}

// GetAll datas in the storage
func (dao *DataDAORedis) GetAll() ([]model.Data, error) {
	var datas []model.Data

	// Collect all datas identifiers
	keys := dao.redisCli.Keys("*").Val()
	if len(keys) == 0 {
		return nil, errors.New("no datas")
	}

	for i := 0; i < len(keys); i++ {
		// Collect data by identifier
		data, err := dao.Get(keys[i])
		if err != nil {
			return nil, err
		}

		datas = append(datas, *data)
	}

	return datas, nil
}

// GetAll datas in the storage
func (dao *DataDAORedis) DeleteAll() {
	dao.redisCli.FlushAll()
}

// AddServer add a server
func (dao *DataDAORedis) AddServer(server *model.Server) ([]model.Server, error) {
	return nil, nil
}
