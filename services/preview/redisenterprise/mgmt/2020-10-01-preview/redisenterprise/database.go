package redisenterprise

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DatabaseClient is the REST API for managing Redis Enterprise resources in Azure.
type DatabaseClient struct {
	BaseClient
}

// NewDatabaseClient creates an instance of the DatabaseClient client.
func NewDatabaseClient(subscriptionID string) DatabaseClient {
	return NewDatabaseClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseClientWithBaseURI creates an instance of the DatabaseClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDatabaseClientWithBaseURI(baseURI string, subscriptionID string) DatabaseClient {
	return DatabaseClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Export exports a database file from target database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - storage information for exporting into the cluster
func (client DatabaseClient) Export(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ExportClusterParameters) (result DatabaseExportFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseClient.Export")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SasURI", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redisenterprise.DatabaseClient", "Export", err.Error())
	}

	req, err := client.ExportPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "Export", nil, "Failure preparing request")
		return
	}

	result, err = client.ExportSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "Export", result.Response(), "Failure sending request")
		return
	}

	return
}

// ExportPreparer prepares the Export request.
func (client DatabaseClient) ExportPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ExportClusterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/export", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExportSender sends the Export request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseClient) ExportSender(req *http.Request) (future DatabaseExportFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ExportResponder handles the response to the Export request. The method always
// closes the http.Response Body.
func (client DatabaseClient) ExportResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Import imports a database file to target database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - storage information for importing into the cluster
func (client DatabaseClient) Import(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ImportClusterParameters) (result DatabaseImportFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseClient.Import")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SasURI", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redisenterprise.DatabaseClient", "Import", err.Error())
	}

	req, err := client.ImportPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "Import", nil, "Failure preparing request")
		return
	}

	result, err = client.ImportSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "Import", result.Response(), "Failure sending request")
		return
	}

	return
}

// ImportPreparer prepares the Import request.
func (client DatabaseClient) ImportPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ImportClusterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/import", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImportSender sends the Import request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseClient) ImportSender(req *http.Request) (future DatabaseImportFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ImportResponder handles the response to the Import request. The method always
// closes the http.Response Body.
func (client DatabaseClient) ImportResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListKeys retrieves the access keys for the RedisEnterprise database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
func (client DatabaseClient) ListKeys(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result AccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseClient.ListKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client DatabaseClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client DatabaseClient) ListKeysResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKey regenerates the RedisEnterprise database's access keys.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - specifies which key to regenerate.
func (client DatabaseClient) RegenerateKey(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters RegenerateKeyParameters) (result DatabaseRegenerateKeyFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	result, err = client.RegenerateKeySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabaseClient", "RegenerateKey", result.Response(), "Failure sending request")
		return
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client DatabaseClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/regenerateKey", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseClient) RegenerateKeySender(req *http.Request) (future DatabaseRegenerateKeyFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client DatabaseClient) RegenerateKeyResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
