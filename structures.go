
type LeadCRM struct {
	Result []struct {
		ID                  string      `json:"ID"`
		Title               string      `json:"TITLE"`
		Honorific           interface{} `json:"HONORIFIC"`
		Name                string      `json:"NAME"`
		SecondName          interface{} `json:"SECOND_NAME"`
		LastName            string      `json:"LAST_NAME"`
		CompanyTitle        interface{} `json:"COMPANY_TITLE"`
		CompanyID           interface{} `json:"COMPANY_ID"`
		ContactID           string      `json:"CONTACT_ID"`
		IsReturnCustomer    string      `json:"IS_RETURN_CUSTOMER"`
		Birthdate           string      `json:"BIRTHDATE"`
		SourceID            string      `json:"SOURCE_ID"`
		SourceDescription   interface{} `json:"SOURCE_DESCRIPTION"`
		StatusID            string      `json:"STATUS_ID"`
		StatusDescription   interface{} `json:"STATUS_DESCRIPTION"`
		Post                interface{} `json:"POST"`
		Comments            string      `json:"COMMENTS"`
		CurrencyID          string      `json:"CURRENCY_ID"`
		Opportunity         string      `json:"OPPORTUNITY"`
		IsManualOpportunity string      `json:"IS_MANUAL_OPPORTUNITY"`
		HasPhone            string      `json:"HAS_PHONE"`
		HasEmail            string      `json:"HAS_EMAIL"`
		HasImol             string      `json:"HAS_IMOL"`
		AssignedByID        string      `json:"ASSIGNED_BY_ID"`
		CreatedByID         string      `json:"CREATED_BY_ID"`
		ModifyByID          string      `json:"MODIFY_BY_ID"`
		DateCreate          time.Time   `json:"DATE_CREATE"`
		DateModify          time.Time   `json:"DATE_MODIFY"`
		DateClosed          time.Time   `json:"DATE_CLOSED"`
		StatusSemanticID    string      `json:"STATUS_SEMANTIC_ID"`
		Opened              string      `json:"OPENED"`
		OriginatorID        interface{} `json:"ORIGINATOR_ID"`
		OriginID            interface{} `json:"ORIGIN_ID"`
		MovedByID           string      `json:"MOVED_BY_ID"`
		MovedTime           time.Time   `json:"MOVED_TIME"`
		Address             interface{} `json:"ADDRESS"`
		Address2            interface{} `json:"ADDRESS_2"`
		AddressCity         interface{} `json:"ADDRESS_CITY"`
		AddressPostalCode   interface{} `json:"ADDRESS_POSTAL_CODE"`
		AddressRegion       interface{} `json:"ADDRESS_REGION"`
		AddressProvince     interface{} `json:"ADDRESS_PROVINCE"`
		AddressCountry      interface{} `json:"ADDRESS_COUNTRY"`
		AddressCountryCode  interface{} `json:"ADDRESS_COUNTRY_CODE"`
		AddressLocAddrID    interface{} `json:"ADDRESS_LOC_ADDR_ID"`
		UtmSource           interface{} `json:"UTM_SOURCE"`
		UtmMedium           interface{} `json:"UTM_MEDIUM"`
		UtmCampaign         interface{} `json:"UTM_CAMPAIGN"`
		UtmContent          interface{} `json:"UTM_CONTENT"`
		UtmTerm             interface{} `json:"UTM_TERM"`
	} `json:"result"`
	Next  int `json:"next"`
	Total int `json:"total"`
	Time  struct {
		Start            float64   `json:"start"`
		Finish           float64   `json:"finish"`
		Duration         float64   `json:"duration"`
		Processing       float64   `json:"processing"`
		DateStart        time.Time `json:"date_start"`
		DateFinish       time.Time `json:"date_finish"`
		OperatingResetAt int       `json:"operating_reset_at"`
		Operating        int       `json:"operating"`
	} `json:"time"`
}
type PhoneCRM struct {
	Result []struct {
		ID    string `json:"ID"`
		Phone []struct {
			ID        string `json:"ID"`
			ValueType string `json:"VALUE_TYPE"`
			Value     string `json:"VALUE"`
			TypeID    string `json:"TYPE_ID"`
		} `json:"PHONE,omitempty"`
	} `json:"result"`
	Next  int `json:"next"`
	Total int `json:"total"`
	Time  struct {
		Start            float64   `json:"start"`
		Finish           float64   `json:"finish"`
		Duration         float64   `json:"duration"`
		Processing       float64   `json:"processing"`
		DateStart        time.Time `json:"date_start"`
		DateFinish       time.Time `json:"date_finish"`
		OperatingResetAt int       `json:"operating_reset_at"`
		Operating        int       `json:"operating"`
	} `json:"time"`
}
