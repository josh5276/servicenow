package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableCMDB defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCMDB = "cmdb_ci"

// GetCMDBItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmdb_ci table and query as the arguments, then format the response into a list of CMDBItem types
func (c Client) GetCMDBItems(query url.Values) ([]CMDBItem, error) {
	var res struct {
		Records []CMDBItem
	}
	err := c.GetRecordsFor(TableCMDB, query, &res)
	return res.Records, err
}

// CMDBItem is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type CMDBItem struct {
	Status                    string      `json:"__status"`
	Asset                     string      `json:"asset"`
	AssetTag                  string      `json:"asset_tag"`
	Assigned                  string      `json:"assigned"`
	AssignedTo                string      `json:"assigned_to"`
	AssignmentGroup           string      `json:"assignment_group"`
	Attributes                string      `json:"attributes"`
	CanPrint                  string      `json:"can_print"`
	Category                  string      `json:"category"`
	CDRom                     bool        `json:"cd_rom,string"`
	CDSpeed                   string      `json:"cd_speed"`
	ChangeControl             string      `json:"change_control"`
	ChassisType               string      `json:"chassis_type"`
	CheckedIn                 string      `json:"checked_in"`
	CheckedOut                string      `json:"checked_out"`
	Classification            string      `json:"classification"`
	Comments                  string      `json:"comments"`
	CorrelationID             string      `json:"correlation_id"`
	Cost                      string      `json:"cost"`
	CostCC                    string      `json:"cost_cc"`
	CPUCoreCount              string      `json:"cpu_core_count"`
	CPUCoreThread             string      `json:"cpu_core_thread"`
	CPUCount                  string      `json:"cpu_count"`
	CPUName                   string      `json:"cpu_name"`
	CPUSpeed                  string      `json:"cpu_speed"`
	CPUType                   string      `json:"cpu_type"`
	DefaultGateway            string      `json:"default_gateway"`
	DeliveryDate              string      `json:"delivery_date"`
	Department                string      `json:"department"`
	DiscoverySource           string      `json:"discovery_source"`
	DiskSpace                 string      `json:"disk_space"`
	DNSDomain                 string      `json:"dns_domain"`
	DRBackup                  string      `json:"dr_backup"`
	Due                       string      `json:"due"`
	DueIn                     string      `json:"due_in"`
	FaultCount                json.Number `json:"fault_count"`
	FirewallStatus            string      `json:"firewall_status"`
	FirstDiscovered           SNTime      `json:"first_discovered"`
	Floppy                    string      `json:"floppy"`
	FormFactor                string      `json:"form_factor"`
	FQDN                      string      `json:"fqdn"`
	GLAccount                 string      `json:"gl_account"`
	HardwareStatus            string      `json:"hardware_status"`
	HardwareSubStatus         string      `json:"hardware_substatus"`
	Hostname                  string      `json:"host_name"`
	InstallDate               string      `json:"install_date"`
	InstallStatus             json.Number `json:"install_status"`
	InvoiceNumber             string      `json:"invoice_number"`
	IPAddress                 string      `json:"ip_address"`
	Justification             string      `json:"justification"`
	LastDiscovered            SNTime      `json:"last_discovered"`
	LeaseID                   string      `json:"lease_id"`
	Location                  string      `json:"location"`
	MacAddress                string      `json:"mac_address"`
	MaintenanceSchedule       string      `json:"maintenance_schedule"`
	ManagedBy                 string      `json:"managed_by"`
	ModelNumber               string      `json:"model_number"`
	Monitor                   bool        `json:"monitor,string"`
	Name                      string      `json:"name"`
	ObjectID                  string      `json:"object_id"`
	OperationalStatus         string      `json:"operational_status"`
	OrderDate                 string      `json:"order_date"`
	OS                        string      `json:"os"`
	OSAddressWidth            string      `json:"os_address_width"`
	OSDomain                  string      `json:"os_domain"`
	OSServicePack             string      `json:"os_service_pack"`
	OSVersion                 string      `json:"os_version"`
	OwnedBy                   string      `json:"owned_by"`
	PONumber                  string      `json:"po_number"`
	PurchaseDate              string      `json:"purchase_date"`
	RAM                       string      `json:"ram"`
	Schedule                  string      `json:"schedule"`
	SerialNumber              string      `json:"serial_number"`
	ShortDescription          string      `json:"short_description"`
	SkipSync                  bool        `json:"skip_sync,string"`
	StartDate                 string      `json:"start_date"`
	Subcategory               string      `json:"subcategory"`
	SupportGroup              string      `json:"support_group"`
	SupportedBy               string      `json:"supported_by"`
	SysClassName              string      `json:"sys_class_name"`
	SysClassPath              string      `json:"sys_class_path"`
	SysCreatedBy              string      `json:"sys_created_by"`
	SysCreatedOn              SNTime      `json:"sys_created_on"`
	SysDomain                 string      `json:"sys_domain"`
	SysDomainPath             string      `json:"sys_domain_path"`
	SysID                     string      `json:"sys_id"`
	SysModCount               json.Number `json:"sys_mod_count"`
	SysTags                   string      `json:"sys_tags"`
	USysUpdatedBy             string      `json:"sys_updated_by"`
	USysUpdatedOn             SNTime      `json:"sys_updated_on"`
	UAccountID                json.Number `json:"u_account_id"`
	UAccountNumber            string      `json:"u_account_number"`
	UCITags                   string      `json:"u_ci_tags"`
	UDeviceID                 string      `json:"u_device_id"`
	UHardwareType             string      `json:"u_hardware_type"`
	UIsCriticalServer         bool        `json:"u_is_critical_server,string"`
	ULegacyDeviceID           json.Number `json:"u_legacy_device_id"`
	UManagementResponsibility string      `json:"u_management_responsibility"`
	UNickname                 string      `json:"u_nickname"`
	UStatus                   string      `json:"u_status"`
	UnVerified                bool        `json:"unverified,string"`
	UsedFor                   string      `json:"used_for"`
	Virtual                   bool        `json:"virtual,string"`
	WarrantyExpiration        string      `json:"warranty_expiration"`
}
