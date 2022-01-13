package httpserver

import (
	"net/http"

	"github.com/emicklei/go-restful"

	"nanto.io/application-auto-scaling-service/pkg/common/utils"
	"nanto.io/application-auto-scaling-service/pkg/search/ascontroller"
	"nanto.io/application-auto-scaling-service/pkg/search/rpcclient"
)

type CreateFleetReq struct {
	FleetName string `json:"fleet_name"`
}

type CreateFleetResp struct {
	Result string `json:"result"`
}

func CreateFleet(req *restful.Request, resp *restful.Response) {
	logger.Trace("=== CreateFleet called")
	createFleetReq := &CreateFleetReq{}
	if err := req.ReadEntity(createFleetReq); err != nil {
		logger.Errorf("Request read entity err: %v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusBadRequest, utils.RequestBodyParamInvalid(err.Error()))
		return
	}
	// todo 参数校验
	if createFleetReq.FleetName == "" {
		logger.Errorf("FleetName invalid")
		return
	}
	logger.Infof("CreateFleet req: %+v", createFleetReq)

	// AS创建伸缩组fleet
	fleetId, err := ascontroller.GetAsController().CreateFleet(createFleetReq.FleetName)
	if err != nil {
		logger.Errorf("AsController create fleet err: %+v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalError(err))
		return
	}

	// 通知gateway
	if err = rpcclient.GetAppGwReporter().AddFleet(fleetId, createFleetReq.FleetName); err != nil {
		logger.Errorf("GetAppGwReporter add fleet err: %+v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalError(err))
		return
	}

	resp.WriteHeaderAndEntity(http.StatusOK, CreateFleetResp{Result: "ok"})
}

// todo
func DeleteFleet(req *restful.Request, resp *restful.Response) {
	logger.Trace("=== CreateFleet called")
	createFleetReq := &CreateFleetReq{}
	if err := req.ReadEntity(createFleetReq); err != nil {
		logger.Errorf("Request read entity err: %v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusBadRequest, utils.RequestBodyParamInvalid(err.Error()))
		return
	}
	// todo 参数校验
	if createFleetReq.FleetName == "" {
		logger.Errorf("FleetName invalid")
		return
	}
	logger.Infof("CreateFleet req: %+v", createFleetReq)

	// AS创建伸缩组fleet
	fleetId, err := ascontroller.GetAsController().CreateFleet(createFleetReq.FleetName)
	if err != nil {
		logger.Errorf("AsController create fleet err: %+v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalError(err))
		return
	}

	// 通知gateway
	if err = rpcclient.GetAppGwReporter().AddFleet(fleetId, createFleetReq.FleetName); err != nil {
		logger.Errorf("GetAppGwReporter add fleet err: %+v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalError(err))
		return
	}

	resp.WriteHeaderAndEntity(http.StatusOK, CreateFleetResp{Result: "ok"})
}
