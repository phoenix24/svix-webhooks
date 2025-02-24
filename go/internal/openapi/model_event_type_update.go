/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`. For more information on authentication, please refer to the [authentication token docs](https://docs.svix.com/api-keys).  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.7.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EventTypeUpdate struct for EventTypeUpdate
type EventTypeUpdate struct {
	Archived *bool `json:"archived,omitempty"`
	Description string `json:"description"`
	FeatureFlag NullableString `json:"featureFlag,omitempty"`
	// The schema for the event type for a specific version as a JSON schema.
	Schemas map[string]map[string]interface{} `json:"schemas,omitempty"`
}

// NewEventTypeUpdate instantiates a new EventTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeUpdate(description string) *EventTypeUpdate {
	this := EventTypeUpdate{}
	var archived bool = false
	this.Archived = &archived
	this.Description = description
	return &this
}

// NewEventTypeUpdateWithDefaults instantiates a new EventTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeUpdateWithDefaults() *EventTypeUpdate {
	this := EventTypeUpdate{}
	var archived bool = false
	this.Archived = &archived
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *EventTypeUpdate) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeUpdate) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *EventTypeUpdate) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *EventTypeUpdate) SetArchived(v bool) {
	o.Archived = &v
}

// GetDescription returns the Description field value
func (o *EventTypeUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventTypeUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventTypeUpdate) SetDescription(v string) {
	o.Description = v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeUpdate) GetFeatureFlag() string {
	if o == nil || o.FeatureFlag.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeUpdate) GetFeatureFlagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *EventTypeUpdate) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *EventTypeUpdate) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *EventTypeUpdate) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *EventTypeUpdate) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetSchemas returns the Schemas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeUpdate) GetSchemas() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeUpdate) GetSchemasOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EventTypeUpdate) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]map[string]interface{} and assigns it to the Schemas field.
func (o *EventTypeUpdate) SetSchemas(v map[string]map[string]interface{}) {
	o.Schemas = v
}

func (o EventTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.FeatureFlag.IsSet() {
		toSerialize["featureFlag"] = o.FeatureFlag.Get()
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeUpdate struct {
	value *EventTypeUpdate
	isSet bool
}

func (v NullableEventTypeUpdate) Get() *EventTypeUpdate {
	return v.value
}

func (v *NullableEventTypeUpdate) Set(val *EventTypeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeUpdate(val *EventTypeUpdate) *NullableEventTypeUpdate {
	return &NullableEventTypeUpdate{value: val, isSet: true}
}

func (v NullableEventTypeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


