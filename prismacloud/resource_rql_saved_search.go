package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRqlSavedSearch() *schema.Resource {
	return &schema.Resource{
		Create: createRqlSearch,
		Read:   readRqlSearch,
		Update: updateRqlSearch,
		Delete: deleteRqlSearch,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "RQL config search name",
			},
			"search_id": {
				Type:        schema.TypeString,
				Required:    true,
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
			"search_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "RQL type",
			},
			"saved": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Save status",
			},
			"time_range": timeRangeSchema("resource_rql_historic_search"),
		},
	}
}

func parseRqlSearch(d *schema.ResourceData, id string) history.PostQuery {
	ans := history.PostQuery{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Query:       d.Get("query").(string),
		SearchType:  d.Get("search_type").(string),
		Saved:       true,
	}

	tr := ResourceDataInterfaceMap(d, "time_range")
	if atr := ToInterfaceMap(tr, "absolute"); len(atr) != 0 {
		ans.TimeRange.Value = timerange.Absolute{
			Start: atr["start"].(int),
			End:   atr["end"].(int),
		}
		ans.TimeRange.Type = timerange.TypeAbsolute
	} else if rtr := ToInterfaceMap(tr, "relative"); len(rtr) != 0 {
		ans.TimeRange.Value = timerange.Relative{
			Amount: rtr["amount"].(int),
			Unit:   rtr["unit"].(string),
		}
		ans.TimeRange.Type = timerange.TypeRelative
	} else if tntr := ToInterfaceMap(tr, "to_now"); len(tntr) != 0 {
		ans.TimeRange.Value = timerange.ToNow{
			Unit: tntr["unit"].(string),
		}
		ans.TimeRange.Type = timerange.TypeToNow

	}

	return ans
}

func saveRqlSearch(d *schema.ResourceData, obj history.Query) {
	d.Set("search_id", obj.Id)
	d.Set("search_type", obj.SearchType)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("query", obj.Query)
}

func createRqlSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseRqlSearch(d, "")

	if err := history.Create(client, obj, d.Get("search_id").(string)); err != nil {
		return err
	}

	id, err := history.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	d.SetId(id)
	return readRqlSearch(d, meta)
}

func readRqlSearch(d *schema.ResourceData, meta interface{}) error {
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

	saveRqlSearch(d, obj)

	return nil
}

func updateRqlSearch(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteRqlSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := history.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
