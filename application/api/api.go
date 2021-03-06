package api

type (
	// Response struct for REST API
	Response struct {
		Response Rest        `json:"response"`
		Data     interface{} `json:"data,omitempty"`
	}
	// Rest struct for response Res API
	Rest struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
	// Datas struct for data
	Datas struct {
		RmID      string `json:"rm_id,omitempty"`
		RmName    string `json:"rm_name,omitempty"`
		RmPlace   string `json:"rm_place,omitempty"`
		RmSumpart string `json:"rm_sumpart,omitempty"`
		RmPrice   string `json:"rm_price,omitempty"`
		RmStatus  string `json:"rm_status,omitempty"`
	}
)
