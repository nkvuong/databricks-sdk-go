// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool `json:"confidential,omitempty"`
	// Name of the custom OAuth app
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId string `json:"client_id,omitempty"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret string `json:"client_secret,omitempty"`
	// Unique integration id for the custom OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime string `json:"create_time,omitempty"`
	// ID of the secret
	Id string `json:"id,omitempty"`
	// Secret Value
	Secret string `json:"secret,omitempty"`
	// Secret Hash
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateServicePrincipalSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails string `json:"authorization_details,omitempty"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DataPlaneInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCustomAppIntegrationOutput struct {
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeletePublishedAppIntegrationOutput struct {
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string `json:"-" url:"-"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId string `json:"client_id,omitempty"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool `json:"confidential,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`

	CreatorUsername string `json:"creator_username,omitempty"`
	// ID of this custom app
	IntegrationId string `json:"integration_id,omitempty"`
	// The display name of the custom OAuth app
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	RedirectUrls []string `json:"redirect_urls,omitempty"`

	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCustomAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	AppId string `json:"app_id,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`
	// Display name of the published OAuth app
	Name string `json:"name,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps []PublishedAppOutput `json:"apps,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername bool `json:"-" url:"include_creator_username,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListCustomAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCustomAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListOAuthPublishedAppsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOAuthPublishedAppsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListPublishedAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPublishedAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

type ListServicePrincipalSecretsResponse struct {
	// List of the secrets
	Secrets []SecretInfo `json:"secrets,omitempty"`
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId string `json:"app_id,omitempty"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId string `json:"client_id,omitempty"`
	// Description of the published OAuth app.
	Description string `json:"description,omitempty"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient bool `json:"is_confidential_client,omitempty"`
	// The display name of the published OAuth app.
	Name string `json:"name,omitempty"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Required scopes for the published OAuth app.
	Scopes []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PublishedAppOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedAppOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime string `json:"create_time,omitempty"`
	// ID of the secret
	Id string `json:"id,omitempty"`
	// Secret Hash
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SecretInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes int `json:"access_token_ttl_in_minutes,omitempty"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCustomAppIntegration struct {
	IntegrationId string `json:"-" url:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

type UpdateCustomAppIntegrationOutput struct {
}

type UpdatePublishedAppIntegration struct {
	IntegrationId string `json:"-" url:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

type UpdatePublishedAppIntegrationOutput struct {
}
