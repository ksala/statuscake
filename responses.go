package statuscake

type autheticationErrorResponse struct {
	ErrNo int
	Error string
}

type updateResponse struct {
	Issues   interface{} `json:"Issues"`
	Success  bool        `json:"Success"`
	Message  string      `json:"Message"`
	InsertID int         `json:"InsertID"`
}

type deleteResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"Error"`
}

type detailResponse struct {
	Method          string   `json:"Method"`
	TestID          int      `json:"TestID"`
	TestType        string   `json:"TestType"`
	Paused          bool     `json:"Paused"`
	WebsiteName     string   `json:"WebsiteName"`
	URI             string   `json:"URI"`
	ContactID       int      `json:"ContactID"`
	Status          string   `json:"Status"`
	Tags            []string `json:"Tags"`
	Uptime          float64  `json:"Uptime"`
	CustomHeader    string   `json:"CustomHeader"`
	CheckRate       int      `json:"CheckRate"`
	Timeout         int      `json:"Timeout"`
	LogoImage       string   `json:"LogoImage"`
	Confirmation    int      `json:"Confirmation,string"`
	WebsiteHost     string   `json:"WebsiteHost"`
	NodeLocations   []string `json:"NodeLocations"`
	FindString      string   `json:"FindString"`
	DoNotFind       bool     `json:"DoNotFind"`
	LastTested      string   `json:"LastTested"`
	NextLocation    string   `json:"NextLocation"`
	Port            int      `json:"Port"`
	Processing      bool     `json:"Processing"`
	ProcessingState string   `json:"ProcessingState"`
	ProcessingOn    string   `json:"ProcessingOn"`
	DownTimes       int      `json:"DownTimes,string"`
	Sensitive       bool     `json:"Sensitive"`
	TriggerRate     int      `json:"TriggerRate,string"`
	UseJar          bool     `json:"UseJar"`
}

func (d *detailResponse) test() *Test {
	return &Test{
		TestID:        d.TestID,
		TestType:      d.TestType,
		Paused:        d.Paused,
		WebsiteName:   d.WebsiteName,
		WebsiteURL:    d.URI,
		CustomHeader:  d.CustomHeader,
		ContactID:     d.ContactID,
		Status:        d.Status,
		TestTags:      d.Tags,
		Uptime:        d.Uptime,
		CheckRate:     d.CheckRate,
		Timeout:       d.Timeout,
		LogoImage:     d.LogoImage,
		Confirmation:  d.Confirmation,
		WebsiteHost:   d.WebsiteHost,
		NodeLocations: d.NodeLocations,
		FindString:    d.FindString,
		DoNotFind:     d.DoNotFind,
		Port:          d.Port,
		TriggerRate:   d.TriggerRate,
		UseJar:        d.UseJar,
	}
}
