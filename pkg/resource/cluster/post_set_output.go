package cluster

import (
	"context"
	svcapitypes "github.com/aws-controllers-k8s/memorydb-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/memorydb"
)

func (rm *resourceManager) setShardDetails(
	ctx context.Context,
	r *resource,
	ko *svcapitypes.Cluster,
) (*svcapitypes.Cluster, error) {

	resp, err := rm.sdkFind(ctx, r)

	if err != nil {
		return nil, err
	}

	ko.Status = resp.ko.Status
	ko.Spec.NumReplicasPerShard = resp.ko.Spec.NumReplicasPerShard
	ko.Spec.NumShards = resp.ko.Spec.NumShards

	return ko, nil
}

func (rm *resourceManager) setAllowedNodeTypeUpdates(
	ctx context.Context,
	ko *svcapitypes.Cluster,
) {
	if *ko.Status.Status != "available" {
		return
	}

	response, respErr := rm.sdkapi.ListAllowedNodeTypeUpdatesWithContext(ctx, &svcsdk.ListAllowedNodeTypeUpdatesInput{
		ClusterName: ko.Spec.Name,
	})
	rm.metrics.RecordAPICall("GET", "ListAllowedNodeTypeUpdates", respErr)
	// Ignore the error since the response from this API is used for information purpose only
	if respErr == nil {
		ko.Status.AllowedScaleDownNodeTypes = response.ScaleDownNodeTypes
		ko.Status.AllowedScaleUpNodeTypes = response.ScaleUpNodeTypes
	}
}
