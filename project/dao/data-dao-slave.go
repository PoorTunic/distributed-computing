package dao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/model"
	"strconv"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// GetFromMaster a data in the storage
func GetFromMaster(id string) (*model.Data, error) {
	var dataMaster model.Data
	for _, serv := range servers {
		url := fmt.Sprintf(serv+"%s", id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Error("NewRequest: ", err)
			return nil, nil
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Error("Do: ", err)
			return nil, nil
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&dataMaster); err != nil {
			log.Println(err)
		}
		if dataMaster.ID == "" && dataMaster.Data == "" {
			return nil, nil
		}
	}
	return &dataMaster, nil
}

// UpsertToMaster method
func UpsertToMaster(data *model.Data) (*model.Data, error) {
	url := fmt.Sprintf("http://localhost:8020/%+v", data)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("NewRequest: ", err)
		return nil, nil
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Do: ", err)
		return nil, nil
	}
	defer resp.Body.Close()
	var dataMaster model.Data
	if err := json.NewDecoder(resp.Body).Decode(&dataMaster); err != nil {
		log.Println(err)
	}
	if data.ID == "" {
		tmpID, _ := uuid.NewV4()
		data.ID = tmpID.String()
	}
	return data, nil
}

// GetAllFromServers a data in the storage
func GetAllFromServers() ([]model.Data, error) {
	var datas []model.Data
	for _, serv := range servers {
		log.Info("Server.GetAllFromServers() = Getting data from " + serv)
		resp, err := http.Get(serv)
		if err != nil {
			log.Error("NewRequest: ", err)
			return nil, nil
		}
		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&datas)

		log.Info("Fetched datas : " + strconv.Itoa(len(datas)) + " from " + serv)

		if len(datas) != 0 {
			log.Info("Adding datas from " + serv)
			for _, data := range datas {
				data.ID = "?"
				data.Source = serv
				datas = append(datas, data)
			}
		}
	}
	log.Info("Total of datas from another servers : " + strconv.Itoa(len(datas)))
	return datas, nil
}

func get(id string) bool {
	for _, serv := range servers {
		rqServ := fmt.Sprintf(serv+"%s", id)
		req, err := http.NewRequest("GET", rqServ, nil)
		if err != nil {
			log.Error("NewRequest: ", err)
			return false
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Error("Do: ", err)
			return false
		}
		defer resp.Body.Close()

		return true
	}
	return true
}
