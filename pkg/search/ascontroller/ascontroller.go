package ascontroller

import (
	"os"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	as "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1/region"
	"github.com/pkg/errors"

	"nanto.io/application-auto-scaling-service/pkg/common/utils/logutil"
	"nanto.io/application-auto-scaling-service/pkg/search/types"
)

var (
	logger = logutil.GetLogger()

	ScalingConfigurationId       = "7da2e330-686a-4135-8805-a71703961a02"
	DesireInstanceNumber   int32 = 2
	MaxInstanceNumber      int32 = 12
	CoolDownTime           int32 = 120 // 单位min
	SubnetId                     = "e2f7329d-1660-4980-9b7f-edde8cec2a6b"
	VpcId                        = "20305e52-1e3a-43ce-9d6b-3fd37a13bb6d"
	EnterpriseProjectId          = "5fd11e03-cf64-4c2f-89e7-f6fab115846b"
	Description                  = "aass创建的fleet"
)

var asCtl *AsController

type AsController struct {
	asClient *as.AsClient
}

func GetAsController() *AsController {
	if asCtl == nil {
		logger.Panic("As controller must be initialized")
	}
	return asCtl
}

func InitAsController() {
	ak := os.Getenv("AK")
	sk := os.Getenv("SK")
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	cli := as.NewAsClient(
		as.AsClientBuilder().
			WithRegion(region.ValueOf("ap-southeast-1")).
			WithCredential(auth).
			Build())
	asCtl = &AsController{asClient: cli}
}

func (c *AsController) CreateFleet(fleetName string) (string, error) {
	// 1.创建AS伸缩组
	// todo 创建vpc、subnet
	createScalingGroupReq := &model.CreateScalingGroupRequest{
		Body: &model.CreateScalingGroupOption{
			ScalingGroupName:       fleetName,
			ScalingConfigurationId: &ScalingConfigurationId,
			DesireInstanceNumber:   &DesireInstanceNumber,
			MinInstanceNumber:      nil,
			MaxInstanceNumber:      &MaxInstanceNumber,
			CoolDownTime:           &CoolDownTime,
			Networks: []model.Networks{{
				Id: SubnetId,
			}},
			VpcId:               VpcId,
			EnterpriseProjectId: &EnterpriseProjectId,
			Description:         &Description,
		},
	}
	createScalingGroupResp, err := c.asClient.CreateScalingGroup(createScalingGroupReq)
	if err != nil {
		return "", errors.Wrap(err, "as client create ScalingGroup err")
	}
	fleetId := *createScalingGroupResp.ScalingGroupId
	logger.Infof("AsClient create ScalingGroup success, group id: %s", fleetId)

	// 2.启用弹性伸缩组
	resumeScalingGroupReq := &model.ResumeScalingGroupRequest{
		ScalingGroupId: fleetId,
		Body: &model.ResumeScalingGroupOption{
			Action: model.GetResumeScalingGroupOptionActionEnum().RESUME,
		},
	}
	resumeResp, err := c.asClient.ResumeScalingGroup(resumeScalingGroupReq)
	if err != nil {
		return "", errors.Wrapf(err, "as client resume ScalingGroup[%s] err", fleetId)
	}
	logger.Infof("AsClient resume ScalingGroup resp: %+v", resumeResp)

	// 查询弹性伸缩组详情，等待实例数变为DesireInstanceNumber，后面考虑异步调用gateway create server
	for {
		time.Sleep(10 * time.Second)
		showScalingGroupResp, err := c.asClient.ShowScalingGroup(&model.ShowScalingGroupRequest{
			ScalingGroupId: fleetId})
		if err != nil {
			return "", errors.Wrapf(err, "as client show ScalingGroup[%s] err", fleetId)
		}
		if *showScalingGroupResp.ScalingGroup.CurrentInstanceNumber == DesireInstanceNumber {
			break
		}
		logger.Infof("Waiting instance init(%d/%d)……",
			*showScalingGroupResp.ScalingGroup.CurrentInstanceNumber, DesireInstanceNumber)
	}

	// 3.查询弹性伸缩组中的实例列表
	listResp, err := c.asClient.ListScalingInstances(&model.ListScalingInstancesRequest{
		ScalingGroupId: fleetId})
	if err != nil {
		return "", errors.Wrapf(err, "as client list ScalingGroup[%s] instances err", fleetId)
	}
	logger.Infof("AsClient list instsance resp: %+v", listResp)

	return fleetId, nil
}

func (c *AsController) AddServers(fleetId string, num int) ([]*types.Server, error) {
	showResp, err := c.asClient.ShowScalingGroup(&model.ShowScalingGroupRequest{ScalingGroupId: fleetId})
	if err != nil {
		return nil, errors.Wrapf(err, "as client show ScalingGroup[%s] err", fleetId)
	}
	logger.Infof("AsClient ShowScalingGroup info resp: %+v", showResp.ScalingGroup)
	updateNum := *showResp.ScalingGroup.DesireInstanceNumber + int32(num)

	_, err = c.asClient.UpdateScalingGroup(&model.UpdateScalingGroupRequest{
		ScalingGroupId: fleetId,
		Body: &model.UpdateScalingGroupOption{
			DesireInstanceNumber: &updateNum,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "as client add server to ScalingGroup[%s] err", fleetId)
	}

	// todo 新添加实例的id和ip
	//cache.NewServer()

	return nil, nil
}

func (c *AsController) DelServer(fleetId string, serverId string) error {
	// 临时代码，如果serverId不指定，则删除最近创建的实例
	if serverId == "" {
		resp, err := c.asClient.ListScalingInstances(&model.ListScalingInstancesRequest{
			ScalingGroupId: fleetId})
		if err != nil {
			return errors.Wrapf(err, "as client list ScalingGroup[%s] instances err", fleetId)
		}
		logger.Infof("AsClient list instsance resp: %+v", resp)
		instances := *resp.ScalingGroupInstances
		if len(instances) == 0 {
			return errors.Errorf("fleet[%s] has not instance", fleetId)
		}
		serverId = *(instances[0].InstanceId)
	}

	isDeleteVM := model.GetDeleteScalingInstanceRequestInstanceDeleteEnum().YES
	_, err := c.asClient.DeleteScalingInstance(&model.DeleteScalingInstanceRequest{
		InstanceId:     serverId,
		InstanceDelete: &isDeleteVM,
	})
	if err != nil {
		return errors.Wrapf(err, "as client delete server[%s] err", serverId)
	}
	logger.Infof("As client delete server[%s] for fleetId[%s] success", serverId, fleetId)
	return nil
}
