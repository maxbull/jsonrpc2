// Copyright 2018 Adam S Levy. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package jsonrpc2

import (
	"encoding/json"
)

// Request represents a JSON-RPC 2.0 Request or Notification object.  ID must
// be a numeric or string type. Params must be a structured type: slice, array,
// map or struct.
type Request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      interface{} `json:"id,omitempty"`
}

// NewRequest is a convenience function that returns a new Request with the
// "jsonrpc" field already populated with the required value, "2.0". If nil id
// is provided, it will be considered a Notification object and not receive a
// response. Use NewNotification if you want a simpler function call to form a
// JSON-RPC 2.0 Notification object.
func NewRequest(method string, id, params interface{}) Request {
	return Request{JSONRPC: "2.0", ID: id, Method: method, Params: params}
}

// IsValid returns true if r has a valid JSONRPC value of "2.0" and a non-empty
// Method. Params and ID are not validated
func (r Request) IsValid() bool {
	return r.JSONRPC == "2.0" && len(r.Method) > 0
}

// String returns a JSON string with "--> " prefixed to represent a Request
// object.
func (r Request) String() string {
	b, _ := json.Marshal(r)
	return "--> " + string(b)
}

// BatchRequest is a type that implements String() for a slice of Requests.
type BatchRequest []Request

// String returns a string of the JSON array with "--> " prefixed to represent
// a BatchRequest object.
func (br BatchRequest) String() string {
	s := "--> [\n"
	for i, res := range br {
		s += "  " + res.String()[4:]
		if i < len(br)-1 {
			s += ","
		}
		s += "\n"
	}
	return s + "]"
}