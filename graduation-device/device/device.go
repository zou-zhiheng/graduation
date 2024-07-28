package device

type Device struct {
	Code       string  `json:"code" bson:"code"`
	Data       []*Data `json:"data" bson:"data"`
	CreateTime string  `json:"createTime" bson:"createTime"`
}

type Data struct {
	Key   string      `json:"key" bson:"key"`
	Value interface{} `json:"value" bson:"value"`
	Unit  string      `json:"unit" bson:"unit"`
}

type JsonData struct {
	Data interface{} `json:"data" bson:"data"`
}
