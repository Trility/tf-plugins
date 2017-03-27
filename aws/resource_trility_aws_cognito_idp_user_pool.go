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
            "poolname": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
            },
            "policies": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "passwordpolicy": &schema.Schema{
                            Type: schema.TypeSet,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "minimumlength": &schema.Schema{
                                        Type: schema.TypeInt,
                                        Required: true,
                                    },
                                    "requireuppercase": &schema.Schema{
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "requirelowercase": &schema.Schema{
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "requirenumbers": &schema.Schema{
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                    "requiresymbols": &schema.Schema{
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "lambdaconfig": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "presignup": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "custommessage": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "postconfirmation": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "preauthentication": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "postauthentication": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "defineauthchallenge": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "createauthchallenge": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "verifyauthchallengeresponse": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },
            "autoverifiedattributes": &schema.Schema{
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },
            "aliasattributes": &schema.Schema{
                Type: schema.TypeList,
                Optional: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },
            "smsverificationmessage": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
            },
            "emailverifcationmessage": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
            },
            "emailverificationsubject": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
            },
            "smsauthenticationmessage": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
            },
            "mfaconfiguration": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
            },
            "deviceconfiguration": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "challengerequiredonnewdevice": &schema.Schema{
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                        "deviceonlyrememberedonuserprompt": &schema.Schema{
                            Type: schema.TypeBool,
                            Optional: true,
                        },
                    },
                },
            },
            "emailconfiguration": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "sourcearn": &schema.Schema{
                            Type: schema.TypeString,
                            Required: true,
                        },
                        "replytoemailaddress": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },
            "smsconfiguration": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "snscallerarn": &schema.Schema{
                            Type: schema.TypeString,
                            Required: true,
                        },
                        "externalid": &schema.Schema{
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },
            "userpooltags": &schema.Schema{
                Type: schema.TypeMap,
                Optional: true,
            },
            "admincreateuserconfig": &schema.Schema{
                Type: schema.TypeSet,
                Optional: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "allowadmincreateuseronly": &schema.Schema{
                            Type: schema.TypeBool,
                            Required: true,
                        },
                        "unusedaccountvaliditydays": &schema.Schema{
                            Type: schema.TypeInt,
                            Required: true,
                        },
                        "invitemessagetemplate": &schema.Schema{
                            Type: schema.TypeSet,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "smsmessage": &schema.Schema{
                                        Type: schema.TypeString,
                                        Required: true,
                                    },
                                    "emailmessage": &schema.Schema{
                                        Type: schema.TypeString,
                                        Required: true,
                                    },
                                    "emailsubject": &schema.Schema{
                                        Type: schema.TypeString,
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

    // createResp, err := cidpconn.CreateUserPool(params)
    _, err := cidpconn.CreateUserPool(params)
    if err != nil {
        return fmt.Errorf("Error creating User Pool %s: %s", poolname, err)
    }
    // return resourceCognitoIdentityProviderRead(d, createResp.Group)
    //d.SetId(createResp.UserPoolId)
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
