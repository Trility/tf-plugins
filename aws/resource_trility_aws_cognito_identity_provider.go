package aws

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceTrilityAwsCognitoIdentityProvider() *schema.Resource {
    return &schema.Resource{
        Create: resourceCognitoIdentityProviderCreate,
        Read: resourceCognitoIdentityProviderRead,
        Update: resourceCognitoIdentityProviderUpdate,
        Delete: resourceCognitoIdentityProviderDelete,

        Schema: map[string]*schema.Schema {
            "pool_name": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCognitoIdentityProviderCreate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceCognitoIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceCognitoIdentityProviderUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceCognitoIdentityProviderDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}
