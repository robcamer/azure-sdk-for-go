package managedapplications

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

// ApplicationsClient is the ARM applications
type ApplicationsClient struct {
	BaseClient
}

// NewApplicationsClient creates an instance of the ApplicationsClient client.
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return NewApplicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationsClientWithBaseURI creates an instance of the ApplicationsClient client.
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return ApplicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new managed application.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationName - the name of the managed application.
// parameters - parameters supplied to the create or update a managed application.
func (client ApplicationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters Application) (result ApplicationsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationProperties.ManagedResourceGroupID", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "parameters.Plan", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Plan.Name", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Publisher", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Product", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Version", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Kind", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Pattern, Rule: `^[-\w\._,\(\)]+$`, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, applicationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, applicationName string, parameters Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateOrUpdateSender(req *http.Request) (future ApplicationsCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateOrUpdateResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByID creates a new managed application.
// Parameters:
// applicationID - the fully qualified ID of the managed application, including the managed application name
// and the managed application resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applications/{application-name}
// parameters - parameters supplied to the create or update a managed application.
func (client ApplicationsClient) CreateOrUpdateByID(ctx context.Context, applicationID string, parameters Application) (result ApplicationsCreateOrUpdateByIDFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationProperties.ManagedResourceGroupID", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "parameters.Plan", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Plan.Name", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Publisher", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Product", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Plan.Version", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Kind", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Pattern, Rule: `^[-\w\._,\(\)]+$`, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "CreateOrUpdateByID", err.Error())
	}

	req, err := client.CreateOrUpdateByIDPreparer(ctx, applicationID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "CreateOrUpdateByID", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "CreateOrUpdateByID", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateByIDPreparer prepares the CreateOrUpdateByID request.
func (client ApplicationsClient) CreateOrUpdateByIDPreparer(ctx context.Context, applicationID string, parameters Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationId": applicationID,
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applicationId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByIDSender sends the CreateOrUpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateOrUpdateByIDSender(req *http.Request) (future ApplicationsCreateOrUpdateByIDFuture, err error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateByIDResponder handles the response to the CreateOrUpdateByID request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateOrUpdateByIDResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the managed application.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationName - the name of the managed application.
func (client ApplicationsClient) Delete(ctx context.Context, resourceGroupName string, applicationName string) (result ApplicationsDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) DeleteSender(req *http.Request) (future ApplicationsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByID deletes the managed application.
// Parameters:
// applicationID - the fully qualified ID of the managed application, including the managed application name
// and the managed application resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applications/{application-name}
func (client ApplicationsClient) DeleteByID(ctx context.Context, applicationID string) (result ApplicationsDeleteByIDFuture, err error) {
	req, err := client.DeleteByIDPreparer(ctx, applicationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "DeleteByID", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "DeleteByID", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client ApplicationsClient) DeleteByIDPreparer(ctx context.Context, applicationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationId": applicationID,
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) DeleteByIDSender(req *http.Request) (future ApplicationsDeleteByIDFuture, err error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) DeleteByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the managed application.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationName - the name of the managed application.
func (client ApplicationsClient) Get(ctx context.Context, resourceGroupName string, applicationName string) (result Application, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets the managed application.
// Parameters:
// applicationID - the fully qualified ID of the managed application, including the managed application name
// and the managed application resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applications/{application-name}
func (client ApplicationsClient) GetByID(ctx context.Context, applicationID string) (result Application, err error) {
	req, err := client.GetByIDPreparer(ctx, applicationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client ApplicationsClient) GetByIDPreparer(ctx context.Context, applicationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationId": applicationID,
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetByIDResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all the applications within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ApplicationsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ApplicationListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplicationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListByResourceGroupResponder(resp *http.Response) (result ApplicationListResult, err error) {
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
func (client ApplicationsClient) listByResourceGroupNextResults(lastResults ApplicationListResult) (result ApplicationListResult, err error) {
	req, err := lastResults.applicationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ApplicationListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription gets all the applications within a subscription.
func (client ApplicationsClient) ListBySubscription(ctx context.Context) (result ApplicationListResultPage, err error) {
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ApplicationsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListBySubscriptionResponder(resp *http.Response) (result ApplicationListResult, err error) {
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
func (client ApplicationsClient) listBySubscriptionNextResults(lastResults ApplicationListResult) (result ApplicationListResult, err error) {
	req, err := lastResults.applicationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationsClient) ListBySubscriptionComplete(ctx context.Context) (result ApplicationListResultIterator, err error) {
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Update updates an existing managed application. The only value that can be updated via PATCH currently is the tags.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationName - the name of the managed application.
// parameters - parameters supplied to update an existing managed application.
func (client ApplicationsClient) Update(ctx context.Context, resourceGroupName string, applicationName string, parameters *Application) (result Application, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplicationsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, applicationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ApplicationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, applicationName string, parameters *Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdateResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateByID updates an existing managed application. The only value that can be updated via PATCH currently is the
// tags.
// Parameters:
// applicationID - the fully qualified ID of the managed application, including the managed application name
// and the managed application resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applications/{application-name}
// parameters - parameters supplied to update an existing managed application.
func (client ApplicationsClient) UpdateByID(ctx context.Context, applicationID string, parameters *Application) (result Application, err error) {
	req, err := client.UpdateByIDPreparer(ctx, applicationID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "UpdateByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "UpdateByID", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplicationsClient", "UpdateByID", resp, "Failure responding to request")
	}

	return
}

// UpdateByIDPreparer prepares the UpdateByID request.
func (client ApplicationsClient) UpdateByIDPreparer(ctx context.Context, applicationID string, parameters *Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationId": applicationID,
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateByIDSender sends the UpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdateByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateByIDResponder handles the response to the UpdateByID request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdateByIDResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}