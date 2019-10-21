package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wayfair/terraform-provider-utils/log"
)

const (
	ComputeAttributesEndpointPrefix = "compute_profiles/%d/compute_resources/%d"
)

// -----------------------------------------------------------------------------
// Struct Definition and Helpers
// -----------------------------------------------------------------------------

type ForemanComputeAttribute struct {
	VMAttrs       json.RawMessage `json:"vm_attrs"`
	VMAttrsString string
}

type ForemanComputeAttributes struct {
	// Inherits the base object's attributes
	ForemanObject

	ComputeResourceID int                     `json:"compute_resource_id"`
	ComputeProfileID  int                     `json:"compute_profile_id"`
	ComputeAttribute  ForemanComputeAttribute `json:"compute_attribute"`
}

type foremanComputeAttributesJSON struct {
	ComputeAttributes []ForemanComputeAttributes `json:"compute_attributes"`
}

func (fca *ForemanComputeAttributes) UnmarshalJSON(a []byte) error {
	var jsonDecErr error

	jsonDecErr = json.Unmarshal(a, fca)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	return nil
}

func (fca *ForemanComputeAttributes) MarshalJSON() ([]byte, error) {
	fcaMap := map[string]interface{}{}
	fcaMap["compute_resource_id"] = fca.ComputeResourceID
	fcaMap["compute_profile_id"] = fca.ComputeProfileID
	fcaMap["compute_attribute"] = fca.ComputeAttribute
	return json.Marshal(fcaMap)
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

func (c *Client) CreateComputeAttributes(d *ForemanComputeAttributes) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/compute_attributes.go#Create")

	reqEndpoint := fmt.Sprintf(ComputeAttributesEndpointPrefix, d.ComputeProfileID, d.ComputeResourceID)

	compute_attributesJSONBytes, jsonEncErr := json.Marshal(d)
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

	var createdComputeAttributes ForemanComputeAttributes
	sendErr := c.SendAndParse(req, &createdComputeAttributes)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdComputeAttributes: [%+v]", createdComputeAttributes)

	return &createdComputeAttributes, nil
}

// ReadComputeAttributes reads the attributes of a ForemanComputeAttributes identified by the
// supplied ID and returns a ForemanComputeAttributes reference.
func (c *Client) ReadComputeAttributes(d *ForemanComputeAttributes) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/compute_attributes.go#Read")

	reqEndpoint := fmt.Sprintf(ComputeAttributesEndpointPrefix, d.ComputeProfileID, d.ComputeResourceID)

	req, reqErr := c.NewRequest(
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var listComputeAttributes foremanComputeAttributesJSON
	sendErr := c.SendAndParse(req, &listComputeAttributes)
	if sendErr != nil {
		return nil, sendErr
	}

	var readComputeAttributes *ForemanComputeAttributes
	for _, computeAttribute := range listComputeAttributes.ComputeAttributes {
		if computeAttribute.ComputeResourceID == d.ComputeResourceID &&
			computeAttribute.ComputeProfileID == d.ComputeProfileID {
			readComputeAttributes = &computeAttribute
		}
	}

	log.Debugf("readComputeAttributes: [%+v]", readComputeAttributes)

	return readComputeAttributes, nil
}

// UpdateComputeAttributes updates a ForemanComputeAttributes's attributes.  The compute_attributes with the ID
// of the supplied ForemanComputeAttributes will be updated. A new ForemanComputeAttributes reference
// is returned with the attributes from the result of the update operation.
func (c *Client) UpdateComputeAttributes(d *ForemanComputeAttributes) (*ForemanComputeAttributes, error) {
	log.Tracef("foreman/api/compute_attributes.go#Update")

	reqEndpoint := fmt.Sprintf(ComputeAttributesEndpointPrefix, d.ComputeProfileID, d.ComputeResourceID)

	compute_attributesJSONBytes, jsonEncErr := json.Marshal(d)
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
