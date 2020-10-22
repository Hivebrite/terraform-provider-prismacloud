package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/search"
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSearch() *schema.Resource {
	return &schema.Resource{
		Create: createSearch,
		Read:   readSearch,
		Delete: deleteSearch,

		Schema: map[string]*schema.Schema{
			"search_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RQL ID",
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "RQL query",
			},
			"search_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "RQL type",
			},
			"time_range": timeRangeSchema("resource_rql_historic_search"),
		},
	}
}

func parseSearch(d *schema.ResourceData, id string) history.Query {
	ans := history.Query{
		Query:      d.Get("query").(string),
		SearchType: d.Get("search_type").(string),
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

func saveSearch(d *schema.ResourceData, obj history.Query) {
	d.Set("search_id", obj.Id)
	d.Set("search_type", obj.SearchType)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("query", obj.Query)
}

func createSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseSearch(d, "")

	if err := search.Create(client, obj); err != nil {
		return err
	}

	id, err := search.Identify(client, obj.Query)
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
