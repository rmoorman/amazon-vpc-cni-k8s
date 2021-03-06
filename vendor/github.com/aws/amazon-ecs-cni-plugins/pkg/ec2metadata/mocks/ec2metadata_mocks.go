// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cni-plugins/pkg/ec2metadata (interfaces: EC2Metadata)

package mock_ec2metadata

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of EC2Metadata interface
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *_MockEC2MetadataRecorder
}

// Recorder for MockEC2Metadata (not exported)
type _MockEC2MetadataRecorder struct {
	mock *MockEC2Metadata
}

func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &_MockEC2MetadataRecorder{mock}
	return mock
}

func (_m *MockEC2Metadata) EXPECT() *_MockEC2MetadataRecorder {
	return _m.recorder
}

func (_m *MockEC2Metadata) GetMetadata(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetadata", arg0)
}
