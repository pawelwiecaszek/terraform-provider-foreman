package foreman

import (
	"fmt"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/helper"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceForemanSmartClassParameter() *schema.Resource {
	// copy attributes from resource definition
	r := resourceForemanSmartClassParameter()
	ds := helper.DataSourceSchemaFromResourceSchema(r.Schema)

	// define searchable attributes for the data source
	ds["match"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: fmt.Sprintf("The match of the smart puppet class parameter. %s", autodoc.MetaExample),
	}

	return &schema.Resource{

		Read: dataSourceForemanSmartClassParameterRead,

		// NOTE(ALL): See comments in the corresponding resource file
		Schema: ds,
	}
}

func dataSourceForemanSmartClassParameterRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("data_source_foreman_smartclassparameter.go#Read")

	client := meta.(*api.Client)
	smartclassparameter := buildForemanSmartClassParameter(d)

	log.Debugf("ForemanSmartClassParameter: [%+v]", smartclassparameter)

	queryResponse, queryErr := client.QuerySmartClassParameter(smartclassparameter)
	if queryErr != nil {
		return queryErr
	}

	if queryResponse.Subtotal == 0 {
		return fmt.Errorf("Data source smart puppet class returned no results")
	} else if queryResponse.Subtotal > 1 {
		return fmt.Errorf("Data source smart puppet class returned more than 1 result")
	}

	var querySmartClassParameter api.ForemanSmartClassParameter
	var ok bool
	if querySmartClassParameter, ok = queryResponse.Results[0].(api.ForemanSmartClassParameter); !ok {
		return fmt.Errorf(
			"Data source results contain unexpected type. Expected "+
				"[api.ForemanSmartClassParameter], got [%T]",
			queryResponse.Results[0],
		)
	}
	smartclassparameter = &querySmartClassParameter

	log.Debugf("ForemanSmartClassParameter: [%+v]", smartclassparameter)

	setResourceDataFromForemanSmartClassParameter(d, smartclassparameter)

	return nil
}
