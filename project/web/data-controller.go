package web

import (
	"net/http"
	"project/dao"
	"project/model"

	log "github.com/sirupsen/logrus"
)

// DataController struct
type DataController struct {
	dataDao  dao.DataDao
	Routes   []Route
	Prefix   string
	ServerIP string
}

const (
	prefixData = ""
)

// NewDataController is uset to get the dao controller
func NewDataController(dataDao dao.DataDao) *DataController {
	controller := DataController{
		dataDao: dataDao,
		Prefix:  prefixData,
	}

	var routes []Route

	routes = append(routes, Route{
		Name:        "Get One Data",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: controller.GetData,
	})

	routes = append(routes, Route{
		Name:        "Get All Data",
		Method:      http.MethodGet,
		Pattern:     "/",
		HandlerFunc: controller.GetAllData,
	})

	routes = append(routes, Route{
		Name:        "Create or Update a Data",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: controller.UpdateData,
	})

	routes = append(routes, Route{
		Name:        "Add IP address to server list",
		Method:      http.MethodPost,
		Pattern:     "/servers",
		HandlerFunc: controller.AddServer,
	})

	controller.Routes = routes

	return &controller
}

// GetAllData method
func (ctrl *DataController) GetAllData(w http.ResponseWriter, r *http.Request) {
	log.Info("Server.GetAllData()")
	datas, err := ctrl.dataDao.GetAll()
	if err != nil {
		log.Error("Server.GetAllData() = " + err.Error())
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendJSONOk(w, datas)
}

// GetData method
func (ctrl *DataController) GetData(w http.ResponseWriter, r *http.Request) {
	idData := ParamAsString("id", r)

	resData, err := ctrl.dataDao.Get(idData)

	if err != nil {
		log.Error(err.Error())
	}
	SendJSONOk(w, resData)
}

// UpdateData method
func (ctrl *DataController) UpdateData(w http.ResponseWriter, r *http.Request) {
	idData := ParamAsString("id", r)
	var nData = model.Data{}
	err := GetJSONContent(&nData, r)
	nData.ID = idData
	if err != nil {
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	nData.Source = ctrl.ServerIP
	resData, err := ctrl.dataDao.Upsert(&nData)

	SendJSONOk(w, resData)
}

// AddServer method
func (ctrl *DataController) AddServer(w http.ResponseWriter, r *http.Request) {
	log.Info("New server connection received")
	var nServer = model.Server{}
	err := GetJSONContent(&nServer, r)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Adding new server to the servers list :" + nServer.IP)
	ctrl.dataDao.AddServer(&nServer)
}
