# IntraService5SettingsApiModelsClientLegalEntityModel62ED

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**TimeZoneCode** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Web** | Pointer to **NullableString** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**UserGroupTags** | Pointer to [**[]IntraService5SettingsApiModelsBaseNamedModel10D7**](IntraService5SettingsApiModelsBaseNamedModel10D7.md) |  | [optional] 
**FieldsValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9**](IntraService5SettingsApiModelsAdditionalFieldValueModel86F9.md) |  | [optional] 
**FieldIdValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7**](IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7.md) |  | [optional] 
**ManagerId** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TaxpayerNumber** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PhoneNumbersList** | Pointer to [**[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78**](IntraService5SettingsApiModelsPhoneNumbersListDtoCB78.md) |  | [optional] 
**ContactPersonId** | Pointer to **NullableInt64** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIntraService5SettingsApiModelsClientLegalEntityModel62ED

`func NewIntraService5SettingsApiModelsClientLegalEntityModel62ED() *IntraService5SettingsApiModelsClientLegalEntityModel62ED`

NewIntraService5SettingsApiModelsClientLegalEntityModel62ED instantiates a new IntraService5SettingsApiModelsClientLegalEntityModel62ED object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntraService5SettingsApiModelsClientLegalEntityModel62EDWithDefaults

`func NewIntraService5SettingsApiModelsClientLegalEntityModel62EDWithDefaults() *IntraService5SettingsApiModelsClientLegalEntityModel62ED`

NewIntraService5SettingsApiModelsClientLegalEntityModel62EDWithDefaults instantiates a new IntraService5SettingsApiModelsClientLegalEntityModel62ED object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetWeb() string`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetWebOk() (*string, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetWeb(v string)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### SetWebNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetWebNil(b bool)`

 SetWebNil sets the value for Web to be an explicit nil

### UnsetWeb
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetWeb()`

UnsetWeb ensures that no value is present for Web, not even an explicit nil
### GetDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUserGroupTags() []IntraService5SettingsApiModelsBaseNamedModel10D7`

GetUserGroupTags returns the UserGroupTags field if non-nil, zero value otherwise.

### GetUserGroupTagsOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetUserGroupTagsOk() (*[]IntraService5SettingsApiModelsBaseNamedModel10D7, bool)`

GetUserGroupTagsOk returns a tuple with the UserGroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetUserGroupTags(v []IntraService5SettingsApiModelsBaseNamedModel10D7)`

SetUserGroupTags sets UserGroupTags field to given value.

### HasUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasUserGroupTags() bool`

HasUserGroupTags returns a boolean if a field has been set.

### SetUserGroupTagsNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetUserGroupTagsNil(b bool)`

 SetUserGroupTagsNil sets the value for UserGroupTags to be an explicit nil

### UnsetUserGroupTags
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetUserGroupTags()`

UnsetUserGroupTags ensures that no value is present for UserGroupTags, not even an explicit nil
### GetFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetFieldsValues() []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9`

GetFieldsValues returns the FieldsValues field if non-nil, zero value otherwise.

### GetFieldsValuesOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetFieldsValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9, bool)`

GetFieldsValuesOk returns a tuple with the FieldsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetFieldsValues(v []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9)`

SetFieldsValues sets FieldsValues field to given value.

### HasFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasFieldsValues() bool`

HasFieldsValues returns a boolean if a field has been set.

### SetFieldsValuesNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetFieldsValuesNil(b bool)`

 SetFieldsValuesNil sets the value for FieldsValues to be an explicit nil

### UnsetFieldsValues
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetFieldsValues()`

UnsetFieldsValues ensures that no value is present for FieldsValues, not even an explicit nil
### GetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetFieldIdValues() []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7`

GetFieldIdValues returns the FieldIdValues field if non-nil, zero value otherwise.

### GetFieldIdValuesOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetFieldIdValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7, bool)`

GetFieldIdValuesOk returns a tuple with the FieldIdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetFieldIdValues(v []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7)`

SetFieldIdValues sets FieldIdValues field to given value.

### HasFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasFieldIdValues() bool`

HasFieldIdValues returns a boolean if a field has been set.

### SetFieldIdValuesNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetFieldIdValuesNil(b bool)`

 SetFieldIdValuesNil sets the value for FieldIdValues to be an explicit nil

### UnsetFieldIdValues
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetFieldIdValues()`

UnsetFieldIdValues ensures that no value is present for FieldIdValues, not even an explicit nil
### GetManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetManagerId() int64`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetManagerIdOk() (*int64, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetManagerId(v int64)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetName

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetTaxpayerNumber() string`

GetTaxpayerNumber returns the TaxpayerNumber field if non-nil, zero value otherwise.

### GetTaxpayerNumberOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetTaxpayerNumberOk() (*string, bool)`

GetTaxpayerNumberOk returns a tuple with the TaxpayerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetTaxpayerNumber(v string)`

SetTaxpayerNumber sets TaxpayerNumber field to given value.

### HasTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasTaxpayerNumber() bool`

HasTaxpayerNumber returns a boolean if a field has been set.

### SetTaxpayerNumberNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetTaxpayerNumberNil(b bool)`

 SetTaxpayerNumberNil sets the value for TaxpayerNumber to be an explicit nil

### UnsetTaxpayerNumber
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetTaxpayerNumber()`

UnsetTaxpayerNumber ensures that no value is present for TaxpayerNumber, not even an explicit nil
### GetEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetPhoneNumbersList() []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78`

GetPhoneNumbersList returns the PhoneNumbersList field if non-nil, zero value otherwise.

### GetPhoneNumbersListOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetPhoneNumbersListOk() (*[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78, bool)`

GetPhoneNumbersListOk returns a tuple with the PhoneNumbersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetPhoneNumbersList(v []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78)`

SetPhoneNumbersList sets PhoneNumbersList field to given value.

### HasPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasPhoneNumbersList() bool`

HasPhoneNumbersList returns a boolean if a field has been set.

### SetPhoneNumbersListNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetPhoneNumbersListNil(b bool)`

 SetPhoneNumbersListNil sets the value for PhoneNumbersList to be an explicit nil

### UnsetPhoneNumbersList
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetPhoneNumbersList()`

UnsetPhoneNumbersList ensures that no value is present for PhoneNumbersList, not even an explicit nil
### GetContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetContactPersonId() int64`

GetContactPersonId returns the ContactPersonId field if non-nil, zero value otherwise.

### GetContactPersonIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetContactPersonIdOk() (*int64, bool)`

GetContactPersonIdOk returns a tuple with the ContactPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetContactPersonId(v int64)`

SetContactPersonId sets ContactPersonId field to given value.

### HasContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasContactPersonId() bool`

HasContactPersonId returns a boolean if a field has been set.

### SetContactPersonIdNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetContactPersonIdNil(b bool)`

 SetContactPersonIdNil sets the value for ContactPersonId to be an explicit nil

### UnsetContactPersonId
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetContactPersonId()`

UnsetContactPersonId ensures that no value is present for ContactPersonId, not even an explicit nil
### GetAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *IntraService5SettingsApiModelsClientLegalEntityModel62ED) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


