package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForemanSmartClassParameter() *schema.Resource {
	return &schema.Resource{

		Create: resourceForemanSmartClassParameterCreate,
		Read:   resourceForemanSmartClassParameterRead,
		Update: resourceForemanSmartClassParameterUpdate,
		Delete: resourceForemanSmartClassParameterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"%s Foreman representation of smart class parameter.",
					autodoc.MetaSummary,
				),
			},
			"match": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"value": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"use_puppet_default": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},
			"omit": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},
			"smart_class_parameter_id": &schema.Schema{
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

// buildForemanSmartClassParameter constructs a ForemanSmartClassParameter reference from a resource data
// reference.  The struct's  members are populated from the data populated in
// the resource data.  Missing members will be left to the zero value for that
// member's type.
func buildForemanSmartClassParameter(d *schema.ResourceData) *api.ForemanSmartClassParameter {
	log.Tracef("resource_foreman_smartclassparameter.go#buildForemanSmartClassParameter")

	smartclassparameter := api.ForemanSmartClassParameter{}

	obj := buildForemanObject(d)
	smartclassparameter.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("match"); ok {
		smartclassparameter.Match = attr.(string)
	}
	if attr, ok = d.GetOk("value"); ok {
		smartclassparameter.Value = attr.(string)
	}
	if attr, ok = d.GetOk("use_puppet_default"); ok {
		smartclassparameter.UsePuppetDefault = attr.(bool)
	}
	if attr, ok = d.GetOk("omit"); ok {
		smartclassparameter.Omit = attr.(bool)
	}
	if attr, ok = d.GetOk("smart_class_parameter_id"); ok {
		smartclassparameter.SmartClassParameterId = attr.(string)
	}

	return &smartclassparameter
}

// setResourceDataFromForemanSmartClassParameter sets a ResourceData's attributes from the
// attributes of the supplied ForemanSmartClassParameter reference
func setResourceDataFromForemanSmartClassParameter(d *schema.ResourceData, fd *api.ForemanSmartClassParameter) {
	log.Tracef("resource_foreman_smartclassparameter.go#setResourceDataFromForemanSmartClassParameter")

	d.SetId(strconv.Itoa(fd.Id))
	d.Set("match", fd.Match)
	d.Set("value", fd.Value)
	d.Set("use_puppet_default", fd.UsePuppetDefault)
	d.Set("omit", fd.Omit)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanSmartClassParameterCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_smartclassparameter.go#Create")

	client := meta.(*api.Client)
	m := buildForemanSmartClassParameter(d)

	log.Debugf("ForemanSmartClassParameter: [%+v]", m)

	createdSmartClassParameter, createErr := client.CreateSmartClassParameter(m)
	if createErr != nil {
		return createErr
	}
	log.Debugf("Created ForemanSmartClassParameter: [%+v]", createdSmartClassParameter)

	setResourceDataFromForemanSmartClassParameter(d, createdSmartClassParameter)
	return nil
}

func resourceForemanSmartClassParameterRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_smartclassparameter.go#Read")

	client := meta.(*api.Client)
	smartclassparameter := buildForemanSmartClassParameter(d)

	log.Debugf("ForemanSmartClassParameter: [%+v]", smartclassparameter)

	readSmartClassParameter, readErr := client.ReadSmartClassParameter(smartclassparameter)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanSmartClassParameter: [%+v]", readSmartClassParameter)

	setResourceDataFromForemanSmartClassParameter(d, readSmartClassParameter)

	return nil
}

func resourceForemanSmartClassParameterUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_smartclassparameter.go#Update")

	client := meta.(*api.Client)
	e := buildForemanSmartClassParameter(d)

	log.Debugf("ForemanSmartClassParameter: [%+v]", e)

	updatedSmartClassParameter, updateErr := client.UpdateSmartClassParameter(e)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanSmartClassParameter: [%+v]", updatedSmartClassParameter)

	setResourceDataFromForemanSmartClassParameter(d, updatedSmartClassParameter)
	return nil
}

func resourceForemanSmartClassParameterDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_smartclassparameter.go#Delete")

	client := meta.(*api.Client)
	p := buildForemanSmartClassParameter(d)

	log.Debugf("ForemanSmartClassParameter: [%+v]", p)

	return client.DeleteSmartClassParameter(p)
}
