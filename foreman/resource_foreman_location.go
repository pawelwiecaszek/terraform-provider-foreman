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

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Description: fmt.Sprintf(
					"Description of the location",
					autodoc.MetaExample,
				),
			},

			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the users associated with this location",
			},

			"smart_proxies": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the smart proxies associated with this location",
			},

			"compute_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the compute resources associated with this location",
			},

			"media": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the medias associated with this location",
			},

			"config_templates": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the config templates associated with this location",
			},

			"ptable_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the partition tables associated with this location",
			},

			"provisioning_templates": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the provisioning templates associated with this location",
			},

			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the domains associated with this location",
			},

			"realms": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the realms associated with this location",
			},

			"hostgroups": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the hostgroups associated with this location",
			},

			"environments": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the environments associated with this location",
			},

			"subnets": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the subnets associated with this location",
			},

			"parent_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the parents associated with this location",
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

	if attr, ok = d.GetOk("description"); ok {
		location.Description = attr.(string)
	}

	if attr, ok = d.GetOk("users"); ok {
		attrSet := attr.(*schema.Set)
		location.UserIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("smart_proxies"); ok {
		attrSet := attr.(*schema.Set)
		location.SmartProxyIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("compute_resources"); ok {
		attrSet := attr.(*schema.Set)
		location.ComputeResourceIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("media"); ok {
		attrSet := attr.(*schema.Set)
		location.MediaIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("config_templates"); ok {
		attrSet := attr.(*schema.Set)
		location.ConfigTemplateIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("ptables"); ok {
		attrSet := attr.(*schema.Set)
		location.PtableIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("provisioning_templates"); ok {
		attrSet := attr.(*schema.Set)
		location.ProvisioningTemplateIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("domains"); ok {
		attrSet := attr.(*schema.Set)
		location.DomainIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("realms"); ok {
		attrSet := attr.(*schema.Set)
		location.RealmIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("hostgroups"); ok {
		attrSet := attr.(*schema.Set)
		location.HostgroupIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("environments"); ok {
		attrSet := attr.(*schema.Set)
		location.EnvironmentIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("subnets"); ok {
		attrSet := attr.(*schema.Set)
		location.SubnetIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("parent_id"); ok {
		location.ParentId = attr.(int)
	}

	return &location
}

// setResourceDataFromForemanLocation sets a ResourceData's attributes from
// the attributes of the supplied ForemanLocation reference
func setResourceDataFromForemanLocation(d *schema.ResourceData, fe *api.ForemanLocation) {
	log.Tracef("resource_foreman_location.go#setResourceDataFromForemanLocation")

	d.SetId(strconv.Itoa(fe.Id))
	d.Set("name", fe.Name)
	d.Set("description", fe.Description)
	d.Set("users", fe.UserIds)
	d.Set("smart_proxies", fe.SmartProxyIds)
	d.Set("compute_resources", fe.ComputeResourceIds)
	d.Set("media", fe.MediaIds)
	d.Set("config_templates", fe.ConfigTemplateIds)
	d.Set("ptables", fe.PtableIds)
	d.Set("provisioning_templates", fe.ProvisioningTemplateIds)
	d.Set("domains", fe.DomainIds)
	d.Set("realms", fe.RealmIds)
	d.Set("hostgroups", fe.HostgroupIds)
	d.Set("environments", fe.EnvironmentIds)
	d.Set("subnets", fe.SubnetIds)
	d.Set("parent_id", fe.ParentId)
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
