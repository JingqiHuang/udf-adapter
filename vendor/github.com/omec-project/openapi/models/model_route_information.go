// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * 3gpp-traffic-influence
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RouteInformation struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr" mapstructure:"Ipv4Addr"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty" yaml:"ipv6Addr" bson:"ipv6Addr" mapstructure:"Ipv6Addr"`
	PortNumber int32  `json:"portNumber" yaml:"portNumber" bson:"portNumber" mapstructure:"PortNumber"`
}
