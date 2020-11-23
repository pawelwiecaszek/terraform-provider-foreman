package foreman

import (
	"fmt"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/helper"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceForemanPuppetClass() *schema.Resource {
	// copy attributes from resource definition
	r := resourceForemanPuppetClass()
	ds := helper.DataSourceSchemaFromResourceSchema(r.Schema)

	// define searchable attributes for the data source
	ds["name"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: fmt.Sprintf("The name of the puppet class resource. %s", autodoc.MetaExample),
	}

	return &schema.Resource{

		Read: dataSourceForemanPuppetClassRead,

		// NOTE(ALL): See comments in the corresponding resource file
		Schema: ds,
	}
}

func dataSourceForemanPuppetClassRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("data_source_foreman_puppetclass.go#Read")

	client := meta.(*api.Client)
	puppetclass := buildForemanPuppetClass(d)

	log.Debugf("ForemanPuppetClass: [%+v]", puppetclass)

	queryResponse, queryErr := client.QueryPuppetClass(puppetclass)
	if queryErr != nil {
		return queryErr
	}

	if queryResponse.Subtotal == 0 {
		return fmt.Errorf("Data source puppetclass returned no results")
	} else if queryResponse.Subtotal > 1 {
		return fmt.Errorf("Data source puppetclass returned more than 1 result")
	}

	var queryPuppetClass api.ForemanPuppetClass
	var ok bool
	if queryPuppetClass, ok = queryResponse.Results[0].(api.ForemanPuppetClass); !ok {
		return fmt.Errorf(
			"Data source results contain unexpected type. Expected "+
				"[api.ForemanPuppetClass], got [%T]",
			queryResponse.Results[0],
		)
	}
	puppetclass = &queryPuppetClass

	log.Debugf("ForemanPuppetClass: [%+v]", puppetclass)

	setResourceDataFromForemanPuppetClass(d, puppetclass)

	return nil
}
