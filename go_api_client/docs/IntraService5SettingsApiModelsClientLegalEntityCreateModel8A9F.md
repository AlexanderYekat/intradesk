# IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F

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
**ContactPersonPhoneNumbersList** | Pointer to [**[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78**](IntraService5SettingsApiModelsPhoneNumbersListDtoCB78.md) |  | [optional] 
**ContactPersonEmail** | Pointer to **NullableString** |  | [optional] 
**RoleId** | Pointer to **NullableInt64** |  | [optional] 
**ContactPersonFirstName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonLastName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonMiddleName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonPassword** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F

`func NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F() *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F`

NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F instantiates a new IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9FWithDefaults

`func NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9FWithDefaults() *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F`

NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9FWithDefaults instantiates a new IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetWeb() string`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetWebOk() (*string, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetWeb(v string)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### SetWebNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetWebNil(b bool)`

 SetWebNil sets the value for Web to be an explicit nil

### UnsetWeb
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetWeb()`

UnsetWeb ensures that no value is present for Web, not even an explicit nil
### GetDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUserGroupTags() []IntraService5SettingsApiModelsBaseNamedModel10D7`

GetUserGroupTags returns the UserGroupTags field if non-nil, zero value otherwise.

### GetUserGroupTagsOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetUserGroupTagsOk() (*[]IntraService5SettingsApiModelsBaseNamedModel10D7, bool)`

GetUserGroupTagsOk returns a tuple with the UserGroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetUserGroupTags(v []IntraService5SettingsApiModelsBaseNamedModel10D7)`

SetUserGroupTags sets UserGroupTags field to given value.

### HasUserGroupTags

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasUserGroupTags() bool`

HasUserGroupTags returns a boolean if a field has been set.

### SetUserGroupTagsNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetUserGroupTagsNil(b bool)`

 SetUserGroupTagsNil sets the value for UserGroupTags to be an explicit nil

### UnsetUserGroupTags
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetUserGroupTags()`

UnsetUserGroupTags ensures that no value is present for UserGroupTags, not even an explicit nil
### GetFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetFieldsValues() []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9`

GetFieldsValues returns the FieldsValues field if non-nil, zero value otherwise.

### GetFieldsValuesOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetFieldsValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9, bool)`

GetFieldsValuesOk returns a tuple with the FieldsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetFieldsValues(v []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9)`

SetFieldsValues sets FieldsValues field to given value.

### HasFieldsValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasFieldsValues() bool`

HasFieldsValues returns a boolean if a field has been set.

### SetFieldsValuesNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetFieldsValuesNil(b bool)`

 SetFieldsValuesNil sets the value for FieldsValues to be an explicit nil

### UnsetFieldsValues
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetFieldsValues()`

UnsetFieldsValues ensures that no value is present for FieldsValues, not even an explicit nil
### GetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetFieldIdValues() []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7`

GetFieldIdValues returns the FieldIdValues field if non-nil, zero value otherwise.

### GetFieldIdValuesOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetFieldIdValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7, bool)`

GetFieldIdValuesOk returns a tuple with the FieldIdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetFieldIdValues(v []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7)`

SetFieldIdValues sets FieldIdValues field to given value.

### HasFieldIdValues

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasFieldIdValues() bool`

HasFieldIdValues returns a boolean if a field has been set.

### SetFieldIdValuesNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetFieldIdValuesNil(b bool)`

 SetFieldIdValuesNil sets the value for FieldIdValues to be an explicit nil

### UnsetFieldIdValues
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetFieldIdValues()`

UnsetFieldIdValues ensures that no value is present for FieldIdValues, not even an explicit nil
### GetManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetManagerId() int64`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetManagerIdOk() (*int64, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetManagerId(v int64)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetTaxpayerNumber() string`

GetTaxpayerNumber returns the TaxpayerNumber field if non-nil, zero value otherwise.

### GetTaxpayerNumberOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetTaxpayerNumberOk() (*string, bool)`

GetTaxpayerNumberOk returns a tuple with the TaxpayerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetTaxpayerNumber(v string)`

SetTaxpayerNumber sets TaxpayerNumber field to given value.

### HasTaxpayerNumber

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasTaxpayerNumber() bool`

HasTaxpayerNumber returns a boolean if a field has been set.

### SetTaxpayerNumberNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetTaxpayerNumberNil(b bool)`

 SetTaxpayerNumberNil sets the value for TaxpayerNumber to be an explicit nil

### UnsetTaxpayerNumber
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetTaxpayerNumber()`

UnsetTaxpayerNumber ensures that no value is present for TaxpayerNumber, not even an explicit nil
### GetEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetPhoneNumbersList() []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78`

GetPhoneNumbersList returns the PhoneNumbersList field if non-nil, zero value otherwise.

### GetPhoneNumbersListOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetPhoneNumbersListOk() (*[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78, bool)`

GetPhoneNumbersListOk returns a tuple with the PhoneNumbersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetPhoneNumbersList(v []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78)`

SetPhoneNumbersList sets PhoneNumbersList field to given value.

### HasPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasPhoneNumbersList() bool`

HasPhoneNumbersList returns a boolean if a field has been set.

### SetPhoneNumbersListNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetPhoneNumbersListNil(b bool)`

 SetPhoneNumbersListNil sets the value for PhoneNumbersList to be an explicit nil

### UnsetPhoneNumbersList
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetPhoneNumbersList()`

UnsetPhoneNumbersList ensures that no value is present for PhoneNumbersList, not even an explicit nil
### GetContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonId() int64`

GetContactPersonId returns the ContactPersonId field if non-nil, zero value otherwise.

### GetContactPersonIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonIdOk() (*int64, bool)`

GetContactPersonIdOk returns a tuple with the ContactPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonId(v int64)`

SetContactPersonId sets ContactPersonId field to given value.

### HasContactPersonId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonId() bool`

HasContactPersonId returns a boolean if a field has been set.

### SetContactPersonIdNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonIdNil(b bool)`

 SetContactPersonIdNil sets the value for ContactPersonId to be an explicit nil

### UnsetContactPersonId
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonId()`

UnsetContactPersonId ensures that no value is present for ContactPersonId, not even an explicit nil
### GetAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonPhoneNumbersList() []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78`

GetContactPersonPhoneNumbersList returns the ContactPersonPhoneNumbersList field if non-nil, zero value otherwise.

### GetContactPersonPhoneNumbersListOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonPhoneNumbersListOk() (*[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78, bool)`

GetContactPersonPhoneNumbersListOk returns a tuple with the ContactPersonPhoneNumbersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonPhoneNumbersList(v []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78)`

SetContactPersonPhoneNumbersList sets ContactPersonPhoneNumbersList field to given value.

### HasContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonPhoneNumbersList() bool`

HasContactPersonPhoneNumbersList returns a boolean if a field has been set.

### SetContactPersonPhoneNumbersListNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonPhoneNumbersListNil(b bool)`

 SetContactPersonPhoneNumbersListNil sets the value for ContactPersonPhoneNumbersList to be an explicit nil

### UnsetContactPersonPhoneNumbersList
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonPhoneNumbersList()`

UnsetContactPersonPhoneNumbersList ensures that no value is present for ContactPersonPhoneNumbersList, not even an explicit nil
### GetContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonEmail() string`

GetContactPersonEmail returns the ContactPersonEmail field if non-nil, zero value otherwise.

### GetContactPersonEmailOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonEmailOk() (*string, bool)`

GetContactPersonEmailOk returns a tuple with the ContactPersonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonEmail(v string)`

SetContactPersonEmail sets ContactPersonEmail field to given value.

### HasContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonEmail() bool`

HasContactPersonEmail returns a boolean if a field has been set.

### SetContactPersonEmailNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonEmailNil(b bool)`

 SetContactPersonEmailNil sets the value for ContactPersonEmail to be an explicit nil

### UnsetContactPersonEmail
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonEmail()`

UnsetContactPersonEmail ensures that no value is present for ContactPersonEmail, not even an explicit nil
### GetRoleId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### SetRoleIdNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetRoleIdNil(b bool)`

 SetRoleIdNil sets the value for RoleId to be an explicit nil

### UnsetRoleId
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetRoleId()`

UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil
### GetContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonFirstName() string`

GetContactPersonFirstName returns the ContactPersonFirstName field if non-nil, zero value otherwise.

### GetContactPersonFirstNameOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonFirstNameOk() (*string, bool)`

GetContactPersonFirstNameOk returns a tuple with the ContactPersonFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonFirstName(v string)`

SetContactPersonFirstName sets ContactPersonFirstName field to given value.

### HasContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonFirstName() bool`

HasContactPersonFirstName returns a boolean if a field has been set.

### SetContactPersonFirstNameNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonFirstNameNil(b bool)`

 SetContactPersonFirstNameNil sets the value for ContactPersonFirstName to be an explicit nil

### UnsetContactPersonFirstName
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonFirstName()`

UnsetContactPersonFirstName ensures that no value is present for ContactPersonFirstName, not even an explicit nil
### GetContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonLastName() string`

GetContactPersonLastName returns the ContactPersonLastName field if non-nil, zero value otherwise.

### GetContactPersonLastNameOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonLastNameOk() (*string, bool)`

GetContactPersonLastNameOk returns a tuple with the ContactPersonLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonLastName(v string)`

SetContactPersonLastName sets ContactPersonLastName field to given value.

### HasContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonLastName() bool`

HasContactPersonLastName returns a boolean if a field has been set.

### SetContactPersonLastNameNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonLastNameNil(b bool)`

 SetContactPersonLastNameNil sets the value for ContactPersonLastName to be an explicit nil

### UnsetContactPersonLastName
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonLastName()`

UnsetContactPersonLastName ensures that no value is present for ContactPersonLastName, not even an explicit nil
### GetContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonMiddleName() string`

GetContactPersonMiddleName returns the ContactPersonMiddleName field if non-nil, zero value otherwise.

### GetContactPersonMiddleNameOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonMiddleNameOk() (*string, bool)`

GetContactPersonMiddleNameOk returns a tuple with the ContactPersonMiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonMiddleName(v string)`

SetContactPersonMiddleName sets ContactPersonMiddleName field to given value.

### HasContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonMiddleName() bool`

HasContactPersonMiddleName returns a boolean if a field has been set.

### SetContactPersonMiddleNameNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonMiddleNameNil(b bool)`

 SetContactPersonMiddleNameNil sets the value for ContactPersonMiddleName to be an explicit nil

### UnsetContactPersonMiddleName
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonMiddleName()`

UnsetContactPersonMiddleName ensures that no value is present for ContactPersonMiddleName, not even an explicit nil
### GetContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonPassword() string`

GetContactPersonPassword returns the ContactPersonPassword field if non-nil, zero value otherwise.

### GetContactPersonPasswordOk

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) GetContactPersonPasswordOk() (*string, bool)`

GetContactPersonPasswordOk returns a tuple with the ContactPersonPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonPassword(v string)`

SetContactPersonPassword sets ContactPersonPassword field to given value.

### HasContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) HasContactPersonPassword() bool`

HasContactPersonPassword returns a boolean if a field has been set.

### SetContactPersonPasswordNil

`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) SetContactPersonPasswordNil(b bool)`

 SetContactPersonPasswordNil sets the value for ContactPersonPassword to be an explicit nil

### UnsetContactPersonPassword
`func (o *IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F) UnsetContactPersonPassword()`

UnsetContactPersonPassword ensures that no value is present for ContactPersonPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


