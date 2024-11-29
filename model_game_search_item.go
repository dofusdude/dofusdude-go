/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.2
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the GameSearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameSearchItem{}

// GameSearchItem struct for GameSearchItem
type GameSearchItem struct {
	Type *GameSearchType `json:"type,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	ImageUrls *Images `json:"image_urls,omitempty"`
}

// NewGameSearchItem instantiates a new GameSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameSearchItem() *GameSearchItem {
	this := GameSearchItem{}
	return &this
}

// NewGameSearchItemWithDefaults instantiates a new GameSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameSearchItemWithDefaults() *GameSearchItem {
	this := GameSearchItem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GameSearchItem) GetType() GameSearchType {
	if o == nil || IsNil(o.Type) {
		var ret GameSearchType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearchItem) GetTypeOk() (*GameSearchType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GameSearchItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GameSearchType and assigns it to the Type field.
func (o *GameSearchItem) SetType(v GameSearchType) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameSearchItem) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameSearchItem) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *GameSearchItem) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *GameSearchItem) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *GameSearchItem) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *GameSearchItem) UnsetLevel() {
	o.Level.Unset()
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *GameSearchItem) GetImageUrls() Images {
	if o == nil || IsNil(o.ImageUrls) {
		var ret Images
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearchItem) GetImageUrlsOk() (*Images, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *GameSearchItem) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given Images and assigns it to the ImageUrls field.
func (o *GameSearchItem) SetImageUrls(v Images) {
	o.ImageUrls = &v
}

func (o GameSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameSearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	return toSerialize, nil
}

type NullableGameSearchItem struct {
	value *GameSearchItem
	isSet bool
}

func (v NullableGameSearchItem) Get() *GameSearchItem {
	return v.value
}

func (v *NullableGameSearchItem) Set(val *GameSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGameSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGameSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameSearchItem(val *GameSearchItem) *NullableGameSearchItem {
	return &NullableGameSearchItem{value: val, isSet: true}
}

func (v NullableGameSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


