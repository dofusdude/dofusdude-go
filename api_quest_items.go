/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.2
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// QuestItemsAPIService QuestItemsAPI service
type QuestItemsAPIService service

type ApiGetAllItemsQuestListRequest struct {
	ctx context.Context
	ApiService *QuestItemsAPIService
	language string
	game string
	sortLevel *string
	filterMinLevel *int32
	filterMaxLevel *int32
	acceptEncoding *string
	filterTypeNameId *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetAllItemsQuestListRequest) SortLevel(sortLevel string) ApiGetAllItemsQuestListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results which level is equal or above this value
func (r ApiGetAllItemsQuestListRequest) FilterMinLevel(filterMinLevel int32) ApiGetAllItemsQuestListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetAllItemsQuestListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetAllItemsQuestListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllItemsQuestListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllItemsQuestListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetAllItemsQuestListRequest) FilterTypeNameId(filterTypeNameId []string) ApiGetAllItemsQuestListRequest {
	r.filterTypeNameId = &filterTypeNameId
	return r
}

func (r ApiGetAllItemsQuestListRequest) Execute() (*ListItems, *http.Response, error) {
	return r.ApiService.GetAllItemsQuestListExecute(r)
}

/*
GetAllItemsQuestList List All Quest Items

Retrieve all quest items with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

If you want everything unfiltered, delete the other query parameters.

Be careful with testing or (god forbid) using /all in your browser, the returned json is huge and will slow down the browser!

Tip: set the HTTP Header 'Accept-Encoding: gzip' for saving bandwidth. You will need to uncompress it on your end.
Example with cURL:
```
curl -sH 'Accept-Encoding: gzip' <api-endpoint> | gunzip -
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game dofus3 | dofus3beta
 @return ApiGetAllItemsQuestListRequest
*/
func (a *QuestItemsAPIService) GetAllItemsQuestList(ctx context.Context, language string, game string) ApiGetAllItemsQuestListRequest {
	return ApiGetAllItemsQuestListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListItems
func (a *QuestItemsAPIService) GetAllItemsQuestListExecute(r ApiGetAllItemsQuestListRequest) (*ListItems, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsAPIService.GetAllItemsQuestList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/quest/all"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "form", "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.filterTypeNameId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type.name_id]", r.filterTypeNameId, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptEncoding != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Encoding", r.acceptEncoding, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemQuestSingleRequest struct {
	ctx context.Context
	ApiService *QuestItemsAPIService
	language string
	ankamaId int32
	game string
}

func (r ApiGetItemQuestSingleRequest) Execute() (*Resource, *http.Response, error) {
	return r.ApiService.GetItemQuestSingleExecute(r)
}

/*
GetItemQuestSingle Single Quest Items

Retrieve a specific quest item with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game dofus3 | dofus3beta
 @return ApiGetItemQuestSingleRequest
*/
func (a *QuestItemsAPIService) GetItemQuestSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetItemQuestSingleRequest {
	return ApiGetItemQuestSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Resource
func (a *QuestItemsAPIService) GetItemQuestSingleExecute(r ApiGetItemQuestSingleRequest) (*Resource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Resource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsAPIService.GetItemQuestSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/quest/{ankama_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ankama_id"+"}", url.PathEscape(parameterValueToString(r.ankamaId, "ankamaId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemsQuestListRequest struct {
	ctx context.Context
	ApiService *QuestItemsAPIService
	language string
	game string
	sortLevel *string
	filterMinLevel *int32
	filterMaxLevel *int32
	pageSize *int32
	pageNumber *int32
	fieldsItem *[]string
	filterTypeNameId *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetItemsQuestListRequest) SortLevel(sortLevel string) ApiGetItemsQuestListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsQuestListRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsQuestListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsQuestListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsQuestListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetItemsQuestListRequest) PageSize(pageSize int32) ApiGetItemsQuestListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetItemsQuestListRequest) PageNumber(pageNumber int32) ApiGetItemsQuestListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetItemsQuestListRequest) FieldsItem(fieldsItem []string) ApiGetItemsQuestListRequest {
	r.fieldsItem = &fieldsItem
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetItemsQuestListRequest) FilterTypeNameId(filterTypeNameId []string) ApiGetItemsQuestListRequest {
	r.filterTypeNameId = &filterTypeNameId
	return r
}

func (r ApiGetItemsQuestListRequest) Execute() (*ListItems, *http.Response, error) {
	return r.ApiService.GetItemsQuestListExecute(r)
}

/*
GetItemsQuestList List Quest Items

Retrieve a list of quest items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game dofus3 | dofus3beta
 @return ApiGetItemsQuestListRequest
*/
func (a *QuestItemsAPIService) GetItemsQuestList(ctx context.Context, language string, game string) ApiGetItemsQuestListRequest {
	return ApiGetItemsQuestListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListItems
func (a *QuestItemsAPIService) GetItemsQuestListExecute(r ApiGetItemsQuestListRequest) (*ListItems, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsAPIService.GetItemsQuestList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/quest"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "form", "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "form", "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "form", "")
	}
	if r.fieldsItem != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[item]", r.fieldsItem, "form", "csv")
	}
	if r.filterTypeNameId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type.name_id]", r.filterTypeNameId, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemsQuestSearchRequest struct {
	ctx context.Context
	ApiService *QuestItemsAPIService
	language string
	game string
	query *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	limit *int32
	filterTypeEnum *[]string
}

// case sensitive search query
func (r ApiGetItemsQuestSearchRequest) Query(query string) ApiGetItemsQuestSearchRequest {
	r.query = &query
	return r
}

// only results with the translated type name
func (r ApiGetItemsQuestSearchRequest) FilterTypeName(filterTypeName string) ApiGetItemsQuestSearchRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsQuestSearchRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsQuestSearchRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsQuestSearchRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsQuestSearchRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// maximum number of returned results
func (r ApiGetItemsQuestSearchRequest) Limit(limit int32) ApiGetItemsQuestSearchRequest {
	r.limit = &limit
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetItemsQuestSearchRequest) FilterTypeEnum(filterTypeEnum []string) ApiGetItemsQuestSearchRequest {
	r.filterTypeEnum = &filterTypeEnum
	return r
}

func (r ApiGetItemsQuestSearchRequest) Execute() ([]ListItem, *http.Response, error) {
	return r.ApiService.GetItemsQuestSearchExecute(r)
}

/*
GetItemsQuestSearch Search Quest Items

Search in all names and descriptions of quest items with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game dofus3 | dofus3beta
 @return ApiGetItemsQuestSearchRequest
*/
func (a *QuestItemsAPIService) GetItemsQuestSearch(ctx context.Context, language string, game string) ApiGetItemsQuestSearchRequest {
	return ApiGetItemsQuestSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ListItem
func (a *QuestItemsAPIService) GetItemsQuestSearchExecute(r ApiGetItemsQuestSearchRequest) ([]ListItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsAPIService.GetItemsQuestSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/quest/search"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	if r.filterTypeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type_name]", r.filterTypeName, "form", "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 8
		r.limit = &defaultValue
	}
	if r.filterTypeEnum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type_enum]", r.filterTypeEnum, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
