/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type GlobalAPI interface {
	/*
			Read Retrieve the global configuration.

			Use this method to retrieve the __Global__ configuration object.
		The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GlobalAPIReadRequest
	*/
	Read(ctx context.Context) GlobalAPIReadRequest

	// ReadExecute executes the request
	//  @return ReadGlobalResponse
	ReadExecute(r GlobalAPIReadRequest) (*ReadGlobalResponse, *http.Response, error)
	/*
			ReadById Retrieve the global configuration.

			Use this method to retrieve the __Global__ configuration object.
		The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return GlobalAPIReadByIdRequest
	*/
	ReadById(ctx context.Context, id string) GlobalAPIReadByIdRequest

	// ReadByIdExecute executes the request
	//  @return ReadGlobalResponse
	ReadByIdExecute(r GlobalAPIReadByIdRequest) (*ReadGlobalResponse, *http.Response, error)
	/*
			Update Update the global configuration.

			Use this method to update the __Global__ configuration object.
		The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GlobalAPIUpdateRequest
	*/
	Update(ctx context.Context) GlobalAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateGlobalResponse
	UpdateExecute(r GlobalAPIUpdateRequest) (*UpdateGlobalResponse, *http.Response, error)
	/*
			UpdateById Update the global configuration.

			Use this method to update the __Global__ configuration object.
		The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return GlobalAPIUpdateByIdRequest
	*/
	UpdateById(ctx context.Context, id string) GlobalAPIUpdateByIdRequest

	// UpdateByIdExecute executes the request
	//  @return UpdateGlobalResponse
	UpdateByIdExecute(r GlobalAPIUpdateByIdRequest) (*UpdateGlobalResponse, *http.Response, error)
}

// GlobalAPIService GlobalAPI service
type GlobalAPIService internal.Service

type GlobalAPIReadRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r GlobalAPIReadRequest) Fields(fields string) GlobalAPIReadRequest {
	r.fields = &fields
	return r
}

func (r GlobalAPIReadRequest) Execute() (*ReadGlobalResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Retrieve the global configuration.

Use this method to retrieve the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GlobalAPIReadRequest
*/
func (a *GlobalAPIService) Read(ctx context.Context) GlobalAPIReadRequest {
	return GlobalAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReadGlobalResponse
func (a *GlobalAPIService) ReadExecute(r GlobalAPIReadRequest) (*ReadGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type GlobalAPIReadByIdRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r GlobalAPIReadByIdRequest) Fields(fields string) GlobalAPIReadByIdRequest {
	r.fields = &fields
	return r
}

func (r GlobalAPIReadByIdRequest) Execute() (*ReadGlobalResponse, *http.Response, error) {
	return r.ApiService.ReadByIdExecute(r)
}

/*
ReadById Retrieve the global configuration.

Use this method to retrieve the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return GlobalAPIReadByIdRequest
*/
func (a *GlobalAPIService) ReadById(ctx context.Context, id string) GlobalAPIReadByIdRequest {
	return GlobalAPIReadByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ReadGlobalResponse
func (a *GlobalAPIService) ReadByIdExecute(r GlobalAPIReadByIdRequest) (*ReadGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.ReadById")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type GlobalAPIUpdateRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	body       *Global
}

func (r GlobalAPIUpdateRequest) Body(body Global) GlobalAPIUpdateRequest {
	r.body = &body
	return r
}

func (r GlobalAPIUpdateRequest) Execute() (*UpdateGlobalResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update the global configuration.

Use this method to update the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GlobalAPIUpdateRequest
*/
func (a *GlobalAPIService) Update(ctx context.Context) GlobalAPIUpdateRequest {
	return GlobalAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateGlobalResponse
func (a *GlobalAPIService) UpdateExecute(r GlobalAPIUpdateRequest) (*UpdateGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type GlobalAPIUpdateByIdRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	id         string
	body       *Global
}

func (r GlobalAPIUpdateByIdRequest) Body(body Global) GlobalAPIUpdateByIdRequest {
	r.body = &body
	return r
}

func (r GlobalAPIUpdateByIdRequest) Execute() (*UpdateGlobalResponse, *http.Response, error) {
	return r.ApiService.UpdateByIdExecute(r)
}

/*
UpdateById Update the global configuration.

Use this method to update the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return GlobalAPIUpdateByIdRequest
*/
func (a *GlobalAPIService) UpdateById(ctx context.Context, id string) GlobalAPIUpdateByIdRequest {
	return GlobalAPIUpdateByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateGlobalResponse
func (a *GlobalAPIService) UpdateByIdExecute(r GlobalAPIUpdateByIdRequest) (*UpdateGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.UpdateById")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
