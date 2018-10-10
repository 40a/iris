package filter

import (
	"github.com/olegsu/iris/pkg/kube"
	"github.com/yalp/jsonpath"
)

type podStatus struct {
	baseFilter `yaml:",inline"`
	Status     []string `yaml:"status"`
	kube       kube.Kube
}

func (p *podStatus) Apply(data interface{}) (bool, error) {
	kind, err := jsonpath.Read(data, "$.involvedObject.kind")
	if err != nil {
		return false, nil
	}
	if kind.(string) != "Pod" {
		return false, nil
	}
	name, err := jsonpath.Read(data, "$.involvedObject.name")
	if err != nil {
		return false, nil
	}
	namespace, err := jsonpath.Read(data, "$.involvedObject.namespace")
	if err != nil {
		return false, nil
	}
	return p.kube.MatchPodToStatus(namespace.(string), name.(string), p.Status), nil
}
