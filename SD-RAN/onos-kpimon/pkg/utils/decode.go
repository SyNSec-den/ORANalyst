// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

// DecodePlmnIDToUint32 decodes PLMN ID from byte array to uint32
func DecodePlmnIDToUint32(plmnBytes []byte) uint32 {
	return uint32(plmnBytes[0]) | uint32(plmnBytes[1])<<8 | uint32(plmnBytes[2])<<16
}
