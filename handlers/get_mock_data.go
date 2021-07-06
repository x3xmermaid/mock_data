package handlers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"ketitik/netmonk/mock-app-data/lib"
	nrespwriter "ketitik/netmonk/mock-app-data/lib/responsewriter"
	nmodel "ketitik/netmonk/mock-app-data/models"
	"net/http"
	"time"

	"github.com/beevik/etree"
	wr "github.com/mroth/weightedrand"
	"golang.org/x/net/html/charset"
)

const (
	//OnlineSmall online
	OnlineSmall = "Online"
	//OnlineBig online
	OnlineBig = "ONLINE"
	//OfflineSmall OfflineSmall
	OfflineSmall = "Offline"
	//OfflineBig OfflineBig
	OfflineBig = "OFFLINE"
	//Los los
	Los = "LOS"
	//Stop stop
	Stop         = "Stop"
	disconnected = "Disconnected"
	dyingGasp    = "DYING GASP"
)

// GetXMLMockData get random mock data
func (h *Handler) GetXMLMockData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mermaid")
	rf := &nrespwriter.ResponseFormat{}
	deviceStatus := &nmodel.XMLResponse{}

	rawbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rf.ResponseNOK(http.StatusInternalServerError, "error read body", w)
		return
	}

	reader := bytes.NewReader(rawbody)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&deviceStatus)
	if err != nil {
		rf.ResponseNOK(http.StatusInternalServerError, "error decode body", w)
		return
	}

	jsoned := lib.LoadMockData("mock_data.json")

	xmlMockData := []nmodel.MockData{}
	err = json.Unmarshal(jsoned, &xmlMockData)
	if err != nil {
		rf.ResponseNOK(http.StatusOK, "cannot read json file", w)
		return
	}

	statusHSI := wr.NewChooser(
		wr.Choice{Item: OnlineSmall, Weight: 8},
		wr.Choice{Item: Stop, Weight: 1},
		wr.Choice{Item: time.Now().String(), Weight: 1},
	)

	messageShow := wr.NewChooser(
		wr.Choice{Item: "showMessage", Weight: 1},
		wr.Choice{Item: "showMessage2", Weight: 5},
		wr.Choice{Item: "hiddenMessage", Weight: 4},
	)
	fmt.Println("data")

	statusONT := wr.NewChooser(
		wr.Choice{Item: OnlineBig, Weight: 7},
		wr.Choice{Item: OfflineBig, Weight: 1},
		wr.Choice{Item: dyingGasp, Weight: 1},
		wr.Choice{Item: Los, Weight: 1},
	)

	var xmlDocBytes []byte
	isRegisteredND := false
	for _, value := range xmlMockData {
		if value.ND == deviceStatus.Body.Ukur.Input.ND {
			isRegisteredND = true
			value.StatusHSI = statusHSI.Pick().(string)
			value.StatusONT = statusONT.Pick().(string)
			xmlDocBytes, err = createXMLResponse(&value, messageShow.Pick().(string))
			if err != nil {
				rf.ResponseNOK(http.StatusOK, "get xml bytes response", w)
				return
			}
		}
	}

	if !isRegisteredND {
		rf.ResponseNOK(http.StatusBadRequest, "nd not listed", w)
		return
	}

	w.Write(xmlDocBytes)
	return
}

func createXMLResponse(deviceStatus *nmodel.MockData, messageShow string) ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="ISO-8859-1"`)

	root := doc.CreateElement("SOAP-ENV:Envelope")
	root.CreateAttr("SOAP-ENV:encodingStyle", "http://schemas.xmlsoap.org/soap/encoding/")
	root.CreateAttr("xmlns:SOAP-ENV", "http://schemas.xmlsoap.org/soap/envelope/")
	root.CreateAttr("xmlns:xsd", "http://www.w3.org/2001/XMLSchema")
	root.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	root.CreateAttr("xmlns:SOAP-ENC", "http://schemas.xmlsoap.org/soap/encoding/")
	root.CreateAttr("xmlns:tns", "http://10.62.165.53/soap/iBooster")

	body := root.CreateElement("SOAP-ENV:Body")
	ukur := body.CreateElement("ns1:ukurResponse")
	ukur.CreateAttr("xmlns:ns1", "http://ibooster.telkom.co.id/api/soap.php")

	output := ukur.CreateElement("output")
	output.CreateAttr("xsi:type", "ibo:ukur_out")

	xmlIDUkur := output.CreateElement("id_ukur")
	xmlIDUkur.CreateAttr("xsi:type", "xsd:string")
	xmlIDUkur.CreateText(deviceStatus.IDUkur)

	xmlIPOlt := output.CreateElement("ip_olt")
	xmlIPOlt.CreateAttr("xsi:type", "xsd:string")
	xmlIPOlt.CreateText(deviceStatus.IPOLT)

	xmlHostname := output.CreateElement("hostname_olt")
	xmlHostname.CreateAttr("xsi:type", "xsd:string")
	xmlHostname.CreateText(deviceStatus.HostnameOLT)

	xmlFrame := output.CreateElement("frame")
	xmlFrame.CreateAttr("xsi:type", "xsd:string")
	xmlFrame.CreateText("test frame")

	xmlSlot := output.CreateElement("slot")
	xmlSlot.CreateAttr("xsi:type", "xsd:string")
	xmlSlot.CreateText(deviceStatus.Slot)

	xmlPort := output.CreateElement("port")
	xmlPort.CreateAttr("xsi:type", "xsd:string")
	xmlPort.CreateText(deviceStatus.Port)

	xmlOnuID := output.CreateElement("onu_id")
	xmlOnuID.CreateAttr("xsi:type", "xsd:string")
	xmlOnuID.CreateText(deviceStatus.OnuID)

	xmlOnuRxPWR := output.CreateElement("onu_rx_pwr")
	xmlOnuRxPWR.CreateAttr("xsi:type", "xsd:string")
	xmlOnuRxPWR.CreateText(deviceStatus.OnuRxPwr)

	xmlOnuTxPWR := output.CreateElement("onu_tx_pwr")
	xmlOnuTxPWR.CreateAttr("xsi:type", "xsd:string")
	xmlOnuTxPWR.CreateText(deviceStatus.OnuTxPwr)

	xmlOltTxPWR := output.CreateElement("olt_tx_pwr")
	xmlOltTxPWR.CreateAttr("xsi:type", "xsd:string")
	xmlOltTxPWR.CreateText(deviceStatus.OltTxPwr)

	xmlOltRxPWR := output.CreateElement("olt_rx_pwr")
	xmlOltRxPWR.CreateAttr("xsi:type", "xsd:string")
	xmlOltRxPWR.CreateText(deviceStatus.OltRxPwr)

	xmlSN := output.CreateElement("sn")
	xmlSN.CreateAttr("xsi:type", "xsd:string")
	xmlSN.CreateText(deviceStatus.SN)

	xmlND := output.CreateElement("nd")
	xmlND.CreateAttr("xsi:type", "xsd:string")
	xmlND.CreateText(deviceStatus.ND)

	xmlRealm := output.CreateElement("realm")
	xmlRealm.CreateAttr("xsi:type", "xsd:string")
	xmlRealm.CreateText(deviceStatus.Realm)

	xmlStatusHSI := output.CreateElement("status_hsi")
	xmlStatusHSI.CreateAttr("xsi:type", "xsd:string")
	xmlStatusHSI.CreateText(deviceStatus.StatusHSI)

	xmlStatusONT := output.CreateElement("status_ont")
	xmlStatusONT.CreateAttr("xsi:type", "xsd:string")
	xmlStatusONT.CreateText(deviceStatus.StatusONT)

	if messageShow == "showMessage" {
		xmlMessage := output.CreateElement("MESSAGE")
		xmlMessage.CreateAttr("xsi:type", "xsd:string")
		xmlMessage.CreateText("nomor internet errors")
	}

	if messageShow == "showMessage2" {
		xmlMessage := output.CreateElement("MESSAGE")
		xmlMessage.CreateAttr("xsi:type", "xsd:string")
		xmlMessage.CreateText("Nomor internet122421321601telkom.net tidak memiliki usage")
	}

	xmlDocByte, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}

	return xmlDocByte, nil
}
