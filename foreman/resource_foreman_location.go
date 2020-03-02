package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForemanLocation() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanLocationCreate,
		Read:   resourceForemanLocationRead,
		Update: resourceForemanLocationUpdate,
		Delete: resourceForemanLocationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s A location.",
					autodoc.MetaSummary,
				),
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: fmt.Sprintf(
					"Name of the location",
					autodoc.MetaExample,
				),
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanLocation constructs a ForemanLocation reference from a
// resource data reference.  The struct's  members are populated from the data
// populated in the resource data.  Missing members will be left to the zero
// value for that member's type.
func buildForemanLocation(d *schema.ResourceData) *api.ForemanLocation {
	log.Tracef("resource_foreman_location.go#buildForemanLocation")

	location := api.ForemanLocation{}

	obj := buildForemanObject(d)
	location.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("name"); ok {
		location.Name = attr.(string)
	}

	return &location
}

// setResourceDataFromForemanLocation sets a ResourceData's attributes from
// the attributes of the supplied ForemanLocation reference
func setResourceDataFromForemanLocation(d *schema.ResourceData, fe *api.ForemanLocation) {
	log.Tracef("resource_foreman_location.go#setResourceDataFromForemanLocation")

	d.SetId(strconv.Itoa(fe.Id))
	d.Set("name", fe.Name)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanLocationCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_location.go#Create")

	client := meta.(*api.Client)
	e := buildForemanLocation(d)

	log.Debugf("ForemanLocation: [%+v]", e)

	createdLocation, createErr := client.CreateLocation(e)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanLocation: [%+v]", createdLocation)

	setResourceDataFromForemanLocation(d, createdLocation)

	return nil
}

func resourceForemanLocationRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_location.go#Read")

	client := meta.(*api.Client)
	e := buildForemanLocation(d)

	log.Debugf("ForemanLocation: [%+v]", e)

	readLocation, readErr := client.ReadLocation(e.Id)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanLocation: [%+v]", readLocation)

	setResourceDataFromForemanLocation(d, readLocation)

	return nil
}

func resourceForemanLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_location.go#Update")

	client := meta.(*api.Client)
	e := buildForemanLocation(d)

	log.Debugf("ForemanLocation: [%+v]", e)

	updatedLocation, updateErr := client.UpdateLocation(e)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanLocation: [%+v]", updatedLocation)

	setResourceDataFromForemanLocation(d, updatedLocation)

	return nil
}

func resourceForemanLocationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_location.go#Delete")

	client := meta.(*api.Client)
	e := buildForemanLocation(d)

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors

	return client.DeleteLocation(e.Id)
}