package servicenow

import "net/url"

// GetUsers returns all users that match a query.
func (c Client) GetUsers(query url.Values) ([]User, error) {
	var m struct{ Records []User }
	err := c.GetRecordsFor("sys_user", query, &m)
	return m.Records, err
}

type User struct {
	CalendarIntegration     string `json:"calendar_integration"`
	Country                 string `json:"country"`
	UserPassword            string `json:"user_password"`
	LastLoginTime           string `json:"last_login_time"`
	Source                  string `json:"source"`
	SysUpdatedOn            string `json:"sys_updated_on"`
	Building                string `json:"building"`
	WebServiceAccessOnly    string `json:"web_service_access_only"`
	UIseureka               string `json:"u_iseureka"`
	Notification            string `json:"notification"`
	ManagedDomain           string `json:"managed_domain"`
	EnableMultifactorAuthn  string `json:"enable_multifactor_authn"`
	SysUpdatedBy            string `json:"sys_updated_by"`
	SsoSource               string `json:"sso_source"`
	SysCreatedOn            string `json:"sys_created_on"`
	UPmaNotifications       string `json:"u_pma_notifications"`
	SysDomain               string `json:"sys_domain"`
	UIsduo                  string `json:"u_isduo"`
	State                   string `json:"state"`
	Vip                     string `json:"vip"`
	SysCreatedBy            string `json:"sys_created_by"`
	UEnabeledMfa            string `json:"u_enabeled_mfa"`
	Zip                     string `json:"zip"`
	HomePhone               string `json:"home_phone"`
	TimeFormat              string `json:"time_format"`
	LastLogin               string `json:"last_login"`
	USecEmail               string `json:"u_sec_email"`
	Active                  string `json:"active"`
	USignature              string `json:"u_signature"`
	SysDomainPath           string `json:"sys_domain_path"`
	Phone                   string `json:"phone"`
	Name                    string `json:"name"`
	EmployeeNumber          string `json:"employee_number"`
	UCommercialAuthoriser   string `json:"u_commercial_authoriser"`
	Gender                  string `json:"gender"`
	City                    string `json:"city"`
	FailedAttempts          string `json:"failed_attempts"`
	UserName                string `json:"user_name"`
	Title                   string `json:"title"`
	SysClassName            string `json:"sys_class_name"`
	SysID                   string `json:"sys_id"`
	InternalIntegrationUser string `json:"internal_integration_user"`
	MobilePhone             string `json:"mobile_phone"`
	Street                  string `json:"street"`
	UPin                    string `json:"u_pin"`
	Company                 string `json:"company"`
	Department              string `json:"department"`
	FirstName               string `json:"first_name"`
	Email                   string `json:"email"`
	Introduction            string `json:"introduction"`
	PreferredLanguage       string `json:"preferred_language"`
	UP1Notifications        string `json:"u_p1_notifications"`
	Manager                 string `json:"manager"`
	LockedOut               string `json:"locked_out"`
	SysModCount             string `json:"sys_mod_count"`
	LastName                string `json:"last_name"`
	Photo                   string `json:"photo"`
	UPhone2                 string `json:"u_phone2"`
	MiddleName              string `json:"middle_name"`
	SysTags                 string `json:"sys_tags"`
	TimeZone                string `json:"time_zone"`
	Schedule                string `json:"schedule"`
	URfcApprover            string `json:"u_rfc_approver"`
	DateFormat              string `json:"date_format"`
	Location                string `json:"location"`
}
