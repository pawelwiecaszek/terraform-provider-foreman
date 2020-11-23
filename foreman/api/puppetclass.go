package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	PuppetClassEndpoint = "puppetclasses"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

// The ForemanPuppetClass API model represents the puppet class name.

type ForemanPuppetClass struct {
	// Inherits the base object's attributes
	ForemanObject

	Name string `json:"name"`
}

// Custom JSON unmarshal function. Unmarshal to the unexported JSON struct
// and then convert over to a ForemanPuppetClass struct.
func (fi *ForemanPuppetClass) UnmarshalJSON(b []byte) error {
	var jsonDecErr error

	// Unmarshal the common Foreman object properties
	var fo ForemanObject
	jsonDecErr = json.Unmarshal(b, &fo)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	fi.ForemanObject = fo

	// Unmarshal into mapstructure and set the rest of the struct properties
	// NOTE(ALL): Properties unmarshalled are of type float64 as opposed to int, hence the below testing
	// Without this, properties will define as default values in state file.
	var fiMap map[string]interface{}
	jsonDecErr = json.Unmarshal(b, &fiMap)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	log.Debugf("fiMap: [%v]", fiMap)
	var ok bool

	if fi.Name, ok = fiMap["name"].(string); !ok {
		fi.Name = ""
	}

	return nil
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

// CreatePuppetClass creates a new ForemanPuppetClass with the attributes of the supplied
// ForemanPuppetClass reference and returns the created ForemanPuppetClass reference.
// The returned reference will have its ID and other API default values set by
// this function.
func (c *Client) CreatePuppetClass(d *ForemanPuppetClass) (*ForemanPuppetClass, error) {
	log.Tracef("foreman/api/puppetclass.go#Create")

	reqEndpoint := fmt.Sprintf("%s", PuppetClassEndpoint)

	puppetclassJSONBytes, jsonEncErr := WrapJson("puppetclass", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("puppetclassJSONBytes: [%s]", puppetclassJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(puppetclassJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdPuppetClass ForemanPuppetClass
	sendErr := c.SendAndParse(req, &createdPuppetClass)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdPuppetClass: [%+v]", createdPuppetClass)

	return &createdPuppetClass, nil
}

// ReadPuppetClass reads the attributes of a ForemanPuppetClass identified by the
// supplied ID and returns a ForemanPuppetClass reference.
func (c *Client) ReadPuppetClass(id int) (*ForemanPuppetClass, error) {
	log.Tracef("foreman/api/puppetclass.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%d", PuppetClassEndpoint, id)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var ReadPuppetClass ForemanPuppetClass
	sendErr := c.SendAndParse(req, &ReadPuppetClass)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("ReadPuppetClass: [%+v]", ReadPuppetClass)

	return &ReadPuppetClass, nil
}

// UpdatePuppetClass updates a ForemanPuppetClass's attributes.  The image with the ID
// of the supplied ForemanPuppetClass will be updated. A new ForemanPuppetClass reference
// is returned with the attributes from the result of the update operation.
func (c *Client) UpdatePuppetClass(d *ForemanPuppetClass) (*ForemanPuppetClass, error) {
	log.Tracef("foreman/api/puppetclass.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%d/", PuppetClassEndpoint, d.Id)

	puppetclassJSONBytes, jsonEncErr := WrapJson("puppetclass", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("puppetclassJSONBytes: [%s]", puppetclassJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(puppetclassJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var UpdatedPuppetClass ForemanPuppetClass
	sendErr := c.SendAndParse(req, &UpdatedPuppetClass)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("UpdatedPuppetClass: [%+v]", UpdatedPuppetClass)

	return &UpdatedPuppetClass, nil
}

// DeletePuppetClass deletes the ForemanPuppetClass identified by the supplied ID
func (c *Client) DeletePuppetClass(id int) error {
	log.Tracef("foreman/api/puppetclass.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%d", PuppetClassEndpoint, id)

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

// QueryPuppetClass queries for a ForemanPuppetClass based on the attributes of the
// supplied ForemanPuppetClass reference and returns a QueryResponse struct
// containing query/response metadata and the matching images.
func (c *Client) QueryPuppetClass(d *ForemanPuppetClass) (QueryResponse, error) {
	log.Tracef("foreman/api/puppetclass.go#Search")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("%s", PuppetClassEndpoint)
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
	name := `"` + d.Name + `"`
	reqQuery.Set("search", "name="+name)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanPuppetClass for
	// the results
	results := []ForemanPuppetClass{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanPuppetClass to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}
