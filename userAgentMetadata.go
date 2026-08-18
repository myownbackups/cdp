package cdp

type Brand struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}
type UserAgentData struct {
	Brands          []Brand  `json:"brands,omitempty"`
	FullVersionList []Brand  `json:"fullVersionList,omitempty"`
	Platform        string   `json:"platform"`
	PlatformVersion string   `json:"platformVersion"`
	Architecture    string   `json:"architecture"`
	UaFullVersion   string   `json:"uaFullVersion"`
	Model           string   `json:"model"`
	Mobile          bool     `json:"mobile"`
	Bitness         string   `json:"bitness,omitempty"`
	Wow64           bool     `json:"wow64,omitempty"`
	FormFactors     []string `json:"formFactors,omitempty"`
}
