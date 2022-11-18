/*
Dofusdude

# A project for you - the developer. The free, always-up-to-date, low-latency, insert-buzzword-here Ankama API for your next cool project!  ## Client SDKs Don't write types or functions yourself - I already (kinda) did! 😉 - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related Twitter, RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅ - Monsters ❌ - Spells ❌  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Only the beginning... 🤯 I want this project to be useful and not just add plain GET-categories no one needs.  There is a long list of features I want to add (see the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h)). But they are all focussed on you, the developers. So please let me know what you need. I will change the list based on demand.  # Get started! 🥳 Scroll down and try it for yourself!  Or see how these other awesome projects use it: - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  I highly recommend using the SDKs for quick results. I use them myself for microservices for the API.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply  donate, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).  

API version: 0.7.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// CreateTwitterWebhook 
type CreateTwitterWebhook struct {
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/twitter
	Subscriptions []string `json:"subscriptions"`
	Format string `json:"format"`
	PreviewLength NullableInt32 `json:"preview_length,omitempty"`
	// Discord Webhook URL
	Callback string `json:"callback"`
}

// NewCreateTwitterWebhook instantiates a new CreateTwitterWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTwitterWebhook(subscriptions []string, format string, callback string) *CreateTwitterWebhook {
	this := CreateTwitterWebhook{}
	this.Subscriptions = subscriptions
	this.Format = format
	this.Callback = callback
	return &this
}

// NewCreateTwitterWebhookWithDefaults instantiates a new CreateTwitterWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTwitterWebhookWithDefaults() *CreateTwitterWebhook {
	this := CreateTwitterWebhook{}
	return &this
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || isNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasWhitelist() bool {
	if o != nil && isNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *CreateTwitterWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || isNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasBlacklist() bool {
	if o != nil && isNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *CreateTwitterWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *CreateTwitterWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *CreateTwitterWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value
func (o *CreateTwitterWebhook) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CreateTwitterWebhook) SetFormat(v string) {
	o.Format = v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetPreviewLength() int32 {
	if o == nil || isNil(o.PreviewLength.Get()) {
		var ret int32
		return ret
	}
	return *o.PreviewLength.Get()
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewLength.Get(), o.PreviewLength.IsSet()
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasPreviewLength() bool {
	if o != nil && o.PreviewLength.IsSet() {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given NullableInt32 and assigns it to the PreviewLength field.
func (o *CreateTwitterWebhook) SetPreviewLength(v int32) {
	o.PreviewLength.Set(&v)
}
// SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil
func (o *CreateTwitterWebhook) SetPreviewLengthNil() {
	o.PreviewLength.Set(nil)
}

// UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil
func (o *CreateTwitterWebhook) UnsetPreviewLength() {
	o.PreviewLength.Unset()
}

// GetCallback returns the Callback field value
func (o *CreateTwitterWebhook) GetCallback() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetCallbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Callback, true
}

// SetCallback sets field value
func (o *CreateTwitterWebhook) SetCallback(v string) {
	o.Callback = v
}

func (o CreateTwitterWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if o.PreviewLength.IsSet() {
		toSerialize["preview_length"] = o.PreviewLength.Get()
	}
	if true {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTwitterWebhook struct {
	value *CreateTwitterWebhook
	isSet bool
}

func (v NullableCreateTwitterWebhook) Get() *CreateTwitterWebhook {
	return v.value
}

func (v *NullableCreateTwitterWebhook) Set(val *CreateTwitterWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTwitterWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTwitterWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTwitterWebhook(val *CreateTwitterWebhook) *NullableCreateTwitterWebhook {
	return &NullableCreateTwitterWebhook{value: val, isSet: true}
}

func (v NullableCreateTwitterWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTwitterWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


