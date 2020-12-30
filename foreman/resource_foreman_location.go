package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/conv"
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
			"realms": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the realms.",
			},
			"compute_resource_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the compute_resource_ids.",
			},
			"domain_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the domain_ids.",
			},
			"subnet_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the subnet_ids.",
			},
			"environment_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the environment_ids.",
			},
			"hostgroup_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the hostgroup_ids.",
			},
			"provisioning_templates": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the provisioning_templates.",
			},
			"smart_proxy_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the smart_proxy_ids.",
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the users.",
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

	if attr, ok = d.GetOk("realms"); ok {
		attrSet := attr.(*schema.Set)
		location.Realms = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("compute_resource_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.ComputeResources = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("domain_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.Domains = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("subnet_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.Subnets = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("environment_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.Environments = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("hostgroup_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.Hostgroups = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("provisioning_templates"); ok {
		attrSet := attr.(*schema.Set)
		location.ProvisioningTemplates = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("smart_proxy_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.SmartProxies = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("users"); ok {
		attrSet := attr.(*schema.Set)
		location.Users = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	return &location
}

// setResourceDataFromForemanLocation sets a ResourceData's attributes from
// the attributes of the supplied ForemanLocation reference
func setResourceDataFromForemanLocation(d *schema.ResourceData, fe *api.ForemanLocation) {
	log.Tracef("resource_foreman_location.go#setResourceDataFromForemanLocation")

	d.SetId(strconv.Itoa(fe.Id))
	d.Set("name", fe.Name)
	d.Set("realms", fe.Realms)
	d.Set("compute_resource_ids", fe.ComputeResources)
	d.Set("domain_ids", fe.Domains)
	d.Set("subnet_ids", fe.Subnets)
	d.Set("environment_ids", fe.Environments)
	d.Set("hostgroup_ids", fe.Hostgroups)
	d.Set("provisioning_templates", fe.ProvisioningTemplates)
	d.Set("smart_proxy_ids", fe.SmartProxies)
	d.Set("users", fe.Users)
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
