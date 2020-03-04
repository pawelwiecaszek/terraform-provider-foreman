package foreman

import (
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-foreman/foreman/api"
	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
					"%s Compute attributes.",
					autodoc.MetaSummary,
				),
			},

			"compute_resource_id": &schema.Schema{
				Type:         schema.TypeInt,
				Required: 	  true,
				ForceNew:     true,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"compute_profile_id": &schema.Schema{
				Type:         schema.TypeInt,
				Required: 	  true,
				ForceNew:     true,
				ValidateFunc: validation.IntAtLeast(0),
			},

			"compute_attribute": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceForemanComputeAttributeInternal(),
				Set:      schema.HashResource(resourceForemanComputeAttributeInternal()),
				Description: "Compute attribute",
			},
		},
	}
}

// resourceForemanComputeAttributeInternal is a nested resource that
// represents a valid template combination attribute.  The "id" of this
// resource is computed and assigned by Foreman at the time of creation.
//
// NOTE(ALL): See comments in ResourceData's "compute_attribute"
//   attribute definition above
func resourceForemanComputeAttributeInternal() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vm_attrs": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Hypervisor specific options",
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanComputeAttributes constructs a ForemanComputeAttributes
// struct from a resource data reference.  The struct's members are populated
// with the data populated in the resource data.  Missing members will be left
// to the zero value for that member's type.
func buildForemanComputeAttributes(d *schema.ResourceData) *api.ForemanComputeAttributes {
	log.Tracef("resource_foreman_computeattributes.go#buildForemanComputeAttributes")

	computeattributes := api.ForemanComputeAttributes{}

	obj := buildForemanObject(d)
	computeattributes.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("compute_resource_id"); ok {
		computeattributes.ComputeResourceId = attr.(int)
	}
	if attr, ok = d.GetOk("compute_profile_id"); ok {
		computeattributes.ComputeProfileId = attr.(int)
	}

	computeattributes.ForemanComputeAttributeInternal = buildForemanForemanComputeAttributeInternal(d)

	return &computeattributes
}

// buildForemanForemanComputeAttributeInternal constructs an array of
// ForemanComputeAttributeInternal structs from a resource data reference.
// The struct's members are populated with the data populated in the resource
// data. Missing members will be left to the zero value for that member's type.
func buildForemanForemanComputeAttributeInternal(d *schema.ResourceData) []api.ForemanComputeAttributeInternal {
	log.Tracef("resource_foreman_computeattributes.go#buildForemanForemanComputeAttributeInternal")

	tempComboCAttr := []api.ForemanComputeAttributeInternal{}
	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("compute_attribute"); !ok {
		return tempComboCAttr
	}

	// type assert the underlying *schema.Set and convert to a list
	attrSet := attr.(*schema.Set)
	attrList := attrSet.List()
	attrListLen := len(attrList)
	tempComboCAttr = make([]api.ForemanComputeAttributeInternal, attrListLen)
	// iterate over each of the map structure entires in the set and convert that
	// to a concrete struct implementation to append to the template combinations
	// attributes list.
	for idx, attrMap := range attrList {
		tempComboCAttrMap := attrMap.(map[string]interface{})
		tempComboCAttr[idx] = mapToForemanComputeAttributeInternal(tempComboCAttrMap)
	}

	return tempComboCAttr
}

// mapToForemanComputeAttributeInternal converts a map[string]interface{}
// to a ForemanComputeAttributeInternal struct.  The supplied map comes
// from an entry in the *schema.Set for the "compute_attribute"
// property of the resource, since *schema.Set stores its entries as this map
// structure.
//
// The map should have the following keys. Omitted or invalid map values will
// result in the struct receiving the zero value for that property.
//
//   id (int)
//   hostgroup_id (int)
//   environment_id (int)
//   _destroy (bool)
func mapToForemanComputeAttributeInternal(m map[string]interface{}) api.ForemanComputeAttributeInternal {
	log.Tracef("mapToForemanComputeAttributeInternal")

	tempComboCAttr := api.ForemanComputeAttributeInternal{}
	var ok bool
	
	if tempComboCAttr.VMAttributes, ok = m["vm_attrs"].(map[string]interface{}); !ok {
		tempComboCAttr.VMAttributes = nil
	}

	log.Debugf("m: [%v], tempComboCAttr: [%+v]", m, tempComboCAttr)
	return tempComboCAttr
}

// setResourceDataFromForemanComputeAttributes sets a ResourceData's
// attributes from the attributes of the supplied ForemanComputeAttributes
// struct
func setResourceDataFromForemanComputeAttributes(d *schema.ResourceData, ft *api.ForemanComputeAttributes) {
	log.Tracef("resource_foreman_computeattributes.go#setResourceDataFromForemanComputeAttributes")

	d.SetId(strconv.Itoa(ft.Id))

	d.Set("compute_resource_id", ft.ComputeResourceId)
	d.Set("compute_profile_id", ft.ComputeProfileId)

	setResourceDataFromForemanForemanComputeAttributeInternal(d, ft.ForemanComputeAttributeInternal)

}

// setResourceDataFromForemanForemanComputeAttributeInternal sets a
// ResourceData's "compute_attribute" attribute to the value of
// the supplied array of ForemanComputeAttributeInternal structs
func setResourceDataFromForemanForemanComputeAttributeInternal(d *schema.ResourceData, ftca []api.ForemanComputeAttributeInternal) {
	log.Tracef("resource_foreman_computeattributes.go#setResourceDataFromForemanTemplateCombinationsAttributes")

	// this attribute is a *schema.Set.  In order to construct a set, we need to
	// supply a hash function so the set can differentiate for uniqueness of
	// entries.  The hash function will be based on the resource definition
	hashFunc := schema.HashResource(resourceForemanComputeAttributeInternal())
	// underneath, a *schema.Set stores an array of map[string]interface{} entries.
	// convert each ForemanTemplateCombination struct in the supplied array to a
	// mapstructure and then add it to the set
	ifaceArr := make([]interface{}, len(ftca))
	for idx, val := range ftca {
		// NOTE(ALL): we ommit the "_destroy" property here - this does not need
		//   to be stored by terraform in the state file. That is a hidden key that
		//   is only used in updates.  Anything that exists will always have it
		//   set to "false".
		ifaceMap := map[string]interface{}{
			"vm_attrs": val.VMAttributes,
		}
		ifaceArr[idx] = ifaceMap
	}
	// with the array set up, create the *schema.Set and set the ResourceData's
	// "compute_attribute" property
	tempComboCAttrSet := schema.NewSet(hashFunc, ifaceArr)
	d.Set("compute_attribute", tempComboCAttrSet)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanComputeAttributesCreate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattributes.go#Create")

	client := meta.(*api.Client)
	t := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", t)

	createdTemplate, createErr := client.CreateComputeAttributes(t)
	if createErr != nil {
		return createErr
	}

	log.Debugf("Created ForemanComputeAttributes: [%+v]", createdTemplate)

	setResourceDataFromForemanComputeAttributes(d, createdTemplate)

	return nil
}

func resourceForemanComputeAttributesRead(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattributes.go#Read")

	client := meta.(*api.Client)
	t := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", t)

	readTemplate, readErr := client.ReadComputeAttributes(t.Id)
	if readErr != nil {
		return readErr
	}

	log.Debugf("Read ForemanComputeAttributes: [%+v]", readTemplate)

	log.Tracef("BeforeSet: %v", d.Get("operatingsystem_ids"))
	setResourceDataFromForemanComputeAttributes(d, readTemplate)
	log.Tracef("AfterSet: %v", d.Get("operatingsystem_ids"))

	return nil
}

func resourceForemanComputeAttributesUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattributes.go#Update")

	client := meta.(*api.Client)
	t := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", t)

	// NOTE(ALL): Handling the removal of a template combination.  See the note
	//   in ForemanComputeAttributeInternal's Destroy property
	if d.HasChange("compute_attribute") {
		oldVal, newVal := d.GetChange("compute_attribute")
		oldValSet, newValSet := oldVal.(*schema.Set), newVal.(*schema.Set)

		// NOTE(ALL): The set difference operation is anticommutative (because math)
		//   ie: [A - B] =/= [B - A].
		//
		//   When performing an update, we need to figure out which template
		//   combinations were removed from the set and tag the destroy property
		//   to true and instruct Foreman which ones to delete from the list. We do
		//   this by performing a set difference between the old set and the new
		//   set (ie: [old - new]) which will return the items that used to be in
		//   the set but are no longer included.
		//
		//   The values that were added to the set or remained unchanged are already
		//   part of the template combinations.  They are present in the
		//   ResourceData and already exist from the
		//   buildForemanComputeAttributes() call.

		setDiff := oldValSet.Difference(newValSet)
		setDiffList := setDiff.List()
		log.Debugf("setDiffList: [%v]", setDiffList)

		// iterate over the removed items, add them back to the template's
		// combination array, but tag them for removal.
		//
		// each of the set's items is stored as a map[string]interface{} - use
		// type assertion and construct the struct
		for _, rmVal := range setDiffList {
			// construct, tag for deletion from list of combinations
			rmValMap := rmVal.(map[string]interface{})
			rmCombination := mapToForemanComputeAttributeInternal(rmValMap)
			// append back to template's list
			t.ForemanComputeAttributeInternal = append(t.ForemanComputeAttributeInternal, rmCombination)
		}

		log.Debugf("ForemanComputeAttributes: [%+v]", t)

	} // end HasChange("compute_attribute")

	updatedTemplate, updateErr := client.UpdateComputeAttributes(t)
	if updateErr != nil {
		return updateErr
	}

	log.Debugf("Updated ForemanComputeAttributes: [%+v]", t)

	setResourceDataFromForemanComputeAttributes(d, updatedTemplate)

	return nil
}

func resourceForemanComputeAttributesDelete(d *schema.ResourceData, meta interface{}) error {
	log.Tracef("resource_foreman_computeattributes.go#Delete")

	client := meta.(*api.Client)
	t := buildForemanComputeAttributes(d)

	log.Debugf("ForemanComputeAttributes: [%+v]", t)

	// NOTE(ALL): The Foreman API will return a '422: Unprocessable Entity' error
	//   if you try to delete a provisioning template with template combinations.
	//   First, you must update the provisioning template to remove the combinations,
	//   then proceed with deletion.
	if len(t.ForemanComputeAttributeInternal) > 0 {
		log.Debugf("deleting template that has combinations set")
		// iterate through each of the template combinations and tag them for
		// removal from the list
		updatedTemplate, updateErr := client.UpdateComputeAttributes(t)
		if updateErr != nil {
			return updateErr
		}

		log.Debugf("Updated ForemanComputeAttributes: [%+v]", updatedTemplate)

		// NOTE(ALL): set the resource data's properties to what comes back from
		//   the update call. This allows us to recover from a partial state if
		//   delete encounters an error after this point - at least the resource's
		//   state will be saved with the correct template combinations.
		setResourceDataFromForemanComputeAttributes(d, updatedTemplate)

		log.Debugf("completed the template combination deletion")

	} // end if len(template.ForemanComputeAttributeInternal) > 0

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors
	return client.DeleteComputeAttributes(t.Id)
}
