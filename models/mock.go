package models

import "encoding/xml"

//XMLResponse main struct arrangement for ibooster xml response
type XMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

//Body define content xml struct response
type Body struct {
	XMLName xml.Name `xml:"Body"`
	Ukur    Ukur     `xml:"ukur"`
}

//Ukur define struct of ukur response
type Ukur struct {
	XMLName xml.Name        `xml:"ukur"`
	Input   xmlInputRequest `xml:"input"`
}

// MockData struct for mock data
type MockData struct {
	CustomerID  string `xml:"customer_id" json:"customer_id"`
	IDUkur      string `xml:"id_ukur" json:"id_ukur"`
	IPOLT       string `xml:"ip_olt" json:"ip_olt"`
	HostnameOLT string `xml:"hostname_olt" json:"hostname_olt"`
	Frame       string `xml:"frame" json:"frame"`
	Slot        string `xml:"slot" json:"slot"`
	Port        string `xml:"port" json:"port"`
	OnuID       string `xml:"onu_id" json:"onu_id"`
	OnuRxPwr    string `xml:"onu_rx_pwr" json:"onu_rx_pwr"`
	OnuTxPwr    string `xml:"onu_tx_pwr" json:"onu_tx_pwr"`
	OltRxPwr    string `xml:"olt_rx_pwr" json:"olt_rx_pwr"`
	OltTxPwr    string `xml:"olt_tx_pwr" json:"olt_tx_pwr"`
	SN          string `xml:"sn" json:"sn"`
	ND          string `xml:"nd"  json:"nd"`
	Realm       string `xml:"realm" json:"realm"`
	StatusHSI   string `xml:"status_hsi" json:"status_hsi"`
	StatusONT   string `xml:"status_ont" json:"status_ont"`
}

// xmlInputRequest struct for xml request mock data
type xmlInputRequest struct {
	ReqID string `xml:"req_id" json:"id_ukur"`
	ND    string `xml:"nd"  json:"nd"`
	Realm string `xml:"realm" json:"realm"`
}
