/*
Dofusdude

# A project for you - the developer. The free, always-up-to-date, low-latency, insert-buzzword-here Ankama API for your next cool project!  ## Client SDKs Don't write types or functions yourself - I already (kinda) did! 😉 - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅ - Monsters ❌ - Spells ❌  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Only the beginning... 🤯 I want this project to be useful and not just add plain GET-categories no one needs.  There is a long list of features I want to add (see the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h)). But they are all focussed on you, the developers. So please let me know what you need. I will change the list based on demand.  # Get started! 🥳 Scroll down and try it for yourself!  Or see how these other awesome projects use it: - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  I highly recommend using the SDKs for quick results. I use them myself for microservices for the API.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply  donate, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).  

API version: 0.6.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// ItemListEntry struct for ItemListEntry
type ItemListEntry struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *ItemsListEntryTypedType `json:"type,omitempty"`
	Level *int32 `json:"level,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Recipe []RecipeEntry `json:"recipe,omitempty"`
}

// NewItemListEntry instantiates a new ItemListEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemListEntry() *ItemListEntry {
	this := ItemListEntry{}
	return &this
}

// NewItemListEntryWithDefaults instantiates a new ItemListEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemListEntryWithDefaults() *ItemListEntry {
	this := ItemListEntry{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *ItemListEntry) GetAnkamaId() int32 {
	if o == nil || o.AnkamaId == nil {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || o.AnkamaId == nil {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *ItemListEntry) HasAnkamaId() bool {
	if o != nil && o.AnkamaId != nil {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *ItemListEntry) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ItemListEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemListEntry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ItemListEntry) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ItemListEntry) GetType() ItemsListEntryTypedType {
	if o == nil || o.Type == nil {
		var ret ItemsListEntryTypedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetTypeOk() (*ItemsListEntryTypedType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ItemListEntry) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ItemsListEntryTypedType and assigns it to the Type field.
func (o *ItemListEntry) SetType(v ItemsListEntryTypedType) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ItemListEntry) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ItemListEntry) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *ItemListEntry) SetLevel(v int32) {
	o.Level = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *ItemListEntry) GetImageUrls() ImageUrls {
	if o == nil || o.ImageUrls == nil {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || o.ImageUrls == nil {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *ItemListEntry) HasImageUrls() bool {
	if o != nil && o.ImageUrls != nil {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *ItemListEntry) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetRecipe returns the Recipe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetRecipe() []RecipeEntry {
	if o == nil {
		var ret []RecipeEntry
		return ret
	}
	return o.Recipe
}

// GetRecipeOk returns a tuple with the Recipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetRecipeOk() ([]RecipeEntry, bool) {
	if o == nil || o.Recipe == nil {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *ItemListEntry) HasRecipe() bool {
	if o != nil && o.Recipe != nil {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []RecipeEntry and assigns it to the Recipe field.
func (o *ItemListEntry) SetRecipe(v []RecipeEntry) {
	o.Recipe = v
}

func (o ItemListEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnkamaId != nil {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.ImageUrls != nil {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	return json.Marshal(toSerialize)
}

type NullableItemListEntry struct {
	value *ItemListEntry
	isSet bool
}

func (v NullableItemListEntry) Get() *ItemListEntry {
	return v.value
}

func (v *NullableItemListEntry) Set(val *ItemListEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableItemListEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableItemListEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemListEntry(val *ItemListEntry) *NullableItemListEntry {
	return &NullableItemListEntry{value: val, isSet: true}
}

func (v NullableItemListEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemListEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


