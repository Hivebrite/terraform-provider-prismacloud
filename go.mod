module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.9
)

replace github.com/paloaltonetworks/prisma-cloud-go => github.com/hivebrite/prisma-cloud-go v0.3.10-0.20201026134347-d8022fca5581
// replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go
go 1.13
