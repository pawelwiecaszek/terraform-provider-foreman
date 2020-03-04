package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForemanComputeProfile() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanComputeProfileCreate,
		Read:   resourceForemanComputeProfileRead,
		Update: resourceForemanComputeProfileUpdate,
		Delete: resourceForemanComputeProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s A ComputeProfile.",
					autodoc.MetaSummary,
				),
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: fmt.Sprintf(
					"Name of the ComputeProfile",
					autodoc.MetaExample,
				),
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanComputeProfile constructs a ForemanComputeProfile reference from a
// resource data reference.  The struct's  members are populated from the data
// populated in the resource data.  Missing members will be left to the zero
// value for that member's type.
func buildForemanComputeProfile(d *schema.ResourceData) *api.ForemanComputeProfile {
	log.Tracef("resource_foreman_ComputeProfile.go#buildForemanComputeProfile")

	ComputeProfile := api.ForemanComputeProfile{}

	obj := buildForemanObject(d)
	ComputeProfile.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("name"); ok {
		ComputeProfile.Name = attr.(string)
	}

	return &ComputeProfile
}

// setResourceDataFromForemanComputeProfile sets a ResourceData's attributes from
// the attributes of the supplied ForemanComputeProfile reference
func setResourceDataFromForemanComputeProfile(d *schema.ResourceData, fe *api.ForemanComputeProfile) {
	log.Tracef("resource_foreman_ComputeProfile.go#setResourceDataFromForemanComputeProfile")

	d.SetId(strconv.Itoa(fe.Id))
	d.Set("name", fe.Name)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanComputeProfileCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_ComputeProfile.go#Create")

	client := meta.(*api.Client)
	e := buildForemanComputeProfile(d)

	log.Debugf("ForemanComputeProfile: [%+v]", e)

	createdComputeProfile, createErr := client.CreateComputeProfile(e)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanComputeProfile: [%+v]", createdComputeProfile)

	setResourceDataFromForemanComputeProfile(d, createdComputeProfile)

	return nil
}

func resourceForemanComputeProfileRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_ComputeProfile.go#Read")

	client := meta.(*api.Client)
	e := buildForemanComputeProfile(d)

	log.Debugf("ForemanComputeProfile: [%+v]", e)

	readComputeProfile, readErr := client.ReadComputeProfile(e.Id)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanComputeProfile: [%+v]", readComputeProfile)

	setResourceDataFromForemanComputeProfile(d, readComputeProfile)

	return nil
}

func resourceForemanComputeProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_ComputeProfile.go#Update")

	client := meta.(*api.Client)
	e := buildForemanComputeProfile(d)

	log.Debugf("ForemanComputeProfile: [%+v]", e)

	updatedComputeProfile, updateErr := client.UpdateComputeProfile(e)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanComputeProfile: [%+v]", updatedComputeProfile)

	setResourceDataFromForemanComputeProfile(d, updatedComputeProfile)

	return nil
}

func resourceForemanComputeProfileDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_ComputeProfile.go#Delete")

	client := meta.(*api.Client)
	e := buildForemanComputeProfile(d)

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors

	return client.DeleteComputeProfile(e.Id)
}