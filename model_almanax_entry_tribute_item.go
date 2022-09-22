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

// AlmanaxEntryTributeItem struct for AlmanaxEntryTributeItem
type AlmanaxEntryTributeItem struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Name *string `json:"name,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
}

// NewAlmanaxEntryTributeItem instantiates a new AlmanaxEntryTributeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxEntryTributeItem() *AlmanaxEntryTributeItem {
	this := AlmanaxEntryTributeItem{}
	return &this
}

// NewAlmanaxEntryTributeItemWithDefaults instantiates a new AlmanaxEntryTributeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxEntryTributeItemWithDefaults() *AlmanaxEntryTributeItem {
	this := AlmanaxEntryTributeItem{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetAnkamaId() int32 {
	if o == nil || o.AnkamaId == nil {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || o.AnkamaId == nil {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasAnkamaId() bool {
	if o != nil && o.AnkamaId != nil {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *AlmanaxEntryTributeItem) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetImageUrls() ImageUrls {
	if o == nil || o.ImageUrls == nil {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || o.ImageUrls == nil {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasImageUrls() bool {
	if o != nil && o.ImageUrls != nil {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *AlmanaxEntryTributeItem) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlmanaxEntryTributeItem) SetName(v string) {
	o.Name = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *AlmanaxEntryTributeItem) SetSubtype(v string) {
	o.Subtype = &v
}

func (o AlmanaxEntryTributeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnkamaId != nil {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if o.ImageUrls != nil {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	return json.Marshal(toSerialize)
}

type NullableAlmanaxEntryTributeItem struct {
	value *AlmanaxEntryTributeItem
	isSet bool
}

func (v NullableAlmanaxEntryTributeItem) Get() *AlmanaxEntryTributeItem {
	return v.value
}

func (v *NullableAlmanaxEntryTributeItem) Set(val *AlmanaxEntryTributeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxEntryTributeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxEntryTributeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxEntryTributeItem(val *AlmanaxEntryTributeItem) *NullableAlmanaxEntryTributeItem {
	return &NullableAlmanaxEntryTributeItem{value: val, isSet: true}
}

func (v NullableAlmanaxEntryTributeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxEntryTributeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


