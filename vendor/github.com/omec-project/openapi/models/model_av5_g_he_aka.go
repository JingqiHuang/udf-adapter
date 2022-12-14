// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Av5GHeAka struct {
	AvType   AvType `json:"avType" yaml:"avType" bson:"avType" mapstructure:"AvType"`
	Rand     string `json:"rand" yaml:"rand" bson:"rand" mapstructure:"Rand"`
	XresStar string `json:"xresStar" yaml:"xresStar" bson:"xresStar" mapstructure:"XresStar"`
	Autn     string `json:"autn" yaml:"autn" bson:"autn" mapstructure:"Autn"`
	Kausf    string `json:"kausf" yaml:"kausf" bson:"kausf" mapstructure:"Kausf"`
}
