package cre

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
	vnet "github.com/skurhse/chitin/generated/hashicorp/azurerm/virtualnetwork"
)

type CoreMelody struct {
	Stack_          cdktf.TerraformStack
	StackName_      *string
	Client_         cnf.DataAzurermClientConfig
	VirtualNetwork_ vnet.VirtualNetwork
}

func (c CoreMelody) Stack() cdktf.TerraformStack {
	return c.Stack_
}

func (c CoreMelody) StackName() *string {
	return c.StackName_
}

func (c CoreMelody) Client() cnf.DataAzurermClientConfig {
	return c.Client_
}

func (c CoreMelody) VirtualNetwork() vnet.VirtualNetwork {
	return c.VirtualNetwork_
}
