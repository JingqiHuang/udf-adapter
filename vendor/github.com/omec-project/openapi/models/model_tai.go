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

type Tai struct {
	PlmnId *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	Tac    string  `json:"tac" yaml:"tac" bson:"tac" mapstructure:"Tac"`
}
