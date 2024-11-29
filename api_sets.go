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


// SetsAPIService SetsAPI service
type SetsAPIService service

type ApiGetAllSetsListRequest struct {
	ctx context.Context
	ApiService *SetsAPIService
	language string
	game string
	sortLevel *string
	filterMinHighestEquipmentLevel *int32
	filterMaxHighestEquipmentLevel *int32
	acceptEncoding *string
	filterContainsCosmeticsOnly *bool
	filterContainsCosmetics *bool
}

// sort the resulting list by level, default unsorted
func (r ApiGetAllSetsListRequest) SortLevel(sortLevel string) ApiGetAllSetsListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results where the equipment with the highest level is above or equal to this value
func (r ApiGetAllSetsListRequest) FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel int32) ApiGetAllSetsListRequest {
	r.filterMinHighestEquipmentLevel = &filterMinHighestEquipmentLevel
	return r
}

// only results where the equipment with the highest level is below or equal to this value
func (r ApiGetAllSetsListRequest) FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel int32) ApiGetAllSetsListRequest {
	r.filterMaxHighestEquipmentLevel = &filterMaxHighestEquipmentLevel
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllSetsListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllSetsListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

// filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment.
func (r ApiGetAllSetsListRequest) FilterContainsCosmeticsOnly(filterContainsCosmeticsOnly bool) ApiGetAllSetsListRequest {
	r.filterContainsCosmeticsOnly = &filterContainsCosmeticsOnly
	return r
}

// filter sets based on if they got cosmetic items in it.
func (r ApiGetAllSetsListRequest) FilterContainsCosmetics(filterContainsCosmetics bool) ApiGetAllSetsListRequest {
	r.filterContainsCosmetics = &filterContainsCosmetics
	return r
}

func (r ApiGetAllSetsListRequest) Execute() (*ListSets, *http.Response, error) {
	return r.ApiService.GetAllSetsListExecute(r)
}

/*
GetAllSetsList List All Sets

Retrieve all sets with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

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
 @return ApiGetAllSetsListRequest
*/
func (a *SetsAPIService) GetAllSetsList(ctx context.Context, language string, game string) ApiGetAllSetsListRequest {
	return ApiGetAllSetsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListSets
func (a *SetsAPIService) GetAllSetsListExecute(r ApiGetAllSetsListRequest) (*ListSets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListSets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetsAPIService.GetAllSetsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/sets/all"
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
	if r.filterMinHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_highest_equipment_level]", r.filterMinHighestEquipmentLevel, "form", "")
	}
	if r.filterMaxHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_highest_equipment_level]", r.filterMaxHighestEquipmentLevel, "form", "")
	}
	if r.filterContainsCosmeticsOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[contains_cosmetics_only]", r.filterContainsCosmeticsOnly, "form", "")
	}
	if r.filterContainsCosmetics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[contains_cosmetics]", r.filterContainsCosmetics, "form", "")
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

type ApiGetSetsListRequest struct {
	ctx context.Context
	ApiService *SetsAPIService
	language string
	game string
	sortLevel *string
	filterMinHighestEquipmentLevel *int32
	filterMaxHighestEquipmentLevel *int32
	pageSize *int32
	pageNumber *int32
	fieldsSet *[]string
	filterContainsCosmeticsOnly *bool
	filterContainsCosmetics *bool
}

// sort the resulting list by level, default unsorted
func (r ApiGetSetsListRequest) SortLevel(sortLevel string) ApiGetSetsListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results where the equipment with the highest level is above or equal to this value
func (r ApiGetSetsListRequest) FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel int32) ApiGetSetsListRequest {
	r.filterMinHighestEquipmentLevel = &filterMinHighestEquipmentLevel
	return r
}

// only results where the equipment with the highest level is below or equal to this value
func (r ApiGetSetsListRequest) FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel int32) ApiGetSetsListRequest {
	r.filterMaxHighestEquipmentLevel = &filterMaxHighestEquipmentLevel
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetSetsListRequest) PageSize(pageSize int32) ApiGetSetsListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetSetsListRequest) PageNumber(pageNumber int32) ApiGetSetsListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetSetsListRequest) FieldsSet(fieldsSet []string) ApiGetSetsListRequest {
	r.fieldsSet = &fieldsSet
	return r
}

// filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment.
func (r ApiGetSetsListRequest) FilterContainsCosmeticsOnly(filterContainsCosmeticsOnly bool) ApiGetSetsListRequest {
	r.filterContainsCosmeticsOnly = &filterContainsCosmeticsOnly
	return r
}

// filter sets based on if they got cosmetic items in it.
func (r ApiGetSetsListRequest) FilterContainsCosmetics(filterContainsCosmetics bool) ApiGetSetsListRequest {
	r.filterContainsCosmetics = &filterContainsCosmetics
	return r
}

func (r ApiGetSetsListRequest) Execute() (*ListSets, *http.Response, error) {
	return r.ApiService.GetSetsListExecute(r)
}

/*
GetSetsList List Sets

Retrieve a list of sets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game dofus3 | dofus3beta
 @return ApiGetSetsListRequest
*/
func (a *SetsAPIService) GetSetsList(ctx context.Context, language string, game string) ApiGetSetsListRequest {
	return ApiGetSetsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListSets
func (a *SetsAPIService) GetSetsListExecute(r ApiGetSetsListRequest) (*ListSets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListSets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetsAPIService.GetSetsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/sets"
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
	if r.filterMinHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_highest_equipment_level]", r.filterMinHighestEquipmentLevel, "form", "")
	}
	if r.filterMaxHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_highest_equipment_level]", r.filterMaxHighestEquipmentLevel, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "form", "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "form", "")
	}
	if r.fieldsSet != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[set]", r.fieldsSet, "form", "csv")
	}
	if r.filterContainsCosmeticsOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[contains_cosmetics_only]", r.filterContainsCosmeticsOnly, "form", "")
	}
	if r.filterContainsCosmetics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[contains_cosmetics]", r.filterContainsCosmetics, "form", "")
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

type ApiGetSetsSearchRequest struct {
	ctx context.Context
	ApiService *SetsAPIService
	language string
	game string
	query *string
	filterMinHighestEquipmentLevel *int32
	filterMaxHighestEquipmentLevel *int32
	limit *int32
	filterIsCosmetic *bool
}

// case sensitive search query
func (r ApiGetSetsSearchRequest) Query(query string) ApiGetSetsSearchRequest {
	r.query = &query
	return r
}

// only results where the equipment with the highest level is above or equal to this value
func (r ApiGetSetsSearchRequest) FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel int32) ApiGetSetsSearchRequest {
	r.filterMinHighestEquipmentLevel = &filterMinHighestEquipmentLevel
	return r
}

// only results where the equipment with the highest level is below or equal to this value
func (r ApiGetSetsSearchRequest) FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel int32) ApiGetSetsSearchRequest {
	r.filterMaxHighestEquipmentLevel = &filterMaxHighestEquipmentLevel
	return r
}

// maximum number of returned results
func (r ApiGetSetsSearchRequest) Limit(limit int32) ApiGetSetsSearchRequest {
	r.limit = &limit
	return r
}

// filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment.
func (r ApiGetSetsSearchRequest) FilterIsCosmetic(filterIsCosmetic bool) ApiGetSetsSearchRequest {
	r.filterIsCosmetic = &filterIsCosmetic
	return r
}

func (r ApiGetSetsSearchRequest) Execute() ([]ListSet, *http.Response, error) {
	return r.ApiService.GetSetsSearchExecute(r)
}

/*
GetSetsSearch Search Sets

Search in all names and descriptions of sets with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game dofus3 | dofus3beta
 @return ApiGetSetsSearchRequest
*/
func (a *SetsAPIService) GetSetsSearch(ctx context.Context, language string, game string) ApiGetSetsSearchRequest {
	return ApiGetSetsSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ListSet
func (a *SetsAPIService) GetSetsSearchExecute(r ApiGetSetsSearchRequest) ([]ListSet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetsAPIService.GetSetsSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/sets/search"
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
	if r.filterMinHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_highest_equipment_level]", r.filterMinHighestEquipmentLevel, "form", "")
	}
	if r.filterMaxHighestEquipmentLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_highest_equipment_level]", r.filterMaxHighestEquipmentLevel, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 8
		r.limit = &defaultValue
	}
	if r.filterIsCosmetic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[is_cosmetic]", r.filterIsCosmetic, "form", "")
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

type ApiGetSetsSingleRequest struct {
	ctx context.Context
	ApiService *SetsAPIService
	language string
	ankamaId int32
	game string
}

func (r ApiGetSetsSingleRequest) Execute() (*Set, *http.Response, error) {
	return r.ApiService.GetSetsSingleExecute(r)
}

/*
GetSetsSingle Single Sets

Retrieve a specific set with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game dofus3 | dofus3beta
 @return ApiGetSetsSingleRequest
*/
func (a *SetsAPIService) GetSetsSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetSetsSingleRequest {
	return ApiGetSetsSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Set
func (a *SetsAPIService) GetSetsSingleExecute(r ApiGetSetsSingleRequest) (*Set, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Set
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetsAPIService.GetSetsSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/sets/{ankama_id}"
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
