package request

type Api struct {
	Page       int                 `json:"page"`
	Size int                 `json:"size"`
	Sn string                 `json:"sn"`
}

type DataInfo struct {
	Data       []string                 `json:"data"`
}