// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfCond struct {
	AmfSetId    string `json:"amfSetId,omitempty" yaml:"amfSetId" bson:"amfSetId" mapstructure:"AmfSetId"`
	AmfRegionId string `json:"amfRegionId,omitempty" yaml:"amfRegionId" bson:"amfRegionId" mapstructure:"AmfRegionId"`
}