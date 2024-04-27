package main

import "encoding/xml"

func GetRouterEntities() any {
    return &RouterEntities{}
}

type Interfaces struct {
	Name        string `xml:"interface-name"`
	Ipv4Address string `xml:"ipv4-network>addresses>primary>address"`
}
type RouterEntities struct {
	XMLName    xml.Name     `xml:"rpc-reply"`
	Hostname   string       `xml:"data>host-names>host-name"`
	Interfaces []Interfaces `xml:"data>interface-configurations>interface-configuration"`
}
