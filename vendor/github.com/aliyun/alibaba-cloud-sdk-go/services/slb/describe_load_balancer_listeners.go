package slb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeLoadBalancerListeners invokes the slb.DescribeLoadBalancerListeners API synchronously
func (client *Client) DescribeLoadBalancerListeners(request *DescribeLoadBalancerListenersRequest) (response *DescribeLoadBalancerListenersResponse, err error) {
	response = CreateDescribeLoadBalancerListenersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerListenersWithChan invokes the slb.DescribeLoadBalancerListeners API asynchronously
func (client *Client) DescribeLoadBalancerListenersWithChan(request *DescribeLoadBalancerListenersRequest) (<-chan *DescribeLoadBalancerListenersResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerListenersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerListeners(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeLoadBalancerListenersWithCallback invokes the slb.DescribeLoadBalancerListeners API asynchronously
func (client *Client) DescribeLoadBalancerListenersWithCallback(request *DescribeLoadBalancerListenersRequest, callback func(response *DescribeLoadBalancerListenersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerListenersResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerListeners(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeLoadBalancerListenersRequest is the request struct for api DescribeLoadBalancerListeners
type DescribeLoadBalancerListenersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ListenerProtocol     string           `position:"Query" name:"ListenerProtocol"`
	LoadBalancerId       *[]string        `position:"Query" name:"LoadBalancerId"  type:"Repeated"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeLoadBalancerListenersResponse is the response struct for api DescribeLoadBalancerListeners
type DescribeLoadBalancerListenersResponse struct {
	*responses.BaseResponse
	MaxResults int                                       `json:"MaxResults" xml:"MaxResults"`
	NextToken  string                                    `json:"NextToken" xml:"NextToken"`
	RequestId  string                                    `json:"RequestId" xml:"RequestId"`
	TotalCount int                                       `json:"TotalCount" xml:"TotalCount"`
	Listeners  []ListenerInDescribeLoadBalancerListeners `json:"Listeners" xml:"Listeners"`
}

// CreateDescribeLoadBalancerListenersRequest creates a request to invoke DescribeLoadBalancerListeners API
func CreateDescribeLoadBalancerListenersRequest() (request *DescribeLoadBalancerListenersRequest) {
	request = &DescribeLoadBalancerListenersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerListeners", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerListenersResponse creates a response to parse from DescribeLoadBalancerListeners response
func CreateDescribeLoadBalancerListenersResponse() (response *DescribeLoadBalancerListenersResponse) {
	response = &DescribeLoadBalancerListenersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
