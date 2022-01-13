package httpserver

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"

	"nanto.io/application-auto-scaling-service/pkg/common/utils"
	"nanto.io/application-auto-scaling-service/pkg/search/ascontroller"
)

const (
	FleetIdLabel = "fleet_id"

	AlertNameFleetOverload = "fleet_overload"
	AlertNameFleetLowLoad  = "fleet_low_load"
)

type AlertReq struct {
	Alerts model.Alerts `json:"alerts"`
}

type AlertResp struct {
	Result string `json:"result"`
}

func Alerts(req *restful.Request, resp *restful.Response) {
	// todo 冷冻期
	logger.Info("=== Handle alerts start")
	alertReq := &AlertReq{}
	if err := req.ReadEntity(alertReq); err != nil {
		logger.Errorf("Request read entity err: %v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusBadRequest, utils.RequestBodyParamInvalid(err.Error()))
		return
	}
	for _, alert := range alertReq.Alerts {
		if alert.Resolved() {
			logger.Infof("Alert[name: %s, fleet: %s, fingerprint: %s] resolved",
				alert.Name(), alert.Labels[FleetIdLabel], alert.Fingerprint())
			continue
		}
		if err := handleAlert(alert); err != nil {
			logger.Errorf("Handle alert[fingerprint: %s] err: %+v", alert.Fingerprint(), err)
		}
	}
	logger.Info("=== Handle alerts end")
}

// 处理“firing”状态的Alert
func handleAlert(alert *model.Alert) error {
	logger.Infof("handle alert[name: %s, fingerprint: %s]", alert.Name(), alert.Fingerprint())
	if alert.Status() != model.AlertFiring {
		return errors.Errorf("invalid alert status[%s]", alert.Status())
	}
	fleetId := string(alert.Labels[FleetIdLabel])
	if fleetId == "" {
		return errors.New("alert has not fleetId field")
	}

	switch alert.Name() {
	case AlertNameFleetOverload:
		_, err := ascontroller.GetAsController().AddServers(fleetId, 1)
		if err != nil {
			return err
		}
	case AlertNameFleetLowLoad:
		// todo 请求普罗米修斯获取、决策需要删除的serverId
		if err := ascontroller.GetAsController().DelServer(fleetId, ""); err != nil {
			return err
		}
	default:
		return errors.Errorf("invalid alert name[%s]", alert.Name())
	}

	return nil
}
