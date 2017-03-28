package aws

import (
    "bytes"
    "fmt"

    "github.com/hashicorp/terraform/helper/hashcode"
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

        Schema: map[string]*schema.Schema{
            "poolname": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "policies": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Set: policiesHash,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "password_policy": {
                            Type: schema.TypeSet,
                            Required: true,
                            Set: passwordPolicyHash,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "minimum_length": {
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "require_uppercase": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "require_lowercase": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "require_numbers": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "require_symbols": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                },
                            },
                        },
                    },
                },
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

    if _, ok := d.GetOk("policies"); ok {
        params.Policies = expandPolicies(d.Get("policies").(*schema.Set).List()[0].(map[string]interface{}))
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

func resourceCognitoIDPUserPoolUpdate(d *schema.ResourceData, meta interface{}) error {
    cidpconn := meta.(*AWSClient).cidpconn
    id := d.Id()

    params := &cognitoidentityprovider.UpdateUserPoolInput{
        UserPoolId: aws.String(id),
    }

    if _, ok := d.GetOk("policies"); ok {
        params.Policies = expandPolicies(d.Get("policies").(*schema.Set).List()[0].(map[string]interface{}))
    }

    _, err := cidpconn.UpdateUserPool(params)
    if err != nil {
        return fmt.Errorf("Error updating User Pool %s: %s", id, err)
    }

    return nil
}

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

// Used https://github.com/hashicorp/terraform/blob/master/builtin/providers/aws/cloudfront_distribution_configuration_structure.go
// for examples on using complicated structures in the resource definition
func expandPolicies(m map[string]interface{}) *cognitoidentityprovider.UserPoolPolicyType {
    pt := &cognitoidentityprovider.UserPoolPolicyType{
        PasswordPolicy: expandPasswordPolicy(m["password_policy"].(*schema.Set).List()[0].(map[string]interface{})),
    }
    return pt
}

func expandPasswordPolicy(m map[string]interface{}) *cognitoidentityprovider.PasswordPolicyType{
    ppt := &cognitoidentityprovider.PasswordPolicyType{
        MinimumLength: aws.Int64(int64(m["minimum_length"].(int))),
        RequireLowercase: aws.Bool(m["require_lowercase"].(bool)),
        RequireUppercase: aws.Bool(m["require_uppercase"].(bool)),
        RequireNumbers: aws.Bool(m["require_numbers"].(bool)),
        RequireSymbols: aws.Bool(m["require_symbols"].(bool)),
    }
    return ppt
}

// TypeSet Attribute
func policiesHash(v interface{}) int {
    var buf bytes.Buffer
    m := v.(map[string]interface{})
    buf.WriteString(fmt.Sprintf("%t-", passwordPolicyHash(m["password_policy"].(*schema.Set).List()[0].(map[string]interface{}))))
    return hashcode.String(buf.String())
}

// TypeSet Attribute
func passwordPolicyHash(v interface{}) int {
    var buf bytes.Buffer
    m := v.(map[string]interface{})
    buf.WriteString(fmt.Sprintf("%t-", m["minimum_length"].(int)))
    buf.WriteString(fmt.Sprintf("%t-", m["require_uppercase"].(bool)))
    buf.WriteString(fmt.Sprintf("%t-", m["require_lowercase"].(bool)))
    buf.WriteString(fmt.Sprintf("%t-", m["require_numbers"].(bool)))
    buf.WriteString(fmt.Sprintf("%t-", m["require_symbols"].(bool)))
    return hashcode.String(buf.String())
}
