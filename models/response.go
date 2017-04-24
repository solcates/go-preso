package models

type RestResponse struct {
	Message    string
	StatusCode int64
	ErrorCode  int64
	Data       interface{}
	Items      []interface{}
	Success    bool
}

func (f *RestResponse) init() {
	//if f.Success == nil {
	f.Success = true
	//}
}
