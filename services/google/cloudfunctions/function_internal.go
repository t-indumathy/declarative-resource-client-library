// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cloudfunctions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Function) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"SourceArchiveUrl", "SourceRepository"}, r.SourceArchiveUrl, r.SourceRepository); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"EventTrigger", "HttpsTrigger"}, r.EventTrigger, r.HttpsTrigger); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Region, "Region"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SourceRepository) {
		if err := r.SourceRepository.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HttpsTrigger) {
		if err := r.HttpsTrigger.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EventTrigger) {
		if err := r.EventTrigger.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FunctionSourceRepository) validate() error {
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *FunctionHttpsTrigger) validate() error {
	return nil
}
func (r *FunctionEventTrigger) validate() error {
	if err := dcl.Required(r, "eventType"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resource"); err != nil {
		return err
	}
	return nil
}

func functionGetURL(userBasePath string, r *Function) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil
}

func functionListURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil

}

func functionCreateURL(userBasePath, project, region string) (string, error) {
	params := map[string]interface{}{
		"project": project,
		"region":  region,
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil

}

func functionDeleteURL(userBasePath string, r *Function) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"region":  dcl.ValueOrEmptyString(r.Region),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, params), nil
}

// functionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type functionApiOperation interface {
	do(context.Context, *Function, *Client) error
}

// newUpdateFunctionUpdateRequest creates a request for an
// Function resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFunctionUpdateRequest(ctx context.Context, f *Function, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		req["runtime"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		req["timeout"] = v
	}
	if v := f.AvailableMemoryMb; !dcl.IsEmptyValueIndirect(v) {
		req["availableMemoryMb"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.EnvironmentVariables; !dcl.IsEmptyValueIndirect(v) {
		req["environmentVariables"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		req["maxInstances"] = v
	}
	if v := f.VPCConnectorEgressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; !dcl.IsEmptyValueIndirect(v) {
		req["ingressSettings"] = v
	}
	return req, nil
}

// marshalUpdateFunctionUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateFunctionUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFunctionUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFunctionUpdateOperation) do(ctx context.Context, r *Function, c *Client) error {
	_, err := c.GetFunction(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}
	mask := strings.Join([]string{"description", "runtime", "timeout", "availableMemoryMb", "labels", "environmentVariables", "maxInstances", "vpcConnectorEgressSettings", "ingressSettings"}, ",")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateFunctionUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFunctionUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listFunctionRaw(ctx context.Context, project, region, pageToken string, pageSize int32) ([]byte, error) {
	u, err := functionListURL(c.Config.BasePath, project, region)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FunctionMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listFunctionOperation struct {
	Functions []map[string]interface{} `json:"functions"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listFunction(ctx context.Context, project, region, pageToken string, pageSize int32) ([]*Function, string, error) {
	b, err := c.listFunctionRaw(ctx, project, region, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFunctionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Function
	for _, v := range m.Functions {
		res := flattenFunction(c, v)
		res.Project = &project
		res.Region = &region
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllFunction(ctx context.Context, f func(*Function) bool, resources []*Function) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFunction(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteFunctionOperation struct{}

func (op *deleteFunctionOperation) do(ctx context.Context, r *Function, c *Client) error {

	_, err := c.GetFunction(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Function not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetFunction checking for existence. error: %v", err)
		return err
	}

	u, err := functionDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetFunction(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFunctionOperation struct {
	response map[string]interface{}
}

func (op *createFunctionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFunctionOperation) do(ctx context.Context, r *Function, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, region := r.createFields()
	u, err := functionCreateURL(c.Config.BasePath, project, region)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://cloudfunctions.googleapis.com/v1/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetFunction(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFunctionRaw(ctx context.Context, r *Function) ([]byte, error) {

	u, err := functionGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) functionDiffsForRawDesired(ctx context.Context, rawDesired *Function, opts ...dcl.ApplyOption) (initial, desired *Function, diffs []functionDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Function
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Function); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Function, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFunction(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Function resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Function resource: %v", err)
		}
		c.Config.Logger.Info("Found that Function resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFunctionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Function: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Function: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFunctionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Function: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFunctionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Function: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFunction(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFunctionInitialState(rawInitial, rawDesired *Function) (*Function, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if dcl.IsZeroValue(rawInitial.SourceArchiveUrl) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.SourceRepository) {
			rawInitial.SourceArchiveUrl = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawInitial.SourceRepository) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.SourceArchiveUrl) {
			rawInitial.SourceRepository = EmptyFunctionSourceRepository
		}
	}

	if dcl.IsZeroValue(rawInitial.EventTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.HttpsTrigger) {
			rawInitial.EventTrigger = EmptyFunctionEventTrigger
		}
	}

	if dcl.IsZeroValue(rawInitial.HttpsTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawInitial.EventTrigger) {
			rawInitial.HttpsTrigger = EmptyFunctionHttpsTrigger
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFunctionDesiredState(rawDesired, rawInitial *Function, opts ...dcl.ApplyOption) (*Function, error) {

	if dcl.IsZeroValue(rawDesired.SourceArchiveUrl) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.SourceRepository) {
			rawDesired.SourceArchiveUrl = dcl.String("")
		}
	}

	if dcl.IsZeroValue(rawDesired.SourceRepository) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.SourceArchiveUrl) {
			rawDesired.SourceRepository = EmptyFunctionSourceRepository
		}
	}

	if dcl.IsZeroValue(rawDesired.EventTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.HttpsTrigger) {
			rawDesired.EventTrigger = EmptyFunctionEventTrigger
		}
	}

	if dcl.IsZeroValue(rawDesired.HttpsTrigger) {
		// check if anything else is set
		if dcl.AnySet(rawDesired.EventTrigger) {
			rawDesired.HttpsTrigger = EmptyFunctionHttpsTrigger
		}
	}

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SourceRepository = canonicalizeFunctionSourceRepository(rawDesired.SourceRepository, nil, opts...)
		rawDesired.HttpsTrigger = canonicalizeFunctionHttpsTrigger(rawDesired.HttpsTrigger, nil, opts...)
		rawDesired.EventTrigger = canonicalizeFunctionEventTrigger(rawDesired.EventTrigger, nil, opts...)

		return rawDesired, nil
	}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	if dcl.StringCanonicalize(rawDesired.SourceArchiveUrl, rawInitial.SourceArchiveUrl) {
		rawDesired.SourceArchiveUrl = rawInitial.SourceArchiveUrl
	}
	rawDesired.SourceRepository = canonicalizeFunctionSourceRepository(rawDesired.SourceRepository, rawInitial.SourceRepository, opts...)
	rawDesired.HttpsTrigger = canonicalizeFunctionHttpsTrigger(rawDesired.HttpsTrigger, rawInitial.HttpsTrigger, opts...)
	rawDesired.EventTrigger = canonicalizeFunctionEventTrigger(rawDesired.EventTrigger, rawInitial.EventTrigger, opts...)
	if dcl.StringCanonicalize(rawDesired.EntryPoint, rawInitial.EntryPoint) {
		rawDesired.EntryPoint = rawInitial.EntryPoint
	}
	if dcl.StringCanonicalize(rawDesired.Runtime, rawInitial.Runtime) {
		rawDesired.Runtime = rawInitial.Runtime
	}
	if dcl.StringCanonicalize(rawDesired.Timeout, rawInitial.Timeout) {
		rawDesired.Timeout = rawInitial.Timeout
	}
	if dcl.IsZeroValue(rawDesired.AvailableMemoryMb) {
		rawDesired.AvailableMemoryMb = rawInitial.AvailableMemoryMb
	}
	if dcl.NameToSelfLink(rawDesired.ServiceAccountEmail, rawInitial.ServiceAccountEmail) {
		rawDesired.ServiceAccountEmail = rawInitial.ServiceAccountEmail
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	if dcl.IsZeroValue(rawDesired.EnvironmentVariables) {
		rawDesired.EnvironmentVariables = rawInitial.EnvironmentVariables
	}
	if dcl.NameToSelfLink(rawDesired.Network, rawInitial.Network) {
		rawDesired.Network = rawInitial.Network
	}
	if dcl.IsZeroValue(rawDesired.MaxInstances) {
		rawDesired.MaxInstances = rawInitial.MaxInstances
	}
	if dcl.NameToSelfLink(rawDesired.VPCConnector, rawInitial.VPCConnector) {
		rawDesired.VPCConnector = rawInitial.VPCConnector
	}
	if dcl.IsZeroValue(rawDesired.VPCConnectorEgressSettings) {
		rawDesired.VPCConnectorEgressSettings = rawInitial.VPCConnectorEgressSettings
	}
	if dcl.IsZeroValue(rawDesired.IngressSettings) {
		rawDesired.IngressSettings = rawInitial.IngressSettings
	}
	if dcl.NameToSelfLink(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}

	return rawDesired, nil
}

func canonicalizeFunctionNewState(c *Client, rawNew, rawDesired *Function) (*Function, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceArchiveUrl) && dcl.IsEmptyValueIndirect(rawDesired.SourceArchiveUrl) {
		rawNew.SourceArchiveUrl = rawDesired.SourceArchiveUrl
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceArchiveUrl, rawNew.SourceArchiveUrl) {
			rawNew.SourceArchiveUrl = rawDesired.SourceArchiveUrl
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceRepository) && dcl.IsEmptyValueIndirect(rawDesired.SourceRepository) {
		rawNew.SourceRepository = rawDesired.SourceRepository
	} else {
		rawNew.SourceRepository = canonicalizeNewFunctionSourceRepository(c, rawDesired.SourceRepository, rawNew.SourceRepository)
	}

	if dcl.IsEmptyValueIndirect(rawNew.HttpsTrigger) && dcl.IsEmptyValueIndirect(rawDesired.HttpsTrigger) {
		rawNew.HttpsTrigger = rawDesired.HttpsTrigger
	} else {
		rawNew.HttpsTrigger = canonicalizeNewFunctionHttpsTrigger(c, rawDesired.HttpsTrigger, rawNew.HttpsTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EventTrigger) && dcl.IsEmptyValueIndirect(rawDesired.EventTrigger) {
		rawNew.EventTrigger = rawDesired.EventTrigger
	} else {
		rawNew.EventTrigger = canonicalizeNewFunctionEventTrigger(c, rawDesired.EventTrigger, rawNew.EventTrigger)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EntryPoint) && dcl.IsEmptyValueIndirect(rawDesired.EntryPoint) {
		rawNew.EntryPoint = rawDesired.EntryPoint
	} else {
		if dcl.StringCanonicalize(rawDesired.EntryPoint, rawNew.EntryPoint) {
			rawNew.EntryPoint = rawDesired.EntryPoint
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Runtime) && dcl.IsEmptyValueIndirect(rawDesired.Runtime) {
		rawNew.Runtime = rawDesired.Runtime
	} else {
		if dcl.StringCanonicalize(rawDesired.Runtime, rawNew.Runtime) {
			rawNew.Runtime = rawDesired.Runtime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Timeout) && dcl.IsEmptyValueIndirect(rawDesired.Timeout) {
		rawNew.Timeout = rawDesired.Timeout
	} else {
		if dcl.StringCanonicalize(rawDesired.Timeout, rawNew.Timeout) {
			rawNew.Timeout = rawDesired.Timeout
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AvailableMemoryMb) && dcl.IsEmptyValueIndirect(rawDesired.AvailableMemoryMb) {
		rawNew.AvailableMemoryMb = rawDesired.AvailableMemoryMb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmail) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmail) {
		rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
	} else {
		if dcl.NameToSelfLink(rawDesired.ServiceAccountEmail, rawNew.ServiceAccountEmail) {
			rawNew.ServiceAccountEmail = rawDesired.ServiceAccountEmail
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
		if dcl.StringCanonicalize(rawDesired.UpdateTime, rawNew.UpdateTime) {
			rawNew.UpdateTime = rawDesired.UpdateTime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VersionId) && dcl.IsEmptyValueIndirect(rawDesired.VersionId) {
		rawNew.VersionId = rawDesired.VersionId
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnvironmentVariables) && dcl.IsEmptyValueIndirect(rawDesired.EnvironmentVariables) {
		rawNew.EnvironmentVariables = rawDesired.EnvironmentVariables
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
		if dcl.NameToSelfLink(rawDesired.Network, rawNew.Network) {
			rawNew.Network = rawDesired.Network
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxInstances) && dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) {
		rawNew.MaxInstances = rawDesired.MaxInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnector) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnector) {
		rawNew.VPCConnector = rawDesired.VPCConnector
	} else {
		if dcl.NameToSelfLink(rawDesired.VPCConnector, rawNew.VPCConnector) {
			rawNew.VPCConnector = rawDesired.VPCConnector
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.VPCConnectorEgressSettings) && dcl.IsEmptyValueIndirect(rawDesired.VPCConnectorEgressSettings) {
		rawNew.VPCConnectorEgressSettings = rawDesired.VPCConnectorEgressSettings
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IngressSettings) && dcl.IsEmptyValueIndirect(rawDesired.IngressSettings) {
		rawNew.IngressSettings = rawDesired.IngressSettings
	} else {
	}

	rawNew.Region = rawDesired.Region

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeFunctionSourceRepository(des, initial *FunctionSourceRepository, opts ...dcl.ApplyOption) *FunctionSourceRepository {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		des.Url = initial.Url
	}

	return des
}

func canonicalizeNewFunctionSourceRepository(c *Client, des, nw *FunctionSourceRepository) *FunctionSourceRepository {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	if dcl.StringCanonicalize(des.DeployedUrl, nw.DeployedUrl) {
		nw.DeployedUrl = des.DeployedUrl
	}

	return nw
}

func canonicalizeNewFunctionSourceRepositorySet(c *Client, des, nw []FunctionSourceRepository) []FunctionSourceRepository {
	if des == nil {
		return nw
	}
	var reorderedNew []FunctionSourceRepository
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFunctionSourceRepositoryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFunctionSourceRepositorySlice(c *Client, des, nw []FunctionSourceRepository) []FunctionSourceRepository {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionSourceRepository
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionSourceRepository(c, &d, &n))
	}

	return items
}

func canonicalizeFunctionHttpsTrigger(des, initial *FunctionHttpsTrigger, opts ...dcl.ApplyOption) *FunctionHttpsTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	return des
}

func canonicalizeNewFunctionHttpsTrigger(c *Client, des, nw *FunctionHttpsTrigger) *FunctionHttpsTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewFunctionHttpsTriggerSet(c *Client, des, nw []FunctionHttpsTrigger) []FunctionHttpsTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []FunctionHttpsTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFunctionHttpsTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFunctionHttpsTriggerSlice(c *Client, des, nw []FunctionHttpsTrigger) []FunctionHttpsTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionHttpsTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionHttpsTrigger(c, &d, &n))
	}

	return items
}

func canonicalizeFunctionEventTrigger(des, initial *FunctionEventTrigger, opts ...dcl.ApplyOption) *FunctionEventTrigger {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.EventType, initial.EventType) || dcl.IsZeroValue(des.EventType) {
		des.EventType = initial.EventType
	}
	if dcl.NameToSelfLink(des.Resource, initial.Resource) || dcl.IsZeroValue(des.Resource) {
		des.Resource = initial.Resource
	}
	if dcl.StringCanonicalize(des.Service, initial.Service) || dcl.IsZeroValue(des.Service) {
		des.Service = initial.Service
	}
	if dcl.BoolCanonicalize(des.FailurePolicy, initial.FailurePolicy) || dcl.IsZeroValue(des.FailurePolicy) {
		des.FailurePolicy = initial.FailurePolicy
	}

	return des
}

func canonicalizeNewFunctionEventTrigger(c *Client, des, nw *FunctionEventTrigger) *FunctionEventTrigger {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.EventType, nw.EventType) {
		nw.EventType = des.EventType
	}
	if dcl.NameToSelfLink(des.Resource, nw.Resource) {
		nw.Resource = des.Resource
	}
	if dcl.StringCanonicalize(des.Service, nw.Service) {
		nw.Service = des.Service
	}
	if dcl.BoolCanonicalize(des.FailurePolicy, nw.FailurePolicy) {
		nw.FailurePolicy = des.FailurePolicy
	}

	return nw
}

func canonicalizeNewFunctionEventTriggerSet(c *Client, des, nw []FunctionEventTrigger) []FunctionEventTrigger {
	if des == nil {
		return nw
	}
	var reorderedNew []FunctionEventTrigger
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareFunctionEventTriggerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewFunctionEventTriggerSlice(c *Client, des, nw []FunctionEventTrigger) []FunctionEventTrigger {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FunctionEventTrigger
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFunctionEventTrigger(c, &d, &n))
	}

	return items
}

type functionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         functionApiOperation
	Diffs            []*dcl.FieldDiff
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffFunction(c *Client, desired, actual *Function, opts ...dcl.ApplyOption) ([]functionDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []functionDiff
	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SourceArchiveUrl, actual.SourceArchiveUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceArchiveUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.SourceRepository, actual.SourceRepository, dcl.Info{ObjectFunction: compareFunctionSourceRepositoryNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceRepository")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.HttpsTrigger, actual.HttpsTrigger, dcl.Info{ObjectFunction: compareFunctionHttpsTriggerNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsTrigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.EventTrigger, actual.EventTrigger, dcl.Info{ObjectFunction: compareFunctionEventTriggerNewStyle, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EventTrigger")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.EntryPoint, actual.EntryPoint, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EntryPoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Runtime, actual.Runtime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Runtime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Timeout, actual.Timeout, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Timeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.AvailableMemoryMb, actual.AvailableMemoryMb, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("AvailableMemoryMb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmail, actual.ServiceAccountEmail, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.VersionId, actual.VersionId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VersionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.EnvironmentVariables, actual.EnvironmentVariables, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("EnvironmentVariables")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.Info{OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.VPCConnector, actual.VPCConnector, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VPCConnector")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.VPCConnectorEgressSettings, actual.VPCConnectorEgressSettings, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("VPCConnectorEgressSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.IngressSettings, actual.IngressSettings, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFunctionUpdateOperation")}, fn.AddNest("IngressSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Region")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)

		dsOld, err := convertFieldDiffToFunctionDiff(ds, opts...)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, dsOld...)
	}

	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []functionDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareFunctionSourceRepositoryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionSourceRepository)
	if !ok {
		desiredNotPointer, ok := d.(FunctionSourceRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionSourceRepository or *FunctionSourceRepository", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionSourceRepository)
	if !ok {
		actualNotPointer, ok := a.(FunctionSourceRepository)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionSourceRepository", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployedUrl, actual.DeployedUrl, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeployedUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFunctionHttpsTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionHttpsTrigger)
	if !ok {
		desiredNotPointer, ok := d.(FunctionHttpsTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionHttpsTrigger or *FunctionHttpsTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionHttpsTrigger)
	if !ok {
		actualNotPointer, ok := a.(FunctionHttpsTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionHttpsTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFunctionEventTriggerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FunctionEventTrigger)
	if !ok {
		desiredNotPointer, ok := d.(FunctionEventTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionEventTrigger or *FunctionEventTrigger", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FunctionEventTrigger)
	if !ok {
		actualNotPointer, ok := a.(FunctionEventTrigger)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FunctionEventTrigger", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EventType, actual.EventType, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EventType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resource, actual.Resource, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Service, actual.Service, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Service")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FailurePolicy, actual.FailurePolicy, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FailurePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Function) urlNormalized() *Function {
	normalized := dcl.Copy(*r).(Function)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceArchiveUrl = dcl.SelfLinkToName(r.SourceArchiveUrl)
	normalized.EntryPoint = dcl.SelfLinkToName(r.EntryPoint)
	normalized.Runtime = dcl.SelfLinkToName(r.Runtime)
	normalized.Timeout = dcl.SelfLinkToName(r.Timeout)
	normalized.ServiceAccountEmail = dcl.SelfLinkToName(r.ServiceAccountEmail)
	normalized.UpdateTime = dcl.SelfLinkToName(r.UpdateTime)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.VPCConnector = dcl.SelfLinkToName(r.VPCConnector)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Function) getFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Function) createFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region)
}

func (r *Function) deleteFields() (string, string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Region), dcl.ValueOrEmptyString(n.Name)
}

func (r *Function) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"region":  dcl.ValueOrEmptyString(n.Region),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{region}}/functions/{{name}}", "https://cloudfunctions.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Function resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Function) marshal(c *Client) ([]byte, error) {
	m, err := expandFunction(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Function: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFunction decodes JSON responses into the Function resource schema.
func unmarshalFunction(b []byte, c *Client) (*Function, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFunction(m, c)
}

func unmarshalMapFunction(m map[string]interface{}, c *Client) (*Function, error) {

	return flattenFunction(c, m), nil
}

// expandFunction expands Function into a JSON request object.
func expandFunction(c *Client, f *Function) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/functions/%s", f.Name, f.Project, f.Region, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.SourceArchiveUrl; !dcl.IsEmptyValueIndirect(v) {
		m["sourceArchiveUrl"] = v
	}
	if v, err := expandFunctionSourceRepository(c, f.SourceRepository); err != nil {
		return nil, fmt.Errorf("error expanding SourceRepository into sourceRepository: %w", err)
	} else if v != nil {
		m["sourceRepository"] = v
	}
	if v, err := expandFunctionHttpsTrigger(c, f.HttpsTrigger); err != nil {
		return nil, fmt.Errorf("error expanding HttpsTrigger into httpsTrigger: %w", err)
	} else {
		m["httpsTrigger"] = v
	}
	if v, err := expandFunctionEventTrigger(c, f.EventTrigger); err != nil {
		return nil, fmt.Errorf("error expanding EventTrigger into eventTrigger: %w", err)
	} else if v != nil {
		m["eventTrigger"] = v
	}
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.EntryPoint; !dcl.IsEmptyValueIndirect(v) {
		m["entryPoint"] = v
	}
	if v := f.Runtime; !dcl.IsEmptyValueIndirect(v) {
		m["runtime"] = v
	}
	if v := f.Timeout; !dcl.IsEmptyValueIndirect(v) {
		m["timeout"] = v
	}
	if v := f.AvailableMemoryMb; !dcl.IsEmptyValueIndirect(v) {
		m["availableMemoryMb"] = v
	}
	if v := f.ServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmail"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.VersionId; !dcl.IsEmptyValueIndirect(v) {
		m["versionId"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.EnvironmentVariables; !dcl.IsEmptyValueIndirect(v) {
		m["environmentVariables"] = v
	}
	if v := f.Network; !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.MaxInstances; !dcl.IsEmptyValueIndirect(v) {
		m["maxInstances"] = v
	}
	if v := f.VPCConnector; !dcl.IsEmptyValueIndirect(v) {
		m["vpcConnector"] = v
	}
	if v := f.VPCConnectorEgressSettings; !dcl.IsEmptyValueIndirect(v) {
		m["vpcConnectorEgressSettings"] = v
	}
	if v := f.IngressSettings; !dcl.IsEmptyValueIndirect(v) {
		m["ingressSettings"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Region into region: %w", err)
	} else if v != nil {
		m["region"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenFunction flattens Function from a JSON request object into the
// Function type.
func flattenFunction(c *Client, i interface{}) *Function {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Function{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.SourceArchiveUrl = dcl.FlattenString(m["sourceArchiveUrl"])
	res.SourceRepository = flattenFunctionSourceRepository(c, m["sourceRepository"])
	res.HttpsTrigger = flattenFunctionHttpsTrigger(c, m["httpsTrigger"])
	res.EventTrigger = flattenFunctionEventTrigger(c, m["eventTrigger"])
	res.Status = flattenFunctionStatusEnum(m["status"])
	res.EntryPoint = dcl.FlattenString(m["entryPoint"])
	res.Runtime = dcl.FlattenString(m["runtime"])
	res.Timeout = dcl.FlattenString(m["timeout"])
	res.AvailableMemoryMb = dcl.FlattenInteger(m["availableMemoryMb"])
	res.ServiceAccountEmail = dcl.FlattenString(m["serviceAccountEmail"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.VersionId = dcl.FlattenInteger(m["versionId"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.EnvironmentVariables = dcl.FlattenKeyValuePairs(m["environmentVariables"])
	res.Network = dcl.FlattenString(m["network"])
	res.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	res.VPCConnector = dcl.FlattenString(m["vpcConnector"])
	res.VPCConnectorEgressSettings = flattenFunctionVPCConnectorEgressSettingsEnum(m["vpcConnectorEgressSettings"])
	res.IngressSettings = flattenFunctionIngressSettingsEnum(m["ingressSettings"])
	res.Region = dcl.FlattenString(m["region"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// expandFunctionSourceRepositoryMap expands the contents of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepositoryMap(c *Client, f map[string]FunctionSourceRepository) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionSourceRepository(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionSourceRepositorySlice expands the contents of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepositorySlice(c *Client, f []FunctionSourceRepository) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionSourceRepository(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionSourceRepositoryMap flattens the contents of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepositoryMap(c *Client, i interface{}) map[string]FunctionSourceRepository {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionSourceRepository{}
	}

	if len(a) == 0 {
		return map[string]FunctionSourceRepository{}
	}

	items := make(map[string]FunctionSourceRepository)
	for k, item := range a {
		items[k] = *flattenFunctionSourceRepository(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFunctionSourceRepositorySlice flattens the contents of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepositorySlice(c *Client, i interface{}) []FunctionSourceRepository {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionSourceRepository{}
	}

	if len(a) == 0 {
		return []FunctionSourceRepository{}
	}

	items := make([]FunctionSourceRepository, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionSourceRepository(c, item.(map[string]interface{})))
	}

	return items
}

// expandFunctionSourceRepository expands an instance of FunctionSourceRepository into a JSON
// request object.
func expandFunctionSourceRepository(c *Client, f *FunctionSourceRepository) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v := f.DeployedUrl; !dcl.IsEmptyValueIndirect(v) {
		m["deployedUrl"] = v
	}

	return m, nil
}

// flattenFunctionSourceRepository flattens an instance of FunctionSourceRepository from a JSON
// response object.
func flattenFunctionSourceRepository(c *Client, i interface{}) *FunctionSourceRepository {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionSourceRepository{}
	r.Url = dcl.FlattenString(m["url"])
	r.DeployedUrl = dcl.FlattenString(m["deployedUrl"])

	return r
}

// expandFunctionHttpsTriggerMap expands the contents of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTriggerMap(c *Client, f map[string]FunctionHttpsTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionHttpsTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionHttpsTriggerSlice expands the contents of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTriggerSlice(c *Client, f []FunctionHttpsTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionHttpsTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionHttpsTriggerMap flattens the contents of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTriggerMap(c *Client, i interface{}) map[string]FunctionHttpsTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return map[string]FunctionHttpsTrigger{}
	}

	items := make(map[string]FunctionHttpsTrigger)
	for k, item := range a {
		items[k] = *flattenFunctionHttpsTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFunctionHttpsTriggerSlice flattens the contents of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTriggerSlice(c *Client, i interface{}) []FunctionHttpsTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionHttpsTrigger{}
	}

	if len(a) == 0 {
		return []FunctionHttpsTrigger{}
	}

	items := make([]FunctionHttpsTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionHttpsTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandFunctionHttpsTrigger expands an instance of FunctionHttpsTrigger into a JSON
// request object.
func expandFunctionHttpsTrigger(c *Client, f *FunctionHttpsTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenFunctionHttpsTrigger flattens an instance of FunctionHttpsTrigger from a JSON
// response object.
func flattenFunctionHttpsTrigger(c *Client, i interface{}) *FunctionHttpsTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionHttpsTrigger{}
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandFunctionEventTriggerMap expands the contents of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTriggerMap(c *Client, f map[string]FunctionEventTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFunctionEventTrigger(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFunctionEventTriggerSlice expands the contents of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTriggerSlice(c *Client, f []FunctionEventTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFunctionEventTrigger(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFunctionEventTriggerMap flattens the contents of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTriggerMap(c *Client, i interface{}) map[string]FunctionEventTrigger {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FunctionEventTrigger{}
	}

	if len(a) == 0 {
		return map[string]FunctionEventTrigger{}
	}

	items := make(map[string]FunctionEventTrigger)
	for k, item := range a {
		items[k] = *flattenFunctionEventTrigger(c, item.(map[string]interface{}))
	}

	return items
}

// flattenFunctionEventTriggerSlice flattens the contents of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTriggerSlice(c *Client, i interface{}) []FunctionEventTrigger {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionEventTrigger{}
	}

	if len(a) == 0 {
		return []FunctionEventTrigger{}
	}

	items := make([]FunctionEventTrigger, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionEventTrigger(c, item.(map[string]interface{})))
	}

	return items
}

// expandFunctionEventTrigger expands an instance of FunctionEventTrigger into a JSON
// request object.
func expandFunctionEventTrigger(c *Client, f *FunctionEventTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EventType; !dcl.IsEmptyValueIndirect(v) {
		m["eventType"] = v
	}
	if v := f.Resource; !dcl.IsEmptyValueIndirect(v) {
		m["resource"] = v
	}
	if v := f.Service; !dcl.IsEmptyValueIndirect(v) {
		m["service"] = v
	}
	if v, err := ExpandFunctionEventRetry(f, f.FailurePolicy); err != nil {
		return nil, fmt.Errorf("error expanding FailurePolicy into failurePolicy: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failurePolicy"] = v
	}

	return m, nil
}

// flattenFunctionEventTrigger flattens an instance of FunctionEventTrigger from a JSON
// response object.
func flattenFunctionEventTrigger(c *Client, i interface{}) *FunctionEventTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FunctionEventTrigger{}
	r.EventType = dcl.FlattenString(m["eventType"])
	r.Resource = dcl.FlattenString(m["resource"])
	r.Service = dcl.FlattenString(m["service"])
	r.FailurePolicy = FlattenFunctionEventRetry(m["failurePolicy"])

	return r
}

// flattenFunctionStatusEnumSlice flattens the contents of FunctionStatusEnum from a JSON
// response object.
func flattenFunctionStatusEnumSlice(c *Client, i interface{}) []FunctionStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionStatusEnum{}
	}

	if len(a) == 0 {
		return []FunctionStatusEnum{}
	}

	items := make([]FunctionStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionStatusEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionStatusEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionStatusEnum with the same value as that string.
func flattenFunctionStatusEnum(i interface{}) *FunctionStatusEnum {
	s, ok := i.(string)
	if !ok {
		return FunctionStatusEnumRef("")
	}

	return FunctionStatusEnumRef(s)
}

// flattenFunctionVPCConnectorEgressSettingsEnumSlice flattens the contents of FunctionVPCConnectorEgressSettingsEnum from a JSON
// response object.
func flattenFunctionVPCConnectorEgressSettingsEnumSlice(c *Client, i interface{}) []FunctionVPCConnectorEgressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionVPCConnectorEgressSettingsEnum{}
	}

	if len(a) == 0 {
		return []FunctionVPCConnectorEgressSettingsEnum{}
	}

	items := make([]FunctionVPCConnectorEgressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionVPCConnectorEgressSettingsEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionVPCConnectorEgressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionVPCConnectorEgressSettingsEnum with the same value as that string.
func flattenFunctionVPCConnectorEgressSettingsEnum(i interface{}) *FunctionVPCConnectorEgressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return FunctionVPCConnectorEgressSettingsEnumRef("")
	}

	return FunctionVPCConnectorEgressSettingsEnumRef(s)
}

// flattenFunctionIngressSettingsEnumSlice flattens the contents of FunctionIngressSettingsEnum from a JSON
// response object.
func flattenFunctionIngressSettingsEnumSlice(c *Client, i interface{}) []FunctionIngressSettingsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FunctionIngressSettingsEnum{}
	}

	if len(a) == 0 {
		return []FunctionIngressSettingsEnum{}
	}

	items := make([]FunctionIngressSettingsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFunctionIngressSettingsEnum(item.(interface{})))
	}

	return items
}

// flattenFunctionIngressSettingsEnum asserts that an interface is a string, and returns a
// pointer to a *FunctionIngressSettingsEnum with the same value as that string.
func flattenFunctionIngressSettingsEnum(i interface{}) *FunctionIngressSettingsEnum {
	s, ok := i.(string)
	if !ok {
		return FunctionIngressSettingsEnumRef("")
	}

	return FunctionIngressSettingsEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Function) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFunction(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Region == nil && ncr.Region == nil {
			c.Config.Logger.Info("Both Region fields null - considering equal.")
		} else if nr.Region == nil || ncr.Region == nil {
			c.Config.Logger.Info("Only one Region field is null - considering unequal.")
			return false
		} else if *nr.Region != *ncr.Region {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

func convertFieldDiffToFunctionDiff(fds []*dcl.FieldDiff, opts ...dcl.ApplyOption) ([]functionDiff, error) {
	var diffs []functionDiff
	for _, fd := range fds {
		for _, op := range fd.ResultingOperation {
			diff := functionDiff{Diffs: []*dcl.FieldDiff{fd}, FieldName: fd.FieldName}
			if op == "Recreate" {
				diff.RequiresRecreate = true
			} else {
				op, err := convertOpNameTofunctionApiOperation(op, opts...)
				if err != nil {
					return nil, err
				}
				diff.UpdateOp = op
			}
			diffs = append(diffs, diff)
		}
	}
	return diffs, nil
}

func convertOpNameTofunctionApiOperation(op string, opts ...dcl.ApplyOption) (functionApiOperation, error) {
	switch op {

	case "updateFunctionUpdateOperation":
		return &updateFunctionUpdateOperation{}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
