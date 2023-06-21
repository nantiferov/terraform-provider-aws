// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package signer

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	signer_sdkv1 "github.com/aws/aws-sdk-go/service/signer"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  DataSourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  ResourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
			Name:     "Signing Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSigningProfilePermission,
			TypeName: "aws_signer_signing_profile_permission",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Signer
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*signer_sdkv1.Signer, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return signer_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
