package utils

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/xco-sk/eck-custom-resources/apis/es.eck/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
)

func DeleteSnapshotLifecyclePolicy(esClient *elasticsearch.Client, snapshotLifecyclePolicyName string) (ctrl.Result, error) {
	_, err := esClient.SlmDeleteLifecycle(snapshotLifecyclePolicyName)
	if err != nil {
		return GetRequeueResult(), err
	}
	return ctrl.Result{}, nil
}

func UpsertSnapshotLifecyclePolicy(esClient *elasticsearch.Client, snapshotLifecyclePolicy v1alpha1.SnapshotLifecyclePolicy) (ctrl.Result, error) {
	_, err := esClient.SlmPutLifecycle(snapshotLifecyclePolicy.Name, esClient.SlmPutLifecycle.WithBody(strings.NewReader(snapshotLifecyclePolicy.Spec.Body)))
	if err != nil {
		return GetRequeueResult(), err
	}
	return ctrl.Result{}, nil
}
