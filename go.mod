module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.9
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

replace github.com/paloaltonetworks/prisma-cloud-go => github.com/Hivebrite/prisma-cloud-go v0.3.10-0.20200803141410-26f8a6410f75

go 1.13
