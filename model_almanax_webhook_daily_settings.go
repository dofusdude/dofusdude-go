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
)

// checks if the AlmanaxWebhookDailySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlmanaxWebhookDailySettings{}

// AlmanaxWebhookDailySettings struct for AlmanaxWebhookDailySettings
type AlmanaxWebhookDailySettings struct {
	// Timezone of your community to determine midnight.
	Timezone *string `json:"timezone,omitempty"`
	// Hours added to midnight at the selected timezone. 8 = 8:00 in the morning.
	MidnightOffset *int32 `json:"midnight_offset,omitempty"`
}

// NewAlmanaxWebhookDailySettings instantiates a new AlmanaxWebhookDailySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxWebhookDailySettings() *AlmanaxWebhookDailySettings {
	this := AlmanaxWebhookDailySettings{}
	var timezone string = "Europe/Paris"
	this.Timezone = &timezone
	var midnightOffset int32 = 0
	this.MidnightOffset = &midnightOffset
	return &this
}

// NewAlmanaxWebhookDailySettingsWithDefaults instantiates a new AlmanaxWebhookDailySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxWebhookDailySettingsWithDefaults() *AlmanaxWebhookDailySettings {
	this := AlmanaxWebhookDailySettings{}
	var timezone string = "Europe/Paris"
	this.Timezone = &timezone
	var midnightOffset int32 = 0
	this.MidnightOffset = &midnightOffset
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AlmanaxWebhookDailySettings) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhookDailySettings) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AlmanaxWebhookDailySettings) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AlmanaxWebhookDailySettings) SetTimezone(v string) {
	o.Timezone = &v
}

// GetMidnightOffset returns the MidnightOffset field value if set, zero value otherwise.
func (o *AlmanaxWebhookDailySettings) GetMidnightOffset() int32 {
	if o == nil || IsNil(o.MidnightOffset) {
		var ret int32
		return ret
	}
	return *o.MidnightOffset
}

// GetMidnightOffsetOk returns a tuple with the MidnightOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhookDailySettings) GetMidnightOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.MidnightOffset) {
		return nil, false
	}
	return o.MidnightOffset, true
}

// HasMidnightOffset returns a boolean if a field has been set.
func (o *AlmanaxWebhookDailySettings) HasMidnightOffset() bool {
	if o != nil && !IsNil(o.MidnightOffset) {
		return true
	}

	return false
}

// SetMidnightOffset gets a reference to the given int32 and assigns it to the MidnightOffset field.
func (o *AlmanaxWebhookDailySettings) SetMidnightOffset(v int32) {
	o.MidnightOffset = &v
}

func (o AlmanaxWebhookDailySettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlmanaxWebhookDailySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.MidnightOffset) {
		toSerialize["midnight_offset"] = o.MidnightOffset
	}
	return toSerialize, nil
}

type NullableAlmanaxWebhookDailySettings struct {
	value *AlmanaxWebhookDailySettings
	isSet bool
}

func (v NullableAlmanaxWebhookDailySettings) Get() *AlmanaxWebhookDailySettings {
	return v.value
}

func (v *NullableAlmanaxWebhookDailySettings) Set(val *AlmanaxWebhookDailySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxWebhookDailySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxWebhookDailySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxWebhookDailySettings(val *AlmanaxWebhookDailySettings) *NullableAlmanaxWebhookDailySettings {
	return &NullableAlmanaxWebhookDailySettings{value: val, isSet: true}
}

func (v NullableAlmanaxWebhookDailySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxWebhookDailySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


