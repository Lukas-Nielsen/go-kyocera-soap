package kyocera_soap

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	create_group_address_enumerationRequest  = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/create_group_address_enumeration</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:create_group_address_enumerationRequest><ns1:number>@@list_count@@</ns1:number></ns1:create_group_address_enumerationRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
	get_group_address_listRequest            = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/get_group_address_list</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:get_group_address_listRequest><ns1:enumeration>@@request_id@@</ns1:enumeration></ns1:get_group_address_listRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
	destroy_group_address_enumerationRequest = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/destroy_group_address_enumeration</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:destroy_group_address_enumerationRequest><ns1:enumeration>@@request_id@@</ns1:enumeration></ns1:destroy_group_address_enumerationRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
)

type create_group_address_enumerationRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                                  string `xml:",chardata"`
		Space                                 string `xml:"space,attr"`
		CreateGroupAddressEnumerationResponse struct {
			Text        string `xml:",chardata"`
			Result      string `xml:"result"`
			Enumeration string `xml:"enumeration"`
		} `xml:"create_group_address_enumerationResponse"`
	} `xml:"Body"`
}

type get_group_address_listRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                        string `xml:",chardata"`
		Space                       string `xml:"space,attr"`
		GetGroupAddressListResponse struct {
			Text         string         `xml:",chardata"`
			Result       string         `xml:"result"`
			GroupAddress []groupAddress `xml:"group_address"`
		} `xml:"get_group_address_listResponse"`
	} `xml:"Body"`
}

type groupAddress struct {
	Text            string `xml:",chardata"`
	NameInformation struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name"`
		Furigana string `xml:"furigana"`
		ID       string `xml:"id"`
	} `xml:"name_information"`
	AddressNumber                       string `xml:"address_number"`
	GroupAddressRegistrationInformation struct {
		Text            string            `xml:",chardata"`
		Type            string            `xml:"type"`
		PersonalAddress []personalAddress `xml:"personal_address"`
	} `xml:"group_address_registration_information"`
}

type destroy_group_address_enumerationRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                                   string `xml:",chardata"`
		Space                                  string `xml:"space,attr"`
		DestroyGroupAddressEnumerationResponse struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
		} `xml:"destroy_group_address_enumerationResponse"`
	} `xml:"Body"`
}

func (c *Config) GetAddressbookGroup() []groupAddress {
	var create create_group_address_enumerationRequestResponse
	var data get_group_address_listRequestResponse
	var destroy destroy_group_address_enumerationRequestResponse
	prepare := true
	if c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(create_group_address_enumerationRequest, "@@list_count@@", "100", 1), &create) && create.Body.CreateGroupAddressEnumerationResponse.Result == "SUCCESS" {
		for prepare {
			get := c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(get_group_address_listRequest, "@@request_id@@", create.Body.CreateGroupAddressEnumerationResponse.Enumeration, 1), &data)
			if get && data.Body.GetGroupAddressListResponse.Result == "ALL_GET_COMPLETE" {
				if c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(destroy_group_address_enumerationRequest, "@@request_id@@", create.Body.CreateGroupAddressEnumerationResponse.Enumeration, 1), &destroy) && destroy.Body.DestroyGroupAddressEnumerationResponse.Result == "SUCCESS" {
					return data.Body.GetGroupAddressListResponse.GroupAddress
				}
			} else if get && data.Body.GetGroupAddressListResponse.Result == "PREPARING_NOW" {
				prepare = true
			} else {
				prepare = false
			}
			time.Sleep(500 * time.Millisecond)
		}
		return []groupAddress{}
	}
	return []groupAddress{}
}
