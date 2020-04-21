package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"kingfisher/kf/common"
	"kingfisher/kf/common/handle"
	"kingfisher/kf/common/log"
	"kingfisher/kf/kit"
	"kingfisher/king-istio/common/chart"
	grpcClient "kingfisher/king-istio/common/grpc"
	"kingfisher/king-istio/pkg/apis/networking/v1alpha3"
	networkingv1alpha3 "kingfisher/king-istio/pkg/client/clientset/versioned/typed/networking/v1alpha3"
	"kingfisher/king-k8s/util"
)

type DestinationRulesResource struct {
	Params   *handle.Resources
	PostData *v1alpha3.DestinationRule
	Access   *networkingv1alpha3.NetworkingV1alpha3Client
}

func (r *DestinationRulesResource) Get() (*v1alpha3.DestinationRule, error) {
	i, err := r.Access.DestinationRules(r.Params.Namespace).Get(r.Params.Name, metav1.GetOptions{})
	if err == nil {
		i.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{Kind: "DestinationRule", Version: "networking.istio.io/v1alpha3"})
	}
	return i, err
}

func (r *DestinationRulesResource) List() (*v1alpha3.DestinationRuleList, error) {
	return r.Access.DestinationRules(r.Params.Namespace).List(metav1.ListOptions{})
}

func (r *DestinationRulesResource) Delete() (err error) {
	if err = r.Access.DestinationRules(r.Params.Namespace).Delete(r.Params.Name, &metav1.DeleteOptions{}); err != nil {
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.DestinationRules,
		ActionType: common.Delete,
		Resources:  r.Params,
		Name:       r.Params.Name,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *DestinationRulesResource) Patch() (res *v1alpha3.DestinationRule, err error) {
	var data []byte
	if data, err = json.Marshal(r.Params.PatchData.Patches); err != nil {
		return
	}
	if res, err = r.Access.DestinationRules(r.Params.Namespace).Patch(r.Params.Name, types.JSONPatchType, data); err != nil {
		log.Errorf("%s patch error:%s; Json:%+v; Name:%s", common.DestinationRules, err, string(data), r.Params.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.DestinationRules,
		ActionType: common.Patch,
		Resources:  r.Params,
		Name:       r.Params.Name,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *DestinationRulesResource) Update() (res *v1alpha3.DestinationRule, err error) {
	if res, err = r.Access.DestinationRules(r.Params.Namespace).Update(r.PostData); err != nil {
		log.Errorf("%s update error:%s; Json:%+v; Name:%s", common.DestinationRules, err, r.PostData, r.PostData.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.DestinationRules,
		ActionType: common.Update,
		PostType:   common.ActionType(r.Params.PostType),
		Resources:  r.Params,
		Name:       r.PostData.Name,
		PostData:   r.PostData,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *DestinationRulesResource) Create() (res *v1alpha3.DestinationRule, err error) {
	if res, err = r.Access.DestinationRules(r.Params.Namespace).Create(r.PostData); err != nil {
		log.Errorf("%s create error:%s; Json:%+v; Name:%s", common.DestinationRules, err, r.PostData, r.PostData.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.DestinationRules,
		ActionType: common.Create,
		PostType:   common.ActionType(r.Params.PostType),
		Resources:  r.Params,
		Name:       r.PostData.Name,
		PostData:   r.PostData,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *DestinationRulesResource) Chart() (interface{}, error) {
	chartData := chart.Chart{}
	if dr, err := r.Get(); err != nil {
		return nil, err
	} else {
		label := make(map[string]string)
		// get service
		service, err := grpcClient.GetService(r.Params.Cluster, r.Params.Namespace, dr.Spec.Host)
		if err != nil {
			log.Errorf("get service by grpc error: %v", err)
			return nil, err
		}
		// 获取service标签
		label = service.Spec.Selector

		count := 0
		chartData.Name = r.Params.Name
		chartData.Rank = "dr"
		for _, v := range dr.Spec.Subsets {
			subsets := chart.Chart{}
			subsets.Name = fmt.Sprintf("版本: %s", v.Name)
			subsets.Rank = "subsets"
			deployment := &appsv1.DeploymentList{}
			// 把DR的标签添进去，就能找到对应的Deployment了
			for key, v := range v.Labels {
				label[key] = v
			}
			// get deployment
			labelSelector := util.GenerateLabelSelector(label)
			deployment, err := grpcClient.GetDeployment(r.Params.Cluster, r.Params.Namespace, labelSelector)
			if err != nil {
				log.Errorf("get deployment by grpc error: %v", err)
			}
			for _, d := range deployment.Items {
				deployment := chart.Chart{}
				deployment.Name = fmt.Sprintf("部署: %s", d.Name)
				deployment.Rank = "deployment"
				subsets.Children = append(subsets.Children, deployment)
				// 计算Pod数量
				count++
			}
			chartData.Children = append(chartData.Children, subsets)
		}
		chartData.Count = count
	}
	return &chartData, nil
}

func (r *DestinationRulesResource) GenerateCreateData(c *gin.Context) (err error) {
	switch r.Params.DataType {
	case "yaml":
		var j []byte
		create := common.PostType{}
		if err = c.BindJSON(&create); err != nil {
			return
		}
		if j, _, err = kit.YamlToJson(create.Context); err != nil {
			return
		}
		if err = json.Unmarshal(j, &r.PostData); err != nil {
			return
		}
	case "json":
		if err = c.BindJSON(&r.PostData); err != nil {
			return
		}
	default:
		return errors.New(common.ContentTypeError)
	}
	return nil
}
