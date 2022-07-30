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

type UeAuthenticationCtx struct {
	AuthType           AuthType                    `json:"authType" yaml:"authType" bson:"authType"`
	Var5gAuthData      interface{}                 `json:"5gAuthData" yaml:"5gAuthData" bson:"5gAuthData"`
	Links              map[string]LinksValueSchema `json:"_links" yaml:"_links" bson:"_links"`
	ServingNetworkName string                      `json:"servingNetworkName,omitempty" yaml:"servingNetworkName" bson:"servingNetworkName"`
}
