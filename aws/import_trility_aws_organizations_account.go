package aws

import (
    "github.com/hashicorp/terraform/helper/schema"

    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
)

func resourceTrilityAwsOrganizationsAccountImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
    orgconn := meta.(*AWSClient).orgconn
    id := d.Id()

    params := &organizations.DescribeAccountInput{
        AccountId: aws.String(id),
    }

    out, err := orgconn.DescribeAccount(params)
    if err != nil {
        return nil, err
    }

    account := out.Account
    results := make([]*schema.ResourceData, 1)

    d.Set("name", account.Name)
    d.Set("arn", account.Arn)
    d.Set("status", account.Status)

    results[0] = d
    return results, nil
}
