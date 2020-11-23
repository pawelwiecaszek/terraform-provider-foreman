package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	SmartClassParameterEndpoint = "smart_class_parameters"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

// The ForemanSmartClassParameter API model represents the puppet class name.

type ForemanSmartClassParameter struct {
	// Inherits the base object's attributes
	ForemanObject

	Match                 string `json:"match"`
	Value                 string `json:"value"`
	UsePuppetDefault      bool   `json:use_puppet_default`
	Omit                  bool   `json:omit`
	SmartClassParameterId string `json:smart_class_parameter_id`
}

// Implement the Marshaler interface
func (fh ForemanSmartClassParameter) MarshalJSON() ([]byte, error) {
	log.Tracef("foreman/api/smartclassparameter.go#MarshalJSON")

	fhMap := map[string]interface{}{}

	fhMap["match"] = fh.Match
	fhMap["value"] = fh.Value
	fhMap["use_puppet_default"] = fh.UsePuppetDefault
	fhMap["omit"] = fh.Omit
	log.Debugf("fhMap: [%+v]", fhMap)

	return json.Marshal(fhMap)
}

// Custom JSON unmarshal function. Unmarshal to the unexported JSON struct
// and then convert over to a ForemanSmartClassParameter struct.
func (fi *ForemanSmartClassParameter) UnmarshalJSON(b []byte) error {
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

	if fi.Value, ok = fiMap["value"].(string); !ok {
		fi.Value = "error_value"
	}

	if fi.Match, ok = fiMap["match"].(string); !ok {
		fi.Value = "error_match"
	}

	if fi.UsePuppetDefault, ok = fiMap["use_puppet_default"].(bool); !ok {
		fi.UsePuppetDefault = false
	}

	if fi.Omit, ok = fiMap["omit"].(bool); !ok {
		fi.Omit = false
	}

	return nil
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

// CreateSmartClassParameter creates a new ForemanSmartClassParameter with the attributes of the supplied
// ForemanSmartClassParameter reference and returns the created ForemanSmartClassParameter reference.
// The returned reference will have its ID and other API default values set by
// this function.
func (c *Client) CreateSmartClassParameter(d *ForemanSmartClassParameter) (*ForemanSmartClassParameter, error) {
	log.Tracef("foreman/api/smartclassparameter.go#Create")

	reqEndpoint := fmt.Sprintf("%s/%s/override_values", SmartClassParameterEndpoint, d.SmartClassParameterId)

	smartclassparameterJSONBytes, jsonEncErr := WrapJson("override_value", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("smartclassparameterJSONBytes: [%s]", smartclassparameterJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(smartclassparameterJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdSmartClassParameter ForemanSmartClassParameter
	sendErr := c.SendAndParse(req, &createdSmartClassParameter)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdSmartClassParameter: [%+v]", createdSmartClassParameter)

	return &createdSmartClassParameter, nil
}

// ReadSmartClassParameter reads the attributes of a ForemanSmartClassParameter identified by the
// supplied ID and returns a ForemanSmartClassParameter reference.
func (c *Client) ReadSmartClassParameter(d *ForemanSmartClassParameter) (*ForemanSmartClassParameter, error) {
	log.Tracef("foreman/api/smartclassparameter.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%s/override_values/%d", SmartClassParameterEndpoint, d.SmartClassParameterId, d.Id)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var ReadSmartClassParameter ForemanSmartClassParameter
	sendErr := c.SendAndParse(req, &ReadSmartClassParameter)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("ReadSmartClassParameter: [%+v]", ReadSmartClassParameter)

	return &ReadSmartClassParameter, nil
}

// UpdateSmartClassParameter updates a ForemanSmartClassParameter's attributes.  The image with the ID
// of the supplied ForemanSmartClassParameter will be updated. A new ForemanSmartClassParameter reference
// is returned with the attributes from the result of the update operation.
func (c *Client) UpdateSmartClassParameter(d *ForemanSmartClassParameter) (*ForemanSmartClassParameter, error) {
	log.Tracef("foreman/api/smartclassparameter.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%s/override_values/%d", SmartClassParameterEndpoint, d.SmartClassParameterId, d.Id)

	smartclassparameterJSONBytes, jsonEncErr := WrapJson("override_values", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("smartclassparameterJSONBytes: [%s]", smartclassparameterJSONBytes)

	req, reqErr := c.NewRequest(
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(smartclassparameterJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var UpdatedSmartClassParameter ForemanSmartClassParameter
	sendErr := c.SendAndParse(req, &UpdatedSmartClassParameter)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("UpdatedSmartClassParameter: [%+v]", UpdatedSmartClassParameter)

	return &UpdatedSmartClassParameter, nil
}

// DeleteSmartClassParameter deletes the ForemanSmartClassParameter identified by the supplied ID
func (c *Client) DeleteSmartClassParameter(d *ForemanSmartClassParameter) error {
	log.Tracef("foreman/api/smartclassparameter.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%s/override_values/%d", SmartClassParameterEndpoint, d.SmartClassParameterId, d.Id)

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

// QuerySmartClassParameter queries for a ForemanSmartClassParameter based on the attributes of the
// supplied ForemanSmartClassParameter reference and returns a QueryResponse struct
// containing query/response metadata and the matching images.
func (c *Client) QuerySmartClassParameter(d *ForemanSmartClassParameter) (QueryResponse, error) {
	log.Tracef("foreman/api/smartclassparameter.go#Search")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("%s/%s/override_values", SmartClassParameterEndpoint, d.SmartClassParameterId)
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
	match := `"` + d.Match + `"`
	reqQuery.Set("search", "match="+match)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanSmartClassParameter for
	// the results
	results := []ForemanSmartClassParameter{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanSmartClassParameter to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}
