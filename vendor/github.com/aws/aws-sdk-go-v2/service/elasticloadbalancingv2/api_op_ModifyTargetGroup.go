// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/ModifyTargetGroupInput
type ModifyTargetGroupInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether health checks are enabled.
	HealthCheckEnabled *bool `type:"boolean"`

	// The approximate amount of time, in seconds, between health checks of an individual
	// target. For Application Load Balancers, the range is 5–300 seconds. For
	// Network Load Balancers, the supported values are 10 or 30 seconds.
	//
	// If the protocol of the target group is TCP, you can't modify this setting.
	HealthCheckIntervalSeconds *int64 `min:"5" type:"integer"`

	// [HTTP/HTTPS health checks] The ping path that is the destination for the
	// health check request.
	HealthCheckPath *string `min:"1" type:"string"`

	// The port the load balancer uses when performing health checks on targets.
	HealthCheckPort *string `type:"string"`

	// The protocol the load balancer uses when performing health checks on targets.
	// The TCP protocol is supported for health checks only if the protocol of the
	// target group is TCP or TLS. The TLS protocol is not supported for health
	// checks.
	//
	// If the protocol of the target group is TCP, you can't modify this setting.
	HealthCheckProtocol ProtocolEnum `type:"string" enum:"true"`

	// [HTTP/HTTPS health checks] The amount of time, in seconds, during which no
	// response means a failed health check.
	//
	// If the protocol of the target group is TCP, you can't modify this setting.
	HealthCheckTimeoutSeconds *int64 `min:"2" type:"integer"`

	// The number of consecutive health checks successes required before considering
	// an unhealthy target healthy.
	HealthyThresholdCount *int64 `min:"2" type:"integer"`

	// [HTTP/HTTPS health checks] The HTTP codes to use when checking for a successful
	// response from a target.
	//
	// If the protocol of the target group is TCP, you can't modify this setting.
	Matcher *Matcher `type:"structure"`

	// The Amazon Resource Name (ARN) of the target group.
	//
	// TargetGroupArn is a required field
	TargetGroupArn *string `type:"string" required:"true"`

	// The number of consecutive health check failures required before considering
	// the target unhealthy. For Network Load Balancers, this value must be the
	// same as the healthy threshold count.
	UnhealthyThresholdCount *int64 `min:"2" type:"integer"`
}

// String returns the string representation
func (s ModifyTargetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTargetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyTargetGroupInput"}
	if s.HealthCheckIntervalSeconds != nil && *s.HealthCheckIntervalSeconds < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckIntervalSeconds", 5))
	}
	if s.HealthCheckPath != nil && len(*s.HealthCheckPath) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HealthCheckPath", 1))
	}
	if s.HealthCheckTimeoutSeconds != nil && *s.HealthCheckTimeoutSeconds < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckTimeoutSeconds", 2))
	}
	if s.HealthyThresholdCount != nil && *s.HealthyThresholdCount < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthyThresholdCount", 2))
	}

	if s.TargetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupArn"))
	}
	if s.UnhealthyThresholdCount != nil && *s.UnhealthyThresholdCount < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("UnhealthyThresholdCount", 2))
	}
	if s.Matcher != nil {
		if err := s.Matcher.Validate(); err != nil {
			invalidParams.AddNested("Matcher", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/ModifyTargetGroupOutput
type ModifyTargetGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the modified target group.
	TargetGroups []TargetGroup `type:"list"`
}

// String returns the string representation
func (s ModifyTargetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyTargetGroup = "ModifyTargetGroup"

// ModifyTargetGroupRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Modifies the health checks used when evaluating the health state of the targets
// in the specified target group.
//
// To monitor the health of the targets, use DescribeTargetHealth.
//
//    // Example sending a request using ModifyTargetGroupRequest.
//    req := client.ModifyTargetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/ModifyTargetGroup
func (c *Client) ModifyTargetGroupRequest(input *ModifyTargetGroupInput) ModifyTargetGroupRequest {
	op := &aws.Operation{
		Name:       opModifyTargetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTargetGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyTargetGroupOutput{})
	return ModifyTargetGroupRequest{Request: req, Input: input, Copy: c.ModifyTargetGroupRequest}
}

// ModifyTargetGroupRequest is the request type for the
// ModifyTargetGroup API operation.
type ModifyTargetGroupRequest struct {
	*aws.Request
	Input *ModifyTargetGroupInput
	Copy  func(*ModifyTargetGroupInput) ModifyTargetGroupRequest
}

// Send marshals and sends the ModifyTargetGroup API request.
func (r ModifyTargetGroupRequest) Send(ctx context.Context) (*ModifyTargetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTargetGroupResponse{
		ModifyTargetGroupOutput: r.Request.Data.(*ModifyTargetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTargetGroupResponse is the response type for the
// ModifyTargetGroup API operation.
type ModifyTargetGroupResponse struct {
	*ModifyTargetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTargetGroup request.
func (r *ModifyTargetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
