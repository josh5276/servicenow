package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableCompany defines the name of the table withing the JSONv2 web service to interface with
// SNOW Companies
const TableCompany = "core_company"

// GetCompanies method will take a url.Value type argument and call the GetRecordsFor method with
// the task table and query as the arguments, then format the response into the TaskRequest type
func (c Client) GetCompanies(query url.Values) ([]Company, error) {
	var res struct {
		Records []Company `json:"records"`
	}
	err := c.GetRecordsFor(TableCompany, query, &res)
	return res.Records, err
}

// Company is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type Company struct {
	Active                                 bool        `json:"active,string"`
	AppleIcon                              string      `json:"apple_icon"`
	BannerImage                            string      `json:"banner_image"`
	BannerImageLight                       string      `json:"banner_image_light"`
	BannerText                             string      `json:"banner_text"`
	City                                   string      `json:"city"`
	Contact                                string      `json:"contact"`
	Country                                string      `json:"country"`
	Customer                               bool        `json:"customer,string"`
	Discount                               string      `json:"discount"`
	FaxPhone                               string      `json:"fax_phone"`
	FiscalYear                             string      `json:"fiscal_year"`
	Latitude                               string      `json:"latitude"`
	LatLongError                           string      `json:"lat_long_error"`
	Longitude                              string      `json:"longitude"`
	Manufacturer                           bool        `json:"manufacturer,string"`
	MarketCap                              string      `json:"market_cap"`
	Name                                   string      `json:"name"`
	Notes                                  string      `json:"notes"`
	NumEmployees                           string      `json:"num_employees"`
	Parent                                 string      `json:"parent"`
	Phone                                  string      `json:"phone"`
	Primary                                string      `json:"primary"`
	Profits                                string      `json:"profits"`
	PubliclyTraded                         string      `json:"publicly_traded"`
	RankTier                               string      `json:"rank_tier"`
	RevenuePerYear                         json.Number `json:"revenue_per_year"`
	SsoSource                              string      `json:"sso_source"`
	State                                  string      `json:"state"`
	Status                                 string      `json:"__status"`
	StockPrice                             string      `json:"stock_price"`
	StockSymbol                            string      `json:"stock_symbol"`
	Street                                 string      `json:"street"`
	SysClassName                           string      `json:"sys_class_name"`
	SysCreatedBy                           string      `json:"sys_created_by"`
	SysCreatedOn                           string      `json:"sys_created_on"`
	SysDomain                              string      `json:"sys_domain"`
	SysDomainPath                          string      `json:"sys_domain_path"`
	SysID                                  string      `json:"sys_id"`
	SysModCount                            string      `json:"sys_mod_count"`
	SysTags                                string      `json:"sys_tags"`
	SysUpdatedBy                           string      `json:"sys_updated_by"`
	SysUpdatedOn                           string      `json:"sys_updated_on"`
	UAccountCustomization                  string      `json:"u_account_customization"`
	UAdapt                                 bool        `json:"u_adapt,string"`
	UAxCustomerID                          string      `json:"u_ax_customer_id"`
	UBannerAlert                           string      `json:"u_banner_alert"`
	UBridgedAccount                        bool        `json:"u_bridged_account,string"`
	UChangeManagement                      string      `json:"u_change_management"`
	UCoiExpirationDate                     string      `json:"u_coi_expiration_date"`
	UCoiReceived                           string      `json:"u_coi_received"`
	UCoiRequired                           bool        `json:"u_coi_required,string"`
	UCompanyLogo                           string      `json:"u_company_logo"`
	UCorrelationid                         string      `json:"u_correlationid"`
	UCustomerID                            string      `json:"u_customer_id"`
	UDPAccountManager                      string      `json:"u_dp_account_manager"`
	UDPSalesEngineer                       string      `json:"u_dp_sales_engineer"`
	UDPSecondaryPm                         string      `json:"u_dp_secondary_pm"`
	UEinSs                                 string      `json:"u_ein_ss"`
	UEm7Monitoring                         bool        `json:"u_em7_monitoring,string"`
	UEntity                                bool        `json:"u_entity,string"`
	UErpCustomerID                         string      `json:"u_erp_customer_id"`
	UHasChangeControl                      bool        `json:"u_has_change_control,string"`
	UIsInternational                       bool        `json:"u_is_international,string"`
	UIsPci                                 bool        `json:"u_is_pci,string"`
	UManagedRep                            string      `json:"u_managed_rep"`
	UManagementLevel                       string      `json:"u_mangement_level"`
	UMfa                                   bool        `json:"u_mfa,string"`
	UObserverMonitoring                    bool        `json:"u_observer_monitoring,string"`
	UPotentialAccessToDatapipeCustomerData string      `json:"u_potential_access_to_datapipe_customer_data"`
	UPuppet                                bool        `json:"u_puppet,string"`
	USalesRep                              string      `json:"u_sales_rep"`
	UTaxID                                 string      `json:"u_tax_id"`
	UUseLogoForCustomerPortal              bool        `json:"u_use_logo_for_customer_portal,string"`
	UW9OnFile                              bool        `json:"u_w_9_on_file,string"`
	Vendor                                 bool        `json:"vendor,string"`
	VendorManager                          string      `json:"vendor_manager"`
	VendorType                             string      `json:"vendor_type"`
	Website                                string      `json:"website"`
	Zip                                    string      `json:"zip"`
}
