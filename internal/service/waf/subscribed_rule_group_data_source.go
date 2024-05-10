// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	awstypes "github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tfslices "github.com/hashicorp/terraform-provider-aws/internal/slices"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_waf_subscribed_rule_group", name="Subscribed Rule Group")
func dataSourceSubscribedRuleGroup() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceSubscribedRuleGroupRead,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{names.AttrName, "metric_name"},
			},
			names.AttrName: {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{names.AttrName, "metric_name"},
			},
		},
	}
}

func dataSourceSubscribedRuleGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFClient(ctx)

	var filter tfslices.Predicate[*awstypes.SubscribedRuleGroupSummary]

	if v, ok := d.GetOk("metric_name"); ok {
		name := v.(string)
		filter = func(v *awstypes.SubscribedRuleGroupSummary) bool {
			return aws.ToString(v.MetricName) == name
		}
	}

	if v, ok := d.GetOk(names.AttrName); ok {
		name := v.(string)
		f := func(v *awstypes.SubscribedRuleGroupSummary) bool {
			return aws.ToString(v.Name) == name
		}

		if filter != nil {
			filter = tfslices.PredicateAnd(filter, f)
		} else {
			filter = f
		}
	}

	input := &waf.ListSubscribedRuleGroupsInput{}
	output, err := findSubscribedRuleGroup(ctx, conn, input, filter)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("WAF Rate Based Rule", err))
	}

	d.SetId(aws.ToString(output.RuleGroupId))
	d.Set("metric_name", output.MetricName)
	d.Set(names.AttrName, output.Name)

	return diags
}

func findSubscribedRuleGroup(ctx context.Context, conn *waf.Client, input *waf.ListSubscribedRuleGroupsInput, filter tfslices.Predicate[*awstypes.SubscribedRuleGroupSummary]) (*awstypes.SubscribedRuleGroupSummary, error) {
	output, err := findSubscribedRuleGroups(ctx, conn, input, filter)

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output)
}

func findSubscribedRuleGroups(ctx context.Context, conn *waf.Client, input *waf.ListSubscribedRuleGroupsInput, filter tfslices.Predicate[*awstypes.SubscribedRuleGroupSummary]) ([]awstypes.SubscribedRuleGroupSummary, error) {
	var output []awstypes.SubscribedRuleGroupSummary

	err := listSubscribedRuleGroupsPages(ctx, conn, input, func(page *waf.ListSubscribedRuleGroupsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, v := range page.RuleGroups {
			if filter(&v) {
				output = append(output, v)
			}
		}

		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
