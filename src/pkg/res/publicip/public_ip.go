package publicip

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	ip "github.com/skurhse/chitin/generated/hashicorp/azurerm/publicip"
	"github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

func NewPublicIP(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup) ip.PublicIp {

	input := ip.PublicIpConfig{
		Name:                 naming.PublicIpOutput(),
		Location:             cfg.Regions().Primary(),
		ResourceGroupName:    rg.Name(),
		Sku:                  jsii.String("Basic"),
		AllocationMethod:     jsii.String("Dynamic"),
		IpVersion:            jsii.String("IPv4"),
		DomainNameLabel:      jsii.String(cfg.Name()),
		IdleTimeoutInMinutes: jsii.Number(4),
	}

	return ip.NewPublicIp(stack, res.Ids.PublicIP, &input)
}
