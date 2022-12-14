// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SubscribedDefaultQos struct {
	Var5qi        int32 `json:"5qi" yaml:"5qi" bson:"5qi" mapstructure:"Var5qi"`
	Arp           *Arp  `json:"arp" yaml:"arp" bson:"arp" mapstructure:"Arp"`
	PriorityLevel int32 `json:"priorityLevel,omitempty" yaml:"priorityLevel" bson:"priorityLevel" mapstructure:"PriorityLevel"`
}
