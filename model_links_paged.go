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

// checks if the LinksPaged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksPaged{}

// LinksPaged struct for LinksPaged
type LinksPaged struct {
	First NullableString `json:"first,omitempty"`
	Prev NullableString `json:"prev,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Last NullableString `json:"last,omitempty"`
}

// NewLinksPaged instantiates a new LinksPaged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksPaged() *LinksPaged {
	this := LinksPaged{}
	return &this
}

// NewLinksPagedWithDefaults instantiates a new LinksPaged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksPagedWithDefaults() *LinksPaged {
	this := LinksPaged{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinksPaged) GetFirst() string {
	if o == nil || IsNil(o.First.Get()) {
		var ret string
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinksPaged) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *LinksPaged) HasFirst() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableString and assigns it to the First field.
func (o *LinksPaged) SetFirst(v string) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *LinksPaged) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *LinksPaged) UnsetFirst() {
	o.First.Unset()
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinksPaged) GetPrev() string {
	if o == nil || IsNil(o.Prev.Get()) {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinksPaged) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *LinksPaged) HasPrev() bool {
	if o != nil && o.Prev.IsSet() {
		return true
	}

	return false
}

// SetPrev gets a reference to the given NullableString and assigns it to the Prev field.
func (o *LinksPaged) SetPrev(v string) {
	o.Prev.Set(&v)
}
// SetPrevNil sets the value for Prev to be an explicit nil
func (o *LinksPaged) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil
func (o *LinksPaged) UnsetPrev() {
	o.Prev.Unset()
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinksPaged) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinksPaged) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *LinksPaged) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *LinksPaged) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *LinksPaged) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *LinksPaged) UnsetNext() {
	o.Next.Unset()
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinksPaged) GetLast() string {
	if o == nil || IsNil(o.Last.Get()) {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinksPaged) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *LinksPaged) HasLast() bool {
	if o != nil && o.Last.IsSet() {
		return true
	}

	return false
}

// SetLast gets a reference to the given NullableString and assigns it to the Last field.
func (o *LinksPaged) SetLast(v string) {
	o.Last.Set(&v)
}
// SetLastNil sets the value for Last to be an explicit nil
func (o *LinksPaged) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil
func (o *LinksPaged) UnsetLast() {
	o.Last.Unset()
}

func (o LinksPaged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksPaged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Prev.IsSet() {
		toSerialize["prev"] = o.Prev.Get()
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	return toSerialize, nil
}

type NullableLinksPaged struct {
	value *LinksPaged
	isSet bool
}

func (v NullableLinksPaged) Get() *LinksPaged {
	return v.value
}

func (v *NullableLinksPaged) Set(val *LinksPaged) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksPaged) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksPaged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksPaged(val *LinksPaged) *NullableLinksPaged {
	return &NullableLinksPaged{value: val, isSet: true}
}

func (v NullableLinksPaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksPaged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


