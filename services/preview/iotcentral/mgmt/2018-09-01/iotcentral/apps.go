package iotcentral

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// AppsClient is the use this API to manage IoT Central Applications in your Azure subscription.
type AppsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// NewAppsClient creates an instance of the AppsClient client.
func NewAppsClient(subscriptionID string) AppsClient {
	return NewAppsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// NewAppsClientWithBaseURI creates an instance of the AppsClient client.
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return AppsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CheckNameAvailability check if an IoT Central application name is available.
// Parameters:
// operationInputs - set the name parameter in the OperationInputs structure to the name of the IoT Central
// application to check.
func (client AppsClient) CheckNameAvailability(ctx context.Context, operationInputs OperationInputs) (result AppNameAvailabilityInfo, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: operationInputs,
			Constraints: []validation.Constraint{{Target: "operationInputs.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.AppsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, operationInputs)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client AppsClient) CheckNameAvailabilityPreparer(ctx context.Context, operationInputs OperationInputs) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/checkNameAvailability", pathParameters),
		autorest.WithJSON(operationInputs),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client AppsClient) CheckNameAvailabilityResponder(resp *http.Response) (result AppNameAvailabilityInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CreateOrUpdate create or update the metadata of an IoT Central application. The usual pattern to modify a property
// is to retrieve the IoT Central application metadata and security metadata, and then combine them with the modified
// values in a new body to update the IoT Central application.
// Parameters:
// resourceGroupName - the name of the resource group that contains the IoT Central application.
// resourceName - the ARM resource name of the IoT Central application.
// app - the IoT Central application metadata and security metadata.
func (client AppsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, app App) (result AppsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: app,
			Constraints: []validation.Constraint{{Target: "app.AppProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "app.AppProperties.DisplayName", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "app.AppProperties.DisplayName", Name: validation.Pattern, Rule: `^.{1,200}$`, Chain: nil}}},
					{Target: "app.AppProperties.Subdomain", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "app.AppProperties.Subdomain", Name: validation.Pattern, Rule: `^[a-z0-9-]{1,63}$`, Chain: nil}}},
				}},
				{Target: "app.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.AppsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, app)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AppsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, app App) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", pathParameters),
		autorest.WithJSON(app),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) CreateOrUpdateSender(req *http.Request) (future AppsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AppsClient) CreateOrUpdateResponder(resp *http.Response) (result App, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// Delete delete an IoT Central application.
// Parameters:
// resourceGroupName - the name of the resource group that contains the IoT Central application.
// resourceName - the ARM resource name of the IoT Central application.
func (client AppsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result AppsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// DeletePreparer prepares the Delete request.
func (client AppsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) DeleteSender(req *http.Request) (future AppsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AppsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// Get get the metadata of an IoT Central application.
// Parameters:
// resourceGroupName - the name of the resource group that contains the IoT Central application.
// resourceName - the ARM resource name of the IoT Central application.
func (client AppsClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result App, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// GetPreparer prepares the Get request.
func (client AppsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AppsClient) GetResponder(resp *http.Response) (result App, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListByResourceGroup get all the IoT Central Applications in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the IoT Central application.
func (client AppsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result AppListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AppsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AppsClient) ListByResourceGroupResponder(resp *http.Response) (result AppListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AppsClient) listByResourceGroupNextResults(lastResults AppListResult) (result AppListResult, err error) {
	req, err := lastResults.appListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AppsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result AppListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListBySubscription get all IoT Central Applications in a subscription.
func (client AppsClient) ListBySubscription(ctx context.Context) (result AppListResultPage, err error) {
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client AppsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/IoTApps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client AppsClient) ListBySubscriptionResponder(resp *http.Response) (result AppListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client AppsClient) listBySubscriptionNextResults(lastResults AppListResult) (result AppListResult, err error) {
	req, err := lastResults.appListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client AppsClient) ListBySubscriptionComplete(ctx context.Context) (result AppListResultIterator, err error) {
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// Update update the metadata of an IoT Central application.
// Parameters:
// resourceGroupName - the name of the resource group that contains the IoT Central application.
// resourceName - the ARM resource name of the IoT Central application.
// appPatch - the IoT Central application metadata and security metadata.
func (client AppsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, appPatch AppPatch) (result AppsUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, appPatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.AppsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// UpdatePreparer prepares the Update request.
func (client AppsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, appPatch AppPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", pathParameters),
		autorest.WithJSON(appPatch),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) UpdateSender(req *http.Request) (future AppsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral instead.
// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AppsClient) UpdateResponder(resp *http.Response) (result App, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
