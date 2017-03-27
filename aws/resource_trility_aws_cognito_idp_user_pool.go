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
        // Update: resourceCognitoIDPUserPoolUpdate,
        Delete: resourceCognitoIDPUserPoolDelete,

        Schema: map[string]*schema.Schema {
            "poolname": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
        },
    }
}

func resourceCognitoIDPUserPoolCreate(d *schema.ResourceData, meta interface{}) error {
    cidpconn := meta.(*AWSClient).cidpconn
    poolname := d.Get("poolname").(string)

    params := &cognitoidentityprovider.CreateUserPoolInput{
        PoolName: aws.String(poolname),
    }

    resp, err := cidpconn.CreateUserPool(params)
    if err != nil {
        return fmt.Errorf("Error creating User Pool %s: %s", poolname, err)
    }

    d.SetId(*resp.UserPool.Id)
    return nil
}

func resourceCognitoIDPUserPoolRead(d *schema.ResourceData, meta interface{}) error {
    return nil
}

// func resourceCognitoIDPUserPoolUpdate(d *schema.ResourceData, meta interface{}) error {
//     return nil
// }

func resourceCognitoIDPUserPoolDelete(d *schema.ResourceData, meta interface{}) error {
    cidpconn := meta.(*AWSClient).cidpconn
    id := d.Id()

    params := &cognitoidentityprovider.DeleteUserPoolInput{
        UserPoolId: aws.String(id),
    }

    _, err := cidpconn.DeleteUserPool(params)
    if err != nil {
        return fmt.Errorf("Error removing user pool id %s: %s", id, err)
    }

    return nil
}
