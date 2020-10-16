package prismacloud

import (
	pc "github.com/hivebrite/prisma-cloud-go"
	"github.com/hivebrite/prisma-cloud-go/compliance/standard"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceComplianceStandard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComplianceStandardRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cs_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "Compliance standard ID",
				AtLeastOneOf: []string{"cs_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Compliance standard name",
				AtLeastOneOf: []string{"cs_id", "name"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"system_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "System default",
			},
			"policies_assigned_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of assigned policies",
			},
			"cloud_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cloud type (determined based on policies assigned)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceComplianceStandardRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	var err error
	csId := d.Get("cs_id").(string)

	if csId == "" {
		name := d.Get("name").(string)
		csId, err = standard.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	o, err := standard.Get(client, csId)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(o.Id)
	saveComplianceStandard(d, o)

	return nil
}
