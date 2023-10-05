package control

type Response struct {
	// response code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
}

type HealthRsp struct {
	// response code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	// timestamp
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp" form:"timestamp"`
}
