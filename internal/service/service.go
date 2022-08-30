package service

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/mod/semver"
)

const (
	ResourceIdLen           = 20
	ClusterResourcesTagName = "percona-cluster-stack-id"
)

const (
	KeyPairName          = "key_pair_name"
	PathToKeyPairStorage = "path_to_key_pair_storage"
	ClusterSize          = "cluster_size"
	ConfigFilePath       = "config_file_path"
	InstanceType         = "instance_type"
	Version              = "version"
	VolumeType           = "volume_type"
	VolumeSize           = "volume_size"
	VolumeIOPS           = "volume_iops"
)

func DefaultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		KeyPairName: {
			Type:     schema.TypeString,
			Required: true,
		},
		PathToKeyPairStorage: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  ".",
		},
		ClusterSize: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		ConfigFilePath: {
			Type:     schema.TypeString,
			Optional: true,
		},
		InstanceType: {
			Type:     schema.TypeString,
			Required: true,
		},
		Version: {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(v interface{}, path cty.Path) diag.Diagnostics {
				version := v.(string)
				if version == "" {
					return diag.Errorf("empty version provided, use semantic versioning (MAJOR.MINOR.PATCH)")
				}
				if !semver.IsValid("v" + version) {
					return diag.Errorf("invalid version: %s, use semantic versioning (MAJOR.MINOR.PATCH)", version)
				}
				return nil
			},
		},
		VolumeType: {
			Type:     schema.TypeString,
			Optional: true,
		},
		VolumeSize: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		VolumeIOPS: {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

const (
	LogArgMasterIP   = "percona_master_ip"
	LogArgReplicaIP  = "percona_replica_ip"
	LogArgVersion    = "percona_version"
	LogArgInstanceIP = "percona_instance_ip"
)
