package axios

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaClient is a client to use Axios on AWS Lambda
type LambdaClient struct {
	lambdaSvc    *lambda.Lambda
	functionName string
}

// Do sends an Axios Request to lambda function and returns an Axios response
func (s *LambdaClient) Do(req *RequestConfig) (*Response, error) {
	reqByte, err1 := json.Marshal(req)

	if err1 != nil {
		return nil, err1
	}

	invokeOutput, err2 := s.lambdaSvc.Invoke(&lambda.InvokeInput{
		FunctionName: &s.functionName,
		Payload:      reqByte,
	})

	if err2 != nil {
		return nil, err2
	}

	if invokeOutput.FunctionError != nil && *invokeOutput.FunctionError == "Handled" {
		var lambdaErrObject lambdaError
		errErr := json.Unmarshal(invokeOutput.Payload, &lambdaErrObject)
		if errErr != nil {
			return nil, errErr
		}
		return nil, fmt.Errorf("%+v", lambdaErrObject)
	}

	var resp Response
	err3 := json.Unmarshal(invokeOutput.Payload, &resp)

	if err3 != nil {
		return nil, err3
	}

	return &resp, nil
}

// NewLambdaClient creats a new LambdaClient
func NewLambdaClient(sess *session.Session, lambdaFunctionName string) *LambdaClient {
	return &LambdaClient{
		lambdaSvc:    lambda.New(sess),
		functionName: lambdaFunctionName,
	}
}
