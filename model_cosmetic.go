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

// Cosmetic struct for Cosmetic
type Cosmetic struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *ItemsListEntryTypedType `json:"type,omitempty"`
	Level *int32 `json:"level,omitempty"`
	Pods *int32 `json:"pods,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Effects []EffectsEntry `json:"effects,omitempty"`
	Conditions []ConditionEntry `json:"conditions,omitempty"`
	Recipe []RecipeEntry `json:"recipe,omitempty"`
}

// NewCosmetic instantiates a new Cosmetic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCosmetic() *Cosmetic {
	this := Cosmetic{}
	return &this
}

// NewCosmeticWithDefaults instantiates a new Cosmetic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCosmeticWithDefaults() *Cosmetic {
	this := Cosmetic{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *Cosmetic) GetAnkamaId() int32 {
	if o == nil || o.AnkamaId == nil {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || o.AnkamaId == nil {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *Cosmetic) HasAnkamaId() bool {
	if o != nil && o.AnkamaId != nil {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *Cosmetic) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Cosmetic) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Cosmetic) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Cosmetic) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cosmetic) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cosmetic) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cosmetic) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Cosmetic) GetType() ItemsListEntryTypedType {
	if o == nil || o.Type == nil {
		var ret ItemsListEntryTypedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetTypeOk() (*ItemsListEntryTypedType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Cosmetic) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ItemsListEntryTypedType and assigns it to the Type field.
func (o *Cosmetic) SetType(v ItemsListEntryTypedType) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Cosmetic) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Cosmetic) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *Cosmetic) SetLevel(v int32) {
	o.Level = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *Cosmetic) GetPods() int32 {
	if o == nil || o.Pods == nil {
		var ret int32
		return ret
	}
	return *o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetPodsOk() (*int32, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *Cosmetic) HasPods() bool {
	if o != nil && o.Pods != nil {
		return true
	}

	return false
}

// SetPods gets a reference to the given int32 and assigns it to the Pods field.
func (o *Cosmetic) SetPods(v int32) {
	o.Pods = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *Cosmetic) GetImageUrls() ImageUrls {
	if o == nil || o.ImageUrls == nil {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmetic) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || o.ImageUrls == nil {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *Cosmetic) HasImageUrls() bool {
	if o != nil && o.ImageUrls != nil {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *Cosmetic) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cosmetic) GetEffects() []EffectsEntry {
	if o == nil {
		var ret []EffectsEntry
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cosmetic) GetEffectsOk() ([]EffectsEntry, bool) {
	if o == nil || o.Effects == nil {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *Cosmetic) HasEffects() bool {
	if o != nil && o.Effects != nil {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []EffectsEntry and assigns it to the Effects field.
func (o *Cosmetic) SetEffects(v []EffectsEntry) {
	o.Effects = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cosmetic) GetConditions() []ConditionEntry {
	if o == nil {
		var ret []ConditionEntry
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cosmetic) GetConditionsOk() ([]ConditionEntry, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Cosmetic) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionEntry and assigns it to the Conditions field.
func (o *Cosmetic) SetConditions(v []ConditionEntry) {
	o.Conditions = v
}

// GetRecipe returns the Recipe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cosmetic) GetRecipe() []RecipeEntry {
	if o == nil {
		var ret []RecipeEntry
		return ret
	}
	return o.Recipe
}

// GetRecipeOk returns a tuple with the Recipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cosmetic) GetRecipeOk() ([]RecipeEntry, bool) {
	if o == nil || o.Recipe == nil {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *Cosmetic) HasRecipe() bool {
	if o != nil && o.Recipe != nil {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []RecipeEntry and assigns it to the Recipe field.
func (o *Cosmetic) SetRecipe(v []RecipeEntry) {
	o.Recipe = v
}

func (o Cosmetic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnkamaId != nil {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}
	if o.ImageUrls != nil {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	return json.Marshal(toSerialize)
}

type NullableCosmetic struct {
	value *Cosmetic
	isSet bool
}

func (v NullableCosmetic) Get() *Cosmetic {
	return v.value
}

func (v *NullableCosmetic) Set(val *Cosmetic) {
	v.value = val
	v.isSet = true
}

func (v NullableCosmetic) IsSet() bool {
	return v.isSet
}

func (v *NullableCosmetic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCosmetic(val *Cosmetic) *NullableCosmetic {
	return &NullableCosmetic{value: val, isSet: true}
}

func (v NullableCosmetic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCosmetic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


