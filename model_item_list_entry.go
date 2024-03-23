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

// checks if the ItemListEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemListEntry{}

// ItemListEntry struct for ItemListEntry
type ItemListEntry struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *ItemsListEntryTypedType `json:"type,omitempty"`
	Level *int32 `json:"level,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Recipe []RecipeEntry `json:"recipe,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// Deprecated
	Conditions []ConditionEntry `json:"conditions,omitempty"`
	ConditionTree *ConditionTreeNode `json:"condition_tree,omitempty"`
	Effects []EffectsEntry `json:"effects,omitempty"`
	IsWeapon NullableBool `json:"is_weapon,omitempty"`
	Pods NullableInt32 `json:"pods,omitempty"`
	ParentSet NullableItemListEntryParentSet `json:"parent_set,omitempty"`
	CriticalHitProbability NullableInt32 `json:"critical_hit_probability,omitempty"`
	CriticalHitBonus NullableInt32 `json:"critical_hit_bonus,omitempty"`
	IsTwoHanded NullableBool `json:"is_two_handed,omitempty"`
	MaxCastPerTurn NullableInt32 `json:"max_cast_per_turn,omitempty"`
	ApCost NullableInt32 `json:"ap_cost,omitempty"`
	Range NullableItemListEntryRange `json:"range,omitempty"`
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
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *ItemListEntry) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemListEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.Type) {
		var ret ItemsListEntryTypedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetTypeOk() (*ItemsListEntryTypedType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ItemListEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ItemListEntry) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
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
	if o == nil || IsNil(o.ImageUrls) {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *ItemListEntry) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
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
	if o == nil || IsNil(o.Recipe) {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *ItemListEntry) HasRecipe() bool {
	if o != nil && !IsNil(o.Recipe) {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []RecipeEntry and assigns it to the Recipe field.
func (o *ItemListEntry) SetRecipe(v []RecipeEntry) {
	o.Recipe = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ItemListEntry) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ItemListEntry) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ItemListEntry) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ItemListEntry) UnsetDescription() {
	o.Description.Unset()
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ItemListEntry) GetConditions() []ConditionEntry {
	if o == nil {
		var ret []ConditionEntry
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ItemListEntry) GetConditionsOk() ([]ConditionEntry, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ItemListEntry) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionEntry and assigns it to the Conditions field.
// Deprecated
func (o *ItemListEntry) SetConditions(v []ConditionEntry) {
	o.Conditions = v
}

// GetConditionTree returns the ConditionTree field value if set, zero value otherwise.
func (o *ItemListEntry) GetConditionTree() ConditionTreeNode {
	if o == nil || IsNil(o.ConditionTree) {
		var ret ConditionTreeNode
		return ret
	}
	return *o.ConditionTree
}

// GetConditionTreeOk returns a tuple with the ConditionTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntry) GetConditionTreeOk() (*ConditionTreeNode, bool) {
	if o == nil || IsNil(o.ConditionTree) {
		return nil, false
	}
	return o.ConditionTree, true
}

// HasConditionTree returns a boolean if a field has been set.
func (o *ItemListEntry) HasConditionTree() bool {
	if o != nil && !IsNil(o.ConditionTree) {
		return true
	}

	return false
}

// SetConditionTree gets a reference to the given ConditionTreeNode and assigns it to the ConditionTree field.
func (o *ItemListEntry) SetConditionTree(v ConditionTreeNode) {
	o.ConditionTree = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetEffects() []EffectsEntry {
	if o == nil {
		var ret []EffectsEntry
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetEffectsOk() ([]EffectsEntry, bool) {
	if o == nil || IsNil(o.Effects) {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *ItemListEntry) HasEffects() bool {
	if o != nil && !IsNil(o.Effects) {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []EffectsEntry and assigns it to the Effects field.
func (o *ItemListEntry) SetEffects(v []EffectsEntry) {
	o.Effects = v
}

// GetIsWeapon returns the IsWeapon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetIsWeapon() bool {
	if o == nil || IsNil(o.IsWeapon.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWeapon.Get()
}

// GetIsWeaponOk returns a tuple with the IsWeapon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetIsWeaponOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWeapon.Get(), o.IsWeapon.IsSet()
}

// HasIsWeapon returns a boolean if a field has been set.
func (o *ItemListEntry) HasIsWeapon() bool {
	if o != nil && o.IsWeapon.IsSet() {
		return true
	}

	return false
}

// SetIsWeapon gets a reference to the given NullableBool and assigns it to the IsWeapon field.
func (o *ItemListEntry) SetIsWeapon(v bool) {
	o.IsWeapon.Set(&v)
}
// SetIsWeaponNil sets the value for IsWeapon to be an explicit nil
func (o *ItemListEntry) SetIsWeaponNil() {
	o.IsWeapon.Set(nil)
}

// UnsetIsWeapon ensures that no value is present for IsWeapon, not even an explicit nil
func (o *ItemListEntry) UnsetIsWeapon() {
	o.IsWeapon.Unset()
}

// GetPods returns the Pods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetPods() int32 {
	if o == nil || IsNil(o.Pods.Get()) {
		var ret int32
		return ret
	}
	return *o.Pods.Get()
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetPodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pods.Get(), o.Pods.IsSet()
}

// HasPods returns a boolean if a field has been set.
func (o *ItemListEntry) HasPods() bool {
	if o != nil && o.Pods.IsSet() {
		return true
	}

	return false
}

// SetPods gets a reference to the given NullableInt32 and assigns it to the Pods field.
func (o *ItemListEntry) SetPods(v int32) {
	o.Pods.Set(&v)
}
// SetPodsNil sets the value for Pods to be an explicit nil
func (o *ItemListEntry) SetPodsNil() {
	o.Pods.Set(nil)
}

// UnsetPods ensures that no value is present for Pods, not even an explicit nil
func (o *ItemListEntry) UnsetPods() {
	o.Pods.Unset()
}

// GetParentSet returns the ParentSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetParentSet() ItemListEntryParentSet {
	if o == nil || IsNil(o.ParentSet.Get()) {
		var ret ItemListEntryParentSet
		return ret
	}
	return *o.ParentSet.Get()
}

// GetParentSetOk returns a tuple with the ParentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetParentSetOk() (*ItemListEntryParentSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSet.Get(), o.ParentSet.IsSet()
}

// HasParentSet returns a boolean if a field has been set.
func (o *ItemListEntry) HasParentSet() bool {
	if o != nil && o.ParentSet.IsSet() {
		return true
	}

	return false
}

// SetParentSet gets a reference to the given NullableItemListEntryParentSet and assigns it to the ParentSet field.
func (o *ItemListEntry) SetParentSet(v ItemListEntryParentSet) {
	o.ParentSet.Set(&v)
}
// SetParentSetNil sets the value for ParentSet to be an explicit nil
func (o *ItemListEntry) SetParentSetNil() {
	o.ParentSet.Set(nil)
}

// UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil
func (o *ItemListEntry) UnsetParentSet() {
	o.ParentSet.Unset()
}

// GetCriticalHitProbability returns the CriticalHitProbability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetCriticalHitProbability() int32 {
	if o == nil || IsNil(o.CriticalHitProbability.Get()) {
		var ret int32
		return ret
	}
	return *o.CriticalHitProbability.Get()
}

// GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetCriticalHitProbabilityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CriticalHitProbability.Get(), o.CriticalHitProbability.IsSet()
}

// HasCriticalHitProbability returns a boolean if a field has been set.
func (o *ItemListEntry) HasCriticalHitProbability() bool {
	if o != nil && o.CriticalHitProbability.IsSet() {
		return true
	}

	return false
}

// SetCriticalHitProbability gets a reference to the given NullableInt32 and assigns it to the CriticalHitProbability field.
func (o *ItemListEntry) SetCriticalHitProbability(v int32) {
	o.CriticalHitProbability.Set(&v)
}
// SetCriticalHitProbabilityNil sets the value for CriticalHitProbability to be an explicit nil
func (o *ItemListEntry) SetCriticalHitProbabilityNil() {
	o.CriticalHitProbability.Set(nil)
}

// UnsetCriticalHitProbability ensures that no value is present for CriticalHitProbability, not even an explicit nil
func (o *ItemListEntry) UnsetCriticalHitProbability() {
	o.CriticalHitProbability.Unset()
}

// GetCriticalHitBonus returns the CriticalHitBonus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetCriticalHitBonus() int32 {
	if o == nil || IsNil(o.CriticalHitBonus.Get()) {
		var ret int32
		return ret
	}
	return *o.CriticalHitBonus.Get()
}

// GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetCriticalHitBonusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CriticalHitBonus.Get(), o.CriticalHitBonus.IsSet()
}

// HasCriticalHitBonus returns a boolean if a field has been set.
func (o *ItemListEntry) HasCriticalHitBonus() bool {
	if o != nil && o.CriticalHitBonus.IsSet() {
		return true
	}

	return false
}

// SetCriticalHitBonus gets a reference to the given NullableInt32 and assigns it to the CriticalHitBonus field.
func (o *ItemListEntry) SetCriticalHitBonus(v int32) {
	o.CriticalHitBonus.Set(&v)
}
// SetCriticalHitBonusNil sets the value for CriticalHitBonus to be an explicit nil
func (o *ItemListEntry) SetCriticalHitBonusNil() {
	o.CriticalHitBonus.Set(nil)
}

// UnsetCriticalHitBonus ensures that no value is present for CriticalHitBonus, not even an explicit nil
func (o *ItemListEntry) UnsetCriticalHitBonus() {
	o.CriticalHitBonus.Unset()
}

// GetIsTwoHanded returns the IsTwoHanded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetIsTwoHanded() bool {
	if o == nil || IsNil(o.IsTwoHanded.Get()) {
		var ret bool
		return ret
	}
	return *o.IsTwoHanded.Get()
}

// GetIsTwoHandedOk returns a tuple with the IsTwoHanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetIsTwoHandedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTwoHanded.Get(), o.IsTwoHanded.IsSet()
}

// HasIsTwoHanded returns a boolean if a field has been set.
func (o *ItemListEntry) HasIsTwoHanded() bool {
	if o != nil && o.IsTwoHanded.IsSet() {
		return true
	}

	return false
}

// SetIsTwoHanded gets a reference to the given NullableBool and assigns it to the IsTwoHanded field.
func (o *ItemListEntry) SetIsTwoHanded(v bool) {
	o.IsTwoHanded.Set(&v)
}
// SetIsTwoHandedNil sets the value for IsTwoHanded to be an explicit nil
func (o *ItemListEntry) SetIsTwoHandedNil() {
	o.IsTwoHanded.Set(nil)
}

// UnsetIsTwoHanded ensures that no value is present for IsTwoHanded, not even an explicit nil
func (o *ItemListEntry) UnsetIsTwoHanded() {
	o.IsTwoHanded.Unset()
}

// GetMaxCastPerTurn returns the MaxCastPerTurn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetMaxCastPerTurn() int32 {
	if o == nil || IsNil(o.MaxCastPerTurn.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxCastPerTurn.Get()
}

// GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetMaxCastPerTurnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCastPerTurn.Get(), o.MaxCastPerTurn.IsSet()
}

// HasMaxCastPerTurn returns a boolean if a field has been set.
func (o *ItemListEntry) HasMaxCastPerTurn() bool {
	if o != nil && o.MaxCastPerTurn.IsSet() {
		return true
	}

	return false
}

// SetMaxCastPerTurn gets a reference to the given NullableInt32 and assigns it to the MaxCastPerTurn field.
func (o *ItemListEntry) SetMaxCastPerTurn(v int32) {
	o.MaxCastPerTurn.Set(&v)
}
// SetMaxCastPerTurnNil sets the value for MaxCastPerTurn to be an explicit nil
func (o *ItemListEntry) SetMaxCastPerTurnNil() {
	o.MaxCastPerTurn.Set(nil)
}

// UnsetMaxCastPerTurn ensures that no value is present for MaxCastPerTurn, not even an explicit nil
func (o *ItemListEntry) UnsetMaxCastPerTurn() {
	o.MaxCastPerTurn.Unset()
}

// GetApCost returns the ApCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetApCost() int32 {
	if o == nil || IsNil(o.ApCost.Get()) {
		var ret int32
		return ret
	}
	return *o.ApCost.Get()
}

// GetApCostOk returns a tuple with the ApCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetApCostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApCost.Get(), o.ApCost.IsSet()
}

// HasApCost returns a boolean if a field has been set.
func (o *ItemListEntry) HasApCost() bool {
	if o != nil && o.ApCost.IsSet() {
		return true
	}

	return false
}

// SetApCost gets a reference to the given NullableInt32 and assigns it to the ApCost field.
func (o *ItemListEntry) SetApCost(v int32) {
	o.ApCost.Set(&v)
}
// SetApCostNil sets the value for ApCost to be an explicit nil
func (o *ItemListEntry) SetApCostNil() {
	o.ApCost.Set(nil)
}

// UnsetApCost ensures that no value is present for ApCost, not even an explicit nil
func (o *ItemListEntry) UnsetApCost() {
	o.ApCost.Unset()
}

// GetRange returns the Range field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemListEntry) GetRange() ItemListEntryRange {
	if o == nil || IsNil(o.Range.Get()) {
		var ret ItemListEntryRange
		return ret
	}
	return *o.Range.Get()
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemListEntry) GetRangeOk() (*ItemListEntryRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Range.Get(), o.Range.IsSet()
}

// HasRange returns a boolean if a field has been set.
func (o *ItemListEntry) HasRange() bool {
	if o != nil && o.Range.IsSet() {
		return true
	}

	return false
}

// SetRange gets a reference to the given NullableItemListEntryRange and assigns it to the Range field.
func (o *ItemListEntry) SetRange(v ItemListEntryRange) {
	o.Range.Set(&v)
}
// SetRangeNil sets the value for Range to be an explicit nil
func (o *ItemListEntry) SetRangeNil() {
	o.Range.Set(nil)
}

// UnsetRange ensures that no value is present for Range, not even an explicit nil
func (o *ItemListEntry) UnsetRange() {
	o.Range.Unset()
}

func (o ItemListEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemListEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ConditionTree) {
		toSerialize["condition_tree"] = o.ConditionTree
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	if o.IsWeapon.IsSet() {
		toSerialize["is_weapon"] = o.IsWeapon.Get()
	}
	if o.Pods.IsSet() {
		toSerialize["pods"] = o.Pods.Get()
	}
	if o.ParentSet.IsSet() {
		toSerialize["parent_set"] = o.ParentSet.Get()
	}
	if o.CriticalHitProbability.IsSet() {
		toSerialize["critical_hit_probability"] = o.CriticalHitProbability.Get()
	}
	if o.CriticalHitBonus.IsSet() {
		toSerialize["critical_hit_bonus"] = o.CriticalHitBonus.Get()
	}
	if o.IsTwoHanded.IsSet() {
		toSerialize["is_two_handed"] = o.IsTwoHanded.Get()
	}
	if o.MaxCastPerTurn.IsSet() {
		toSerialize["max_cast_per_turn"] = o.MaxCastPerTurn.Get()
	}
	if o.ApCost.IsSet() {
		toSerialize["ap_cost"] = o.ApCost.Get()
	}
	if o.Range.IsSet() {
		toSerialize["range"] = o.Range.Get()
	}
	return toSerialize, nil
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


