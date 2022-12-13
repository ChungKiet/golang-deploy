package request

type GetChemistryReq struct {
	TypeChemical string `json:"typeChemical" form:"typeChemical"`
	GroupName    string `json:"groupName" form:"groupName"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type GetRefDocument struct {
	Type string `json:"type" form:"type"`
}

// write func post, put, delete to use in backend
type GetMenu struct {
	TypeChemical string `json:"typeChemical,omitempty" form:"typeChemical"`
	GroupName    string `json:"groupName,omitempty" form:"groupName"`
	Chemical     string `json:"chemical,omitempty" form:"chemical"`
}
