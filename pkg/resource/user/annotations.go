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

package user

import (
	svcapitypes "github.com/aws-controllers-k8s/memorydb-controller/apis/v1alpha1"
)

const (
	// AnnotationLastRequestedAccessString is an annotation whose value is the string
	// passed in as input to either the create or modify API called most recently
	AnnotationLastRequestedAccessString = svcapitypes.AnnotationPrefix + "last-requested-access-string"
)

// setAnnotationsFields copies the desired object's annotations, populates any
// relevant fields, and sets the latest object's annotations to this newly populated map.
// Fields that are handled by custom modify implementation are not set here.
// This should only be called upon a successful create or modify call.
func (rm *resourceManager) setAnnotationsFields(
	r *resource,
	ko *svcapitypes.User,
) {
	annotations := getAnnotationsFields(r, ko)
	setLastRequestedAccessString(r, annotations)
	ko.ObjectMeta.Annotations = annotations
}

// getAnnotationsFields return the annotations map that would be used to set the fields
func getAnnotationsFields(
	r *resource,
	ko *svcapitypes.User) map[string]string {

	if ko.ObjectMeta.Annotations != nil {
		return ko.ObjectMeta.Annotations
	}

	desiredAnnotations := r.ko.ObjectMeta.GetAnnotations()
	annotations := make(map[string]string)
	for k, v := range desiredAnnotations {
		annotations[k] = v
	}

	ko.ObjectMeta.Annotations = annotations
	return annotations
}

// setLastRequestedAccessString copies desired.Spec.AccessString into the annotation
// of the object.
func setLastRequestedAccessString(
	r *resource,
	annotations map[string]string,
) {
	if r.ko.Spec.AccessString != nil {
		annotations[AnnotationLastRequestedAccessString] = *r.ko.Spec.AccessString
	}
}
