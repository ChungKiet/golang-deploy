package request

type GetChemistryReq struct {
	TypeChemical string `json:"typeChemical" form:"typeChemical"`
	GroupName    string `json:"groupName" form:"groupName"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type GetRefDocument struct {
	Type string `json:"type"`
}
