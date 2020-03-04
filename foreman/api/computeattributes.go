package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	ComputeAttributesEndpointPrefix = "compute_attributes"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

// The ForemanComputeAttributes API model represents a provisioning template.
// Provisioning templates are scripts used to describe how to boostrap and
// install the operating system on a host.
type ForemanComputeAttributes struct {
	// Inherits the base object's attributes
	ForemanObject

	// ComputeResourceId specifies the Hypervisor to deploy on
	ComputeResourceId int `json:"compute_resource_id"`
	// ComputeProfileId specifies the Attributes via the Profile Id on the Hypervisor
	ComputeProfileId int `json:"compute_profile_id"`
	
	ForemanComputeAttributeInternal []ForemanComputeAttributeInternal `json:"compute_attribute"`
}

// See the comment in ForemanComputeAttributes.ForemanComputeAttributeInternal
type ForemanComputeAttributeInternal struct {
	VMAttributes map[string]interface{} `json:"vm_attrs"`
}

type foremanComputeAttributeJSON struct {
	ForemanComputeAttributeInternal []ForemanComputeAttributeInternal `json:"attributes"`
}

// Custom JSON marshal function for provisioning temmplates.  The Foreman API
// expects all parameters to be enclosed in double quotes, with the exception
// of boolean and slice values.
func (ft ForemanComputeAttributes) MarshalJSON() ([]byte, error) {
	log.Tracef("Compute attributes marshal")

	// map structure representation of the passed ForemanComputeAttributes
	// for ease of marshalling - essentially convert over to a map then call
	// json.Marshal() on the mapstructure
	ftMap := map[string]interface{}{}

	ftMap["compute_resource_id"] = intIdToJSONString(ft.ComputeResourceId)
	ftMap["compute_profile_id"] = intIdToJSONString(ft.ComputeProfileId)

	if len(ft.ForemanComputeAttributeInternal) > 0 {
		ftMap["compute_attribute"] = ft.ForemanComputeAttributeInternal
	}

	log.Debugf("ftMap: [%v]", ftMap)

	return json.Marshal(ftMap)
}

// Custom JSON unmarshal function. Unmarshal to the unexported JSON struct
// and then convert over to a ForemanComputeAttributes struct.
func (ft *ForemanComputeAttributes) UnmarshalJSON(b []byte) error {
	var jsonDecErr error

	// Unmarshal the common Foreman object properties
	var fo ForemanObject
	jsonDecErr = json.Unmarshal(b, &fo)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	ft.ForemanObject = fo

	// Unmarshal to temporary JSON struct to get the properties with differently
	// named keys
	var ftJSON foremanComputeAttributeJSON
	jsonDecErr = json.Unmarshal(b, &ftJSON)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	ft.ForemanComputeAttributeInternal = ftJSON.ForemanComputeAttributeInternal

	// Unmarshal into mapstructure and set the rest of the struct properties
	var ftMap map[string]interface{}
	jsonDecErr = json.Unmarshal(b, &ftMap)
	if jsonDecErr != nil {
		return jsonDecErr
	}

	ft.ComputeResourceId = unmarshalInteger(ftMap["compute_resource_id"])
	ft.ComputeProfileId = unmarshalInteger(ftMap["compute_profile_id"])

	return nil
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

// CreateComputeAttributes creates a new ForemanComputeAttributes with
// the attributes of the supplied ForemanComputeAttributes reference and
// returns the created ForemanComputeAttributes reference.  The returned
// reference will have its ID and other API default values set by this
// function.
func (c *Client) CreateComputeAttributes(t *ForemanComputeAttributes) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/computeattributes.go#Create")

	reqEndpoint := fmt.Sprintf("/%s", ComputeAttributesEndpointPrefix)

	compute_attributesJSONBytes, jsonEncErr := json.Marshal(t)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("compute_attributesJSONBytes: [%s]", compute_attributesJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(compute_attributesJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdComputeAttributes ForemanComputeAttributes
	sendErr := c.SendAndParse(req, &createdComputeAttributes)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdComputeAttributes: [%+v]", createdComputeAttributes)

	return &createdComputeAttributes, nil
}

// ReadComputeAttributes reads the attributes of a
// ForemanComputeAttributes identified by the supplied ID and returns a
// ForemanComputeAttributes reference.
func (c *Client) ReadComputeAttributes(id int) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/computeattributes.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeAttributesEndpointPrefix, id)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var readComputeAttributes ForemanComputeAttributes
	sendErr := c.SendAndParse(req, &readComputeAttributes)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("readComputeAttributes: [%+v]", readComputeAttributes)

	return &readComputeAttributes, nil
}

// UpdateComputeAttributes updates a ForemanComputeAttributes's
// attributes.  The template with the ID of the supplied
// ForemanComputeAttributes will be updated. A new
// ForemanComputeAttributes reference is returned with the attributes from
// the result of the update operation.
func (c *Client) UpdateComputeAttributes(t *ForemanComputeAttributes) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/computeattributes.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeAttributesEndpointPrefix, t.Id)

	compute_attributesJSONBytes, jsonEncErr := json.Marshal(t)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("compute_attributesJSONBytes: [%s]", compute_attributesJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(compute_attributesJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var updatedComputeAttributes ForemanComputeAttributes
	sendErr := c.SendAndParse(req, &updatedComputeAttributes)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("updatedComputeAttributes: [%+v]", updatedComputeAttributes)

	return &updatedComputeAttributes, nil
}

// DeleteComputeAttributes deletes the ForemanComputeAttributes
// identified by the supplied ID
func (c *Client) DeleteComputeAttributes(id int) error {
	log.Tracef("foreman/api/computeattributes.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%d", ComputeAttributesEndpointPrefix, id)

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

// QueryComputeAttributes queries for a ForemanComputeAttributes based on
// the attributes of the supplied ForemanComputeAttributes reference and
// returns a QueryResponse struct containing query/response metadata and the
// matching templates.
func (c *Client) QueryComputeAttributes(t *ForemanComputeAttributes) (QueryResponse, error) {
	log.Tracef("foreman/api/computeattributes.go#Query")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("/%s", ComputeAttributesEndpointPrefix)
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
	name := "\"" + t.Name + "\""
	reqQuery.Set("search", "name="+name)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanComputeAttributes for
	// the results
	results := []ForemanComputeAttributes{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanComputeAttributes to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}
