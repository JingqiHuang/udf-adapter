// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UncertaintyEllipse struct {
	SemiMajor        float32 `json:"semiMajor" yaml:"semiMajor" bson:"semiMajor"`
	SemiMinor        float32 `json:"semiMinor" yaml:"semiMinor" bson:"semiMinor"`
	OrientationMajor int32   `json:"orientationMajor" yaml:"orientationMajor" bson:"orientationMajor"`
}
