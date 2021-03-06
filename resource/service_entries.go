package resource

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/open-kingfisher/king-utils/common"
	"github.com/open-kingfisher/king-utils/common/handle"
	"github.com/open-kingfisher/king-utils/common/log"
	"github.com/open-kingfisher/king-utils/kit"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type ServiceEntriesResource struct {
	Params   *handle.Resources
	PostData *v1alpha3.ServiceEntry
	Access   *versionedclient.Clientset
}

func (r *ServiceEntriesResource) Get() (*v1alpha3.ServiceEntry, error) {
	i, err := r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Get(context.TODO(), r.Params.Name, metav1.GetOptions{})
	if err == nil {
		i.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{Kind: "ServiceEntry", Version: "networking.istio.io/v1alpha3"})
	}
	return i, err
}

func (r *ServiceEntriesResource) List() (*v1alpha3.ServiceEntryList, error) {
	return r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).List(context.TODO(), metav1.ListOptions{})
}

func (r *ServiceEntriesResource) Delete() (err error) {
	if err = r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Delete(context.TODO(), r.Params.Name, metav1.DeleteOptions{}); err != nil {
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.ServiceEntries,
		ActionType: common.Delete,
		Resources:  r.Params,
		Name:       r.Params.Name,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *ServiceEntriesResource) Patch() (res *v1alpha3.ServiceEntry, err error) {
	var data []byte
	if data, err = json.Marshal(r.Params.PatchData.Patches); err != nil {
		return
	}
	if res, err = r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Patch(context.TODO(), r.Params.Name, types.JSONPatchType, data, metav1.PatchOptions{}); err != nil {
		log.Errorf("%s patch error:%s; Json:%+v; Name:%s", common.ServiceEntries, err, string(data), r.Params.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.ServiceEntries,
		ActionType: common.Patch,
		Resources:  r.Params,
		Name:       r.Params.Name,
	}
	if err = auditLog.InsertAuditLog(); err != nil {
		return
	}
	return
}

func (r *ServiceEntriesResource) Update() (res *v1alpha3.ServiceEntry, err error) {
	if r.Params.PostType == "form" {
		if se, err := r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Get(context.TODO(), r.PostData.Name, metav1.GetOptions{}); err != nil {
			log.Errorf("%s get error:%s; Json:%+v; Name:%s", common.ServiceEntries, err, r.PostData, r.PostData.Name)
			return nil, err
		} else {
			r.PostData.Spec.ExportTo = se.Spec.ExportTo
			r.PostData.Spec.SubjectAltNames = se.Spec.SubjectAltNames
		}
	}
	if res, err = r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Update(context.TODO(), r.PostData, metav1.UpdateOptions{}); err != nil {
		log.Errorf("%s update error:%s; Json:%+v; Name:%s", common.ServiceEntries, err, r.PostData, r.PostData.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.ServiceEntries,
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

func (r *ServiceEntriesResource) Create() (res *v1alpha3.ServiceEntry, err error) {
	if res, err = r.Access.NetworkingV1alpha3().ServiceEntries(r.Params.Namespace).Create(context.TODO(), r.PostData, metav1.CreateOptions{}); err != nil {
		log.Errorf("%s create error:%s; Json:%+v; Name:%s", common.ServiceEntries, err, r.PostData, r.PostData.Name)
		return
	}
	auditLog := handle.AuditLog{
		Kind:       common.ServiceEntries,
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

func (r *ServiceEntriesResource) GenerateCreateData(c *gin.Context) (err error) {
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
