package foreman

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/wayfair/terraform-provider-utils/autodoc"
	"github.com/wayfair/terraform-provider-utils/log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceForemanComputeAttributes() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanComputeAttributesCreate,
		Read:   resourceForemanComputeAttributesRead,
		Update: resourceForemanComputeAttributesUpdate,
		Delete: resourceForemanComputeAttributesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s Foreman representation of compute attributes. Compute Attributes define"+
						"the actual values the compute resources use to create virtual machines. "+
						"This includes but is not limited to CPU shares, memory, hard drives and "+
						"network interfaces",
					autodoc.MetaSummary,
				),
			},

			"compute_resource_id": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the Compute Resource to create the Attributes on",
			},

			"compute_profile_id": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "ID of the Compute Profile to add the Attributes to",
			},

			"compute_attributes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "Freestyle JSON Map of the VM Attributes to associate with" +
					"the Compute Resource and Profile",
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanComputeAttributes constructs a ForemanComputeAttributes reference from a resource data
// reference.  The struct's  members are populated from the data populated in
// the resource data.  Missing members will be left to the zero value for that
// member's type.
func buildForemanComputeAttributes(d *schema.ResourceData) *api.ForemanComputeAttributes {
	log.Tracef("resource_foreman_computeresource.go#buildForemanComputeAttributes")

	computeresource := api.ForemanComputeAttributes{}

	obj := buildForemanObject(d)
	computeresource.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("compute_resource_id"); ok {
		computeresource.ComputeResourceID = attr.(int)
	}
	if attr, ok = d.GetOk("compute_profile_id"); ok {
		computeresource.ComputeProfileID = attr.(int)
	}
	if attr, ok = d.GetOk("compute_attributes"); ok {
		computeresource.ComputeAttribute.VMAttrs = json.RawMessage(attr.(string))
	}
	return &computeresource
}

// setResourceDataFromForemanComputeAttributes sets a ResourceData's attributes from the
// attributes of the supplied ForemanComputeAttributes reference
func setResourceDataFromForemanComputeAttributes(d *schema.ResourceData, fd *api.ForemanComputeAttributes) {
	log.Tracef("resource_foreman_computeresource.go#setResourceDataFromForemanComputeAttributes")

	d.SetId(strconv.Itoa(fd.Id))
	d.Set("compute_resource_id", fd.ComputeResourceID)
	d.Set("compute_profile_id", fd.ComputeProfileID)
	d.Set("compute_attributes", string(fd.ComputeAttribute.VMAttrs))
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanComputeAttributesCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattribute.go#Create")

	client := meta.(*api.Client)
	m := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", m)

	createdComputeAttributes, createErr := client.CreateComputeAttributes(m)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanComputeAttributes: [%+v]", createdComputeAttributes)

	setResourceDataFromForemanComputeAttributes(d, createdComputeAttributes)

	return nil
}

func resourceForemanComputeAttributesRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattribute.go#Read")

	client := meta.(*api.Client)
	m := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", m)

	readComputeAttributes, readErr := client.ReadComputeAttributes(m)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanComputeAttributes: [%+v]", readComputeAttributes)

	setResourceDataFromForemanComputeAttributes(d, readComputeAttributes)

	return nil
}

func resourceForemanComputeAttributesUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattribute.go#Update")

	client := meta.(*api.Client)
	m := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", m)

	updatedComputeAttributes, updateErr := client.UpdateComputeAttributes(m)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanComputeAttributes: [%+v]", updatedComputeAttributes)

	setResourceDataFromForemanComputeAttributes(d, updatedComputeAttributes)

	return nil
}

func resourceForemanComputeAttributesDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattribute.go#Delete")

	return nil
}
