// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type EapAuthMethodResponse200 struct {
	// contains an EAP packet
	EapPayload string `json:"eapPayload" yaml:"eapPayload" bson:"eapPayload"`
	// URI : /{eapSessionUri}
	Links map[string]LinksValueSchema `json:"_links" yaml:"_links" bson:"_links"`
}