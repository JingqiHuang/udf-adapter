// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type DnnInfo struct {
	Dnn                 string `json:"dnn" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	DefaultDnnIndicator bool   `json:"defaultDnnIndicator,omitempty" yaml:"defaultDnnIndicator" bson:"defaultDnnIndicator" mapstructure:"DefaultDnnIndicator"`
	LboRoamingAllowed   bool   `json:"lboRoamingAllowed,omitempty" yaml:"lboRoamingAllowed" bson:"lboRoamingAllowed" mapstructure:"LboRoamingAllowed"`
	IwkEpsInd           bool   `json:"iwkEpsInd,omitempty" yaml:"iwkEpsInd" bson:"iwkEpsInd" mapstructure:"IwkEpsInd"`
}
