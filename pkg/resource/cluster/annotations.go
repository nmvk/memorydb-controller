package cluster

import (
	svcapitypes "github.com/aws-controllers-k8s/memorydb-controller/apis/v1alpha1"
	"strconv"
)

const (
	// AnnotationLastRequestedNumShards is an annotation whose value is the value of NumShards
	// passed in as input to either the create or modify API called most recently
	AnnotationLastRequestedNumShards = svcapitypes.AnnotationPrefix + "last-requested-num-shards"
	// AnnotationLastRequestedNumReplicasPerShard is an annotation whose value is the value of NumReplicasPerShard
	// passed in as input to either the create or modify API called most recently
	AnnotationLastRequestedNumReplicasPerShard = svcapitypes.AnnotationPrefix + "last-requested-num-replicas-per-shard"
	// AnnotationLastRequestedNodeType is an annotation whose value is the value of NodeType
	// passed in as input to either the create or modify API called most recently
	AnnotationLastRequestedNodeType = svcapitypes.AnnotationPrefix + "last-requested-node-type"
)

// setNumShardAnnotation sets the AnnotationLastRequestedNumShards annotation for cluster resource
// This should only be called upon a successful create or modify call.
func (rm *resourceManager) setNumShardAnnotation(
	numShards *int64,
	ko *svcapitypes.Cluster,
) {
	if ko.ObjectMeta.Annotations == nil {
		ko.ObjectMeta.Annotations = make(map[string]string)
	}

	annotations := ko.ObjectMeta.Annotations

	if numShards != nil {
		annotations[AnnotationLastRequestedNumShards] = strconv.FormatInt(*numShards, 10)
	}
}

// setNumReplicasPerShardAnnotation sets the AnnotationLastRequestedNumReplicasPerShard annotation for cluster resource
// This should only be called upon a successful create or modify call.
func (rm *resourceManager) setNumReplicasPerShardAnnotation(
	numReplicasPerShard *int64,
	ko *svcapitypes.Cluster,
) {
	if ko.ObjectMeta.Annotations == nil {
		ko.ObjectMeta.Annotations = make(map[string]string)
	}

	annotations := ko.ObjectMeta.Annotations

	if numReplicasPerShard != nil {
		annotations[AnnotationLastRequestedNumReplicasPerShard] = strconv.FormatInt(*numReplicasPerShard, 10)
	}
}

// setNodeTypeAnnotation sets the AnnotationLastRequestedNodeType annotation for cluster resource
// This should only be called upon a successful create or modify call.
func (rm *resourceManager) setNodeTypeAnnotation(
	nodeType *string,
	ko *svcapitypes.Cluster,
) {
	if ko.ObjectMeta.Annotations == nil {
		ko.ObjectMeta.Annotations = make(map[string]string)
	}

	annotations := ko.ObjectMeta.Annotations

	if nodeType != nil {
		annotations[AnnotationLastRequestedNodeType] = *nodeType
	}
}
