// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PpData struct {
	CommunicationCharacteristics *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty" yaml:"communicationCharacteristics" bson:"communicationCharacteristics" mapstructure:"CommunicationCharacteristics"`
	SupportedFeatures            string                        `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}