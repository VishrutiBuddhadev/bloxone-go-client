/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsdata

import (
	"encoding/json"
)

// checks if the SOASerialIncrementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SOASerialIncrementRequest{}

// SOASerialIncrementRequest struct for SOASerialIncrementRequest
type SOASerialIncrementRequest struct {
	Fields *ProtobufFieldMask `json:"fields,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// Default increment SOA serial number by 1,000.
	SerialIncrement      *int64 `json:"serial_increment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SOASerialIncrementRequest SOASerialIncrementRequest

// NewSOASerialIncrementRequest instantiates a new SOASerialIncrementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSOASerialIncrementRequest() *SOASerialIncrementRequest {
	this := SOASerialIncrementRequest{}
	return &this
}

// NewSOASerialIncrementRequestWithDefaults instantiates a new SOASerialIncrementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSOASerialIncrementRequestWithDefaults() *SOASerialIncrementRequest {
	this := SOASerialIncrementRequest{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SOASerialIncrementRequest) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SOASerialIncrementRequest) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SOASerialIncrementRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *SOASerialIncrementRequest) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SOASerialIncrementRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SOASerialIncrementRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SOASerialIncrementRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SOASerialIncrementRequest) SetId(v string) {
	o.Id = &v
}

// GetSerialIncrement returns the SerialIncrement field value if set, zero value otherwise.
func (o *SOASerialIncrementRequest) GetSerialIncrement() int64 {
	if o == nil || IsNil(o.SerialIncrement) {
		var ret int64
		return ret
	}
	return *o.SerialIncrement
}

// GetSerialIncrementOk returns a tuple with the SerialIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SOASerialIncrementRequest) GetSerialIncrementOk() (*int64, bool) {
	if o == nil || IsNil(o.SerialIncrement) {
		return nil, false
	}
	return o.SerialIncrement, true
}

// HasSerialIncrement returns a boolean if a field has been set.
func (o *SOASerialIncrementRequest) HasSerialIncrement() bool {
	if o != nil && !IsNil(o.SerialIncrement) {
		return true
	}

	return false
}

// SetSerialIncrement gets a reference to the given int64 and assigns it to the SerialIncrement field.
func (o *SOASerialIncrementRequest) SetSerialIncrement(v int64) {
	o.SerialIncrement = &v
}

func (o SOASerialIncrementRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SOASerialIncrementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SerialIncrement) {
		toSerialize["serial_increment"] = o.SerialIncrement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SOASerialIncrementRequest) UnmarshalJSON(data []byte) (err error) {
	varSOASerialIncrementRequest := _SOASerialIncrementRequest{}

	err = json.Unmarshal(data, &varSOASerialIncrementRequest)

	if err != nil {
		return err
	}

	*o = SOASerialIncrementRequest(varSOASerialIncrementRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "id")
		delete(additionalProperties, "serial_increment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSOASerialIncrementRequest struct {
	value *SOASerialIncrementRequest
	isSet bool
}

func (v NullableSOASerialIncrementRequest) Get() *SOASerialIncrementRequest {
	return v.value
}

func (v *NullableSOASerialIncrementRequest) Set(val *SOASerialIncrementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSOASerialIncrementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSOASerialIncrementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSOASerialIncrementRequest(val *SOASerialIncrementRequest) *NullableSOASerialIncrementRequest {
	return &NullableSOASerialIncrementRequest{value: val, isSet: true}
}

func (v NullableSOASerialIncrementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSOASerialIncrementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
