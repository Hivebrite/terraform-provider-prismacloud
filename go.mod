module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.9
)

replace github.com/paloaltonetworks/prisma-cloud-go => github.com/hivebrite/prisma-cloud-go v0.3.10-0.20201021160311-01bf0c7a74f9
//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go
go 1.13
