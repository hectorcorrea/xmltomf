package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

// <resource resourceClass="Project" resourceID="10.34770/az09-0001" resourceIDType="DOI">
//
//	<projectID projectIDType="DOI" inherited="false" discoverable="true" trackingLevel="ResourceRecord">10.34770/az09-0001</projectID>
//	<dataSponsor userID="mjc12" userIDType="NetID" discoverable="true" inherited="true" trackingLevel="ResourceRecord"></dataSponsor>
//
// </resource>
type Resource struct {
	XmlName        xml.Name    `xml:"resource"`
	ResourceClass  string      `xml:"resourceClass,attr"`
	ResourceID     string      `xml:"resourceID,attr"`
	ResourceIDType string      `xml:"resourceIDType,attr"`
	ProjectID      Project     `xml:"projectID"`
	DataSponsor    DataSponsor `xml:"dataSponsor"`
}

type Project struct {
	ProjectIDType string `xml:"projectIDType,attr"`
	Inherited     string `xml:"inherited,attr"`
	Discoverable  string `xml:"discoverable,attr"`
	TrackingLevel string `xml:"trackingLevel,attr"`
	Value         string `xml:",chardata"`
}

type DataSponsor struct {
	UserID        string `xml:"userID,attr"`
	UserIDType    string `xml:"userIDType,attr"`
	Inherited     string `xml:"inherited,attr"`
	Discoverable  string `xml:"discoverable,attr"`
	TrackingLevel string `xml:"trackingLevel,attr"`
}

var xmlFile string
var metadataNamespace string

func init() {
	flag.StringVar(&xmlFile, "file", "", "XML file to parse")
	flag.StringVar(&metadataNamespace, "namespace", "tigerdataX:resourceDoc", "Fullnamespace of the metadata schema to use")
	flag.Parse()
}

func main() {
	if xmlFile == "" {
		fmt.Printf("Must provide a file\r\n")
		return
	}
	reader, err := os.Open(xmlFile)
	if err != nil {
		fmt.Printf("Error reading XML file %s: %s\r\n", xmlFile, err)
		return
	}
	defer reader.Close()

	byteValue, _ := io.ReadAll(reader)
	var resource Resource
	xml.Unmarshal(byteValue, &resource)
	fmt.Printf("====================\r\n")
	fmt.Printf("DEBUG: %#v\r\n", resource)
	fmt.Printf("\r\n")

	fmt.Printf("resource\r\n")
	fmt.Printf("  resourceClass: %s\r\n", resource.ResourceClass)
	fmt.Printf("  resourceID: %s\r\n", resource.ResourceID)
	fmt.Printf("  resourceIDType: %s\r\n", resource.ResourceIDType)

	fmt.Printf("project\r\n")
	fmt.Printf("  projectIDType: %s\r\n", resource.ProjectID.ProjectIDType)
	fmt.Printf("  inherited: %s\r\n", resource.ProjectID.Inherited)
	fmt.Printf("  discoverable: %s\r\n", resource.ProjectID.Discoverable)
	fmt.Printf("  trackingLevel: %s\r\n", resource.ProjectID.TrackingLevel)
	fmt.Printf("  value: %s\r\n", resource.ProjectID.Value)

	fmt.Printf("dataSponsor\r\n")
	fmt.Printf("  userID: %s\r\n", resource.DataSponsor.UserID)
	fmt.Printf("  userIDType: %s\r\n", resource.DataSponsor.UserIDType)
	fmt.Printf("  inherited: %s\r\n", resource.DataSponsor.Inherited)
	fmt.Printf("  discoverable: %s\r\n", resource.DataSponsor.Discoverable)
	fmt.Printf("  trackingLevel: %s\r\n", resource.DataSponsor.TrackingLevel)

	fmt.Printf("====================\r\n")
	fmt.Printf(":meta < \\ \r\n")
	fmt.Printf("  %s < \\ \r\n", metadataNamespace)
	fmt.Printf("    :resource -resourceClass \"%s\" -resourceID \"%s\" -resourceIDTYpe \"%s\" < \\ \r\n", resource.ResourceClass, resource.ResourceID, resource.ResourceIDType)
	fmt.Printf("      :projectID -projectIDType \"%s\" \"%s\" \\ \r\n", resource.ProjectID.ProjectIDType, resource.ProjectID.Value)
	fmt.Printf("    > \r\n")
	fmt.Printf("  > \r\n")
	fmt.Printf("> \r\n")
}
