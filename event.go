package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableEvent defines the name of the table withing the JSONv2 web service to interface with
// SNOW IM Events
const TableEvent = "u_im_event"

// GetEvents method will take a url.Value type argument and call the GetRecordsFor method with
// the cmdb_ci table and query as the arguments, then format the response into a list of CMDBItem types
func (c Client) GetEvents(query url.Values) ([]Event, error) {
	var res struct {
		Records []Event
	}
	err := c.GetRecordsFor(TableEvent, query, &res)
	return res.Records, err
}

// Event is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type Event struct {
	UReportedByGroupNew struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_reported_by_group_new"`
	USnowChangeRequestNumber string `json:"u_snow_change_request_number__"`
	UIncidentOverview        string `json:"u_incident_overview"`
	URcaOwnerNew             struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_rca_owner_new"`
	SysUpdatedOn              SNTime      `json:"sys_updated_on"`
	UOpenDuration             SNTime      `json:"u_open_duration"`
	UReportedBy               string      `json:"u_reported_by"`
	UMethodOfEscalation       string      `json:"u_method_of_escalation"`
	UBusinessImpact           string      `json:"u_business_impact"`
	UIncidentManager          string      `json:"u_incident_manager"`
	UImNotification           SNTime      `json:"u_im_notification"`
	UClosed                   SNTime      `json:"u_closed"`
	UFirstCustomerFacingComm  SNTime      `json:"u_first_customer_facing_commun"`
	SysCreatedBy              string      `json:"sys_created_by"`
	UCustomerTemplate         string      `json:"u_customer_template"`
	URootCauseExplanation     string      `json:"u_root_cause_explanation"`
	UCommMetric               bool        `json:"u_comm_metric,string"`
	URcaStartDate             SNTime      `json:"u_rca_start_date"`
	UIncidentRecap            string      `json:"u_incident_recap"`
	UChatVolume               string      `json:"u_chat_volume"`
	UNumberOfSubscribers      json.Number `json:"u_number_of_subscribers"`
	UCirUploaded              string      `json:"u_cir_uploaded"`
	UServersDown              string      `json:"u_servers_down"`
	UTemplate                 string      `json:"u_template"`
	UNextUpdateDue            string      `json:"u_next_update_due"`
	UOverrideEntireCompany    bool        `json:"u_override__entire_company_,string"`
	URcaCompleted             SNTime      `json:"u_rca_completed"`
	UIncidentLevel            string      `json:"u_incident_level"`
	UImpactType               string      `json:"u_impact_type"`
	USendToAlert              bool        `json:"u_send_to_alert,string"`
	UAdditionalPocs           string      `json:"u_additional_pocs"`
	URackerTemplateTwo        string      `json:"u_racker_template_two"`
	UPlatform                 string      `json:"u_platform"`
	UEventUpdates             string      `json:"u_event_updates"`
	UImpactedSystemApp        string      `json:"u_impacted_system_app"`
	USendIncidentRecap        bool        `json:"u_send_incident_recap,string"`
	UTechnicalImpact          string      `json:"u_technical_impact"`
	UAllUsaRackers            bool        `json:"u_all_usa_rackers,string"`
	UOutageUpdate             string      `json:"u_outage_update"`
	URackerInternalComments   string      `json:"u_racker_internal_comments"`
	UReturnToService          SNTime      `json:"u_return_to_service"`
	URootCausePreventionActio string      `json:"u_root_cause_prevention__actio"`
	UNetNumber                string      `json:"u_net_number"`
	UPhoneVolume              string      `json:"u_phone_volume"`
	UImpactDistribution       string      `json:"u_impact_distribution"`
	UIncidentTimeline         string      `json:"u_incident_timeline"`
	UCustomersImpacted        string      `json:"u_customers_impacted"`
	USendAfterActionReport    bool        `json:"u_send_after_action_report_,string"`
	UCommentList              string      `json:"u_comment_list"`
	UImpactedSystem           string      `json:"u_impacted_system"`
	UTicketVolume             string      `json:"u_ticket_volume"`
	UIncidentSeverityNumber   json.Number `json:"u_incident_severity_number"`
	UHTMLComment              string      `json:"u_html_comment"`
	USubRootCauseCode         string      `json:"u_sub_root_cause_code"`
	UDevicesImpacted          string      `json:"u_devices_impacted"`
	UImNotificationDuration   SNTime      `json:"u_im_notification_duration"`
	UEventLeader              string      `json:"u_event_leader"`
	URootCauseCode            string      `json:"u_root_cause_code"`
	URootCauseSystem          struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_root_cause_system"`
	UCirMetric                bool   `json:"u_cir_metric,string"`
	UCustomerApprovedComments string `json:"u_customer_approved_comments"`
	UWorknotes                string `json:"u_worknotes"`
	SysUpdatedBy              string `json:"sys_updated_by"`
	UIncidentNumber           string `json:"u_incident_number"`
	SysCreatedOn              SNTime `json:"sys_created_on"`
	UImpactDuration           SNTime `json:"u_impact_duration"`
	UActive                   bool   `json:"u_active,string"`
	UIncidentSummary          string `json:"u_incident_summary"`
	UOpenedBy                 struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_opened_by"`
	UAdditionalInformation  string      `json:"u_additional_information"`
	UInitialEmailRecipients string      `json:"u_initial_email_recipients"`
	ULessonsLearned         string      `json:"u_lessons_learned"`
	UOtherKeyCommentary     string      `json:"u_other_key_commentary"`
	UIncidentStatus         json.Number `json:"u_incident_status"`
	UDl                     string      `json:"u_dl"`
	UChangeControl          string      `json:"u_change_control"`
	UNewSubscribers         string      `json:"u_new_subscribers"`
	UEventImpact            string      `json:"u_event_impact"`
	UPrivate                string      `json:"u_private"`
	UIncidentDuration       SNTime      `json:"u_incident_duration"`
	URootCauseChangeRelated bool        `json:"u_root_cause_change_related_,string"`
	UProblem                string      `json:"u_problem"`
	URcaOwner               struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_rca_owner"`
	USubscribers    string `json:"u_subscribers"`
	UIncidentOrigin string `json:"u_incident_origin"`
	UClosedBy       struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_closed_by"`
	SysID                  string      `json:"sys_id"`
	UCreateEtherpadPage    string      `json:"u_create_etherpad_page"`
	URootCauseCodeNew      string      `json:"u_root_cause_code_new"`
	UOs                    string      `json:"u_os"`
	UIncidentEnd           SNTime      `json:"u_incident_end"`
	UReasonForOther        string      `json:"u_reason_for_other"`
	UChangeControlNumber   string      `json:"u_change_control_number"`
	UIncidentTitle         string      `json:"u_incident_title"`
	UPoc                   string      `json:"u_poc"`
	SysModCount            json.Number `json:"sys_mod_count"`
	UImpact                string      `json:"u_impact"`
	SysTags                string      `json:"sys_tags"`
	UIncidentSeverityLevel string      `json:"u_incident_severity_level"`
	UNextAction            string      `json:"u_next_action"`
	UTeamsEngaged          string      `json:"u_teams_engaged"`
	UIncidentStart         SNTime      `json:"u_incident_start"`
	UIncidentAlert         bool        `json:"u_incident_alert,string"`
	UOrganization          string      `json:"u_organization"`
	UTeamsEngagedBcc       string      `json:"u_teams_engaged_bcc"`
	URootCauseOrganization struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_root_cause_organization"`
}
