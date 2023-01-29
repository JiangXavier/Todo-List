package serializer

//基础序列化器
type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}

//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}


//BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}