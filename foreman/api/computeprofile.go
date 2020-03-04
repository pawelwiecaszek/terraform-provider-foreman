package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	ComputeProfileEndpointPrefix = "compute_profiles"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

type ForemanComputeProfile struct {
	// Inherits the base object's attributes
	ForemanObject
	Name        string `json:"name"`
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

func (c *Client) CreateComputeProfile(e *ForemanComputeProfile) (*ForemanComputeProfile, error) {
	log.Tracef("foreman/api/computeprofile.go#Create")

	reqEndpoint := fmt.Sprintf("/%s", ComputeProfileEndpointPrefix)

	ComputeProfileJSONBytes, jsonEncErr := WrapJson("compute_profile", e)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("ComputeProfileJSONBytes: [%s]", ComputeProfileJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(ComputeProfileJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdComputeProfile ForemanComputeProfile
	sendErr := c.SendAndParse(req, &createdComputeProfile)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdComputeProfile: [%+v]", createdComputeProfile)

	return &createdComputeProfile, nil
}

// ReadComputeProfile reads the attributes of a ForemanComputeProfile identified by
// the supplied ID and returns a ForemanComputeProfile reference.
func (c *Client) ReadComputeProfile(id int) (*ForemanComputeProfile, error) {
	log.Tracef("foreman/api/templatekind.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeProfileEndpointPrefix, id)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var readComputeProfile ForemanComputeProfile
	sendErr := c.SendAndParse(req, &readComputeProfile)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("readComputeProfile: [%+v]", readComputeProfile)

	return &readComputeProfile, nil
}

func (c *Client) UpdateComputeProfile(e *ForemanComputeProfile) (*ForemanComputeProfile, error) {
	log.Tracef("foreman/api/ComputeProfile.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeProfileEndpointPrefix, e.Id)

	ComputeProfileJSONBytes, jsonEncErr := WrapJson("ComputeProfile", e)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("ComputeProfileJSONBytes: [%s]", ComputeProfileJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(ComputeProfileJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var updatedComputeProfile ForemanComputeProfile
	sendErr := c.SendAndParse(req, &updatedComputeProfile)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("updatedComputeProfile: [%+v]", updatedComputeProfile)

	return &updatedComputeProfile, nil
}

func (c *Client) DeleteComputeProfile(id int) error {
	log.Tracef("foreman/api/ComputeProfile.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeProfileEndpointPrefix, id)

	req, reqErr := c.NewRequest(
		http.MethodDelete,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return reqErr
	}

	return c.SendAndParse(req, nil)
}

// -----------------------------------------------------------------------------
// Query Implementation
// -----------------------------------------------------------------------------

// QueryComputeProfile queries for a ForemanComputeProfile based on the attributes
// of the supplied ForemanComputeProfile reference and returns a QueryResponse
// struct containing query/response metadata and the matching template kinds
func (c *Client) QueryComputeProfile(t *ForemanComputeProfile) (QueryResponse, error) {
	log.Tracef("foreman/api/templatekind.go#Search")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("/%s", ComputeProfileEndpointPrefix)
	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return queryResponse, reqErr
	}

	// dynamically build the query based on the attributes
	reqQuery := req.URL.Query()
	name := `"` + t.Name + `"`
	reqQuery.Set("search", "name="+name)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanComputeProfile for
	// the results
	results := []ForemanComputeProfile{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanComputeProfile to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}