// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LocationQoS struct {
	HAccuracy         float32      `json:"hAccuracy,omitempty" yaml:"hAccuracy" bson:"hAccuracy"`
	VAccuracy         float32      `json:"vAccuracy,omitempty" yaml:"vAccuracy" bson:"vAccuracy"`
	VerticalRequested bool         `json:"verticalRequested,omitempty" yaml:"verticalRequested" bson:"verticalRequested"`
	ResponseTime      ResponseTime `json:"responseTime,omitempty" yaml:"responseTime" bson:"responseTime"`
}
