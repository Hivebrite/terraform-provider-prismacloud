package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"
	"github.com/spf13/cast"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUserProfile() *schema.Resource {
	return &schema.Resource{
		Create: createUserProfile,
		Read:   readUserProfile,
		Update: updateUserProfile,
		Delete: deleteUserProfile,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"firstname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Firstname of the user",
			},
			"lastname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "lastname of the user",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "email of the user",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "account enabled",
			},
			"access_keys_allowed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "can use API",
			},
			"default_role_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "default role",
			},
			"time_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Europe/Paris",
				Description: "timezone",
			},
			"role_ids": {
				Type:        schema.TypeList,
				Description: "Roles",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func parseUserProfile(d *schema.ResourceData) *profile.UserProfile {
	return &profile.UserProfile{
		Firstname:         d.Get("firstname").(string),
		Lastname:          d.Get("lastname").(string),
		Email:             d.Get("email").(string),
		Enabled:           d.Get("enabled").(bool),
		Timezone:          d.Get("time_zone").(string),
		AccessKeysAllowed: d.Get("access_keys_allowed").(bool),
		DefaultRoleId:     d.Get("default_role_id").(string),
		RoleIds:           cast.ToStringSlice(d.Get("role_ids")),
	}
}

func saveUserProfile(d *schema.ResourceData, obj profile.UserProfile) {
	d.Set("firstname", obj.Firstname)
	d.Set("lastname", obj.Lastname)
	d.Set("email", obj.Email)
	d.Set("time_zone", obj.Timezone)
	d.Set("enabled", obj.Enabled)
	d.Set("access_keys_allowed", obj.AccessKeysAllowed)
	d.Set("default_role_id", obj.DefaultRoleId)
	d.Set("role_ids", obj.RoleIds)
}

func createUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseUserProfile(d)

	if err := profile.Create(client, *obj); err != nil {
		return err
	}

	d.SetId(obj.Email)
	return readUserProfile(d, meta)
}

func readUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := profile.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveUserProfile(d, obj)

	return nil
}

func updateUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseUserProfile(d)

	if err := profile.Update(client, *obj); err != nil {
		return err
	}

	return readUserProfile(d, meta)
}

func deleteUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := profile.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
