package dns

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// HTTPStatusCode enumerates the values for http status code.
type HTTPStatusCode string

const (
	// Accepted ...
	Accepted HTTPStatusCode = "Accepted"
	// Ambiguous ...
	Ambiguous HTTPStatusCode = "Ambiguous"
	// BadGateway ...
	BadGateway HTTPStatusCode = "BadGateway"
	// BadRequest ...
	BadRequest HTTPStatusCode = "BadRequest"
	// Conflict ...
	Conflict HTTPStatusCode = "Conflict"
	// Continue ...
	Continue HTTPStatusCode = "Continue"
	// Created ...
	Created HTTPStatusCode = "Created"
	// ExpectationFailed ...
	ExpectationFailed HTTPStatusCode = "ExpectationFailed"
	// Forbidden ...
	Forbidden HTTPStatusCode = "Forbidden"
	// Found ...
	Found HTTPStatusCode = "Found"
	// GatewayTimeout ...
	GatewayTimeout HTTPStatusCode = "GatewayTimeout"
	// Gone ...
	Gone HTTPStatusCode = "Gone"
	// HTTPVersionNotSupported ...
	HTTPVersionNotSupported HTTPStatusCode = "HttpVersionNotSupported"
	// InternalServerError ...
	InternalServerError HTTPStatusCode = "InternalServerError"
	// LengthRequired ...
	LengthRequired HTTPStatusCode = "LengthRequired"
	// MethodNotAllowed ...
	MethodNotAllowed HTTPStatusCode = "MethodNotAllowed"
	// Moved ...
	Moved HTTPStatusCode = "Moved"
	// MovedPermanently ...
	MovedPermanently HTTPStatusCode = "MovedPermanently"
	// MultipleChoices ...
	MultipleChoices HTTPStatusCode = "MultipleChoices"
	// NoContent ...
	NoContent HTTPStatusCode = "NoContent"
	// NonAuthoritativeInformation ...
	NonAuthoritativeInformation HTTPStatusCode = "NonAuthoritativeInformation"
	// NotAcceptable ...
	NotAcceptable HTTPStatusCode = "NotAcceptable"
	// NotFound ...
	NotFound HTTPStatusCode = "NotFound"
	// NotImplemented ...
	NotImplemented HTTPStatusCode = "NotImplemented"
	// NotModified ...
	NotModified HTTPStatusCode = "NotModified"
	// OK ...
	OK HTTPStatusCode = "OK"
	// PartialContent ...
	PartialContent HTTPStatusCode = "PartialContent"
	// PaymentRequired ...
	PaymentRequired HTTPStatusCode = "PaymentRequired"
	// PreconditionFailed ...
	PreconditionFailed HTTPStatusCode = "PreconditionFailed"
	// ProxyAuthenticationRequired ...
	ProxyAuthenticationRequired HTTPStatusCode = "ProxyAuthenticationRequired"
	// Redirect ...
	Redirect HTTPStatusCode = "Redirect"
	// RedirectKeepVerb ...
	RedirectKeepVerb HTTPStatusCode = "RedirectKeepVerb"
	// RedirectMethod ...
	RedirectMethod HTTPStatusCode = "RedirectMethod"
	// RequestedRangeNotSatisfiable ...
	RequestedRangeNotSatisfiable HTTPStatusCode = "RequestedRangeNotSatisfiable"
	// RequestEntityTooLarge ...
	RequestEntityTooLarge HTTPStatusCode = "RequestEntityTooLarge"
	// RequestTimeout ...
	RequestTimeout HTTPStatusCode = "RequestTimeout"
	// RequestURITooLong ...
	RequestURITooLong HTTPStatusCode = "RequestUriTooLong"
	// ResetContent ...
	ResetContent HTTPStatusCode = "ResetContent"
	// SeeOther ...
	SeeOther HTTPStatusCode = "SeeOther"
	// ServiceUnavailable ...
	ServiceUnavailable HTTPStatusCode = "ServiceUnavailable"
	// SwitchingProtocols ...
	SwitchingProtocols HTTPStatusCode = "SwitchingProtocols"
	// TemporaryRedirect ...
	TemporaryRedirect HTTPStatusCode = "TemporaryRedirect"
	// Unauthorized ...
	Unauthorized HTTPStatusCode = "Unauthorized"
	// UnsupportedMediaType ...
	UnsupportedMediaType HTTPStatusCode = "UnsupportedMediaType"
	// Unused ...
	Unused HTTPStatusCode = "Unused"
	// UpgradeRequired ...
	UpgradeRequired HTTPStatusCode = "UpgradeRequired"
	// UseProxy ...
	UseProxy HTTPStatusCode = "UseProxy"
)

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// Failed ...
	Failed OperationStatus = "Failed"
	// InProgress ...
	InProgress OperationStatus = "InProgress"
	// Succeeded ...
	Succeeded OperationStatus = "Succeeded"
)

// RecordType enumerates the values for record type.
type RecordType string

const (
	// A ...
	A RecordType = "A"
	// AAAA ...
	AAAA RecordType = "AAAA"
	// CNAME ...
	CNAME RecordType = "CNAME"
	// MX ...
	MX RecordType = "MX"
	// NS ...
	NS RecordType = "NS"
	// PTR ...
	PTR RecordType = "PTR"
	// SOA ...
	SOA RecordType = "SOA"
	// SRV ...
	SRV RecordType = "SRV"
	// TXT ...
	TXT RecordType = "TXT"
)

// AaaaRecord an AAAA record.
type AaaaRecord struct {
	// Ipv6Address - The IPv6 address of this AAAA record.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// ARecord an A record.
type ARecord struct {
	// Ipv4Address - The IPv4 address of this A record.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// CloudError ...
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody ...
type CloudErrorBody struct {
	Code    *string           `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
	Target  *string           `json:"target,omitempty"`
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// CnameRecord a CNAME record.
type CnameRecord struct {
	// Cname - The canonical name for this CNAME record.
	Cname *string `json:"cname,omitempty"`
}

// MxRecord an MX record.
type MxRecord struct {
	// Preference - The preference value for this MX record.
	Preference *int32 `json:"preference,omitempty"`
	// Exchange - The domain name of the mail host for this MX record.
	Exchange *string `json:"exchange,omitempty"`
}

// NsRecord an NS record.
type NsRecord struct {
	// Nsdname - The name server name for this NS record.
	Nsdname *string `json:"nsdname,omitempty"`
}

// PtrRecord a PTR record.
type PtrRecord struct {
	// Ptrdname - The PTR target domain name for this PTR record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// RecordSet describes a DNS record set (a collection of DNS records with the same name and type).
type RecordSet struct {
	autorest.Response `json:"-"`
	// ID - The ID of the record set.
	ID *string `json:"id,omitempty"`
	// Name - The name of the record set.
	Name *string `json:"name,omitempty"`
	// Type - The type of the record set.
	Type *string `json:"type,omitempty"`
	// Etag - The etag of the record set.
	Etag *string `json:"etag,omitempty"`
	// RecordSetProperties - The properties of the record set.
	*RecordSetProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for RecordSet struct.
func (rs *RecordSet) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		rs.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		rs.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		rs.Type = &typeVar
	}

	v = m["etag"]
	if v != nil {
		var etag string
		err = json.Unmarshal(*m["etag"], &etag)
		if err != nil {
			return err
		}
		rs.Etag = &etag
	}

	v = m["properties"]
	if v != nil {
		var properties RecordSetProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		rs.RecordSetProperties = &properties
	}

	return nil
}

// RecordSetListResult the response to a record set List operation.
type RecordSetListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the record sets in the response.
	Value *[]RecordSet `json:"value,omitempty"`
	// NextLink - The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// RecordSetListResultIterator provides access to a complete listing of RecordSet values.
type RecordSetListResultIterator struct {
	i    int
	page RecordSetListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RecordSetListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RecordSetListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RecordSetListResultIterator) Response() RecordSetListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RecordSetListResultIterator) Value() RecordSet {
	if !iter.page.NotDone() {
		return RecordSet{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (rslr RecordSetListResult) IsEmpty() bool {
	return rslr.Value == nil || len(*rslr.Value) == 0
}

// recordSetListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rslr RecordSetListResult) recordSetListResultPreparer() (*http.Request, error) {
	if rslr.NextLink == nil || len(to.String(rslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rslr.NextLink)))
}

// RecordSetListResultPage contains a page of RecordSet values.
type RecordSetListResultPage struct {
	fn   func(RecordSetListResult) (RecordSetListResult, error)
	rslr RecordSetListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RecordSetListResultPage) Next() error {
	next, err := page.fn(page.rslr)
	if err != nil {
		return err
	}
	page.rslr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RecordSetListResultPage) NotDone() bool {
	return !page.rslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RecordSetListResultPage) Response() RecordSetListResult {
	return page.rslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RecordSetListResultPage) Values() []RecordSet {
	if page.rslr.IsEmpty() {
		return nil
	}
	return *page.rslr.Value
}

// RecordSetProperties represents the properties of the records in the record set.
type RecordSetProperties struct {
	// Metadata - The metadata attached to the record set.
	Metadata *map[string]*string `json:"metadata,omitempty"`
	// TTL - The TTL (time-to-live) of the records in the record set.
	TTL *int64 `json:"TTL,omitempty"`
	// ARecords - The list of A records in the record set.
	ARecords *[]ARecord `json:"ARecords,omitempty"`
	// AaaaRecords - The list of AAAA records in the record set.
	AaaaRecords *[]AaaaRecord `json:"AAAARecords,omitempty"`
	// MxRecords - The list of MX records in the record set.
	MxRecords *[]MxRecord `json:"MXRecords,omitempty"`
	// NsRecords - The list of NS records in the record set.
	NsRecords *[]NsRecord `json:"NSRecords,omitempty"`
	// PtrRecords - The list of PTR records in the record set.
	PtrRecords *[]PtrRecord `json:"PTRRecords,omitempty"`
	// SrvRecords - The list of SRV records in the record set.
	SrvRecords *[]SrvRecord `json:"SRVRecords,omitempty"`
	// TxtRecords - The list of TXT records in the record set.
	TxtRecords *[]TxtRecord `json:"TXTRecords,omitempty"`
	// CnameRecord - The CNAME record in the  record set.
	CnameRecord *CnameRecord `json:"CNAMERecord,omitempty"`
	// SoaRecord - The SOA record in the record set.
	SoaRecord *SoaRecord `json:"SOARecord,omitempty"`
}

// RecordSetUpdateParameters parameters supplied to update a record set.
type RecordSetUpdateParameters struct {
	// RecordSet - Specifies information about the record set being updated.
	RecordSet *RecordSet `json:"RecordSet,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
}

// SoaRecord an SOA record.
type SoaRecord struct {
	// Host - The domain name of the authoritative name server for this SOA record.
	Host *string `json:"host,omitempty"`
	// Email - The email contact for this SOA record.
	Email *string `json:"email,omitempty"`
	// SerialNumber - The serial number for this SOA record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
	// RefreshTime - The refresh value for this SOA record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// RetryTime - The retry time for this SOA record.
	RetryTime *int64 `json:"retryTime,omitempty"`
	// ExpireTime - The expire time for this SOA record.
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// MinimumTTL - The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTTL *int64 `json:"minimumTTL,omitempty"`
}

// SrvRecord an SRV record.
type SrvRecord struct {
	// Priority - The priority value for this SRV record.
	Priority *int32 `json:"priority,omitempty"`
	// Weight - The weight value for this SRV record.
	Weight *int32 `json:"weight,omitempty"`
	// Port - The port value for this SRV record.
	Port *int32 `json:"port,omitempty"`
	// Target - The target domain name for this SRV record.
	Target *string `json:"target,omitempty"`
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
}

// TxtRecord a TXT record.
type TxtRecord struct {
	// Value - The text value of this TXT record.
	Value *[]string `json:"value,omitempty"`
}

// Zone describes a DNS zone.
type Zone struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Etag - The etag of the zone.
	Etag *string `json:"etag,omitempty"`
	// ZoneProperties - The properties of the zone.
	*ZoneProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Zone struct.
func (z *Zone) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["etag"]
	if v != nil {
		var etag string
		err = json.Unmarshal(*m["etag"], &etag)
		if err != nil {
			return err
		}
		z.Etag = &etag
	}

	v = m["properties"]
	if v != nil {
		var properties ZoneProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		z.ZoneProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		z.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		z.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		z.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		z.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		z.Tags = &tags
	}

	return nil
}

// ZoneDeleteResult the response to a Zone Delete operation.
type ZoneDeleteResult struct {
	autorest.Response `json:"-"`
	// AzureAsyncOperation - Users can perform a Get on Azure-AsyncOperation to get the status of their delete Zone operations.
	AzureAsyncOperation *string `json:"azureAsyncOperation,omitempty"`
	// Status - Possible values include: 'InProgress', 'Succeeded', 'Failed'
	Status OperationStatus `json:"status,omitempty"`
	// StatusCode - Possible values include: 'Continue', 'SwitchingProtocols', 'OK', 'Created', 'Accepted', 'NonAuthoritativeInformation', 'NoContent', 'ResetContent', 'PartialContent', 'MultipleChoices', 'Ambiguous', 'MovedPermanently', 'Moved', 'Found', 'Redirect', 'SeeOther', 'RedirectMethod', 'NotModified', 'UseProxy', 'Unused', 'TemporaryRedirect', 'RedirectKeepVerb', 'BadRequest', 'Unauthorized', 'PaymentRequired', 'Forbidden', 'NotFound', 'MethodNotAllowed', 'NotAcceptable', 'ProxyAuthenticationRequired', 'RequestTimeout', 'Conflict', 'Gone', 'LengthRequired', 'PreconditionFailed', 'RequestEntityTooLarge', 'RequestURITooLong', 'UnsupportedMediaType', 'RequestedRangeNotSatisfiable', 'ExpectationFailed', 'UpgradeRequired', 'InternalServerError', 'NotImplemented', 'BadGateway', 'ServiceUnavailable', 'GatewayTimeout', 'HTTPVersionNotSupported'
	StatusCode HTTPStatusCode `json:"statusCode,omitempty"`
	RequestID  *string        `json:"requestId,omitempty"`
}

// ZoneListResult the response to a Zone List or ListAll operation.
type ZoneListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the DNS zones.
	Value *[]Zone `json:"value,omitempty"`
	// NextLink - The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ZoneListResultIterator provides access to a complete listing of Zone values.
type ZoneListResultIterator struct {
	i    int
	page ZoneListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ZoneListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ZoneListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ZoneListResultIterator) Response() ZoneListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ZoneListResultIterator) Value() Zone {
	if !iter.page.NotDone() {
		return Zone{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (zlr ZoneListResult) IsEmpty() bool {
	return zlr.Value == nil || len(*zlr.Value) == 0
}

// zoneListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (zlr ZoneListResult) zoneListResultPreparer() (*http.Request, error) {
	if zlr.NextLink == nil || len(to.String(zlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(zlr.NextLink)))
}

// ZoneListResultPage contains a page of Zone values.
type ZoneListResultPage struct {
	fn  func(ZoneListResult) (ZoneListResult, error)
	zlr ZoneListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ZoneListResultPage) Next() error {
	next, err := page.fn(page.zlr)
	if err != nil {
		return err
	}
	page.zlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ZoneListResultPage) NotDone() bool {
	return !page.zlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ZoneListResultPage) Response() ZoneListResult {
	return page.zlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ZoneListResultPage) Values() []Zone {
	if page.zlr.IsEmpty() {
		return nil
	}
	return *page.zlr.Value
}

// ZoneProperties represents the properties of the zone.
type ZoneProperties struct {
	// MaxNumberOfRecordSets - The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty"`
	// NumberOfRecordSets - The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty"`
	// NameServers - The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers *[]string `json:"nameServers,omitempty"`
}

// ZonesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type ZonesDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future ZonesDeleteFuture) Result(client ZonesClient) (zdr ZoneDeleteResult, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return zdr, autorest.NewError("dns.ZonesDeleteFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		zdr, err = client.DeleteResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	zdr, err = client.DeleteResponder(resp)
	return
}
