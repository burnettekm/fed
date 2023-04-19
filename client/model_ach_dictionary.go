/*
 * FED API
 *
 * FED API is designed to create FEDACH and FEDWIRE dictionaries.  The FEDACH dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive ACH entries.  The FEDWIRE dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive WIRE entries.  This project implements a modern REST HTTP API for FEDACH Dictionary and FEDWIRE Dictionary.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// AchDictionary Search results containing ACHDictionary of Participants
type AchDictionary struct {
	ACHParticipants []AchParticipant `json:"ACHParticipants,omitempty"`
}
