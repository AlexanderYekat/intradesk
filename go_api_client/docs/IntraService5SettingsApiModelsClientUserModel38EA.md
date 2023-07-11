# IntraService5SettingsApiModelsClientUserModel38EA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**MiddleName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PhoneNumbersList** | Pointer to [**[]IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D**](IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D.md) |  | [optional] 
**AdditionalEmails** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**IntraService5SharedEnumsUserTypeE5AC**](IntraService5SharedEnumsUserTypeE5AC.md) |  | [optional] 
**RoleId** | Pointer to **NullableInt64** |  | [optional] 
**TimeZoneCode** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**UserGroups** | Pointer to [**[]IntraService5SettingsApiModelsUserGroupDto2E44**](IntraService5SettingsApiModelsUserGroupDto2E44.md) |  | [optional] 
**UserTags** | Pointer to [**[]IntraService5SettingsApiModelsBaseNamedModel10D7**](IntraService5SettingsApiModelsBaseNamedModel10D7.md) |  | [optional] 
**FieldsValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9**](IntraService5SettingsApiModelsAdditionalFieldValueModel86F9.md) |  | [optional] 
**FotoImage** | Pointer to **NullableString** |  | [optional] 
**IsNewFoto** | Pointer to **bool** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailConfirmed** | Pointer to **NullableBool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**IsArchived** | Pointer to **NullableBool** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**SetContactPerson** | Pointer to **bool** |  | [optional] 
**PhoneNumbers** | Pointer to [**[]IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D**](IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D.md) |  | [optional] 
**FieldIdValues** | Pointer to [**[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7**](IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7.md) |  | [optional] 

## Methods

### NewIntraService5SettingsApiModelsClientUserModel38EA

`func NewIntraService5SettingsApiModelsClientUserModel38EA() *IntraService5SettingsApiModelsClientUserModel38EA`

NewIntraService5SettingsApiModelsClientUserModel38EA instantiates a new IntraService5SettingsApiModelsClientUserModel38EA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntraService5SettingsApiModelsClientUserModel38EAWithDefaults

`func NewIntraService5SettingsApiModelsClientUserModel38EAWithDefaults() *IntraService5SettingsApiModelsClientUserModel38EA`

NewIntraService5SettingsApiModelsClientUserModel38EAWithDefaults instantiates a new IntraService5SettingsApiModelsClientUserModel38EA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFirstName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMiddleName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetEmail

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPhoneNumbersList() []IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D`

GetPhoneNumbersList returns the PhoneNumbersList field if non-nil, zero value otherwise.

### GetPhoneNumbersListOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPhoneNumbersListOk() (*[]IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D, bool)`

GetPhoneNumbersListOk returns a tuple with the PhoneNumbersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPhoneNumbersList(v []IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D)`

SetPhoneNumbersList sets PhoneNumbersList field to given value.

### HasPhoneNumbersList

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasPhoneNumbersList() bool`

HasPhoneNumbersList returns a boolean if a field has been set.

### SetPhoneNumbersListNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPhoneNumbersListNil(b bool)`

 SetPhoneNumbersListNil sets the value for PhoneNumbersList to be an explicit nil

### UnsetPhoneNumbersList
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetPhoneNumbersList()`

UnsetPhoneNumbersList ensures that no value is present for PhoneNumbersList, not even an explicit nil
### GetAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetAdditionalEmails() []string`

GetAdditionalEmails returns the AdditionalEmails field if non-nil, zero value otherwise.

### GetAdditionalEmailsOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetAdditionalEmailsOk() (*[]string, bool)`

GetAdditionalEmailsOk returns a tuple with the AdditionalEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetAdditionalEmails(v []string)`

SetAdditionalEmails sets AdditionalEmails field to given value.

### HasAdditionalEmails

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasAdditionalEmails() bool`

HasAdditionalEmails returns a boolean if a field has been set.

### SetAdditionalEmailsNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetAdditionalEmailsNil(b bool)`

 SetAdditionalEmailsNil sets the value for AdditionalEmails to be an explicit nil

### UnsetAdditionalEmails
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetAdditionalEmails()`

UnsetAdditionalEmails ensures that no value is present for AdditionalEmails, not even an explicit nil
### GetType

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetType() IntraService5SharedEnumsUserTypeE5AC`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetTypeOk() (*IntraService5SharedEnumsUserTypeE5AC, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetType(v IntraService5SharedEnumsUserTypeE5AC)`

SetType sets Type field to given value.

### HasType

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoleId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### SetRoleIdNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetRoleIdNil(b bool)`

 SetRoleIdNil sets the value for RoleId to be an explicit nil

### UnsetRoleId
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetRoleId()`

UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil
### GetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetTimeZoneCode() string`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetTimeZoneCodeOk() (*string, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetTimeZoneCode(v string)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetLanguage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetPassword

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUserGroups

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserGroups() []IntraService5SettingsApiModelsUserGroupDto2E44`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserGroupsOk() (*[]IntraService5SettingsApiModelsUserGroupDto2E44, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserGroups(v []IntraService5SettingsApiModelsUserGroupDto2E44)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroupsNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserGroupsNil(b bool)`

 SetUserGroupsNil sets the value for UserGroups to be an explicit nil

### UnsetUserGroups
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetUserGroups()`

UnsetUserGroups ensures that no value is present for UserGroups, not even an explicit nil
### GetUserTags

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserTags() []IntraService5SettingsApiModelsBaseNamedModel10D7`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserTagsOk() (*[]IntraService5SettingsApiModelsBaseNamedModel10D7, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserTags(v []IntraService5SettingsApiModelsBaseNamedModel10D7)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### SetUserTagsNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserTagsNil(b bool)`

 SetUserTagsNil sets the value for UserTags to be an explicit nil

### UnsetUserTags
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetUserTags()`

UnsetUserTags ensures that no value is present for UserTags, not even an explicit nil
### GetFieldsValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFieldsValues() []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9`

GetFieldsValues returns the FieldsValues field if non-nil, zero value otherwise.

### GetFieldsValuesOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFieldsValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldValueModel86F9, bool)`

GetFieldsValuesOk returns a tuple with the FieldsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFieldsValues(v []IntraService5SettingsApiModelsAdditionalFieldValueModel86F9)`

SetFieldsValues sets FieldsValues field to given value.

### HasFieldsValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasFieldsValues() bool`

HasFieldsValues returns a boolean if a field has been set.

### SetFieldsValuesNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFieldsValuesNil(b bool)`

 SetFieldsValuesNil sets the value for FieldsValues to be an explicit nil

### UnsetFieldsValues
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetFieldsValues()`

UnsetFieldsValues ensures that no value is present for FieldsValues, not even an explicit nil
### GetFotoImage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFotoImage() string`

GetFotoImage returns the FotoImage field if non-nil, zero value otherwise.

### GetFotoImageOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFotoImageOk() (*string, bool)`

GetFotoImageOk returns a tuple with the FotoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFotoImage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFotoImage(v string)`

SetFotoImage sets FotoImage field to given value.

### HasFotoImage

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasFotoImage() bool`

HasFotoImage returns a boolean if a field has been set.

### SetFotoImageNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFotoImageNil(b bool)`

 SetFotoImageNil sets the value for FotoImage to be an explicit nil

### UnsetFotoImage
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetFotoImage()`

UnsetFotoImage ensures that no value is present for FotoImage, not even an explicit nil
### GetIsNewFoto

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetIsNewFoto() bool`

GetIsNewFoto returns the IsNewFoto field if non-nil, zero value otherwise.

### GetIsNewFotoOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetIsNewFotoOk() (*bool, bool)`

GetIsNewFotoOk returns a tuple with the IsNewFoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewFoto

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetIsNewFoto(v bool)`

SetIsNewFoto sets IsNewFoto field to given value.

### HasIsNewFoto

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasIsNewFoto() bool`

HasIsNewFoto returns a boolean if a field has been set.

### GetAdditionalData

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### SetEmailConfirmedNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetEmailConfirmedNil(b bool)`

 SetEmailConfirmedNil sets the value for EmailConfirmed to be an explicit nil

### UnsetEmailConfirmed
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetEmailConfirmed()`

UnsetEmailConfirmed ensures that no value is present for EmailConfirmed, not even an explicit nil
### GetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsArchived

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetUserName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetSetContactPerson

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetSetContactPerson() bool`

GetSetContactPerson returns the SetContactPerson field if non-nil, zero value otherwise.

### GetSetContactPersonOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetSetContactPersonOk() (*bool, bool)`

GetSetContactPersonOk returns a tuple with the SetContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetContactPerson

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetSetContactPerson(v bool)`

SetSetContactPerson sets SetContactPerson field to given value.

### HasSetContactPerson

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasSetContactPerson() bool`

HasSetContactPerson returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPhoneNumbers() []IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetPhoneNumbersOk() (*[]IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPhoneNumbers(v []IntraService5SharedSqlServerStorageDataModelPhoneNumberValue925D)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### SetPhoneNumbersNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetPhoneNumbersNil(b bool)`

 SetPhoneNumbersNil sets the value for PhoneNumbers to be an explicit nil

### UnsetPhoneNumbers
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetPhoneNumbers()`

UnsetPhoneNumbers ensures that no value is present for PhoneNumbers, not even an explicit nil
### GetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFieldIdValues() []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7`

GetFieldIdValues returns the FieldIdValues field if non-nil, zero value otherwise.

### GetFieldIdValuesOk

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) GetFieldIdValuesOk() (*[]IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7, bool)`

GetFieldIdValuesOk returns a tuple with the FieldIdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIdValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFieldIdValues(v []IntraService5SettingsApiModelsAdditionalFieldIdValueModel15E7)`

SetFieldIdValues sets FieldIdValues field to given value.

### HasFieldIdValues

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) HasFieldIdValues() bool`

HasFieldIdValues returns a boolean if a field has been set.

### SetFieldIdValuesNil

`func (o *IntraService5SettingsApiModelsClientUserModel38EA) SetFieldIdValuesNil(b bool)`

 SetFieldIdValuesNil sets the value for FieldIdValues to be an explicit nil

### UnsetFieldIdValues
`func (o *IntraService5SettingsApiModelsClientUserModel38EA) UnsetFieldIdValues()`

UnsetFieldIdValues ensures that no value is present for FieldIdValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


