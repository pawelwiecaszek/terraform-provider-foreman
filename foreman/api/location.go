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

	Name                    string `json:"name"`
	Description             string `json:"description"`
	UserIds                 []int  `json:"user_ids"`
	SmartProxyIds           []int  `json:"smart_proxy_ids"`
	ComputeResourceIds      []int  `json:"compute_resource_ids"`
	MediaIds                []int  `json:"media_ids"`
	ConfigTemplateIds       []int  `json:"config_template_ids"`
	PtableIds               []int  `json:"ptable_ids"`
	ProvisioningTemplateIds []int  `json:"provisiong_template_ids"`
	DomainIds               []int  `json:"domain_ids"`
	RealmIds                []int  `json:"realm_ids"`
	HostgroupIds            []int  `json:"hostgroup_ids"`
	EnvironmentIds          []int  `json:"environment_ids"`
	SubnetIds               []int  `json:"subnet_ids"`
	ParentId                int    `json:"parent_id"`
}

type foremanLocationRespJSON struct {
	Users                 []ForemanObject `json:"users"`
	SmartProxies          []ForemanObject `json:"smart_proxies"`
	ComputeResources      []ForemanObject `json:"compute_resources"`
	Media                 []ForemanObject `json:"media"`
	ConfigTemplates       []ForemanObject `json:"config_templates"`
	Ptables               []ForemanObject `json:"ptables"`
	ProvisioningTemplates []ForemanObject `json:"provisioning_templates"`
	Domains               []ForemanObject `json:"domains"`
	Realms                []ForemanObject `json:"realms"`
	Hostgroups            []ForemanObject `json:"hostgroups"`
	Environments          []ForemanObject `json:"environments"`
	Subnets               []ForemanObject `json:"subnets"`
}

// Implement the Unmarshaler interface
func (l *ForemanLocation) UnmarshalJSON(b []byte) error {
	var jsonDecErr error

	// Unmarshal the common Foreman object properties
	var fo ForemanObject
	jsonDecErr = json.Unmarshal(b, &fo)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	l.ForemanObject = fo

	var foJSON foremanLocationRespJSON
	jsonDecErr = json.Unmarshal(b, &foJSON)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	l.UserIds = foremanObjectArrayToIdIntArray(foJSON.Users)
	l.SmartProxyIds = foremanObjectArrayToIdIntArray(foJSON.SmartProxies)
	l.ComputeResourceIds = foremanObjectArrayToIdIntArray(foJSON.ComputeResources)
	l.MediaIds = foremanObjectArrayToIdIntArray(foJSON.Media)
	l.ConfigTemplateIds = foremanObjectArrayToIdIntArray(foJSON.ConfigTemplates)
	l.PtableIds = foremanObjectArrayToIdIntArray(foJSON.Ptables)
	l.ProvisioningTemplateIds = foremanObjectArrayToIdIntArray(foJSON.ProvisioningTemplates)
	l.DomainIds = foremanObjectArrayToIdIntArray(foJSON.Domains)
	l.RealmIds = foremanObjectArrayToIdIntArray(foJSON.Realms)
	l.HostgroupIds = foremanObjectArrayToIdIntArray(foJSON.Hostgroups)
	l.EnvironmentIds = foremanObjectArrayToIdIntArray(foJSON.Environments)
	l.SubnetIds = foremanObjectArrayToIdIntArray(foJSON.Subnets)

	var foMap map[string]interface{}
	jsonDecErr = json.Unmarshal(b, &foMap)
	if jsonDecErr != nil {
		return jsonDecErr
	}
	log.Debugf("foMap: [%v]", foMap)

	var ok bool
	if l.Name, ok = foMap["name"].(string); !ok {
		l.Name = ""
	}
	if l.Description, ok = foMap["description"].(string); !ok {
		l.Description = ""
	}

	return nil
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
