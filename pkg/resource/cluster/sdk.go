// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package cluster

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/memorydb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/memorydb-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.MemoryDB{}
	_ = &svcapitypes.Cluster{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeClustersOutput
	resp, err = rm.sdkapi.DescribeClustersWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeClusters", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ClusterNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Clusters {
		if elem.ACLName != nil {
			ko.Spec.ACLName = elem.ACLName
		} else {
			ko.Spec.ACLName = nil
		}
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AutoMinorVersionUpgrade != nil {
			ko.Spec.AutoMinorVersionUpgrade = elem.AutoMinorVersionUpgrade
		} else {
			ko.Spec.AutoMinorVersionUpgrade = nil
		}
		if elem.AvailabilityMode != nil {
			ko.Status.AvailabilityMode = elem.AvailabilityMode
		} else {
			ko.Status.AvailabilityMode = nil
		}
		if elem.ClusterEndpoint != nil {
			f4 := &svcapitypes.Endpoint{}
			if elem.ClusterEndpoint.Address != nil {
				f4.Address = elem.ClusterEndpoint.Address
			}
			if elem.ClusterEndpoint.Port != nil {
				f4.Port = elem.ClusterEndpoint.Port
			}
			ko.Status.ClusterEndpoint = f4
		} else {
			ko.Status.ClusterEndpoint = nil
		}
		if elem.Description != nil {
			ko.Spec.Description = elem.Description
		} else {
			ko.Spec.Description = nil
		}
		if elem.EnginePatchVersion != nil {
			ko.Status.EnginePatchVersion = elem.EnginePatchVersion
		} else {
			ko.Status.EnginePatchVersion = nil
		}
		if elem.EngineVersion != nil {
			ko.Spec.EngineVersion = elem.EngineVersion
		} else {
			ko.Spec.EngineVersion = nil
		}
		if elem.KmsKeyId != nil {
			ko.Spec.KMSKeyID = elem.KmsKeyId
		} else {
			ko.Spec.KMSKeyID = nil
		}
		if elem.MaintenanceWindow != nil {
			ko.Spec.MaintenanceWindow = elem.MaintenanceWindow
		} else {
			ko.Spec.MaintenanceWindow = nil
		}
		if elem.Name != nil {
			ko.Spec.Name = elem.Name
		} else {
			ko.Spec.Name = nil
		}
		if elem.NodeType != nil {
			ko.Spec.NodeType = elem.NodeType
		} else {
			ko.Spec.NodeType = nil
		}
		if elem.NumberOfShards != nil {
			ko.Status.NumberOfShards = elem.NumberOfShards
		} else {
			ko.Status.NumberOfShards = nil
		}
		if elem.ParameterGroupName != nil {
			ko.Spec.ParameterGroupName = elem.ParameterGroupName
		} else {
			ko.Spec.ParameterGroupName = nil
		}
		if elem.ParameterGroupStatus != nil {
			ko.Status.ParameterGroupStatus = elem.ParameterGroupStatus
		} else {
			ko.Status.ParameterGroupStatus = nil
		}
		if elem.PendingUpdates != nil {
			f15 := &svcapitypes.ClusterPendingUpdates{}
			if elem.PendingUpdates.ACLs != nil {
				f15f0 := &svcapitypes.ACLsUpdateStatus{}
				if elem.PendingUpdates.ACLs.ACLToApply != nil {
					f15f0.ACLToApply = elem.PendingUpdates.ACLs.ACLToApply
				}
				f15.ACLs = f15f0
			}
			if elem.PendingUpdates.Resharding != nil {
				f15f1 := &svcapitypes.ReshardingStatus{}
				if elem.PendingUpdates.Resharding.SlotMigration != nil {
					f15f1f0 := &svcapitypes.SlotMigration{}
					if elem.PendingUpdates.Resharding.SlotMigration.ProgressPercentage != nil {
						f15f1f0.ProgressPercentage = elem.PendingUpdates.Resharding.SlotMigration.ProgressPercentage
					}
					f15f1.SlotMigration = f15f1f0
				}
				f15.Resharding = f15f1
			}
			if elem.PendingUpdates.ServiceUpdates != nil {
				f15f2 := []*svcapitypes.PendingModifiedServiceUpdate{}
				for _, f15f2iter := range elem.PendingUpdates.ServiceUpdates {
					f15f2elem := &svcapitypes.PendingModifiedServiceUpdate{}
					if f15f2iter.ServiceUpdateName != nil {
						f15f2elem.ServiceUpdateName = f15f2iter.ServiceUpdateName
					}
					if f15f2iter.Status != nil {
						f15f2elem.Status = f15f2iter.Status
					}
					f15f2 = append(f15f2, f15f2elem)
				}
				f15.ServiceUpdates = f15f2
			}
			ko.Status.PendingUpdates = f15
		} else {
			ko.Status.PendingUpdates = nil
		}
		if elem.SecurityGroups != nil {
			f16 := []*svcapitypes.SecurityGroupMembership{}
			for _, f16iter := range elem.SecurityGroups {
				f16elem := &svcapitypes.SecurityGroupMembership{}
				if f16iter.SecurityGroupId != nil {
					f16elem.SecurityGroupID = f16iter.SecurityGroupId
				}
				if f16iter.Status != nil {
					f16elem.Status = f16iter.Status
				}
				f16 = append(f16, f16elem)
			}
			ko.Status.SecurityGroups = f16
		} else {
			ko.Status.SecurityGroups = nil
		}
		if elem.Shards != nil {
			f17 := []*svcapitypes.Shard{}
			for _, f17iter := range elem.Shards {
				f17elem := &svcapitypes.Shard{}
				if f17iter.Name != nil {
					f17elem.Name = f17iter.Name
				}
				if f17iter.Nodes != nil {
					f17elemf1 := []*svcapitypes.Node{}
					for _, f17elemf1iter := range f17iter.Nodes {
						f17elemf1elem := &svcapitypes.Node{}
						if f17elemf1iter.AvailabilityZone != nil {
							f17elemf1elem.AvailabilityZone = f17elemf1iter.AvailabilityZone
						}
						if f17elemf1iter.CreateTime != nil {
							f17elemf1elem.CreateTime = &metav1.Time{*f17elemf1iter.CreateTime}
						}
						if f17elemf1iter.Endpoint != nil {
							f17elemf1elemf2 := &svcapitypes.Endpoint{}
							if f17elemf1iter.Endpoint.Address != nil {
								f17elemf1elemf2.Address = f17elemf1iter.Endpoint.Address
							}
							if f17elemf1iter.Endpoint.Port != nil {
								f17elemf1elemf2.Port = f17elemf1iter.Endpoint.Port
							}
							f17elemf1elem.Endpoint = f17elemf1elemf2
						}
						if f17elemf1iter.Name != nil {
							f17elemf1elem.Name = f17elemf1iter.Name
						}
						if f17elemf1iter.Status != nil {
							f17elemf1elem.Status = f17elemf1iter.Status
						}
						f17elemf1 = append(f17elemf1, f17elemf1elem)
					}
					f17elem.Nodes = f17elemf1
				}
				if f17iter.NumberOfNodes != nil {
					f17elem.NumberOfNodes = f17iter.NumberOfNodes
				}
				if f17iter.Slots != nil {
					f17elem.Slots = f17iter.Slots
				}
				if f17iter.Status != nil {
					f17elem.Status = f17iter.Status
				}
				f17 = append(f17, f17elem)
			}
			ko.Status.Shards = f17
		} else {
			ko.Status.Shards = nil
		}
		if elem.SnapshotRetentionLimit != nil {
			ko.Spec.SnapshotRetentionLimit = elem.SnapshotRetentionLimit
		} else {
			ko.Spec.SnapshotRetentionLimit = nil
		}
		if elem.SnapshotWindow != nil {
			ko.Spec.SnapshotWindow = elem.SnapshotWindow
		} else {
			ko.Spec.SnapshotWindow = nil
		}
		if elem.SnsTopicArn != nil {
			ko.Spec.SnsTopicARN = elem.SnsTopicArn
		} else {
			ko.Spec.SnsTopicARN = nil
		}
		if elem.SnsTopicStatus != nil {
			ko.Status.SnsTopicStatus = elem.SnsTopicStatus
		} else {
			ko.Status.SnsTopicStatus = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		if elem.SubnetGroupName != nil {
			ko.Spec.SubnetGroupName = elem.SubnetGroupName
		} else {
			ko.Spec.SubnetGroupName = nil
		}
		if elem.TLSEnabled != nil {
			ko.Spec.TLSEnabled = elem.TLSEnabled
		} else {
			ko.Spec.TLSEnabled = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	if resp.Clusters[0].NumberOfShards != nil {
		ko.Spec.NumShards = resp.Clusters[0].NumberOfShards
	} else {
		ko.Spec.NumShards = nil
	}

	if resp.Clusters[0].Shards != nil && resp.Clusters[0].Shards[0].NumberOfNodes != nil {
		replicas := *resp.Clusters[0].Shards[0].NumberOfNodes - 1
		ko.Spec.NumReplicasPerShard = &replicas
	} else {
		ko.Spec.NumReplicasPerShard = nil
	}

	if resp.Clusters[0].SecurityGroups != nil {
		var securityGroupIds []*string
		for _, securityGroup := range resp.Clusters[0].SecurityGroups {
			if securityGroup.SecurityGroupId != nil {
				securityGroupIds = append(securityGroupIds, securityGroup.SecurityGroupId)
			}
		}
		ko.Spec.SecurityGroupIDs = securityGroupIds
	} else {
		ko.Spec.SecurityGroupIDs = nil
	}

	rm.setAllowedNodeTypeUpdates(ctx, ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Spec.Name == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeClustersInput, error) {
	res := &svcsdk.DescribeClustersInput{}

	if r.ko.Spec.Name != nil {
		res.SetClusterName(*r.ko.Spec.Name)
	}
	res.SetShowShardDetails(true)

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateClusterOutput
	_ = resp
	resp, err = rm.sdkapi.CreateClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateCluster", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.Cluster.ACLName != nil {
		ko.Spec.ACLName = resp.Cluster.ACLName
	} else {
		ko.Spec.ACLName = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Cluster.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Cluster.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Cluster.AutoMinorVersionUpgrade != nil {
		ko.Spec.AutoMinorVersionUpgrade = resp.Cluster.AutoMinorVersionUpgrade
	} else {
		ko.Spec.AutoMinorVersionUpgrade = nil
	}
	if resp.Cluster.AvailabilityMode != nil {
		ko.Status.AvailabilityMode = resp.Cluster.AvailabilityMode
	} else {
		ko.Status.AvailabilityMode = nil
	}
	if resp.Cluster.ClusterEndpoint != nil {
		f4 := &svcapitypes.Endpoint{}
		if resp.Cluster.ClusterEndpoint.Address != nil {
			f4.Address = resp.Cluster.ClusterEndpoint.Address
		}
		if resp.Cluster.ClusterEndpoint.Port != nil {
			f4.Port = resp.Cluster.ClusterEndpoint.Port
		}
		ko.Status.ClusterEndpoint = f4
	} else {
		ko.Status.ClusterEndpoint = nil
	}
	if resp.Cluster.Description != nil {
		ko.Spec.Description = resp.Cluster.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.Cluster.EnginePatchVersion != nil {
		ko.Status.EnginePatchVersion = resp.Cluster.EnginePatchVersion
	} else {
		ko.Status.EnginePatchVersion = nil
	}
	if resp.Cluster.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.Cluster.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.Cluster.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.Cluster.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if resp.Cluster.MaintenanceWindow != nil {
		ko.Spec.MaintenanceWindow = resp.Cluster.MaintenanceWindow
	} else {
		ko.Spec.MaintenanceWindow = nil
	}
	if resp.Cluster.Name != nil {
		ko.Spec.Name = resp.Cluster.Name
	} else {
		ko.Spec.Name = nil
	}
	if resp.Cluster.NodeType != nil {
		ko.Spec.NodeType = resp.Cluster.NodeType
	} else {
		ko.Spec.NodeType = nil
	}
	if resp.Cluster.NumberOfShards != nil {
		ko.Status.NumberOfShards = resp.Cluster.NumberOfShards
	} else {
		ko.Status.NumberOfShards = nil
	}
	if resp.Cluster.ParameterGroupName != nil {
		ko.Spec.ParameterGroupName = resp.Cluster.ParameterGroupName
	} else {
		ko.Spec.ParameterGroupName = nil
	}
	if resp.Cluster.ParameterGroupStatus != nil {
		ko.Status.ParameterGroupStatus = resp.Cluster.ParameterGroupStatus
	} else {
		ko.Status.ParameterGroupStatus = nil
	}
	if resp.Cluster.PendingUpdates != nil {
		f15 := &svcapitypes.ClusterPendingUpdates{}
		if resp.Cluster.PendingUpdates.ACLs != nil {
			f15f0 := &svcapitypes.ACLsUpdateStatus{}
			if resp.Cluster.PendingUpdates.ACLs.ACLToApply != nil {
				f15f0.ACLToApply = resp.Cluster.PendingUpdates.ACLs.ACLToApply
			}
			f15.ACLs = f15f0
		}
		if resp.Cluster.PendingUpdates.Resharding != nil {
			f15f1 := &svcapitypes.ReshardingStatus{}
			if resp.Cluster.PendingUpdates.Resharding.SlotMigration != nil {
				f15f1f0 := &svcapitypes.SlotMigration{}
				if resp.Cluster.PendingUpdates.Resharding.SlotMigration.ProgressPercentage != nil {
					f15f1f0.ProgressPercentage = resp.Cluster.PendingUpdates.Resharding.SlotMigration.ProgressPercentage
				}
				f15f1.SlotMigration = f15f1f0
			}
			f15.Resharding = f15f1
		}
		if resp.Cluster.PendingUpdates.ServiceUpdates != nil {
			f15f2 := []*svcapitypes.PendingModifiedServiceUpdate{}
			for _, f15f2iter := range resp.Cluster.PendingUpdates.ServiceUpdates {
				f15f2elem := &svcapitypes.PendingModifiedServiceUpdate{}
				if f15f2iter.ServiceUpdateName != nil {
					f15f2elem.ServiceUpdateName = f15f2iter.ServiceUpdateName
				}
				if f15f2iter.Status != nil {
					f15f2elem.Status = f15f2iter.Status
				}
				f15f2 = append(f15f2, f15f2elem)
			}
			f15.ServiceUpdates = f15f2
		}
		ko.Status.PendingUpdates = f15
	} else {
		ko.Status.PendingUpdates = nil
	}
	if resp.Cluster.SecurityGroups != nil {
		f16 := []*svcapitypes.SecurityGroupMembership{}
		for _, f16iter := range resp.Cluster.SecurityGroups {
			f16elem := &svcapitypes.SecurityGroupMembership{}
			if f16iter.SecurityGroupId != nil {
				f16elem.SecurityGroupID = f16iter.SecurityGroupId
			}
			if f16iter.Status != nil {
				f16elem.Status = f16iter.Status
			}
			f16 = append(f16, f16elem)
		}
		ko.Status.SecurityGroups = f16
	} else {
		ko.Status.SecurityGroups = nil
	}
	if resp.Cluster.Shards != nil {
		f17 := []*svcapitypes.Shard{}
		for _, f17iter := range resp.Cluster.Shards {
			f17elem := &svcapitypes.Shard{}
			if f17iter.Name != nil {
				f17elem.Name = f17iter.Name
			}
			if f17iter.Nodes != nil {
				f17elemf1 := []*svcapitypes.Node{}
				for _, f17elemf1iter := range f17iter.Nodes {
					f17elemf1elem := &svcapitypes.Node{}
					if f17elemf1iter.AvailabilityZone != nil {
						f17elemf1elem.AvailabilityZone = f17elemf1iter.AvailabilityZone
					}
					if f17elemf1iter.CreateTime != nil {
						f17elemf1elem.CreateTime = &metav1.Time{*f17elemf1iter.CreateTime}
					}
					if f17elemf1iter.Endpoint != nil {
						f17elemf1elemf2 := &svcapitypes.Endpoint{}
						if f17elemf1iter.Endpoint.Address != nil {
							f17elemf1elemf2.Address = f17elemf1iter.Endpoint.Address
						}
						if f17elemf1iter.Endpoint.Port != nil {
							f17elemf1elemf2.Port = f17elemf1iter.Endpoint.Port
						}
						f17elemf1elem.Endpoint = f17elemf1elemf2
					}
					if f17elemf1iter.Name != nil {
						f17elemf1elem.Name = f17elemf1iter.Name
					}
					if f17elemf1iter.Status != nil {
						f17elemf1elem.Status = f17elemf1iter.Status
					}
					f17elemf1 = append(f17elemf1, f17elemf1elem)
				}
				f17elem.Nodes = f17elemf1
			}
			if f17iter.NumberOfNodes != nil {
				f17elem.NumberOfNodes = f17iter.NumberOfNodes
			}
			if f17iter.Slots != nil {
				f17elem.Slots = f17iter.Slots
			}
			if f17iter.Status != nil {
				f17elem.Status = f17iter.Status
			}
			f17 = append(f17, f17elem)
		}
		ko.Status.Shards = f17
	} else {
		ko.Status.Shards = nil
	}
	if resp.Cluster.SnapshotRetentionLimit != nil {
		ko.Spec.SnapshotRetentionLimit = resp.Cluster.SnapshotRetentionLimit
	} else {
		ko.Spec.SnapshotRetentionLimit = nil
	}
	if resp.Cluster.SnapshotWindow != nil {
		ko.Spec.SnapshotWindow = resp.Cluster.SnapshotWindow
	} else {
		ko.Spec.SnapshotWindow = nil
	}
	if resp.Cluster.SnsTopicArn != nil {
		ko.Spec.SnsTopicARN = resp.Cluster.SnsTopicArn
	} else {
		ko.Spec.SnsTopicARN = nil
	}
	if resp.Cluster.SnsTopicStatus != nil {
		ko.Status.SnsTopicStatus = resp.Cluster.SnsTopicStatus
	} else {
		ko.Status.SnsTopicStatus = nil
	}
	if resp.Cluster.Status != nil {
		ko.Status.Status = resp.Cluster.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.Cluster.SubnetGroupName != nil {
		ko.Spec.SubnetGroupName = resp.Cluster.SubnetGroupName
	} else {
		ko.Spec.SubnetGroupName = nil
	}
	if resp.Cluster.TLSEnabled != nil {
		ko.Spec.TLSEnabled = resp.Cluster.TLSEnabled
	} else {
		ko.Spec.TLSEnabled = nil
	}

	rm.setStatusDefaults(ko)
	ko, err = rm.setShardDetails(ctx, desired, ko)

	if err != nil {
		return nil, err
	}

	// Update the annotations to handle async rollback
	rm.setNodeTypeAnnotation(input.NodeType, ko)
	rm.setNumReplicasPerShardAnnotation(input.NumReplicasPerShard, ko)
	rm.setNumShardAnnotation(input.NumShards, ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateClusterInput, error) {
	res := &svcsdk.CreateClusterInput{}

	if r.ko.Spec.ACLName != nil {
		res.SetACLName(*r.ko.Spec.ACLName)
	}
	if r.ko.Spec.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	}
	if r.ko.Spec.Name != nil {
		res.SetClusterName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.MaintenanceWindow != nil {
		res.SetMaintenanceWindow(*r.ko.Spec.MaintenanceWindow)
	}
	if r.ko.Spec.NodeType != nil {
		res.SetNodeType(*r.ko.Spec.NodeType)
	}
	if r.ko.Spec.NumReplicasPerShard != nil {
		res.SetNumReplicasPerShard(*r.ko.Spec.NumReplicasPerShard)
	}
	if r.ko.Spec.NumShards != nil {
		res.SetNumShards(*r.ko.Spec.NumShards)
	}
	if r.ko.Spec.ParameterGroupName != nil {
		res.SetParameterGroupName(*r.ko.Spec.ParameterGroupName)
	}
	if r.ko.Spec.Port != nil {
		res.SetPort(*r.ko.Spec.Port)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f12 := []*string{}
		for _, f12iter := range r.ko.Spec.SecurityGroupIDs {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		res.SetSecurityGroupIds(f12)
	}
	if r.ko.Spec.SnapshotARNs != nil {
		f13 := []*string{}
		for _, f13iter := range r.ko.Spec.SnapshotARNs {
			var f13elem string
			f13elem = *f13iter
			f13 = append(f13, &f13elem)
		}
		res.SetSnapshotArns(f13)
	}
	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}
	if r.ko.Spec.SnapshotRetentionLimit != nil {
		res.SetSnapshotRetentionLimit(*r.ko.Spec.SnapshotRetentionLimit)
	}
	if r.ko.Spec.SnapshotWindow != nil {
		res.SetSnapshotWindow(*r.ko.Spec.SnapshotWindow)
	}
	if r.ko.Spec.SnsTopicARN != nil {
		res.SetSnsTopicArn(*r.ko.Spec.SnsTopicARN)
	}
	if r.ko.Spec.SubnetGroupName != nil {
		res.SetSubnetGroupName(*r.ko.Spec.SubnetGroupName)
	}
	if r.ko.Spec.TLSEnabled != nil {
		res.SetTLSEnabled(*r.ko.Spec.TLSEnabled)
	}
	if r.ko.Spec.Tags != nil {
		f20 := []*svcsdk.Tag{}
		for _, f20iter := range r.ko.Spec.Tags {
			f20elem := &svcsdk.Tag{}
			if f20iter.Key != nil {
				f20elem.SetKey(*f20iter.Key)
			}
			if f20iter.Value != nil {
				f20elem.SetValue(*f20iter.Value)
			}
			f20 = append(f20, f20elem)
		}
		res.SetTags(f20)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	res, err := rm.validateClusterNeedsUpdate(desired, latest, delta)

	if err != nil || res != nil {
		return res, err
	}
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}
	input = rm.newMemoryDBClusterUploadPayload(desired, latest, delta)

	var resp *svcsdk.UpdateClusterOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateCluster", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.Cluster.ACLName != nil {
		ko.Spec.ACLName = resp.Cluster.ACLName
	} else {
		ko.Spec.ACLName = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Cluster.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Cluster.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Cluster.AutoMinorVersionUpgrade != nil {
		ko.Spec.AutoMinorVersionUpgrade = resp.Cluster.AutoMinorVersionUpgrade
	} else {
		ko.Spec.AutoMinorVersionUpgrade = nil
	}
	if resp.Cluster.AvailabilityMode != nil {
		ko.Status.AvailabilityMode = resp.Cluster.AvailabilityMode
	} else {
		ko.Status.AvailabilityMode = nil
	}
	if resp.Cluster.ClusterEndpoint != nil {
		f4 := &svcapitypes.Endpoint{}
		if resp.Cluster.ClusterEndpoint.Address != nil {
			f4.Address = resp.Cluster.ClusterEndpoint.Address
		}
		if resp.Cluster.ClusterEndpoint.Port != nil {
			f4.Port = resp.Cluster.ClusterEndpoint.Port
		}
		ko.Status.ClusterEndpoint = f4
	} else {
		ko.Status.ClusterEndpoint = nil
	}
	if resp.Cluster.Description != nil {
		ko.Spec.Description = resp.Cluster.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.Cluster.EnginePatchVersion != nil {
		ko.Status.EnginePatchVersion = resp.Cluster.EnginePatchVersion
	} else {
		ko.Status.EnginePatchVersion = nil
	}
	if resp.Cluster.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.Cluster.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.Cluster.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.Cluster.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if resp.Cluster.MaintenanceWindow != nil {
		ko.Spec.MaintenanceWindow = resp.Cluster.MaintenanceWindow
	} else {
		ko.Spec.MaintenanceWindow = nil
	}
	if resp.Cluster.Name != nil {
		ko.Spec.Name = resp.Cluster.Name
	} else {
		ko.Spec.Name = nil
	}
	if resp.Cluster.NodeType != nil {
		ko.Spec.NodeType = resp.Cluster.NodeType
	} else {
		ko.Spec.NodeType = nil
	}
	if resp.Cluster.NumberOfShards != nil {
		ko.Status.NumberOfShards = resp.Cluster.NumberOfShards
	} else {
		ko.Status.NumberOfShards = nil
	}
	if resp.Cluster.ParameterGroupName != nil {
		ko.Spec.ParameterGroupName = resp.Cluster.ParameterGroupName
	} else {
		ko.Spec.ParameterGroupName = nil
	}
	if resp.Cluster.ParameterGroupStatus != nil {
		ko.Status.ParameterGroupStatus = resp.Cluster.ParameterGroupStatus
	} else {
		ko.Status.ParameterGroupStatus = nil
	}
	if resp.Cluster.PendingUpdates != nil {
		f15 := &svcapitypes.ClusterPendingUpdates{}
		if resp.Cluster.PendingUpdates.ACLs != nil {
			f15f0 := &svcapitypes.ACLsUpdateStatus{}
			if resp.Cluster.PendingUpdates.ACLs.ACLToApply != nil {
				f15f0.ACLToApply = resp.Cluster.PendingUpdates.ACLs.ACLToApply
			}
			f15.ACLs = f15f0
		}
		if resp.Cluster.PendingUpdates.Resharding != nil {
			f15f1 := &svcapitypes.ReshardingStatus{}
			if resp.Cluster.PendingUpdates.Resharding.SlotMigration != nil {
				f15f1f0 := &svcapitypes.SlotMigration{}
				if resp.Cluster.PendingUpdates.Resharding.SlotMigration.ProgressPercentage != nil {
					f15f1f0.ProgressPercentage = resp.Cluster.PendingUpdates.Resharding.SlotMigration.ProgressPercentage
				}
				f15f1.SlotMigration = f15f1f0
			}
			f15.Resharding = f15f1
		}
		if resp.Cluster.PendingUpdates.ServiceUpdates != nil {
			f15f2 := []*svcapitypes.PendingModifiedServiceUpdate{}
			for _, f15f2iter := range resp.Cluster.PendingUpdates.ServiceUpdates {
				f15f2elem := &svcapitypes.PendingModifiedServiceUpdate{}
				if f15f2iter.ServiceUpdateName != nil {
					f15f2elem.ServiceUpdateName = f15f2iter.ServiceUpdateName
				}
				if f15f2iter.Status != nil {
					f15f2elem.Status = f15f2iter.Status
				}
				f15f2 = append(f15f2, f15f2elem)
			}
			f15.ServiceUpdates = f15f2
		}
		ko.Status.PendingUpdates = f15
	} else {
		ko.Status.PendingUpdates = nil
	}
	if resp.Cluster.SecurityGroups != nil {
		f16 := []*svcapitypes.SecurityGroupMembership{}
		for _, f16iter := range resp.Cluster.SecurityGroups {
			f16elem := &svcapitypes.SecurityGroupMembership{}
			if f16iter.SecurityGroupId != nil {
				f16elem.SecurityGroupID = f16iter.SecurityGroupId
			}
			if f16iter.Status != nil {
				f16elem.Status = f16iter.Status
			}
			f16 = append(f16, f16elem)
		}
		ko.Status.SecurityGroups = f16
	} else {
		ko.Status.SecurityGroups = nil
	}
	if resp.Cluster.Shards != nil {
		f17 := []*svcapitypes.Shard{}
		for _, f17iter := range resp.Cluster.Shards {
			f17elem := &svcapitypes.Shard{}
			if f17iter.Name != nil {
				f17elem.Name = f17iter.Name
			}
			if f17iter.Nodes != nil {
				f17elemf1 := []*svcapitypes.Node{}
				for _, f17elemf1iter := range f17iter.Nodes {
					f17elemf1elem := &svcapitypes.Node{}
					if f17elemf1iter.AvailabilityZone != nil {
						f17elemf1elem.AvailabilityZone = f17elemf1iter.AvailabilityZone
					}
					if f17elemf1iter.CreateTime != nil {
						f17elemf1elem.CreateTime = &metav1.Time{*f17elemf1iter.CreateTime}
					}
					if f17elemf1iter.Endpoint != nil {
						f17elemf1elemf2 := &svcapitypes.Endpoint{}
						if f17elemf1iter.Endpoint.Address != nil {
							f17elemf1elemf2.Address = f17elemf1iter.Endpoint.Address
						}
						if f17elemf1iter.Endpoint.Port != nil {
							f17elemf1elemf2.Port = f17elemf1iter.Endpoint.Port
						}
						f17elemf1elem.Endpoint = f17elemf1elemf2
					}
					if f17elemf1iter.Name != nil {
						f17elemf1elem.Name = f17elemf1iter.Name
					}
					if f17elemf1iter.Status != nil {
						f17elemf1elem.Status = f17elemf1iter.Status
					}
					f17elemf1 = append(f17elemf1, f17elemf1elem)
				}
				f17elem.Nodes = f17elemf1
			}
			if f17iter.NumberOfNodes != nil {
				f17elem.NumberOfNodes = f17iter.NumberOfNodes
			}
			if f17iter.Slots != nil {
				f17elem.Slots = f17iter.Slots
			}
			if f17iter.Status != nil {
				f17elem.Status = f17iter.Status
			}
			f17 = append(f17, f17elem)
		}
		ko.Status.Shards = f17
	} else {
		ko.Status.Shards = nil
	}
	if resp.Cluster.SnapshotRetentionLimit != nil {
		ko.Spec.SnapshotRetentionLimit = resp.Cluster.SnapshotRetentionLimit
	} else {
		ko.Spec.SnapshotRetentionLimit = nil
	}
	if resp.Cluster.SnapshotWindow != nil {
		ko.Spec.SnapshotWindow = resp.Cluster.SnapshotWindow
	} else {
		ko.Spec.SnapshotWindow = nil
	}
	if resp.Cluster.SnsTopicArn != nil {
		ko.Spec.SnsTopicARN = resp.Cluster.SnsTopicArn
	} else {
		ko.Spec.SnsTopicARN = nil
	}
	if resp.Cluster.SnsTopicStatus != nil {
		ko.Status.SnsTopicStatus = resp.Cluster.SnsTopicStatus
	} else {
		ko.Status.SnsTopicStatus = nil
	}
	if resp.Cluster.Status != nil {
		ko.Status.Status = resp.Cluster.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.Cluster.SubnetGroupName != nil {
		ko.Spec.SubnetGroupName = resp.Cluster.SubnetGroupName
	} else {
		ko.Spec.SubnetGroupName = nil
	}
	if resp.Cluster.TLSEnabled != nil {
		ko.Spec.TLSEnabled = resp.Cluster.TLSEnabled
	} else {
		ko.Spec.TLSEnabled = nil
	}

	rm.setStatusDefaults(ko)
	ko, err = rm.setShardDetails(ctx, desired, ko)

	if err != nil {
		return nil, err
	}

	// Do not perform spec patching as these fields eventually get updated
	ko.Spec.NumShards = desired.ko.Spec.NumShards
	ko.Spec.NumReplicasPerShard = desired.ko.Spec.NumReplicasPerShard
	ko.Spec.ACLName = desired.ko.Spec.ACLName
	ko.Spec.NodeType = desired.ko.Spec.NodeType
	ko.Spec.EngineVersion = desired.ko.Spec.EngineVersion

	// Update the annotations to handle async rollback
	rm.setNodeTypeAnnotation(input.NodeType, ko)
	if input.ReplicaConfiguration != nil {
		rm.setNumReplicasPerShardAnnotation(input.ReplicaConfiguration.ReplicaCount, ko)
	}
	if input.ShardConfiguration != nil {
		rm.setNumShardAnnotation(input.ShardConfiguration.ShardCount, ko)
	}
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateClusterInput, error) {
	res := &svcsdk.UpdateClusterInput{}

	if r.ko.Spec.ACLName != nil {
		res.SetACLName(*r.ko.Spec.ACLName)
	}
	if r.ko.Spec.Name != nil {
		res.SetClusterName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.MaintenanceWindow != nil {
		res.SetMaintenanceWindow(*r.ko.Spec.MaintenanceWindow)
	}
	if r.ko.Spec.NodeType != nil {
		res.SetNodeType(*r.ko.Spec.NodeType)
	}
	if r.ko.Spec.ParameterGroupName != nil {
		res.SetParameterGroupName(*r.ko.Spec.ParameterGroupName)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f8 := []*string{}
		for _, f8iter := range r.ko.Spec.SecurityGroupIDs {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetSecurityGroupIds(f8)
	}
	if r.ko.Spec.SnapshotRetentionLimit != nil {
		res.SetSnapshotRetentionLimit(*r.ko.Spec.SnapshotRetentionLimit)
	}
	if r.ko.Spec.SnapshotWindow != nil {
		res.SetSnapshotWindow(*r.ko.Spec.SnapshotWindow)
	}
	if r.ko.Spec.SnsTopicARN != nil {
		res.SetSnsTopicArn(*r.ko.Spec.SnsTopicARN)
	}
	if r.ko.Status.SnsTopicStatus != nil {
		res.SetSnsTopicStatus(*r.ko.Status.SnsTopicStatus)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteClusterOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteCluster", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteClusterInput, error) {
	res := &svcsdk.DeleteClusterInput{}

	if r.ko.Spec.Name != nil {
		res.SetClusterName(*r.ko.Spec.Name)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Cluster,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ClusterAlreadyExistsFault",
		"ClusterQuotaForCustomerExceededFault",
		"NodeQuotaForClusterExceededFault",
		"NodeQuotaForCustomerExceededFault",
		"InsufficientClusterCapacityFault",
		"InvalidVPCNetworkStateFault",
		"ShardsPerClusterQuotaExceededFault",
		"InvalidParameterValueException",
		"InvalidParameterCombinationException",
		"InvalidCredentialsException",
		"TagQuotaPerResourceExceeded",
		"InvalidACLStateFault",
		"InvalidNodeStateFault",
		"InvalidKMSKeyFault",
		"NoOperationFault":
		return true
	default:
		return false
	}
}
