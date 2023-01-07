package kyocera_soap

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	create_personal_address_enumerationRequest  = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/create_personal_address_enumeration</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:create_personal_address_enumerationRequest><ns1:number>@@list_count@@</ns1:number></ns1:create_personal_address_enumerationRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
	get_personal_address_listRequest            = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/get_personal_address_list</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:get_personal_address_listRequest><ns1:enumeration>@@request_id@@</ns1:enumeration></ns1:get_personal_address_listRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
	destroy_personal_address_enumerationRequest = `<?xml version="1.0" encoding="utf-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:ns1="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book"><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand="true">http://www.kyoceramita.com/ws/km-wsdl/setting/address_book/destroy_personal_address_enumeration</wsa:Action></SOAP-ENV:Header><SOAP-ENV:Body><ns1:destroy_personal_address_enumerationRequest><ns1:enumeration>@@request_id@@</ns1:enumeration></ns1:destroy_personal_address_enumerationRequest></SOAP-ENV:Body></SOAP-ENV:Envelope>`
)

type create_personal_address_enumerationRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                                     string `xml:",chardata"`
		Space                                    string `xml:"space,attr"`
		CreatePersonalAddressEnumerationResponse struct {
			Text        string `xml:",chardata"`
			Result      string `xml:"result"`
			Enumeration string `xml:"enumeration"`
		} `xml:"create_personal_address_enumerationResponse"`
	} `xml:"Body"`
}

type get_personal_address_listRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                           string `xml:",chardata"`
		Space                          string `xml:"space,attr"`
		GetPersonalAddressListResponse struct {
			Text            string            `xml:",chardata"`
			Result          string            `xml:"result"`
			PersonalAddress []PersonalAddress `xml:"personal_address"`
		} `xml:"get_personal_address_listResponse"`
	} `xml:"Body"`
}

type PersonalAddress struct {
	Text            string `xml:",chardata"`
	NameInformation struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name"`
		Furigana string `xml:"furigana"`
		ID       string `xml:"id"`
	} `xml:"name_information"`
	EmailInformation struct {
		Text    string `xml:",chardata"`
		Address string `xml:"address"`
	} `xml:"email_information"`
	FtpInformation struct {
		Text          string `xml:",chardata"`
		ServerName    string `xml:"server_name"`
		FilePath      string `xml:"file_path"`
		PortNumber    string `xml:"port_number"`
		LoginName     string `xml:"login_name"`
		LoginPassword string `xml:"login_password"`
	} `xml:"ftp_information"`
	SmbInformation struct {
		Text          string `xml:",chardata"`
		ServerName    string `xml:"server_name"`
		FilePath      string `xml:"file_path"`
		PortNumber    string `xml:"port_number"`
		LoginName     string `xml:"login_name"`
		LoginPassword string `xml:"login_password"`
	} `xml:"smb_information"`
	IfaxInformation struct {
		Text           string   `xml:",chardata"`
		Address        string   `xml:"address"`
		SendResolution []string `xml:"send_resolution"`
		FileSize       string   `xml:"file_size"`
		SendingMode    string   `xml:"sending_mode"`
		Mode           string   `xml:"mode"`
		FileFormat     string   `xml:"file_format"`
	} `xml:"ifax_information"`
	FaxInformation struct {
		Text                    string `xml:",chardata"`
		FaxNumber               string `xml:"fax_number"`
		Ecm                     string `xml:"ecm"`
		CodeKeyID               string `xml:"code_key_id"`
		CodeSendSetting         string `xml:"code_send_setting"`
		CodeBoxNumber           string `xml:"code_box_number"`
		CodeBoxSetting          string `xml:"code_box_setting"`
		Honorific               string `xml:"honorific"`
		ConnectionBeginingSpeed string `xml:"connection_begining_speed"`
	} `xml:"fax_information"`
}

type destroy_personal_address_enumerationRequestResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action string `xml:"Action"`
	} `xml:"Header"`
	Body struct {
		Text                                      string `xml:",chardata"`
		Space                                     string `xml:"space,attr"`
		DestroyPersonalAddressEnumerationResponse struct {
			Text   string `xml:",chardata"`
			Result string `xml:"result"`
		} `xml:"destroy_personal_address_enumerationResponse"`
	} `xml:"Body"`
}

func (c *Config) GetAddressbookEntry() []PersonalAddress {
	var create create_personal_address_enumerationRequestResponse
	var data get_personal_address_listRequestResponse
	var destroy destroy_personal_address_enumerationRequestResponse
	if c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(create_personal_address_enumerationRequest, "@@list_count@@", "1000", 1), &create) && create.Body.CreatePersonalAddressEnumerationResponse.Result == "SUCCESS" {
		time.Sleep(5 * time.Second)
		if c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(get_personal_address_listRequest, "@@request_id@@", create.Body.CreatePersonalAddressEnumerationResponse.Enumeration, 1), &data) && data.Body.GetPersonalAddressListResponse.Result == "ALL_GET_COMPLETE" {
			if c.reuqest("/ws/km-wsdl/setting/address_book", strings.Replace(destroy_personal_address_enumerationRequest, "@@request_id@@", create.Body.CreatePersonalAddressEnumerationResponse.Enumeration, 1), &destroy) && destroy.Body.DestroyPersonalAddressEnumerationResponse.Result == "SUCCESS" {
				return data.Body.GetPersonalAddressListResponse.PersonalAddress
			}
			return []PersonalAddress{}
		}
		return []PersonalAddress{}
	}
	return []PersonalAddress{}
}
