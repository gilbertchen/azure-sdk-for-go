package containerregistry

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// RegistriesClient is the client for the Registries methods of the Containerregistry service.
type RegistriesClient struct {
	ManagementClient
}

// NewRegistriesClient creates an instance of the RegistriesClient client.
func NewRegistriesClient(p pipeline.Pipeline) RegistriesClient {
	return RegistriesClient{NewManagementClient(p)}
}

// CheckNameAvailability checks whether the container registry name is available for use. The name must contain only
// alphanumeric characters, be globally unique, and between 5 and 60 characters in length.
//
// registryNameCheckRequest is the object containing information for the availability request.
func (client RegistriesClient) CheckNameAvailability(ctx context.Context, registryNameCheckRequest RegistryNameCheckRequest) (*RegistryNameStatus, error) {
	if err := validate([]validation{
		{targetValue: registryNameCheckRequest,
			constraints: []constraint{{target: "registryNameCheckRequest.Name", name: null, rule: true, chain: nil},
				{target: "registryNameCheckRequest.Type", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.checkNameAvailabilityPreparer(registryNameCheckRequest)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.checkNameAvailabilityResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RegistryNameStatus), err
}

// checkNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client RegistriesClient) checkNameAvailabilityPreparer(registryNameCheckRequest RegistryNameCheckRequest) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/checkNameAvailability"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(registryNameCheckRequest)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// checkNameAvailabilityResponder handles the response to the CheckNameAvailability request.
func (client RegistriesClient) checkNameAvailabilityResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RegistryNameStatus{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// CreateOrUpdate creates or updates a container registry with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. registry is the parameters for creating or updating a container registry.
func (client RegistriesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, registry Registry) (*Registry, error) {
	if err := validate([]validation{
		{targetValue: registry,
			constraints: []constraint{{target: "registry.RegistryProperties", name: null, rule: false,
				chain: []constraint{{target: "registry.RegistryProperties.StorageAccount", name: null, rule: true,
					chain: []constraint{{target: "registry.RegistryProperties.StorageAccount.Name", name: null, rule: true, chain: nil},
						{target: "registry.RegistryProperties.StorageAccount.AccessKey", name: null, rule: true, chain: nil},
					}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, registryName, registry)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Registry), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegistriesClient) createOrUpdatePreparer(resourceGroupName string, registryName string, registry Registry) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(registry)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client RegistriesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &Registry{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes a container registry.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry.
func (client RegistriesClient) Delete(ctx context.Context, resourceGroupName string, registryName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, registryName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client RegistriesClient) deletePreparer(resourceGroupName string, registryName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client RegistriesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// GetCredentials gets the administrator login credentials for the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry.
func (client RegistriesClient) GetCredentials(ctx context.Context, resourceGroupName string, registryName string) (*RegistryCredentials, error) {
	req, err := client.getCredentialsPreparer(resourceGroupName, registryName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getCredentialsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RegistryCredentials), err
}

// getCredentialsPreparer prepares the GetCredentials request.
func (client RegistriesClient) getCredentialsPreparer(resourceGroupName string, registryName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/getCredentials"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getCredentialsResponder handles the response to the GetCredentials request.
func (client RegistriesClient) getCredentialsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RegistryCredentials{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetProperties gets the properties of the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry.
func (client RegistriesClient) GetProperties(ctx context.Context, resourceGroupName string, registryName string) (*Registry, error) {
	req, err := client.getPropertiesPreparer(resourceGroupName, registryName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Registry), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client RegistriesClient) getPropertiesPreparer(resourceGroupName string, registryName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client RegistriesClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Registry{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List lists all the available container registries under the specified subscription.
func (client RegistriesClient) List(ctx context.Context) (*RegistryListResult, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RegistryListResult), err
}

// listPreparer prepares the List request.
func (client RegistriesClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client RegistriesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RegistryListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListByResourceGroup lists all the available container registries under the specified resource group.
//
// resourceGroupName is the name of the resource group to which the container registry belongs.
func (client RegistriesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*RegistryListResult, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RegistryListResult), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client RegistriesClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client RegistriesClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RegistryListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// RegenerateCredentials regenerates the administrator login credentials for the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry.
func (client RegistriesClient) RegenerateCredentials(ctx context.Context, resourceGroupName string, registryName string) (*RegistryCredentials, error) {
	req, err := client.regenerateCredentialsPreparer(resourceGroupName, registryName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.regenerateCredentialsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RegistryCredentials), err
}

// regenerateCredentialsPreparer prepares the RegenerateCredentials request.
func (client RegistriesClient) regenerateCredentialsPreparer(resourceGroupName string, registryName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/regenerateCredentials"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// regenerateCredentialsResponder handles the response to the RegenerateCredentials request.
func (client RegistriesClient) regenerateCredentialsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RegistryCredentials{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Update updates a container registry with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. registryUpdateParameters is the parameters for updating a container registry.
func (client RegistriesClient) Update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (*Registry, error) {
	req, err := client.updatePreparer(resourceGroupName, registryName, registryUpdateParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Registry), err
}

// updatePreparer prepares the Update request.
func (client RegistriesClient) updatePreparer(resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(registryUpdateParameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateResponder handles the response to the Update request.
func (client RegistriesClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Registry{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
