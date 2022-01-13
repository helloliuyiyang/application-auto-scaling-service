package httpserver

import (
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"nanto.io/application-auto-scaling-service/pkg/common/utils"
	"nanto.io/application-auto-scaling-service/pkg/search/cache"
	"nanto.io/application-auto-scaling-service/pkg/search/types"
)

func Prometheus() {
	var (
		fleetMaxConcurrencyGauge = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "fleetMaxConcurrency",
				Help: "队列支持的最大请求并发数",
			}, []string{"fleet_id"},
		)
		fleetNowConcurrencyGauge = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "fleetNowConcurrency",
				Help: "队列当前请求并发数",
			}, []string{"fleet_id"},
		)
	)

	// Register the fleetMaxConcurrencyGauge and the fleetNowConcurrencyGauge with Prometheus's default registry.
	prometheus.MustRegister(fleetMaxConcurrencyGauge, fleetNowConcurrencyGauge)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())

	go func() {
		for {
			fleets := cache.GetFleetsCache().Snapshot()
			for _, fleet := range fleets {
				fleetMaxConcurrencyGauge.WithLabelValues(fleet.Id).Set(float64(fleet.MaxConcurrency))
				fleetNowConcurrencyGauge.WithLabelValues(fleet.Id).Set(float64(fleet.NowConcurrency))
			}
			time.Sleep(time.Second * 5)
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	addr := ":9090"
	logger.Infof("Start listen[%s] for prometheus metrics", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type MockCacheReq struct {
	FleetInfos []FleetInfo `json:"fleet_infos"`
}

type FleetInfo struct {
	Id             string `json:"id"`
	MaxConcurrency int    `json:"max_concurrency"`
	NowConcurrency int    `json:"now_concurrency"`
}

type MockCacheResp struct {
	Result string `json:"result"`
}

func MockCache(req *restful.Request, resp *restful.Response) {
	mockCacheReq := &MockCacheReq{}
	if err := req.ReadEntity(mockCacheReq); err != nil {
		logger.Errorf("Request read entity err: %v", err)
		utils.WriteFailedJSONResponse(resp, http.StatusBadRequest, utils.RequestBodyParamInvalid(err.Error()))
		return
	}

	logger.Infof("MockCache req: %+v", mockCacheReq)

	for _, info := range mockCacheReq.FleetInfos {
		cache.GetFleetsCache().AddFleet(&types.Fleet{
			Id:             info.Id,
			MaxConcurrency: info.MaxConcurrency,
			NowConcurrency: info.NowConcurrency,
		})
	}

	resp.WriteHeaderAndEntity(http.StatusOK, MockCacheResp{Result: "ok"})
}
