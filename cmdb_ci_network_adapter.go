package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableCMDBNetworkAdapter defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCMDBNetworkAdapter = "cmdb_ci_network_adapter"

// GetCMDBNetworkAdapters method will take a url.Value type argument and call the GetRecordsFor method with
// the cmdb_ci table and query as the arguments, then format the response into a list of CMDBItem types
func (c Client) GetCMDBNetworkAdapters(query url.Values) ([]NetworkAdapter, error) {
	var res struct {
		Records []NetworkAdapter
	}
	err := c.GetRecordsFor(TableCMDBNetworkAdapter, query, &res)
	return res.Records, err
}

// CMDBServer is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type NetworkAdapter struct {
	Status            string      `json:"__status"`
	CanPrint          bool        `json:"can_print,string"`
	Category          string      `json:"category"`
	CMDBCi            string      `json:"cmdb_ci"`
	CostCC            string      `json:"cost_cc"`
	DHCPEnabled       bool        `json:"dhcp_enabled,string"`
	FaultCount        json.Number `json:"fault_count"`
	FirstDiscovered   SNTime      `json:"first_discovered"`
	InstallStatus     json.Number `json:"install_status"`
	IPAddress         string      `json:"ip_address"`
	LastDiscovered    SNTime      `json:"last_discovered"`
	LeaseID           string      `json:"lease_id"`
	MacAddress        string      `json:"mac_address"`
	Monitor           bool        `json:"monitor,string"`
	Name              string      `json:"name"`
	Netmask           string      `json:"netmask"`
	OperationalStatus json.Number `json:"operational_status"`
	SkipSync          bool        `json:"skip_sync,string"`
	SubCategory       string      `json:"subcategory"`
	SysClassName      string      `json:"sys_class_name"`
	SysClassPath      string      `json:"sys_class_path"`
	SysCreatedBy      string      `json:"sys_created_by"`
	SysCreatedOn      SNTime      `json:"sys_created_on"`
	SysDomain         string      `json:"sys_domain"`
	SysDomainPath     string      `json:"sys_domain_path"`
	SysID             string      `json:"sys_id"`
	SysModCount       json.Number `json:"sys_mod_count"`
	SysUpdatedBy      string      `json:"sys_updated_by"`
	SysUpdatedOn      SNTime      `json:"sys_updated_on"`
	UIsCriticalServer bool        `json:"u_is_critical_server,string"`
	Unverified        bool        `json:"unverified,string"`
	Virtual           bool        `json:"virtual,string"`
}
