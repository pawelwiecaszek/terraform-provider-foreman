package foreman

import (
	"fmt"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/helper"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceForemanLocation() *schema.Resource {
	// copy attributes from resource definition
	r := resourceForemanLocation()
	ds := helper.DataSourceSchemaFromResourceSchema(r.Schema)

	// define searchable attributes for the data source
	ds["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		Description: fmt.Sprintf(
			"The name of the location. "+
				"%s \"production\"",
			autodoc.MetaExample,
		),
	}

	return &schema.Resource{

		Read: dataSourceForemanLocationRead,

		// NOTE(ALL): See comments in the corresponding resource file
		Schema: ds,
	}
}

func dataSourceForemanLocationRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("data_source_foreman_location.go#Read")

	client := meta.(*api.Client)
	e := buildForemanLocation(d)

	log.Debugf("ForemanLocation: [%+v]", e)

	queryResponse, queryErr := client.QueryLocation(e)
	if queryErr != nil {
		return queryErr
	}

	if queryResponse.Subtotal == 0 {
		return fmt.Errorf("Data source location returned no results")
	} else if queryResponse.Subtotal > 1 {
		return fmt.Errorf("Data source location returned more than 1 result")
	}

	var queryLocation api.ForemanLocation
	var ok bool
	if queryLocation, ok = queryResponse.Results[0].(api.ForemanLocation); !ok {
		return fmt.Errorf(
			"Data source results contain unexpected type. Expected "+
				"[api.ForemanLocation], got [%T]",
			queryResponse.Results[0],
		)
	}
	e = &queryLocation

	log.Debugf("ForemanLocation: [%+v]", e)

	setResourceDataFromForemanLocation(d, e)

	return nil
}