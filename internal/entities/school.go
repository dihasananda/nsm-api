package entities

type School struct {
	NSM           string `json:"nsm"`
	Name          string `json:"name"`
	SubDistrictID string `json:"sub_district_id"`
	Address       string `json:"address"`
	SchoolType    string `json:"school_type"`
}
