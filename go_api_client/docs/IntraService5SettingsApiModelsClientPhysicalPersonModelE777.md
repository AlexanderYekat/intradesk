# IntraService5SettingsApiModelsClientPhysicalPersonModelE777

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
**ContactPersonFirstName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonLastName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonMiddleName** | Pointer to **NullableString** |  | [optional] 
**ContactPersonPassword** | Pointer to **NullableString** |  | [optional] 
**ContactPersonRoleId** | Pointer to **NullableInt64** |  | [optional] 
**ContactPersonEmail** | Pointer to **NullableString** |  | [optional] 
**ContactPersonAdditionalEmails** | Pointer to **[]string** |  | [optional] 
**ContactPersonPhoneNumbersList** | Pointer to [**[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78**](IntraService5SettingsApiModelsPhoneNumbersListDtoCB78.md) |  | [optional] 
**ContactPersonFieldsValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9**](IntraService5SettingsApiModelsAdditionalFieldValueModel86F9.md) |  | [optional] 
**ContactPersonFieldIdValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7**](IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7.md) |  | [optional] 
**ContactPersonUserTags** | Pointer to [**[]IntraService5SettingsApiModelsBaseNamedModel10D7**](IntraService5SettingsApiModelsBaseNamedModel10D7.md) |  | [optional] 
**ContactPersonFotoImage** | Pointer to **NullableString** |  | [optional] 
**ContactPersonIsNewFoto** | Pointer to **bool** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 
**ContactPersonId** | Pointer to **NullableInt64** |  | [optional] 
**EmailConfirmed** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777

`func NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777() *IntraService5SettingsApiModelsClientPhysicalPersonModelE777`

NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777 instantiates a new IntraService5SettingsApiModelsClientPhysicalPersonModelE777 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777WithDefaults

`func NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777WithDefaults() *IntraService5SettingsApiModelsClientPhysicalPersonModelE777`

NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777WithDefaults instantiates a new IntraService5SettingsApiModelsClientPhysicalPersonModelE777 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsArchived

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetLanguage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAddress

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetWeb

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetWeb() string`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetWebOk() (*string, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetWeb(v string)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### SetWebNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetWebNil(b bool)`

 SetWebNil sets the value for Web to be an explicit nil

### UnsetWeb
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetWeb()`

UnsetWeb ensures that no value is present for Web, not even an explicit nil
### GetDomains

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUserGroupTags() []IntraService5SettingsApiModelsBaseNamedModel10D7`

GetUserGroupTags returns the UserGroupTags field if non-nil, zero value otherwise.

### GetUserGroupTagsOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetUserGroupTagsOk() (*[]IntraService5SettingsApiModelsBaseNamedModel10D7, bool)`

GetUserGroupTagsOk returns a tuple with the UserGroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetUserGroupTags(v []IntraService5SettingsApiModelsBaseNamedModel10D7)`

SetUserGroupTags sets UserGroupTags field to given value.

### HasUserGroupTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasUserGroupTags() bool`

HasUserGroupTags returns a boolean if a field has been set.

### SetUserGroupTagsNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetUserGroupTagsNil(b bool)`

 SetUserGroupTagsNil sets the value for UserGroupTags to be an explicit nil

### UnsetUserGroupTags
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetUserGroupTags()`

UnsetUserGroupTags ensures that no value is present for UserGroupTags, not even an explicit nil
### GetFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetFieldsValues() []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9`

GetFieldsValues returns the FieldsValues field if non-nil, zero value otherwise.

### GetFieldsValuesOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetFieldsValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9, bool)`

GetFieldsValuesOk returns a tuple with the FieldsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetFieldsValues(v []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9)`

SetFieldsValues sets FieldsValues field to given value.

### HasFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasFieldsValues() bool`

HasFieldsValues returns a boolean if a field has been set.

### SetFieldsValuesNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetFieldsValuesNil(b bool)`

 SetFieldsValuesNil sets the value for FieldsValues to be an explicit nil

### UnsetFieldsValues
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetFieldsValues()`

UnsetFieldsValues ensures that no value is present for FieldsValues, not even an explicit nil
### GetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetFieldIdValues() []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7`

GetFieldIdValues returns the FieldIdValues field if non-nil, zero value otherwise.

### GetFieldIdValuesOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetFieldIdValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7, bool)`

GetFieldIdValuesOk returns a tuple with the FieldIdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetFieldIdValues(v []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7)`

SetFieldIdValues sets FieldIdValues field to given value.

### HasFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasFieldIdValues() bool`

HasFieldIdValues returns a boolean if a field has been set.

### SetFieldIdValuesNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetFieldIdValuesNil(b bool)`

 SetFieldIdValuesNil sets the value for FieldIdValues to be an explicit nil

### UnsetFieldIdValues
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetFieldIdValues()`

UnsetFieldIdValues ensures that no value is present for FieldIdValues, not even an explicit nil
### GetManagerId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetManagerId() int64`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetManagerIdOk() (*int64, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetManagerId(v int64)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFirstName() string`

GetContactPersonFirstName returns the ContactPersonFirstName field if non-nil, zero value otherwise.

### GetContactPersonFirstNameOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFirstNameOk() (*string, bool)`

GetContactPersonFirstNameOk returns a tuple with the ContactPersonFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFirstName(v string)`

SetContactPersonFirstName sets ContactPersonFirstName field to given value.

### HasContactPersonFirstName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonFirstName() bool`

HasContactPersonFirstName returns a boolean if a field has been set.

### SetContactPersonFirstNameNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFirstNameNil(b bool)`

 SetContactPersonFirstNameNil sets the value for ContactPersonFirstName to be an explicit nil

### UnsetContactPersonFirstName
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonFirstName()`

UnsetContactPersonFirstName ensures that no value is present for ContactPersonFirstName, not even an explicit nil
### GetContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonLastName() string`

GetContactPersonLastName returns the ContactPersonLastName field if non-nil, zero value otherwise.

### GetContactPersonLastNameOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonLastNameOk() (*string, bool)`

GetContactPersonLastNameOk returns a tuple with the ContactPersonLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonLastName(v string)`

SetContactPersonLastName sets ContactPersonLastName field to given value.

### HasContactPersonLastName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonLastName() bool`

HasContactPersonLastName returns a boolean if a field has been set.

### SetContactPersonLastNameNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonLastNameNil(b bool)`

 SetContactPersonLastNameNil sets the value for ContactPersonLastName to be an explicit nil

### UnsetContactPersonLastName
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonLastName()`

UnsetContactPersonLastName ensures that no value is present for ContactPersonLastName, not even an explicit nil
### GetContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonMiddleName() string`

GetContactPersonMiddleName returns the ContactPersonMiddleName field if non-nil, zero value otherwise.

### GetContactPersonMiddleNameOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonMiddleNameOk() (*string, bool)`

GetContactPersonMiddleNameOk returns a tuple with the ContactPersonMiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonMiddleName(v string)`

SetContactPersonMiddleName sets ContactPersonMiddleName field to given value.

### HasContactPersonMiddleName

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonMiddleName() bool`

HasContactPersonMiddleName returns a boolean if a field has been set.

### SetContactPersonMiddleNameNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonMiddleNameNil(b bool)`

 SetContactPersonMiddleNameNil sets the value for ContactPersonMiddleName to be an explicit nil

### UnsetContactPersonMiddleName
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonMiddleName()`

UnsetContactPersonMiddleName ensures that no value is present for ContactPersonMiddleName, not even an explicit nil
### GetContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonPassword() string`

GetContactPersonPassword returns the ContactPersonPassword field if non-nil, zero value otherwise.

### GetContactPersonPasswordOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonPasswordOk() (*string, bool)`

GetContactPersonPasswordOk returns a tuple with the ContactPersonPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonPassword(v string)`

SetContactPersonPassword sets ContactPersonPassword field to given value.

### HasContactPersonPassword

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonPassword() bool`

HasContactPersonPassword returns a boolean if a field has been set.

### SetContactPersonPasswordNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonPasswordNil(b bool)`

 SetContactPersonPasswordNil sets the value for ContactPersonPassword to be an explicit nil

### UnsetContactPersonPassword
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonPassword()`

UnsetContactPersonPassword ensures that no value is present for ContactPersonPassword, not even an explicit nil
### GetContactPersonRoleId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonRoleId() int64`

GetContactPersonRoleId returns the ContactPersonRoleId field if non-nil, zero value otherwise.

### GetContactPersonRoleIdOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonRoleIdOk() (*int64, bool)`

GetContactPersonRoleIdOk returns a tuple with the ContactPersonRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonRoleId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonRoleId(v int64)`

SetContactPersonRoleId sets ContactPersonRoleId field to given value.

### HasContactPersonRoleId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonRoleId() bool`

HasContactPersonRoleId returns a boolean if a field has been set.

### SetContactPersonRoleIdNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonRoleIdNil(b bool)`

 SetContactPersonRoleIdNil sets the value for ContactPersonRoleId to be an explicit nil

### UnsetContactPersonRoleId
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonRoleId()`

UnsetContactPersonRoleId ensures that no value is present for ContactPersonRoleId, not even an explicit nil
### GetContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonEmail() string`

GetContactPersonEmail returns the ContactPersonEmail field if non-nil, zero value otherwise.

### GetContactPersonEmailOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonEmailOk() (*string, bool)`

GetContactPersonEmailOk returns a tuple with the ContactPersonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonEmail(v string)`

SetContactPersonEmail sets ContactPersonEmail field to given value.

### HasContactPersonEmail

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonEmail() bool`

HasContactPersonEmail returns a boolean if a field has been set.

### SetContactPersonEmailNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonEmailNil(b bool)`

 SetContactPersonEmailNil sets the value for ContactPersonEmail to be an explicit nil

### UnsetContactPersonEmail
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonEmail()`

UnsetContactPersonEmail ensures that no value is present for ContactPersonEmail, not even an explicit nil
### GetContactPersonAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonAdditionalEmails() []string`

GetContactPersonAdditionalEmails returns the ContactPersonAdditionalEmails field if non-nil, zero value otherwise.

### GetContactPersonAdditionalEmailsOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonAdditionalEmailsOk() (*[]string, bool)`

GetContactPersonAdditionalEmailsOk returns a tuple with the ContactPersonAdditionalEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonAdditionalEmails(v []string)`

SetContactPersonAdditionalEmails sets ContactPersonAdditionalEmails field to given value.

### HasContactPersonAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonAdditionalEmails() bool`

HasContactPersonAdditionalEmails returns a boolean if a field has been set.

### SetContactPersonAdditionalEmailsNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonAdditionalEmailsNil(b bool)`

 SetContactPersonAdditionalEmailsNil sets the value for ContactPersonAdditionalEmails to be an explicit nil

### UnsetContactPersonAdditionalEmails
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonAdditionalEmails()`

UnsetContactPersonAdditionalEmails ensures that no value is present for ContactPersonAdditionalEmails, not even an explicit nil
### GetContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonPhoneNumbersList() []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78`

GetContactPersonPhoneNumbersList returns the ContactPersonPhoneNumbersList field if non-nil, zero value otherwise.

### GetContactPersonPhoneNumbersListOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonPhoneNumbersListOk() (*[]IntraService5SettingsApiModelsPhoneNumbersListDtoCB78, bool)`

GetContactPersonPhoneNumbersListOk returns a tuple with the ContactPersonPhoneNumbersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonPhoneNumbersList(v []IntraService5SettingsApiModelsPhoneNumbersListDtoCB78)`

SetContactPersonPhoneNumbersList sets ContactPersonPhoneNumbersList field to given value.

### HasContactPersonPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonPhoneNumbersList() bool`

HasContactPersonPhoneNumbersList returns a boolean if a field has been set.

### SetContactPersonPhoneNumbersListNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonPhoneNumbersListNil(b bool)`

 SetContactPersonPhoneNumbersListNil sets the value for ContactPersonPhoneNumbersList to be an explicit nil

### UnsetContactPersonPhoneNumbersList
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonPhoneNumbersList()`

UnsetContactPersonPhoneNumbersList ensures that no value is present for ContactPersonPhoneNumbersList, not even an explicit nil
### GetContactPersonFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFieldsValues() []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9`

GetContactPersonFieldsValues returns the ContactPersonFieldsValues field if non-nil, zero value otherwise.

### GetContactPersonFieldsValuesOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFieldsValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9, bool)`

GetContactPersonFieldsValuesOk returns a tuple with the ContactPersonFieldsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFieldsValues(v []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9)`

SetContactPersonFieldsValues sets ContactPersonFieldsValues field to given value.

### HasContactPersonFieldsValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonFieldsValues() bool`

HasContactPersonFieldsValues returns a boolean if a field has been set.

### SetContactPersonFieldsValuesNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFieldsValuesNil(b bool)`

 SetContactPersonFieldsValuesNil sets the value for ContactPersonFieldsValues to be an explicit nil

### UnsetContactPersonFieldsValues
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonFieldsValues()`

UnsetContactPersonFieldsValues ensures that no value is present for ContactPersonFieldsValues, not even an explicit nil
### GetContactPersonFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFieldIdValues() []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7`

GetContactPersonFieldIdValues returns the ContactPersonFieldIdValues field if non-nil, zero value otherwise.

### GetContactPersonFieldIdValuesOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFieldIdValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7, bool)`

GetContactPersonFieldIdValuesOk returns a tuple with the ContactPersonFieldIdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFieldIdValues(v []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7)`

SetContactPersonFieldIdValues sets ContactPersonFieldIdValues field to given value.

### HasContactPersonFieldIdValues

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonFieldIdValues() bool`

HasContactPersonFieldIdValues returns a boolean if a field has been set.

### SetContactPersonFieldIdValuesNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFieldIdValuesNil(b bool)`

 SetContactPersonFieldIdValuesNil sets the value for ContactPersonFieldIdValues to be an explicit nil

### UnsetContactPersonFieldIdValues
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonFieldIdValues()`

UnsetContactPersonFieldIdValues ensures that no value is present for ContactPersonFieldIdValues, not even an explicit nil
### GetContactPersonUserTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonUserTags() []IntraService5SettingsApiModelsBaseNamedModel10D7`

GetContactPersonUserTags returns the ContactPersonUserTags field if non-nil, zero value otherwise.

### GetContactPersonUserTagsOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonUserTagsOk() (*[]IntraService5SettingsApiModelsBaseNamedModel10D7, bool)`

GetContactPersonUserTagsOk returns a tuple with the ContactPersonUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonUserTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonUserTags(v []IntraService5SettingsApiModelsBaseNamedModel10D7)`

SetContactPersonUserTags sets ContactPersonUserTags field to given value.

### HasContactPersonUserTags

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonUserTags() bool`

HasContactPersonUserTags returns a boolean if a field has been set.

### SetContactPersonUserTagsNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonUserTagsNil(b bool)`

 SetContactPersonUserTagsNil sets the value for ContactPersonUserTags to be an explicit nil

### UnsetContactPersonUserTags
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonUserTags()`

UnsetContactPersonUserTags ensures that no value is present for ContactPersonUserTags, not even an explicit nil
### GetContactPersonFotoImage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFotoImage() string`

GetContactPersonFotoImage returns the ContactPersonFotoImage field if non-nil, zero value otherwise.

### GetContactPersonFotoImageOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonFotoImageOk() (*string, bool)`

GetContactPersonFotoImageOk returns a tuple with the ContactPersonFotoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonFotoImage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFotoImage(v string)`

SetContactPersonFotoImage sets ContactPersonFotoImage field to given value.

### HasContactPersonFotoImage

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonFotoImage() bool`

HasContactPersonFotoImage returns a boolean if a field has been set.

### SetContactPersonFotoImageNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonFotoImageNil(b bool)`

 SetContactPersonFotoImageNil sets the value for ContactPersonFotoImage to be an explicit nil

### UnsetContactPersonFotoImage
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonFotoImage()`

UnsetContactPersonFotoImage ensures that no value is present for ContactPersonFotoImage, not even an explicit nil
### GetContactPersonIsNewFoto

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonIsNewFoto() bool`

GetContactPersonIsNewFoto returns the ContactPersonIsNewFoto field if non-nil, zero value otherwise.

### GetContactPersonIsNewFotoOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonIsNewFotoOk() (*bool, bool)`

GetContactPersonIsNewFotoOk returns a tuple with the ContactPersonIsNewFoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonIsNewFoto

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonIsNewFoto(v bool)`

SetContactPersonIsNewFoto sets ContactPersonIsNewFoto field to given value.

### HasContactPersonIsNewFoto

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonIsNewFoto() bool`

HasContactPersonIsNewFoto returns a boolean if a field has been set.

### GetAdditionalData

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetContactPersonId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonId() int64`

GetContactPersonId returns the ContactPersonId field if non-nil, zero value otherwise.

### GetContactPersonIdOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetContactPersonIdOk() (*int64, bool)`

GetContactPersonIdOk returns a tuple with the ContactPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonId(v int64)`

SetContactPersonId sets ContactPersonId field to given value.

### HasContactPersonId

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasContactPersonId() bool`

HasContactPersonId returns a boolean if a field has been set.

### SetContactPersonIdNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetContactPersonIdNil(b bool)`

 SetContactPersonIdNil sets the value for ContactPersonId to be an explicit nil

### UnsetContactPersonId
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetContactPersonId()`

UnsetContactPersonId ensures that no value is present for ContactPersonId, not even an explicit nil
### GetEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### SetEmailConfirmedNil

`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) SetEmailConfirmedNil(b bool)`

 SetEmailConfirmedNil sets the value for EmailConfirmed to be an explicit nil

### UnsetEmailConfirmed
`func (o *IntraService5SettingsApiModelsClientPhysicalPersonModelE777) UnsetEmailConfirmed()`

UnsetEmailConfirmed ensures that no value is present for EmailConfirmed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


