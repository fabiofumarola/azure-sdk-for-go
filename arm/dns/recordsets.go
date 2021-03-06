package dns

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// RecordSetsClient is the client for the RecordSets methods of the Dns
// service.
type RecordSetsClient struct {
	ManagementClient
}

// NewRecordSetsClient creates an instance of the RecordSetsClient client.
func NewRecordSetsClient(subscriptionID string) RecordSetsClient {
	return NewRecordSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecordSetsClientWithBaseURI creates an instance of the RecordSetsClient
// client.
func NewRecordSetsClientWithBaseURI(baseURI string, subscriptionID string) RecordSetsClient {
	return RecordSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates a RecordSet within a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. relativeRecordSetName is the name
// of the RecordSet, relative to the name of the zone. recordType is the type
// of DNS record. Possible values include: 'A', 'AAAA', 'CNAME', 'MX', 'NS',
// 'PTR', 'SOA', 'SRV', 'TXT' parameters is parameters supplied to the
// CreateOrUpdate operation. ifMatch is the etag of Recordset. ifNoneMatch is
// defines the If-None-Match condition. Set to '*' to force
// Create-If-Not-Exist. Other values will be ignored.
func (client RecordSetsClient) CreateOrUpdate(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (result RecordSet, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RecordSetsClient) CreateOrUpdatePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) CreateOrUpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes a RecordSet from a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. relativeRecordSetName is the name
// of the RecordSet, relative to the name of the zone. recordType is the type
// of DNS record. Possible values include: 'A', 'AAAA', 'CNAME', 'MX', 'NS',
// 'PTR', 'SOA', 'SRV', 'TXT' ifMatch is defines the If-Match condition. The
// delete operation will be performed only if the ETag of the zone on the
// server matches this value. ifNoneMatch is defines the If-None-Match
// condition. The delete operation will be performed only if the ETag of the
// zone on the server does not match this value.
func (client RecordSetsClient) Delete(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, ifMatch string, ifNoneMatch string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, ifMatch, ifNoneMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RecordSetsClient) DeletePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a RecordSet.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. relativeRecordSetName is the name
// of the RecordSet, relative to the name of the zone. recordType is the type
// of DNS record. Possible values include: 'A', 'AAAA', 'CNAME', 'MX', 'NS',
// 'PTR', 'SOA', 'SRV', 'TXT'
func (client RecordSetsClient) Get(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType) (result RecordSet, err error) {
	req, err := client.GetPreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecordSetsClient) GetPreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) GetResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllInResourceGroup lists all RecordSets in a DNS zone.
//
// resourceGroupName is the name of the resource group that contains the zone.
// zoneName is the name of the zone from which to enumerate RecordSets. top
// is query parameters. If null is passed returns the default number of
// zones.
func (client RecordSetsClient) ListAllInResourceGroup(resourceGroupName string, zoneName string, top string) (result RecordSetListResult, err error) {
	req, err := client.ListAllInResourceGroupPreparer(resourceGroupName, zoneName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListAllInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListAllInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListAllInResourceGroupPreparer prepares the ListAllInResourceGroup request.
func (client RecordSetsClient) ListAllInResourceGroupPreparer(resourceGroupName string, zoneName string, top string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/recordsets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllInResourceGroupSender sends the ListAllInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListAllInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllInResourceGroupResponder handles the response to the ListAllInResourceGroup request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListAllInResourceGroupResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllInResourceGroupNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListAllInResourceGroupNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", resp, "Failure sending next results request request")
	}

	result, err = client.ListAllInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAllInResourceGroup", resp, "Failure responding to next results request request")
	}

	return
}

// ListByType lists the RecordSets of a specified type in a DNS zone.
//
// resourceGroupName is the name of the resource group that contains the zone.
// zoneName is the name of the zone from which to enumerate RecordsSets.
// recordType is the type of record sets to enumerate. Possible values
// include: 'A', 'AAAA', 'CNAME', 'MX', 'NS', 'PTR', 'SOA', 'SRV', 'TXT' top
// is query parameters. If null is passed returns the default number of
// zones.
func (client RecordSetsClient) ListByType(resourceGroupName string, zoneName string, recordType RecordType, top string) (result RecordSetListResult, err error) {
	req, err := client.ListByTypePreparer(resourceGroupName, zoneName, recordType, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", nil, "Failure preparing request")
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure sending request")
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure responding to request")
	}

	return
}

// ListByTypePreparer prepares the ListByType request.
func (client RecordSetsClient) ListByTypePreparer(resourceGroupName string, zoneName string, recordType RecordType, top string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":        autorest.Encode("path", recordType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByTypeSender sends the ListByType request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListByTypeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByTypeResponder handles the response to the ListByType request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListByTypeResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTypeNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListByTypeNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure sending next results request request")
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListByType", resp, "Failure responding to next results request request")
	}

	return
}

// Update updates a RecordSet within a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. relativeRecordSetName is the name
// of the RecordSet, relative to the name of the zone. recordType is the type
// of DNS record. Possible values include: 'A', 'AAAA', 'CNAME', 'MX', 'NS',
// 'PTR', 'SOA', 'SRV', 'TXT' parameters is parameters supplied to the Update
// operation. ifMatch is the etag of Zone. ifNoneMatch is defines the
// If-None-Match condition. Set to '*' to force Create-If-Not-Exist. Other
// values will be ignored.
func (client RecordSetsClient) Update(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (result RecordSet, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RecordSetsClient) UpdatePreparer(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) UpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
