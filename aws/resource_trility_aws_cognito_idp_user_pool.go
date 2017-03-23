package aws

import (
    "fmt"

    "github.com/hashicorp/terraform/helper/schema"

    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func resourceTrilityAwsCognitoIDPUserPool() *schema.Resource {
    return &schema.Resource{
        Create: resourceCognitoIDPUserPoolCreate,
        Read: resourceCognitoIDPUserPoolRead,
        Update: resourceCognitoIDPUserPoolUpdate,
        Delete: resourceCognitoIDPUserPoolDelete,

        Schema: map[string]*schema.Schema {
            "name": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCognitoIDPUserPoolCreate(d *schema.ResourceData, meta interface{}) error {
    cidpconn := meta.(*AWSClient).cidpconn
    name := d.Get("name").(string)

    params := &cognitoidentityprovider.CreateUserPoolInput{
        PoolName: aws.String(name),
    }

    // createResp, err := cidpconn.CreateUserPool(params)
    _, err := cidpconn.CreateUserPool(params)
    if err != nil {
        return fmt.Errorf("Error creating User Pool %s: %s", name, err)
    }
    // return resourceCognitoIdentityProviderRead(d, createResp.Group)
    return nil
}

func resourceCognitoIDPUserPoolRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceCognitoIDPUserPoolUpdate(d *schema.ResourceData, meta interface{}) error {
    return nil
}

func resourceCognitoIDPUserPoolDelete(d *schema.ResourceData, meta interface{}) error {
    return nil
}
