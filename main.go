package main

import (
    "github.com/Trility/tf-plugins/aws"
    "github.com/hashicorp/terraform/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: aws.Provider,
    })
}
