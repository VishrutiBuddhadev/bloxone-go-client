/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the NetworkListUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkListUpdateResponse{}

// NetworkListUpdateResponse The Network List update response.
type NetworkListUpdateResponse struct {
	Results              *NetworkList `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkListUpdateResponse NetworkListUpdateResponse

// NewNetworkListUpdateResponse instantiates a new NetworkListUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkListUpdateResponse() *NetworkListUpdateResponse {
	this := NetworkListUpdateResponse{}
	return &this
}

// NewNetworkListUpdateResponseWithDefaults instantiates a new NetworkListUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkListUpdateResponseWithDefaults() *NetworkListUpdateResponse {
	this := NetworkListUpdateResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *NetworkListUpdateResponse) GetResults() NetworkList {
	if o == nil || IsNil(o.Results) {
		var ret NetworkList
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListUpdateResponse) GetResultsOk() (*NetworkList, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *NetworkListUpdateResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given NetworkList and assigns it to the Results field.
func (o *NetworkListUpdateResponse) SetResults(v NetworkList) {
	o.Results = &v
}

func (o NetworkListUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkListUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varNetworkListUpdateResponse := _NetworkListUpdateResponse{}

	err = json.Unmarshal(data, &varNetworkListUpdateResponse)

	if err != nil {
		return err
	}

	*o = NetworkListUpdateResponse(varNetworkListUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkListUpdateResponse struct {
	value *NetworkListUpdateResponse
	isSet bool
}

func (v NullableNetworkListUpdateResponse) Get() *NetworkListUpdateResponse {
	return v.value
}

func (v *NullableNetworkListUpdateResponse) Set(val *NetworkListUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkListUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkListUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkListUpdateResponse(val *NetworkListUpdateResponse) *NullableNetworkListUpdateResponse {
	return &NullableNetworkListUpdateResponse{value: val, isSet: true}
}

func (v NullableNetworkListUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkListUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
