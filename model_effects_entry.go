/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude - [PHP](https://github.com/dofusdude/dofusdude-php)  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply donate to that cause, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).

API version: 0.8.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the EffectsEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectsEntry{}

// EffectsEntry struct for EffectsEntry
type EffectsEntry struct {
	// minimum int value, can be a single if ignore_int_max and no ignore_int_min
	IntMinimum *int32 `json:"int_minimum,omitempty"`
	// maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value
	IntMaximum *int32 `json:"int_maximum,omitempty"`
	Type *SetEffectsEntryType `json:"type,omitempty"`
	// ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case. 
	IgnoreIntMin *bool `json:"ignore_int_min,omitempty"`
	// ignore the int max field, if ignore_int_min is true, int min is a single value
	IgnoreIntMax *bool `json:"ignore_int_max,omitempty"`
	// all fields from above encoded in a single string
	Formatted *string `json:"formatted,omitempty"`
}

// NewEffectsEntry instantiates a new EffectsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectsEntry() *EffectsEntry {
	this := EffectsEntry{}
	return &this
}

// NewEffectsEntryWithDefaults instantiates a new EffectsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectsEntryWithDefaults() *EffectsEntry {
	this := EffectsEntry{}
	return &this
}

// GetIntMinimum returns the IntMinimum field value if set, zero value otherwise.
func (o *EffectsEntry) GetIntMinimum() int32 {
	if o == nil || IsNil(o.IntMinimum) {
		var ret int32
		return ret
	}
	return *o.IntMinimum
}

// GetIntMinimumOk returns a tuple with the IntMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetIntMinimumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMinimum) {
		return nil, false
	}
	return o.IntMinimum, true
}

// HasIntMinimum returns a boolean if a field has been set.
func (o *EffectsEntry) HasIntMinimum() bool {
	if o != nil && !IsNil(o.IntMinimum) {
		return true
	}

	return false
}

// SetIntMinimum gets a reference to the given int32 and assigns it to the IntMinimum field.
func (o *EffectsEntry) SetIntMinimum(v int32) {
	o.IntMinimum = &v
}

// GetIntMaximum returns the IntMaximum field value if set, zero value otherwise.
func (o *EffectsEntry) GetIntMaximum() int32 {
	if o == nil || IsNil(o.IntMaximum) {
		var ret int32
		return ret
	}
	return *o.IntMaximum
}

// GetIntMaximumOk returns a tuple with the IntMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetIntMaximumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMaximum) {
		return nil, false
	}
	return o.IntMaximum, true
}

// HasIntMaximum returns a boolean if a field has been set.
func (o *EffectsEntry) HasIntMaximum() bool {
	if o != nil && !IsNil(o.IntMaximum) {
		return true
	}

	return false
}

// SetIntMaximum gets a reference to the given int32 and assigns it to the IntMaximum field.
func (o *EffectsEntry) SetIntMaximum(v int32) {
	o.IntMaximum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EffectsEntry) GetType() SetEffectsEntryType {
	if o == nil || IsNil(o.Type) {
		var ret SetEffectsEntryType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetTypeOk() (*SetEffectsEntryType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EffectsEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SetEffectsEntryType and assigns it to the Type field.
func (o *EffectsEntry) SetType(v SetEffectsEntryType) {
	o.Type = &v
}

// GetIgnoreIntMin returns the IgnoreIntMin field value if set, zero value otherwise.
func (o *EffectsEntry) GetIgnoreIntMin() bool {
	if o == nil || IsNil(o.IgnoreIntMin) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMin
}

// GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetIgnoreIntMinOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMin) {
		return nil, false
	}
	return o.IgnoreIntMin, true
}

// HasIgnoreIntMin returns a boolean if a field has been set.
func (o *EffectsEntry) HasIgnoreIntMin() bool {
	if o != nil && !IsNil(o.IgnoreIntMin) {
		return true
	}

	return false
}

// SetIgnoreIntMin gets a reference to the given bool and assigns it to the IgnoreIntMin field.
func (o *EffectsEntry) SetIgnoreIntMin(v bool) {
	o.IgnoreIntMin = &v
}

// GetIgnoreIntMax returns the IgnoreIntMax field value if set, zero value otherwise.
func (o *EffectsEntry) GetIgnoreIntMax() bool {
	if o == nil || IsNil(o.IgnoreIntMax) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMax
}

// GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetIgnoreIntMaxOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMax) {
		return nil, false
	}
	return o.IgnoreIntMax, true
}

// HasIgnoreIntMax returns a boolean if a field has been set.
func (o *EffectsEntry) HasIgnoreIntMax() bool {
	if o != nil && !IsNil(o.IgnoreIntMax) {
		return true
	}

	return false
}

// SetIgnoreIntMax gets a reference to the given bool and assigns it to the IgnoreIntMax field.
func (o *EffectsEntry) SetIgnoreIntMax(v bool) {
	o.IgnoreIntMax = &v
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *EffectsEntry) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntry) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *EffectsEntry) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *EffectsEntry) SetFormatted(v string) {
	o.Formatted = &v
}

func (o EffectsEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectsEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntMinimum) {
		toSerialize["int_minimum"] = o.IntMinimum
	}
	if !IsNil(o.IntMaximum) {
		toSerialize["int_maximum"] = o.IntMaximum
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IgnoreIntMin) {
		toSerialize["ignore_int_min"] = o.IgnoreIntMin
	}
	if !IsNil(o.IgnoreIntMax) {
		toSerialize["ignore_int_max"] = o.IgnoreIntMax
	}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	return toSerialize, nil
}

type NullableEffectsEntry struct {
	value *EffectsEntry
	isSet bool
}

func (v NullableEffectsEntry) Get() *EffectsEntry {
	return v.value
}

func (v *NullableEffectsEntry) Set(val *EffectsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectsEntry(val *EffectsEntry) *NullableEffectsEntry {
	return &NullableEffectsEntry{value: val, isSet: true}
}

func (v NullableEffectsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


