package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	DefaultSubnetCidrBlock = "10.0.1.0/16"

	DefaultSecurityGroupName        = "percona-security-group"
	DefaultSecurityGroupDescription = "Percona Terraform plugin security group"
)

var mapRegionImage = map[string]string{
	"us-east-1":      "ami-04505e74c0741db8d",
	"us-east-2":      "ami-0fb653ca2d3203ac1",
	"us-west-1":      "ami-01f87c43e618bf8f0",
	"us-west-2":      "ami-0892d3c7ee96c0bf7",
	"af-south-1":     "ami-0670428c515903d37",
	"ap-east-1":      "ami-0350928fdb53ae439",
	"ap-southeast-3": "ami-0f06496957d1fe04a",
	"ap-south-1":     "ami-05ba3a39a75be1ec4",
	"ap-northeast-3": "ami-0c2223049202ca738",
	"ap-northeast-2": "ami-0225bc2990c54ce9a",
	"ap-southeast-1": "ami-0750a20e9959e44ff",
	"ap-southeast-2": "ami-0d539270873f66397",
	"ap-northeast-1": "ami-0a3eb6ca097b78895",
	"ca-central-1":   "ami-073c944d45ffb4f27",
	"eu-central-1":   "ami-02584c1c9d05efa69",
	"eu-west-1":      "ami-00e7df8df28dfa791",
	"eu-west-2":      "ami-00826bd51e68b1487",
	"eu-south-1":     "ami-06ea0ad3f5adc2565",
	"eu-west-3":      "ami-0a21d1c76ac56fee7",
	"eu-north-1":     "ami-09f0506c9ef0fb473",
	"me-south-1":     "ami-05b680b37c7917206",
	"sa-east-1":      "ami-077518a464c82703b",
	"us-gov-east-1":  "ami-0eb7ef4cc0594fa04",
	"us-gov-west-1":  "ami-029a634618d6c0300",
}

func Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}