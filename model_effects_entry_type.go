/*
Dofusdude

The last API for everything Dofus 🤯  ### JS Quickstart ```js var dofusdude = require(\"dofusdude-js\");  new dofusdude.AllItemsApi().getItemsAllSearch(   \"en\",   \"dofus2\",   \"nidas\",   { filterTypeName: \"hat\" },   (err, data, response) => {     console.log(data[0]);   } ); ```  ### Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue at the Docs Repo.  ## Main Features - 🥷 **seamless auto-update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **blazingly fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **search by relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅  - Monsters ❌ - Classes ❌ - Spells ❌ - Professions ❌   ### Maybes? I don't know what for 🤷 - Sidekicks ❌ - Haven Bags ❌ - Map ❌   ## Future I want this project to be useful and not just add plain categories no one needs. More and more features will be added to enhance the quality based on the needs of you, the developers.  Examples: _I need to know where I can drop the all the items I need to craft set X!_  _Please get a detailed always up-to-date spell description so I can calculate the damage for my set builder site!_  Nearly everything can be done. But I want to make sure somebody also wants it.  If you have anything or you are just interested in the project, join the [Discord](https://discord.gg/3EtHskZD8h).  ### Versioning Updating an API is a hard problem. This is why we'll keep it simple:  Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted. Ever.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  We can prevent this specific use case with a nice community but even then, it would be hidden behind a feature flag.  ## Get started! 🥳 Scroll down and try it for yourself!  If you are ready to use them in your project, think about [generating a client 🧙](https://github.com/OpenAPITools/openapi-generator) or use one of our pre generated SDKs linked at the top.  Awesome Projects using this API:  - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  

API version: 0.5.4
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// EffectsEntryType struct for EffectsEntryType
type EffectsEntryType struct {
	Name *string `json:"name,omitempty"`
	Id *int32 `json:"id,omitempty"`
	// true if a type is generated from the Api instead of Ankama. In that case, always prefer showing the templated string and omit everything else. The \"name\" field will have an english description of the meta type. An example for such effects are class sets effects.
	IsMeta *bool `json:"is_meta,omitempty"`
}

// NewEffectsEntryType instantiates a new EffectsEntryType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectsEntryType() *EffectsEntryType {
	this := EffectsEntryType{}
	return &this
}

// NewEffectsEntryTypeWithDefaults instantiates a new EffectsEntryType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectsEntryTypeWithDefaults() *EffectsEntryType {
	this := EffectsEntryType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EffectsEntryType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntryType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EffectsEntryType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EffectsEntryType) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EffectsEntryType) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntryType) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EffectsEntryType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EffectsEntryType) SetId(v int32) {
	o.Id = &v
}

// GetIsMeta returns the IsMeta field value if set, zero value otherwise.
func (o *EffectsEntryType) GetIsMeta() bool {
	if o == nil || o.IsMeta == nil {
		var ret bool
		return ret
	}
	return *o.IsMeta
}

// GetIsMetaOk returns a tuple with the IsMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectsEntryType) GetIsMetaOk() (*bool, bool) {
	if o == nil || o.IsMeta == nil {
		return nil, false
	}
	return o.IsMeta, true
}

// HasIsMeta returns a boolean if a field has been set.
func (o *EffectsEntryType) HasIsMeta() bool {
	if o != nil && o.IsMeta != nil {
		return true
	}

	return false
}

// SetIsMeta gets a reference to the given bool and assigns it to the IsMeta field.
func (o *EffectsEntryType) SetIsMeta(v bool) {
	o.IsMeta = &v
}

func (o EffectsEntryType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsMeta != nil {
		toSerialize["is_meta"] = o.IsMeta
	}
	return json.Marshal(toSerialize)
}

type NullableEffectsEntryType struct {
	value *EffectsEntryType
	isSet bool
}

func (v NullableEffectsEntryType) Get() *EffectsEntryType {
	return v.value
}

func (v *NullableEffectsEntryType) Set(val *EffectsEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectsEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectsEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectsEntryType(val *EffectsEntryType) *NullableEffectsEntryType {
	return &NullableEffectsEntryType{value: val, isSet: true}
}

func (v NullableEffectsEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectsEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


