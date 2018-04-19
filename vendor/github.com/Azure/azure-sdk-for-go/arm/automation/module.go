package automation

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ModuleClient is the automation Client
type ModuleClient struct {
	ManagementClient
}

// NewModuleClient creates an instance of the ModuleClient client.
func NewModuleClient(subscriptionID string) ModuleClient {
	return NewModuleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewModuleClientWithBaseURI creates an instance of the ModuleClient client.
func NewModuleClientWithBaseURI(baseURI string, subscriptionID string) ModuleClient {
	return ModuleClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update the module identified by module name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. moduleName is
// the name of module. parameters is the create or update parameters for module.
func (client ModuleClient) CreateOrUpdate(resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters) (result Module, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ModuleCreateOrUpdateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ModuleCreateOrUpdateProperties.ContentLink", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash.Algorithm", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash.Value", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ModuleClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, automationAccountName, moduleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ModuleClient) CreateOrUpdatePreparer(resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ModuleClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ModuleClient) CreateOrUpdateResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the module by name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. moduleName is
// the module name.
func (client ModuleClient) Delete(resourceGroupName string, automationAccountName string, moduleName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ModuleClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, automationAccountName, moduleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ModuleClient) DeletePreparer(resourceGroupName string, automationAccountName string, moduleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ModuleClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ModuleClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the module identified by module name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. moduleName is
// the module name.
func (client ModuleClient) Get(resourceGroupName string, automationAccountName string, moduleName string) (result Module, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ModuleClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, moduleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ModuleClient) GetPreparer(resourceGroupName string, automationAccountName string, moduleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ModuleClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ModuleClient) GetResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of modules.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name.
func (client ModuleClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string) (result ModuleListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ModuleClient", "ListByAutomationAccount")
	}

	req, err := client.ListByAutomationAccountPreparer(resourceGroupName, automationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client ModuleClient) ListByAutomationAccountPreparer(resourceGroupName string, automationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ModuleClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client ModuleClient) ListByAutomationAccountResponder(resp *http.Response) (result ModuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccountNextResults retrieves the next set of results, if any.
func (client ModuleClient) ListByAutomationAccountNextResults(lastResults ModuleListResult) (result ModuleListResult, err error) {
	req, err := lastResults.ModuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", resp, "Failure sending next results request")
	}

	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "ListByAutomationAccount", resp, "Failure responding to next results request")
	}

	return
}

// ListByAutomationAccountComplete gets all elements from the list without paging.
func (client ModuleClient) ListByAutomationAccountComplete(resourceGroupName string, automationAccountName string, cancel <-chan struct{}) (<-chan Module, <-chan error) {
	resultChan := make(chan Module)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByAutomationAccount(resourceGroupName, automationAccountName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByAutomationAccountNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// Update update the module identified by module name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. moduleName is
// the name of module. parameters is the update parameters for module.
func (client ModuleClient) Update(resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleUpdateParameters) (result Module, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ModuleClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, automationAccountName, moduleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ModuleClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ModuleClient) UpdatePreparer(resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ModuleClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ModuleClient) UpdateResponder(resp *http.Response) (result Module, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}