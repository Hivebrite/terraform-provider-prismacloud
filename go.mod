module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.3.12
	github.com/spf13/cast v1.3.1
)

// go mod edit -replace github.com/paloaltonetworks/prisma-cloud-go=github.com/Hivebrite/prisma-cloud-go@SHA

replace github.com/paloaltonetworks/prisma-cloud-go => github.com/Hivebrite/prisma-cloud-go v0.3.10-0.20210325102847-a7c9432777c1

// replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

go 1.13
