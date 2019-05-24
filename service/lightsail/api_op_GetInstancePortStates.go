// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstancePortStatesRequest
type GetInstancePortStatesInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstancePortStatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstancePortStatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstancePortStatesInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstancePortStatesResult
type GetInstancePortStatesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the port states resulting from your request.
	PortStates []InstancePortState `locationName:"portStates" type:"list"`
}

// String returns the string representation
func (s GetInstancePortStatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstancePortStates = "GetInstancePortStates"

// GetInstancePortStatesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the port states for a specific virtual private server, or instance.
//
//    // Example sending a request using GetInstancePortStatesRequest.
//    req := client.GetInstancePortStatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstancePortStates
func (c *Client) GetInstancePortStatesRequest(input *GetInstancePortStatesInput) GetInstancePortStatesRequest {
	op := &aws.Operation{
		Name:       opGetInstancePortStates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstancePortStatesInput{}
	}

	req := c.newRequest(op, input, &GetInstancePortStatesOutput{})
	return GetInstancePortStatesRequest{Request: req, Input: input, Copy: c.GetInstancePortStatesRequest}
}

// GetInstancePortStatesRequest is the request type for the
// GetInstancePortStates API operation.
type GetInstancePortStatesRequest struct {
	*aws.Request
	Input *GetInstancePortStatesInput
	Copy  func(*GetInstancePortStatesInput) GetInstancePortStatesRequest
}

// Send marshals and sends the GetInstancePortStates API request.
func (r GetInstancePortStatesRequest) Send(ctx context.Context) (*GetInstancePortStatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstancePortStatesResponse{
		GetInstancePortStatesOutput: r.Request.Data.(*GetInstancePortStatesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstancePortStatesResponse is the response type for the
// GetInstancePortStates API operation.
type GetInstancePortStatesResponse struct {
	*GetInstancePortStatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstancePortStates request.
func (r *GetInstancePortStatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}