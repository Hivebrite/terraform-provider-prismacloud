module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/Hivebrite/prisma-cloud-go v0.3.10-0.20201016060206-f6452c2702ef // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.12
)

//replace github.com/hivebrite/prisma-cloud-go => ../prisma-cloud-go

go 1.13
