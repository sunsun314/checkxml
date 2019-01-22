package checkxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestAllConfig_R(t *testing.T) {

	data := []byte(`
        <?xml version="1.0" encoding="UTF-8"?>
       <!DOCTYPE dble:rule SYSTEM "rule.dtd">
          <dble:rule>
            <tableRule name="sharding-by-enum" unknow0="ccc">
              <unknow1></unknow1>
              <rule unknow2="unknow2">
                <unknow3>qwe</unknow3>
                <columns unknow4="ggg">id</columns>
                <algorithm>enum</algorithm>
              </rule>
            </tableRule>

            <function name="hashmod" class="Hash">
              <property name="partitionCount">4</property>
              <property name="partitionLength">1</property>
             </function>
            </dble:rule>
        `)

	tv := CoreRuleXml{}
	tags, _, err := UnknownXMLTags(data, tv)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tags)
}

type CoreRuleXml struct {
	XMLName    xml.Name         `xml:"dble:rule"`
	TableRules []*CoreTableRule `xml:"tableRule"`
	Functions  []*CoreFunction  `xml:"function"`
}

type CoreTableRule struct {
	XMLName xml.Name  `xml:"tableRule"`
	Name    string    `xml:"name,attr"`
	Rule    *CoreRule `xml:"rule"`
}

type CoreRule struct {
	XMLName   xml.Name `xml:"rule"`
	Columns   string   `xml:"columns"`
	Algorithm string   `xml:"algorithm"`
}

type CoreFunction struct {
	XMLName    xml.Name        `xml:"function"`
	Name       string          `xml:"name,attr"`
	Class      string          `xml:"class,attr"`
	Properties []*CoreProperty `xml:"property"`
}

type CoreSchemaXml struct {
	XMLName   xml.Name        `xml:"dble:schema"`
	Xmlns     string          `xml:"xmlns:dble,attr"`
	Schemas   []*CoreSchema   `xml:"schema"`
	DataNodes []*CoreDataNode `xml:"dataNode"`
	DataHosts []*CoreDataHost `xml:"dataHost"`
}

type AsyncRootXmlForUnmarshal struct {
	XMLName         xml.Name       `xml:"asyncRoot"`
	Level           string         `xml:"level,attr"`
	IncludeLocation string         `xml:"includeLocation,attr"`
	AppenderRefs    []*AppenderRef `xml:"AppenderRef"`
}

type AppenderRef struct {
	XMLName xml.Name `xml:"AppenderRef"`
	Ref     string   `xml:"ref,attr"`
}

// server.xml typ
type CoreServerXml struct {
	XMLName  xml.Name      `xml:"dble:server"`
	Xmlns    string        `xml:"xmlns:dble,attr"`
	System   *CoreSystem   `xml:"system"`
	Firewall *CoreFirewall `xml:"firewall"`
	Users    []*CoreUser   `xml:"user"`
}

type CoreSystem struct {
	XMLName    xml.Name        `xml:"system"`
	Properties []*CoreProperty `xml:"property"`
}

type CoreProperty struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"#text"`
}

type CoreFirewall struct {
	XMLName   xml.Name                `xml:"firewall"`
	WhiteHost *CoreFirewallWriteHosts `xml:"whitehost"`
	BlackList *CoreFirewallBlackList  `xml:"blacklist"`
}

type CoreFirewallWriteHosts struct {
	XMLName xml.Name                 `xml:"whitehost"`
	Host    []*CoreFirewallWriteHost `xml:"host"`
}

type CoreFirewallWriteHost struct {
	XMLName  xml.Name        `xml:"host"`
	Host     string          `xml:"host,attr"`
	User     string          `xml:"user,attr"`
	Property []*CoreProperty `xml:"property"`
}

type CoreFirewallBlackList struct {
	XMLName  xml.Name        `xml:"blacklist"`
	Check    string          `xml:"check,attr"`
	Property []*CoreProperty `xml:"property"`
}

type CoreUser struct {
	XMLName    xml.Name        `xml:"user"`
	Name       string          `xml:"name,attr"`
	Properties []*CoreProperty `xml:"property"`
	Privileges *CorePrivileges `xml:"privileges"`
}

type CorePrivileges struct {
	XMLName xml.Name               `xml:"privileges"`
	Check   string                 `xml:"check,attr"`
	Schema  []*CorePrivilegeSchema `xml:"schema"`
}

type CorePrivilegeSchema struct {
	XMLName xml.Name              `xml:"schema"`
	Name    string                `xml:"name,attr"`
	Dml     string                `xml:"dml,attr"`
	Tables  []*CorePrivilegeTable `xml:"table"`
}

type CorePrivilegeTable struct {
	XMLName xml.Name `xml:"table"`
	Name    string   `xml:"name,attr"`
	Dml     string   `xml:"dml,attr"`
}

// scheme.xml typ
type CoreSchema struct {
	XMLName     xml.Name     `xml:"schema"`
	Name        string       `xml:"name,attr"`
	SqlMaxLimit *string      `xml:"sqlMaxLimit,attr,omitempty"`
	DataNode    *string      `xml:"dataNode,attr,omitempty"`
	Tables      []*CoreTable `xml:"table"`
}

type CoreTable struct {
	XMLName       xml.Name          `xml:"table"`
	Name          string            `xml:"name,attr"`
	Type          *string           `xml:"type,attr,omitempty"`
	PrimaryKey    *string           `xml:"primaryKey,attr,omitempty"`
	AutoIncrement *string           `xml:"autoIncrement,attr,omitempty"`
	NeedAddLimit  *string           `xml:"needAddLimit,attr,omitempty"`
	Rule          *string           `xml:"rule,attr,omitempty"`
	RuleRequired  *string           `xml:"ruleRequired,attr,omitempty"`
	DataNode      string            `xml:"dataNode,attr"`
	ChildTables   []*CoreChildTable `xml:"childTable"`
}

type CoreChildTable struct {
	XMLName       xml.Name          `xml:"childTable"`
	Name          string            `xml:"name,attr"`
	JoinKey       string            `xml:"joinKey,attr"`
	ParentKey     string            `xml:"parentKey,attr"`
	PrimaryKey    *string           `xml:"primaryKey,attr,omitempty"`
	AutoIncrement *string           `xml:"autoIncrement,attr,omitempty"`
	ChildTables   []*CoreChildTable `xml:"childTable"`
}

type CoreDataNode struct {
	XMLName  xml.Name `xml:"dataNode"`
	Name     string   `xml:"name,attr"`
	DataHost string   `xml:"dataHost,attr"`
	Database string   `xml:"database,attr"`
}

type CoreDataHost struct {
	XMLName               xml.Name         `xml:"dataHost"`
	Name                  string           `xml:"name,attr"`
	MaxCon                string           `xml:"maxCon,attr"`
	MinCon                string           `xml:"minCon,attr"`
	Balance               string           `xml:"balance,attr"`
	SwitchType            *string          `xml:"switchType,attr,omitempty"`
	SlaveThreshold        *string          `xml:"slaveThreshold,attr,omitempty"`
	TempReadHostAvailable *string          `xml:"tempReadHostAvailable,attr,omitempty"`
	KeepOrig              *string          `xml:"keepOrig,attr,omitempty"`
	Heartbeat             string           `xml:"heartbeat"`
	WriteHosts            []*CoreWriteHost `xml:"writeHost,omitempty"`
}

type CoreWriteHost struct {
	XMLName      xml.Name        `xml:"writeHost"`
	Host         string          `xml:"host,attr"`
	Url          string          `xml:"url,attr"`
	User         string          `xml:"user,attr"`
	Password     string          `xml:"password,attr"`
	ReadHosts    []*CoreReadHost `xml:"readHost,allowempty,omitempty"`
	Weight       *string         `xml:"weight,attr,omitempty"`
	Id           *string         `xml:"id,attr,omitempty"`
	UsingDecrypt *string         `xml:"usingDecrypt,attr,omitempty"`
	Disabled     *string         `xml:"disabled,attr,omitempty"`
}

func (c *CoreWriteHost) toReadHost() *CoreReadHost {
	return &CoreReadHost{
		Host:         c.Host,
		Url:          c.Url,
		User:         c.User,
		Password:     c.Password,
		UsingDecrypt: c.UsingDecrypt,
		Id:           c.Id,
		Disabled:     c.Disabled,
	}
}

type CoreReadHost struct {
	XMLName      xml.Name `xml:"readHost"`
	Host         string   `xml:"host,attr"`
	Url          string   `xml:"url,attr"`
	User         string   `xml:"user,attr"`
	Password     string   `xml:"password,attr"`
	Weight       *string  `xml:"weight,attr,omitempty"`
	Id           *string  `xml:"id,attr,omitempty"`
	UsingDecrypt *string  `xml:"usingDecrypt,attr,omitempty"`
	Disabled     *string  `xml:"disabled,attr,omitempty"`
}
