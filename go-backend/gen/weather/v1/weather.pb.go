// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: weather/v1/weather.proto

package weatherv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeatherProvider int32

const (
	WeatherProvider_WEATHER_PROVIDER_UNSPECIFIED WeatherProvider = 0
	WeatherProvider_WEATHER_PROVIDER_OPENWEATHER WeatherProvider = 1
	WeatherProvider_WEATHER_PROVIDER_WEATHERAPI  WeatherProvider = 2
)

// Enum value maps for WeatherProvider.
var (
	WeatherProvider_name = map[int32]string{
		0: "WEATHER_PROVIDER_UNSPECIFIED",
		1: "WEATHER_PROVIDER_OPENWEATHER",
		2: "WEATHER_PROVIDER_WEATHERAPI",
	}
	WeatherProvider_value = map[string]int32{
		"WEATHER_PROVIDER_UNSPECIFIED": 0,
		"WEATHER_PROVIDER_OPENWEATHER": 1,
		"WEATHER_PROVIDER_WEATHERAPI":  2,
	}
)

func (x WeatherProvider) Enum() *WeatherProvider {
	p := new(WeatherProvider)
	*p = x
	return p
}

func (x WeatherProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeatherProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_weather_v1_weather_proto_enumTypes[0].Descriptor()
}

func (WeatherProvider) Type() protoreflect.EnumType {
	return &file_weather_v1_weather_proto_enumTypes[0]
}

func (x WeatherProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeatherProvider.Descriptor instead.
func (WeatherProvider) EnumDescriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

type GetCurrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32         `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32         `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Provider  WeatherProvider `protobuf:"varint,3,opt,name=provider,proto3,enum=weather.v1.WeatherProvider" json:"provider,omitempty"`
}

func (x *GetCurrentRequest) Reset() {
	*x = GetCurrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentRequest) ProtoMessage() {}

func (x *GetCurrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrentRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GetCurrentRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GetCurrentRequest) GetProvider() WeatherProvider {
	if x != nil {
		return x.Provider
	}
	return WeatherProvider_WEATHER_PROVIDER_UNSPECIFIED
}

type GetCurrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    float32 `protobuf:"fixed32,2,opt,name=humidity,proto3" json:"humidity,omitempty"`
	UvIndex     float32 `protobuf:"fixed32,3,opt,name=uv_index,json=uvIndex,proto3" json:"uv_index,omitempty"`
	Visibility  int32   `protobuf:"varint,4,opt,name=visibility,proto3" json:"visibility,omitempty"`
	WindSpeed   float32 `protobuf:"fixed32,5,opt,name=wind_speed,json=windSpeed,proto3" json:"wind_speed,omitempty"`
}

func (x *GetCurrentResponse) Reset() {
	*x = GetCurrentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentResponse) ProtoMessage() {}

func (x *GetCurrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GetCurrentResponse) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *GetCurrentResponse) GetUvIndex() float32 {
	if x != nil {
		return x.UvIndex
	}
	return 0
}

func (x *GetCurrentResponse) GetVisibility() int32 {
	if x != nil {
		return x.Visibility
	}
	return 0
}

func (x *GetCurrentResponse) GetWindSpeed() float32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

type GetForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32         `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32         `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Provider  WeatherProvider `protobuf:"varint,3,opt,name=provider,proto3,enum=weather.v1.WeatherProvider" json:"provider,omitempty"`
}

func (x *GetForecastRequest) Reset() {
	*x = GetForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForecastRequest) ProtoMessage() {}

func (x *GetForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForecastRequest.ProtoReflect.Descriptor instead.
func (*GetForecastRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{2}
}

func (x *GetForecastRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GetForecastRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GetForecastRequest) GetProvider() WeatherProvider {
	if x != nil {
		return x.Provider
	}
	return WeatherProvider_WEATHER_PROVIDER_UNSPECIFIED
}

type GetForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *GetForecastResponse) Reset() {
	*x = GetForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForecastResponse) ProtoMessage() {}

func (x *GetForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForecastResponse.ProtoReflect.Descriptor instead.
func (*GetForecastResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{3}
}

func (x *GetForecastResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_weather_v1_weather_proto protoreflect.FileDescriptor

var file_weather_v1_weather_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22,
	0xac, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x76, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x75, 0x76, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x87,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2a, 0x76, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x45, 0x41, 0x54, 0x48, 0x45, 0x52, 0x5f,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x45, 0x41, 0x54, 0x48, 0x45,
	0x52, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x57,
	0x45, 0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x45, 0x41, 0x54,
	0x48, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x57, 0x45, 0x41,
	0x54, 0x48, 0x45, 0x52, 0x41, 0x50, 0x49, 0x10, 0x02, 0x32, 0xb1, 0x01, 0x0a, 0x0e, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa0, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x65,
	0x72, 0x6b, 0x75, 0x6e, 0x64, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2d, 0x6d, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_v1_weather_proto_rawDescOnce sync.Once
	file_weather_v1_weather_proto_rawDescData = file_weather_v1_weather_proto_rawDesc
)

func file_weather_v1_weather_proto_rawDescGZIP() []byte {
	file_weather_v1_weather_proto_rawDescOnce.Do(func() {
		file_weather_v1_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_v1_weather_proto_rawDescData)
	})
	return file_weather_v1_weather_proto_rawDescData
}

var file_weather_v1_weather_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_weather_v1_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weather_v1_weather_proto_goTypes = []any{
	(WeatherProvider)(0),        // 0: weather.v1.WeatherProvider
	(*GetCurrentRequest)(nil),   // 1: weather.v1.GetCurrentRequest
	(*GetCurrentResponse)(nil),  // 2: weather.v1.GetCurrentResponse
	(*GetForecastRequest)(nil),  // 3: weather.v1.GetForecastRequest
	(*GetForecastResponse)(nil), // 4: weather.v1.GetForecastResponse
}
var file_weather_v1_weather_proto_depIdxs = []int32{
	0, // 0: weather.v1.GetCurrentRequest.provider:type_name -> weather.v1.WeatherProvider
	0, // 1: weather.v1.GetForecastRequest.provider:type_name -> weather.v1.WeatherProvider
	1, // 2: weather.v1.WeatherService.GetCurrent:input_type -> weather.v1.GetCurrentRequest
	3, // 3: weather.v1.WeatherService.GetForecast:input_type -> weather.v1.GetForecastRequest
	2, // 4: weather.v1.WeatherService.GetCurrent:output_type -> weather.v1.GetCurrentResponse
	4, // 5: weather.v1.WeatherService.GetForecast:output_type -> weather.v1.GetForecastResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weather_v1_weather_proto_init() }
func file_weather_v1_weather_proto_init() {
	if File_weather_v1_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_v1_weather_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetCurrentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_v1_weather_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetCurrentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_v1_weather_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetForecastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_v1_weather_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetForecastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weather_v1_weather_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_v1_weather_proto_goTypes,
		DependencyIndexes: file_weather_v1_weather_proto_depIdxs,
		EnumInfos:         file_weather_v1_weather_proto_enumTypes,
		MessageInfos:      file_weather_v1_weather_proto_msgTypes,
	}.Build()
	File_weather_v1_weather_proto = out.File
	file_weather_v1_weather_proto_rawDesc = nil
	file_weather_v1_weather_proto_goTypes = nil
	file_weather_v1_weather_proto_depIdxs = nil
}
