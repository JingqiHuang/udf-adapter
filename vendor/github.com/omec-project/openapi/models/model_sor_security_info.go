// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nausf_SoRProtection Service
 *
 * AUSF SoR Protection Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SorSecurityInfo struct {
	SorMacIausf string `json:"sorMacIausf" yaml:"sorMacIausf" bson:"sorMacIausf"`
	CounterSor  string `json:"counterSor" yaml:"counterSor" bson:"counterSor"`
	SorXmacIue  string `json:"sorXmacIue,omitempty" yaml:"sorXmacIue" bson:"sorXmacIue"`
}
