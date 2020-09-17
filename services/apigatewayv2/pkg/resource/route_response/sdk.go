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

package route_response

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.RouteResponse{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetRouteResponseWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.RouteResponseID == nil || r.ko.Spec.APIID == nil || r.ko.Spec.RouteID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetRouteResponseInput, error) {
	res := &svcsdk.GetRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetRouteResponsesInput, error) {
	res := &svcsdk.GetRouteResponsesInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateRouteResponseWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	}

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateRouteResponseInput, error) {
	res := &svcsdk.CreateRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.ResponseModels != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.ResponseModels {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseModels(f2)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f3 := map[string]*svcsdk.ParameterConstraints{}
		for f3key, f3valiter := range r.ko.Spec.ResponseParameters {
			f3val := &svcsdk.ParameterConstraints{}
			if f3valiter.Required != nil {
				f3val.SetRequired(*f3valiter.Required)
			}
			f3[f3key] = f3val
		}
		res.SetResponseParameters(f3)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Spec.RouteResponseKey != nil {
		res.SetRouteResponseKey(*r.ko.Spec.RouteResponseKey)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	input, err := rm.newUpdateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateRouteResponseWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	}

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateRouteResponseInput, error) {
	res := &svcsdk.UpdateRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.ResponseModels != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.ResponseModels {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseModels(f2)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f3 := map[string]*svcsdk.ParameterConstraints{}
		for f3key, f3valiter := range r.ko.Spec.ResponseParameters {
			f3val := &svcsdk.ParameterConstraints{}
			if f3valiter.Required != nil {
				f3val.SetRequired(*f3valiter.Required)
			}
			f3[f3key] = f3val
		}
		res.SetResponseParameters(f3)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}
	if r.ko.Spec.RouteResponseKey != nil {
		res.SetRouteResponseKey(*r.ko.Spec.RouteResponseKey)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteRouteResponseWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRouteResponseInput, error) {
	res := &svcsdk.DeleteRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}

	return res, nil
}
