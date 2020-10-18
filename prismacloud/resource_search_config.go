package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/policy"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/search/config"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConfigSearch() *schema.Resource {
	return &schema.Resource{
		Create: createSearch,
		Read:   readSearch,
		Update: updateSearch,
		Delete: deleteSearch,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "RQL config search name",
			},
			"search_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RQL ID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "RQL query",
			},
			"time_range": timeRangeSchema("data_source_rql_historic_search"),
		},
	}
}

func parseSearch(d *schema.ResourceData, id string) history.Query {
	ans := history.Query{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Query:       d.Get("query").(string),
	}

	return ans
}

func saveSearch(d *schema.ResourceData, obj history.Query) {
	d.Set("search_id", obj.Id)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("query", obj.Query)
}

func createSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseSearch(d, "")

	if err := config.Create(client, obj); err != nil {
		return err
	}

	id, err := policy.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	d.SetId(id)
	return readSearch(d, meta)
}

func readSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := history.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveSearch(d, obj)

	return nil
}

func updateSearch(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteSearch(d *schema.ResourceData, meta interface{}) error {
	return nil
}
