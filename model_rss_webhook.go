/*
Dofusdude

# A project for you - the developer. The free, always-up-to-date, low-latency, insert-buzzword-here Ankama API for your next cool project!  ## Client SDKs Don't write types or functions yourself - I already (kinda) did! 😉 - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude - [PHP](https://github.com/dofusdude/dofusdude-php)  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related Twitter, RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅ - Monsters ❌ - Spells ❌  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Only the beginning... 🤯 I want this project to be useful and not just add plain GET-categories no one needs.  There is a long list of features I want to add (see the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h)). But they are all focussed on you, the developers. So please let me know what you need. I will change the list based on demand.  # Get started! 🥳 Scroll down and try it for yourself!  Or see how these other awesome projects use it: - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [DofuStuffSimulator](https://dofusstuffsimulator.netlify.app/)  I highly recommend using the SDKs for quick results. I use them myself for parts of the API.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply  donate, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude). 

API version: 0.7.2
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
	"time"
)

// checks if the RssWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RssWebhook{}

// RssWebhook struct for RssWebhook
type RssWebhook struct {
	Id *string `json:"id,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	Format *string `json:"format,omitempty"`
	PreviewLength *int32 `json:"preview_length,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastFiredAt NullableTime `json:"last_fired_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewRssWebhook instantiates a new RssWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRssWebhook() *RssWebhook {
	this := RssWebhook{}
	return &this
}

// NewRssWebhookWithDefaults instantiates a new RssWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRssWebhookWithDefaults() *RssWebhook {
	this := RssWebhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RssWebhook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RssWebhook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RssWebhook) SetId(v string) {
	o.Id = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RssWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RssWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *RssWebhook) HasWhitelist() bool {
	if o != nil && IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *RssWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RssWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RssWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *RssWebhook) HasBlacklist() bool {
	if o != nil && IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *RssWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *RssWebhook) GetSubscriptions() []string {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *RssWebhook) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *RssWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *RssWebhook) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *RssWebhook) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *RssWebhook) SetFormat(v string) {
	o.Format = &v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise.
func (o *RssWebhook) GetPreviewLength() int32 {
	if o == nil || IsNil(o.PreviewLength) {
		var ret int32
		return ret
	}
	return *o.PreviewLength
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PreviewLength) {
		return nil, false
	}
	return o.PreviewLength, true
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *RssWebhook) HasPreviewLength() bool {
	if o != nil && !IsNil(o.PreviewLength) {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given int32 and assigns it to the PreviewLength field.
func (o *RssWebhook) SetPreviewLength(v int32) {
	o.PreviewLength = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RssWebhook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RssWebhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RssWebhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastFiredAt returns the LastFiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RssWebhook) GetLastFiredAt() time.Time {
	if o == nil || IsNil(o.LastFiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastFiredAt.Get()
}

// GetLastFiredAtOk returns a tuple with the LastFiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RssWebhook) GetLastFiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFiredAt.Get(), o.LastFiredAt.IsSet()
}

// HasLastFiredAt returns a boolean if a field has been set.
func (o *RssWebhook) HasLastFiredAt() bool {
	if o != nil && o.LastFiredAt.IsSet() {
		return true
	}

	return false
}

// SetLastFiredAt gets a reference to the given NullableTime and assigns it to the LastFiredAt field.
func (o *RssWebhook) SetLastFiredAt(v time.Time) {
	o.LastFiredAt.Set(&v)
}
// SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil
func (o *RssWebhook) SetLastFiredAtNil() {
	o.LastFiredAt.Set(nil)
}

// UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
func (o *RssWebhook) UnsetLastFiredAt() {
	o.LastFiredAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RssWebhook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RssWebhook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RssWebhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RssWebhook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o RssWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RssWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.PreviewLength) {
		toSerialize["preview_length"] = o.PreviewLength
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastFiredAt.IsSet() {
		toSerialize["last_fired_at"] = o.LastFiredAt.Get()
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableRssWebhook struct {
	value *RssWebhook
	isSet bool
}

func (v NullableRssWebhook) Get() *RssWebhook {
	return v.value
}

func (v *NullableRssWebhook) Set(val *RssWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRssWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRssWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRssWebhook(val *RssWebhook) *NullableRssWebhook {
	return &NullableRssWebhook{value: val, isSet: true}
}

func (v NullableRssWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRssWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


