package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCmppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cmppolicy,
		Read:          read_cmppolicy,
		Update:        update_cmppolicy,
		Delete:        delete_cmppolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_cmppolicy(d *schema.ResourceData) nitro.Cmppolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cmppolicy{
		Name:      d.Get("name").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	return resource
}

func set_cmppolicy(d *schema.ResourceData, resource *nitro.Cmppolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("name", resource.Name)
	d.Set("resaction", resource.Resaction)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_cmppolicy_key(d *schema.ResourceData) nitro.CmppolicyKey {

	key := nitro.CmppolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_cmppolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cmppolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmppolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cmppolicy(d, resource)
	} else {
		err := client.AddCmppolicy(get_cmppolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCmppolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cmppolicy(d, resource)
	}

	return nil
}

func read_cmppolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cmppolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmppolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cmppolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cmppolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cmppolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.CmppolicyUpdate{}
	unset := nitro.CmppolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("resaction") {
		updateFlag = true

		value := d.Get("resaction").(string)
		update.Resaction = value

		if value == "" {
			unsetFlag = true

			unset.Resaction = true
		}

	}
	key := get_cmppolicy_key(d)

	if updateFlag {
		if err := client.UpdateCmppolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetCmppolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetCmppolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_cmppolicy(d, resource)
	}

	return nil
}

func delete_cmppolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cmppolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCmppolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
