module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.12
)

replace github.com/paloaltonetworks/prisma-cloud-go => github.com/hivebrite/prisma-cloud-go v0.3.10-0.20201022080149-9471379f9a37
//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go
go 1.13
