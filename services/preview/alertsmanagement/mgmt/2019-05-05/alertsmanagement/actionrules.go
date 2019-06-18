package alertsmanagement

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ActionRulesClient is the alertsManagement Client
type ActionRulesClient struct {
	BaseClient
}

// NewActionRulesClient creates an instance of the ActionRulesClient client.
func NewActionRulesClient(subscriptionID string) ActionRulesClient {
	return NewActionRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActionRulesClientWithBaseURI creates an instance of the ActionRulesClient client.
func NewActionRulesClientWithBaseURI(baseURI string, subscriptionID string) ActionRulesClient {
	return ActionRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateUpdate creates/Updates a specific action rule
// Parameters:
// resourceGroupName - resource group name where the resource is created.
// actionRuleName - the name of action rule that needs to be created/updated
// actionRule - action rule to be created/updated
func (client ActionRulesClient) CreateUpdate(ctx context.Context, resourceGroupName string, actionRuleName string, actionRule ActionRule) (result ActionRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.CreateUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateUpdatePreparer(ctx, resourceGroupName, actionRuleName, actionRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "CreateUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "CreateUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "CreateUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateUpdatePreparer prepares the CreateUpdate request.
func (client ActionRulesClient) CreateUpdatePreparer(ctx context.Context, resourceGroupName string, actionRuleName string, actionRule ActionRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionRuleName":    autorest.Encode("path", actionRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", pathParameters),
		autorest.WithJSON(actionRule))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateUpdateSender sends the CreateUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) CreateUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateUpdateResponder handles the response to the CreateUpdate request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) CreateUpdateResponder(resp *http.Response) (result ActionRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a given action rule
// Parameters:
// resourceGroupName - resource group name where the resource is created.
// actionRuleName - the name that needs to be deleted
func (client ActionRulesClient) Delete(ctx context.Context, resourceGroupName string, actionRuleName string) (result Bool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, actionRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ActionRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, actionRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionRuleName":    autorest.Encode("path", actionRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) DeleteResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByName get a specific action rule
// Parameters:
// resourceGroupName - resource group name where the resource is created.
// actionRuleName - the name of action rule that needs to be fetched
func (client ActionRulesClient) GetByName(ctx context.Context, resourceGroupName string, actionRuleName string) (result ActionRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.GetByName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByNamePreparer(ctx, resourceGroupName, actionRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "GetByName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "GetByName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "GetByName", resp, "Failure responding to request")
	}

	return
}

// GetByNamePreparer prepares the GetByName request.
func (client ActionRulesClient) GetByNamePreparer(ctx context.Context, resourceGroupName string, actionRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionRuleName":    autorest.Encode("path", actionRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByNameSender sends the GetByName request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) GetByNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByNameResponder handles the response to the GetByName request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) GetByNameResponder(resp *http.Response) (result ActionRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list all action rules of the subscription, created in given resource group and given input
// filters
// Parameters:
// resourceGroupName - resource group name where the resource is created.
// targetResourceGroup - filter by target resource group name
// targetResourceType - filter by target resource type
// targetResource - filter by target resource
// severity - filter by severity
// monitorService - filter by monitor service which is the source of the alert object.
// impactedScope - filter by impacted/target scope (provide comma separated list for multiple scopes). The
// value should be an well constructed ARM id of the scope.
// description - filter by alert rule description
// alertRuleID - filter by alert rule id
// actionGroup - filter by action group configured as part of action rule
// name - filter by action rule name
func (client ActionRulesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (result ActionRulesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.arl.Response.Response != nil {
				sc = result.arl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, targetResourceGroup, targetResourceType, targetResource, severity, monitorService, impactedScope, description, alertRuleID, actionGroup, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.arl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.arl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ActionRulesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{}
	if len(targetResourceGroup) > 0 {
		queryParameters["targetResourceGroup"] = autorest.Encode("query", targetResourceGroup)
	}
	if len(targetResourceType) > 0 {
		queryParameters["targetResourceType"] = autorest.Encode("query", targetResourceType)
	}
	if len(targetResource) > 0 {
		queryParameters["targetResource"] = autorest.Encode("query", targetResource)
	}
	if len(string(severity)) > 0 {
		queryParameters["severity"] = autorest.Encode("query", severity)
	}
	if len(string(monitorService)) > 0 {
		queryParameters["monitorService"] = autorest.Encode("query", monitorService)
	}
	if len(impactedScope) > 0 {
		queryParameters["impactedScope"] = autorest.Encode("query", impactedScope)
	}
	if len(description) > 0 {
		queryParameters["description"] = autorest.Encode("query", description)
	}
	if len(alertRuleID) > 0 {
		queryParameters["alertRuleId"] = autorest.Encode("query", alertRuleID)
	}
	if len(actionGroup) > 0 {
		queryParameters["actionGroup"] = autorest.Encode("query", actionGroup)
	}
	if len(name) > 0 {
		queryParameters["name"] = autorest.Encode("query", name)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) ListByResourceGroupResponder(resp *http.Response) (result ActionRulesList, err error) {
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
func (client ActionRulesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ActionRulesList) (result ActionRulesList, err error) {
	req, err := lastResults.actionRulesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActionRulesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (result ActionRulesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, targetResourceGroup, targetResourceType, targetResource, severity, monitorService, impactedScope, description, alertRuleID, actionGroup, name)
	return
}

// ListBySubscription list all action rules of the subscription and given input filters
// Parameters:
// targetResourceGroup - filter by target resource group name
// targetResourceType - filter by target resource type
// targetResource - filter by target resource
// severity - filter by severity
// monitorService - filter by monitor service which is the source of the alert object.
// impactedScope - filter by impacted/target scope (provide comma separated list for multiple scopes). The
// value should be an well constructed ARM id of the scope.
// description - filter by alert rule description
// alertRuleID - filter by alert rule id
// actionGroup - filter by action group configured as part of action rule
// name - filter by action rule name
func (client ActionRulesClient) ListBySubscription(ctx context.Context, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (result ActionRulesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.arl.Response.Response != nil {
				sc = result.arl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, targetResourceGroup, targetResourceType, targetResource, severity, monitorService, impactedScope, description, alertRuleID, actionGroup, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.arl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.arl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ActionRulesClient) ListBySubscriptionPreparer(ctx context.Context, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{}
	if len(targetResourceGroup) > 0 {
		queryParameters["targetResourceGroup"] = autorest.Encode("query", targetResourceGroup)
	}
	if len(targetResourceType) > 0 {
		queryParameters["targetResourceType"] = autorest.Encode("query", targetResourceType)
	}
	if len(targetResource) > 0 {
		queryParameters["targetResource"] = autorest.Encode("query", targetResource)
	}
	if len(string(severity)) > 0 {
		queryParameters["severity"] = autorest.Encode("query", severity)
	}
	if len(string(monitorService)) > 0 {
		queryParameters["monitorService"] = autorest.Encode("query", monitorService)
	}
	if len(impactedScope) > 0 {
		queryParameters["impactedScope"] = autorest.Encode("query", impactedScope)
	}
	if len(description) > 0 {
		queryParameters["description"] = autorest.Encode("query", description)
	}
	if len(alertRuleID) > 0 {
		queryParameters["alertRuleId"] = autorest.Encode("query", alertRuleID)
	}
	if len(actionGroup) > 0 {
		queryParameters["actionGroup"] = autorest.Encode("query", actionGroup)
	}
	if len(name) > 0 {
		queryParameters["name"] = autorest.Encode("query", name)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/actionRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) ListBySubscriptionResponder(resp *http.Response) (result ActionRulesList, err error) {
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
func (client ActionRulesClient) listBySubscriptionNextResults(ctx context.Context, lastResults ActionRulesList) (result ActionRulesList, err error) {
	req, err := lastResults.actionRulesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActionRulesClient) ListBySubscriptionComplete(ctx context.Context, targetResourceGroup string, targetResourceType string, targetResource string, severity Severity, monitorService MonitorService, impactedScope string, description string, alertRuleID string, actionGroup string, name string) (result ActionRulesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, targetResourceGroup, targetResourceType, targetResource, severity, monitorService, impactedScope, description, alertRuleID, actionGroup, name)
	return
}

// Update update enabled flag and/or tags for the given action rule
// Parameters:
// resourceGroupName - resource group name where the resource is created.
// actionRuleName - the name that needs to be updated
// actionRulePatch - parameters supplied to the operation.
func (client ActionRulesClient) Update(ctx context.Context, resourceGroupName string, actionRuleName string, actionRulePatch PatchObject) (result ActionRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionRulesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, actionRuleName, actionRulePatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertsmanagement.ActionRulesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ActionRulesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, actionRuleName string, actionRulePatch PatchObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionRuleName":    autorest.Encode("path", actionRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", pathParameters),
		autorest.WithJSON(actionRulePatch))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ActionRulesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ActionRulesClient) UpdateResponder(resp *http.Response) (result ActionRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}