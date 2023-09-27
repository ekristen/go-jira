package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CustomFieldOption struct {
	Self  string `json:"self,omitempty"`
	Value string `json:"value,omitempty"`
}

type CustomFieldContextOptionList struct {
	IsLast     bool                       `json:"isLast,omitempty"`
	MaxResults int64                      `json:"maxResults,omitempty"`
	NextPage   string                     `json:"nextPage,omitempty"`
	Self       string                     `json:"self,omitempty"`
	StartsAt   int64                      `json:"startsAt,omitempty"`
	Total      int64                      `json:"total,omitempty"`
	Values     []CustomFieldContextOption `json:"values,omitempty"`
}

type CustomFieldContextOption struct {
	ID       string `json:"id,omitempty" structs:"id,omitempty"`
	Value    string `json:"value,omitempty" structs:"value,omitempty"`
	OptionID string `json:"optionId,omitempty" structs:"optionId,omitempty"`
	Disabled bool   `json:"disabled,omitempty" structs:"disabled,omitempty"`
	Self     string `json:"self,omitempty"`
}

type CustomFieldContextOptions struct {
	Options []CustomFieldContextOption `json:"options,omitempty"`
}

type CustomFieldListOptions struct {
	OptionID    string `url:"optionId,omitempty"`
	OnlyOptions bool   `url:"onlyOptions,omitempty"`
	StartAt     int    `url:"startAt,omitempty"`
	MaxResults  int    `url:"maxResults,omitempty"`
}

func (s *FieldService) GetCustomFieldOptions(ctx context.Context, fieldID string) (*CustomFieldOption, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/customFieldOption/%s", fieldID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	entity := new(CustomFieldOption)
	resp, err := s.client.Do(req, &entity)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return entity, resp, nil
}
func (s *FieldService) GetCustomFieldOptionContext(ctx context.Context, fieldID string, contextID string, options *CustomFieldListOptions) (*CustomFieldContextOptionList, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/field/%s/context/%s/option", fieldID, contextID)

	u := url.URL{
		Path: apiEndpoint,
	}

	uv := url.Values{}
	if options != nil {
		if options.StartAt != 0 {
			uv.Add("startAt", strconv.Itoa(options.StartAt))
		}
		if options.MaxResults != 0 {
			uv.Add("maxResults", strconv.Itoa(options.MaxResults))
		}
	}

	u.RawQuery = uv.Encode()

	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	entity := new(CustomFieldContextOptionList)
	resp, err := s.client.Do(req, &entity)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return entity, resp, nil
}
func (s *FieldService) CreateCustomFieldOptionContext(ctx context.Context, fieldID string, contextID string, options *CustomFieldContextOptions) (*CustomFieldContextOptions, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/field/%s/context/%s/option", fieldID, contextID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	entity := new(CustomFieldContextOptions)
	resp, err := s.client.Do(req, &entity)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return entity, resp, nil
}
func (s *FieldService) UpdateCustomFieldOptionContext(ctx context.Context, fieldID string, contextID string, options *CustomFieldContextOptions) (*CustomFieldContextOptions, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/field/%s/context/%s/option", fieldID, contextID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPut, apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	entity := new(CustomFieldContextOptions)
	resp, err := s.client.Do(req, &entity)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return entity, resp, nil
}
func (s *FieldService) DeleteCustomFieldOptionContext(ctx context.Context, fieldID string, contextID string, optionID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/field/%s/context/%s/option/%s", fieldID, contextID, optionID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodDelete, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}
	return resp, nil
}
