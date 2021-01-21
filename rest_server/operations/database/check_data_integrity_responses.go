// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CheckDataIntegrityAcceptedCode is the HTTP code returned for type CheckDataIntegrityAccepted
const CheckDataIntegrityAcceptedCode int = 202

/*CheckDataIntegrityAccepted Base empty response

swagger:response checkDataIntegrityAccepted
*/
type CheckDataIntegrityAccepted struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewCheckDataIntegrityAccepted creates CheckDataIntegrityAccepted with default headers values
func NewCheckDataIntegrityAccepted() *CheckDataIntegrityAccepted {

	return &CheckDataIntegrityAccepted{}
}

// WithPayload adds the payload to the check data integrity accepted response
func (o *CheckDataIntegrityAccepted) WithPayload(payload *rest_model.Empty) *CheckDataIntegrityAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check data integrity accepted response
func (o *CheckDataIntegrityAccepted) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDataIntegrityAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDataIntegrityUnauthorizedCode is the HTTP code returned for type CheckDataIntegrityUnauthorized
const CheckDataIntegrityUnauthorizedCode int = 401

/*CheckDataIntegrityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response checkDataIntegrityUnauthorized
*/
type CheckDataIntegrityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCheckDataIntegrityUnauthorized creates CheckDataIntegrityUnauthorized with default headers values
func NewCheckDataIntegrityUnauthorized() *CheckDataIntegrityUnauthorized {

	return &CheckDataIntegrityUnauthorized{}
}

// WithPayload adds the payload to the check data integrity unauthorized response
func (o *CheckDataIntegrityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CheckDataIntegrityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check data integrity unauthorized response
func (o *CheckDataIntegrityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDataIntegrityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDataIntegrityTooManyRequestsCode is the HTTP code returned for type CheckDataIntegrityTooManyRequests
const CheckDataIntegrityTooManyRequestsCode int = 429

/*CheckDataIntegrityTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response checkDataIntegrityTooManyRequests
*/
type CheckDataIntegrityTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCheckDataIntegrityTooManyRequests creates CheckDataIntegrityTooManyRequests with default headers values
func NewCheckDataIntegrityTooManyRequests() *CheckDataIntegrityTooManyRequests {

	return &CheckDataIntegrityTooManyRequests{}
}

// WithPayload adds the payload to the check data integrity too many requests response
func (o *CheckDataIntegrityTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CheckDataIntegrityTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check data integrity too many requests response
func (o *CheckDataIntegrityTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDataIntegrityTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
