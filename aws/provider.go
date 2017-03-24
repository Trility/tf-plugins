package aws

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "access_key": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "AWS Access Key",
            },
            "secret_key": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "AWS Secret Key",
            },
            "profile": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "AWS Creds File Profile",
            },
            "max_retries": &schema.Schema{
                Type: schema.TypeInt,
                Optional: true,
                Default: 11,
                Description: "Maximum retries for a request",
            },
            "token": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "AWS Secret Token",
            },
            "region": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                DefaultFunc: schema.MultiEnvDefaultFunc([]string{
                    "AWS_REGION",
                    "AWS_DEFAULT_REGION",
                }, nil),
                Description: "AWS Region",
                InputDefault: "us-east-1",
            },
            "shared_credentials_file": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "Credentials File",
            },
            "insecure": &schema.Schema{
                Type: schema.TypeBool,
                Optional: true,
                Default: false,
                Description: "Allow the provider to perform insecure SSL requests",
            },
            "iam_endpoint": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                Default: "",
                Description: "Use this to override the default endpoint",
            },
        },

        ResourcesMap: map[string]*schema.Resource{
            "trility_aws_cognito_idp_user_pool": resourceTrilityAwsCognitoIDPUserPool(),
            "trility_aws_organizations_account": resourceTrilityAwsOrganizationsAccount(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
        SecretKey: d.Get("secret_key").(string),
        CredsFilename: d.Get("shared_credentials_file").(string),
        Profile: d.Get("profile").(string),
        Token: d.Get("token").(string),
        Region: d.Get("region").(string),
        MaxRetries: d.Get("max_retries").(int),
        Insecure: d.Get("insecure").(bool),
        IamEndpoint: d.Get("iam_endpoint").(string),
	}

	return config.Client()
}
