package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppflowpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowpolicylabel,
		Read:          read_appflowpolicylabel,
		Update:        nil,
		Delete:        delete_appflowpolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_appflowpolicylabel(d *schema.ResourceData) nitro.Appflowpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_appflowpolicylabel(d *schema.ResourceData, resource *nitro.Appflowpolicylabel) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("labelname", resource.Labelname)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func get_appflowpolicylabel_key(d *schema.ResourceData) nitro.AppflowpolicylabelKey {

	key := nitro.AppflowpolicylabelKey{
		d.Get("labelname").(string),
	}
	return key
}

func create_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	} else {
		err := client.AddAppflowpolicylabel(get_appflowpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	}

	return nil
}

func read_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
