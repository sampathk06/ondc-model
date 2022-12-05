package openapi

import (
	// "encoding/json"
	"time"
)

// Context Describes a ONDC message context
type Context struct {
	// Codification of domains supported by ONDC
	Domain string `json:"domain"`
	Country string `json:"country"`
	City string `json:"city"`
	// Defines the ONDC API call. Any actions other than the enumerated actions are not supported by ONDC Protocol
	Action string `json:"action"`
	// Version of ONDC core API specification being used
	CoreVersion string `json:"core_version"`
	// Unique id of the Buyer App. By default it is the fully qualified domain name of the Buyer App
	BapId string `json:"bap_id"`
	// URI of the Seller App for accepting callbacks. Must have the same domain name as the bap_id
	BapUri string `json:"bap_uri"`
	// Unique id of the Seller App. By default it is the fully qualified domain name of the Seller App
	BppId *string `json:"bpp_id,omitempty"`
	// URI of the Seller App. Must have the same domain name as the bap_id
	BppUri *string `json:"bpp_uri,omitempty"`
	// This is a unique value which persists across all API calls from search through confirm
	TransactionId string `json:"transaction_id"`
	// This is a unique value which persists during a request / callback cycle
	MessageId string `json:"message_id"`
	// Time of request generation in RFC3339 format
	Timestamp time.Time `json:"timestamp"`
	// The encryption public key of the sender
	Key *string `json:"key,omitempty"`
	// The duration in ISO8601 format after timestamp for which this message holds valid.
	Ttl *string `json:"ttl,omitempty"`
}

func NewContext(domain string, country string, city string, action string, coreVersion string, bapId string, bapUri string, transactionId string, messageId string, timestamp time.Time) *Context {
	this := Context{}
	this.Domain = domain
	this.Country = country
	this.City = city
	this.Action = action
	this.CoreVersion = coreVersion
	this.BapId = bapId
	this.BapUri = bapUri
	this.TransactionId = transactionId
	this.MessageId = messageId
	this.Timestamp = timestamp
	return &this
}