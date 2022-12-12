package request

type GetChemistryReq struct {
	TypeMaterial string `json:"typeMaterial" form:"typeMaterial"`
	TypeSpectrum string `json:"typeSpectrum" form:"typeSpectrum"`
	Chemical     string `json:"chemical" form:"chemical"`
}

type GetRefDocument struct {
	Type string `json:"type"`
}
