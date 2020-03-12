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

			"user_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the users associated with this location",
			},

			"smart_proxy_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the smart proxies associated with this location",
			},

			"compute_resource_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the compute resources associated with this location",
			},

			"media_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the medias associated with this location",
			},

			"config_template_ids": &schema.Schema{
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

			"provisiong_template_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the provisioning templates associated with this location",
			},

			"domain_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the domains associated with this location",
			},

			"realm_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the realms associated with this location",
			},

			"hostgroup_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the hostgroups associated with this location",
			},

			"environment_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the environments associated with this location",
			},

			"subnet_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the subnets associated with this location",
			},

			"parent_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "IDs of the parents associated with this location",
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

	if attr, ok = d.GetOk("user_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.DomainIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("smart_proxy_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.SmartProxyIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("compute_resource_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.ComputeResourceIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("media_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.MediaIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("config_template_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.ConfigTemplateIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("ptables_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.PtableIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("provisiong_template_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.ProvisioningTemplateIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("domain_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.DomainIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("realm_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.RealmIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("hostgroup_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.HostgroupIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("environment_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.EnvironmentIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("subnet_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.SubnetIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("parent_ids"); ok {
		attrSet := attr.(*schema.Set)
		location.ParentIds = conv.InterfaceSliceToIntSlice(attrSet.List())
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
	d.Set("user_ids", fe.UserIds)
	d.Set("smart_proxy_ids", fe.SmartProxyIds)
	d.Set("compute_resource_ids", fe.ComputeResourceIds)
	d.Set("media_ids", fe.MediaIds)
	d.Set("config_template_ids", fe.ConfigTemplateIds)
	d.Set("ptable_ids", fe.PtableIds)
	d.Set("provisioning_template_ids", fe.ProvisioningTemplateIds)
	d.Set("domain_ids", fe.DomainIds)
	d.Set("realm_ids", fe.RealmIds)
	d.Set("hostgroup_ids", fe.HostgroupIds)
	d.Set("environment_ids", fe.EnvironmentIds)
	d.Set("subnet_ids", fe.SubnetIds)
	d.Set("parent_ids", fe.ParentIds)
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
