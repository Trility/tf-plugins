package aws

import (
    "fmt"

    "github.com/hashicorp/terraform/helper/schema"

    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
)

func resourceTrilityAwsOrganizationsAccount() *schema.Resource {
    return &schema.Resource{
        Create: resourceOrganizationsAccountCreate,
        Read: resourceOrganizationsAccountRead,
        Delete: resourceOrganizationsAccountRemove,
        Importer: &schema.ResourceImporter{
            State: resourceTrilityAwsOrganizationsAccountImport,
        },

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "email": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "role_name": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
        },
    }
}

func resourceOrganizationsAccountCreate(d *schema.ResourceData, meta interface{}) error {
    orgconn := meta.(*AWSClient).orgconn
    name := d.Get("name").(string)
    email := d.Get("email").(string)
    role_name := d.Get("role_name").(string)

    params := &organizations.CreateAccountInput{
        AccountName: aws.String(name),
        Email: aws.String(email),
        RoleName: aws.String(role_name),
    }

    out, err := orgconn.CreateAccount(params)
    if err != nil {
        return err
    }

    d.SetId(*out.CreateAccountStatus.AccountId)
    return nil
}

func resourceOrganizationsAccountRead(d *schema.ResourceData, meta interface{}) error {
    orgconn := meta.(*AWSClient).orgconn
    name := d.Get("name").(string)
    id := d.Id()

    params := &organizations.DescribeAccountInput{
        AccountId: aws.String(id),
    }

    out, err := orgconn.DescribeAccount(params)
    if err != nil {
        return fmt.Errorf("Error reading account %s (%s): %s", name, id, err)
    }

    d.Set("arn", out.Account.Arn)
    d.Set("status", out.Account.Status)
    return nil
}

func resourceOrganizationsAccountRemove(d *schema.ResourceData, meta interface{}) error {
    orgconn := meta.(*AWSClient).orgconn
    name := d.Get("name").(string)
    id := d.Id()

    params := &organizations.RemoveAccountFromOrganizationInput{
        AccountId: aws.String(id),
    }

    _, err := orgconn.RemoveAccountFromOrganization(params)
    if err != nil {
        return fmt.Errorf("Error removing account %s (%s) from the organization: %s", name, id, err)
    }

    return nil
}
