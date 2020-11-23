package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForemanPuppetClass() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanPuppetClassCreate,
		Read:   resourceForemanPuppetClassRead,
		Update: resourceForemanPuppetClassUpdate,
		Delete: resourceForemanPuppetClassDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s Foreman representation of image.",
					autodoc.MetaSummary,
				),
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanPuppetClass constructs a ForemanPuppetClass reference from a resource data
// reference.  The struct's  members are populated from the data populated in
// the resource data.  Missing members will be left to the zero value for that
// member's type.
func buildForemanPuppetClass(d *schema.ResourceData) *api.ForemanPuppetClass {
	log.Tracef("resource_foreman_puppetclass.go#buildForemanPuppetClass")

	puppetclass := api.ForemanPuppetClass{}

	obj := buildForemanObject(d)
	puppetclass.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("name"); ok {
		puppetclass.Name = attr.(string)
	}

	return &puppetclass
}

// setResourceDataFromForemanPuppetClass sets a ResourceData's attributes from the
// attributes of the supplied ForemanPuppetClass reference
func setResourceDataFromForemanPuppetClass(d *schema.ResourceData, fd *api.ForemanPuppetClass) {
	log.Tracef("resource_foreman_puppetclass.go#setResourceDataFromForemanPuppetClass")

	d.SetId(strconv.Itoa(fd.Id))
	d.Set("name", fd.Name)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanPuppetClassCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_puppetclass.go#Create")
	client := meta.(*api.Client)
	m := buildForemanPuppetClass(d)

	log.Debugf("ForemanPuppetClass: [%+v]", m)

	createdPuppetClass, createErr := client.CreatePuppetClass(m)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanPuppetClass: [%+v]", createdPuppetClass)

	setResourceDataFromForemanPuppetClass(d, createdPuppetClass)

	return nil
}

func resourceForemanPuppetClassRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_puppetclass.go#Read")

	client := meta.(*api.Client)
	puppetclass := buildForemanPuppetClass(d)

	log.Debugf("ForemanPuppetClass: [%+v]", puppetclass)

	readPuppetClass, readErr := client.ReadPuppetClass(puppetclass.Id)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanPuppetClass: [%+v]", readPuppetClass)

	setResourceDataFromForemanPuppetClass(d, readPuppetClass)

	return nil
}

func resourceForemanPuppetClassUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_puppetclass.go#Update")
	client := meta.(*api.Client)
	m := buildForemanPuppetClass(d)

	log.Debugf("ForemanPuppetClass: [%+v]", m)

	updatedPuppetClass, updateErr := client.UpdatePuppetClass(m)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanPuppetClass: [%+v]", updatedPuppetClass)

	setResourceDataFromForemanPuppetClass(d, updatedPuppetClass)

	return nil
}

func resourceForemanPuppetClassDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_puppetclass.go#Delete")

	client := meta.(*api.Client)
	m := buildForemanPuppetClass(d)

	log.Debugf("ForemanPuppetClass: [%+v]", m)

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors
	return client.DeletePuppetClass(m.Id)
}
