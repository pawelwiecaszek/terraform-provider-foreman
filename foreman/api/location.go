package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	LocationEndpointPrefix = "locations"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

// The ForemanLocation API model represents a location
type ForemanLocation struct {
	// Inherits the base object's attributes
	ForemanObject
	Realms                []int
	ComputeResources      []int
	Domains               []int
	Subnets               []int
	Environments          []int
	Hostgroups            []int
	ProvisioningTemplates []int
	SmartProxies          []int
	Users                 []int
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

// CreateLocation creates a new ForemanLocation with the attributes of
// the supplied ForemanLocation reference and returns the created
// ForemanLocation reference.  The returned reference will have its ID and
// other API default values set by this function.
func (c *Client) CreateLocation(e *ForemanLocation) (*ForemanLocation, error) {
	log.Tracef("foreman/api/location.go#Create")

	reqEndpoint := fmt.Sprintf("/%s", LocationEndpointPrefix)

	locationJSONBytes, jsonEncErr := WrapJson("location", e)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("locationJSONBytes: [%s]", locationJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(locationJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdLocation ForemanLocation
	sendErr := c.SendAndParse(req, &createdLocation)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdLocation: [%+v]", createdLocation)

	return &createdLocation, nil
}

// ReadLocation reads the attributes of a ForemanLocation identified by
// the supplied ID and returns a ForemanLocation reference.
func (c *Client) ReadLocation(id int) (*ForemanLocation, error) {
	log.Tracef("foreman/api/location.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%d", LocationEndpointPrefix, id)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var readLocation ForemanLocation
	sendErr := c.SendAndParse(req, &readLocation)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("readLocation: [%+v]", readLocation)

	return &readLocation, nil
}

// UpdateLocation updates a ForemanLocation's attributes.  The
// location with the ID of the supplied ForemanLocation will be updated.
// A new ForemanLocation reference is returned with the attributes from the
// result of the update operation.
func (c *Client) UpdateLocation(e *ForemanLocation) (*ForemanLocation, error) {
	log.Tracef("foreman/api/location.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%d", LocationEndpointPrefix, e.Id)

	locationJSONBytes, jsonEncErr := WrapJson("location", e)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("locationJSONBytes: [%s]", locationJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(locationJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var updatedLocation ForemanLocation
	sendErr := c.SendAndParse(req, &updatedLocation)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("updatedLocation: [%+v]", updatedLocation)

	return &updatedLocation, nil
}

// DeleteLocation deletes the ForemanLocation identified by the supplied
// ID
func (c *Client) DeleteLocation(id int) error {
	log.Tracef("foreman/api/location.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%d", LocationEndpointPrefix, id)

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

// QueryLocation queries for a ForemanLocation based on the attributes of
// the supplied ForemanLocation reference and returns a QueryLocation struct
// containing query/response metadata and the matching locations.
func (c *Client) QueryLocation(e *ForemanLocation) (QueryResponse, error) {
	log.Tracef("foreman/api/location.go#Search")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("/%s", LocationEndpointPrefix)
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
	name := `"` + e.Name + `"`
	reqQuery.Set("search", "name="+name)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanLocation for
	// the results
	results := []ForemanLocation{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanLocation to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}
