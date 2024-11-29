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

// checks if the EffectType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectType{}

// EffectType struct for EffectType
type EffectType struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Affects target or source actively.
	IsActive *bool `json:"is_active,omitempty"`
	// true if a type is generated from the Api instead of Ankama. In that case, always prefer showing the templated string and omit everything else. The \"name\" field will have an english description of the meta type. An example for such effects are class sets effects.
	IsMeta *bool `json:"is_meta,omitempty"`
}

// NewEffectType instantiates a new EffectType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectType() *EffectType {
	this := EffectType{}
	return &this
}

// NewEffectTypeWithDefaults instantiates a new EffectType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectTypeWithDefaults() *EffectType {
	this := EffectType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EffectType) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectType) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EffectType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EffectType) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EffectType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EffectType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EffectType) SetName(v string) {
	o.Name = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EffectType) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectType) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EffectType) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EffectType) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsMeta returns the IsMeta field value if set, zero value otherwise.
func (o *EffectType) GetIsMeta() bool {
	if o == nil || IsNil(o.IsMeta) {
		var ret bool
		return ret
	}
	return *o.IsMeta
}

// GetIsMetaOk returns a tuple with the IsMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectType) GetIsMetaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMeta) {
		return nil, false
	}
	return o.IsMeta, true
}

// HasIsMeta returns a boolean if a field has been set.
func (o *EffectType) HasIsMeta() bool {
	if o != nil && !IsNil(o.IsMeta) {
		return true
	}

	return false
}

// SetIsMeta gets a reference to the given bool and assigns it to the IsMeta field.
func (o *EffectType) SetIsMeta(v bool) {
	o.IsMeta = &v
}

func (o EffectType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsMeta) {
		toSerialize["is_meta"] = o.IsMeta
	}
	return toSerialize, nil
}

type NullableEffectType struct {
	value *EffectType
	isSet bool
}

func (v NullableEffectType) Get() *EffectType {
	return v.value
}

func (v *NullableEffectType) Set(val *EffectType) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectType) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectType(val *EffectType) *NullableEffectType {
	return &NullableEffectType{value: val, isSet: true}
}

func (v NullableEffectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


