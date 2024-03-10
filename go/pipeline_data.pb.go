// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pipeline_data.proto

package pipeline_data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A single instruction to be executed by a pipeline stage
type CommandStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name provided by the user (typically the first argument)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// All other values provided to the function by the user
	Args string `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	// Location at which the function will be contacted
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CommandStep) Reset() {
	*x = CommandStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandStep) ProtoMessage() {}

func (x *CommandStep) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandStep.ProtoReflect.Descriptor instead.
func (*CommandStep) Descriptor() ([]byte, []int) {
	return file_pipeline_data_proto_rawDescGZIP(), []int{0}
}

func (x *CommandStep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandStep) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *CommandStep) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// A list of commands to be executed by a pipeline stage
type CommandList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for the query
	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	// The list of commands to be executed
	Commands []*CommandStep `protobuf:"bytes,2,rep,name=commands,proto3" json:"commands,omitempty"`
	// The current step to be executed.
	// Recipients should execute the command at this index.
	Step uint32 `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	// The URL at which errors should be reported
	ErrorUrl string `protobuf:"bytes,4,opt,name=error_url,json=errorUrl,proto3" json:"error_url,omitempty"`
}

func (x *CommandList) Reset() {
	*x = CommandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandList) ProtoMessage() {}

func (x *CommandList) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandList.ProtoReflect.Descriptor instead.
func (*CommandList) Descriptor() ([]byte, []int) {
	return file_pipeline_data_proto_rawDescGZIP(), []int{1}
}

func (x *CommandList) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

func (x *CommandList) GetCommands() []*CommandStep {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *CommandList) GetStep() uint32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *CommandList) GetErrorUrl() string {
	if x != nil {
		return x.ErrorUrl
	}
	return ""
}

// A single event to be processed by a pipeline stage.
// Events are the primary data type used in the pipeline.
// Events are typically generated by a source,
// and are processed by a series of stages.
// All fields are optional,
// however, if no fields are present the event should be dropped / ignored.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw log line which generated the event.
	// Typically generated by the detetion system being monitored.
	Raw *string `protobuf:"bytes,1,opt,name=raw,proto3,oneof" json:"raw,omitempty"`
	// The time at which an event was recieved by the event indexer.
	IndexTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=index_time,json=indexTime,proto3,oneof" json:"index_time,omitempty"`
	// The time that the event occured on the source system.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	// The format of event, such as an nginx access log, or a syslog message.
	EventType *string `protobuf:"bytes,4,opt,name=event_type,json=eventType,proto3,oneof" json:"event_type,omitempty"`
	// The category of the event for indexing / organisation.
	// For example, "main_server", "subnet_b_flow_logs", etc.
	Category *string `protobuf:"bytes,5,opt,name=category,proto3,oneof" json:"category,omitempty"`
	// Additional fields which were extracted from the raw event.
	Derived *structpb.Struct `protobuf:"bytes,6,opt,name=derived,proto3,oneof" json:"derived,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pipeline_data_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetRaw() string {
	if x != nil && x.Raw != nil {
		return *x.Raw
	}
	return ""
}

func (x *Event) GetIndexTime() *timestamppb.Timestamp {
	if x != nil {
		return x.IndexTime
	}
	return nil
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Event) GetEventType() string {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return ""
}

func (x *Event) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *Event) GetDerived() *structpb.Struct {
	if x != nil {
		return x.Derived
	}
	return nil
}

// A batch of events to be processed by a pipeline stage.
// Includes instructions regarding how to handle the events.
type BatchProcessingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The commands to be executed.
	Commands *CommandList `protobuf:"bytes,1,opt,name=commands,proto3" json:"commands,omitempty"`
	// A list of events to be processed according to the relevant command.
	Events []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *BatchProcessingEvent) Reset() {
	*x = BatchProcessingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchProcessingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchProcessingEvent) ProtoMessage() {}

func (x *BatchProcessingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchProcessingEvent.ProtoReflect.Descriptor instead.
func (*BatchProcessingEvent) Descriptor() ([]byte, []int) {
	return file_pipeline_data_proto_rawDescGZIP(), []int{3}
}

func (x *BatchProcessingEvent) GetCommands() *CommandList {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *BatchProcessingEvent) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_pipeline_data_proto protoreflect.FileDescriptor

var file_pipeline_data_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x91, 0x01, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x55, 0x72, 0x6c,
	0x22, 0xe7, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x61,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x88, 0x01,
	0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x02, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48,
	0x05, 0x52, 0x07, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x72, 0x61, 0x77, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x22, 0x7c, 0x0a, 0x14, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipeline_data_proto_rawDescOnce sync.Once
	file_pipeline_data_proto_rawDescData = file_pipeline_data_proto_rawDesc
)

func file_pipeline_data_proto_rawDescGZIP() []byte {
	file_pipeline_data_proto_rawDescOnce.Do(func() {
		file_pipeline_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipeline_data_proto_rawDescData)
	})
	return file_pipeline_data_proto_rawDescData
}

var file_pipeline_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pipeline_data_proto_goTypes = []interface{}{
	(*CommandStep)(nil),           // 0: pipeline_data.CommandStep
	(*CommandList)(nil),           // 1: pipeline_data.CommandList
	(*Event)(nil),                 // 2: pipeline_data.Event
	(*BatchProcessingEvent)(nil),  // 3: pipeline_data.BatchProcessingEvent
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 5: google.protobuf.Struct
}
var file_pipeline_data_proto_depIdxs = []int32{
	0, // 0: pipeline_data.CommandList.commands:type_name -> pipeline_data.CommandStep
	4, // 1: pipeline_data.Event.index_time:type_name -> google.protobuf.Timestamp
	4, // 2: pipeline_data.Event.timestamp:type_name -> google.protobuf.Timestamp
	5, // 3: pipeline_data.Event.derived:type_name -> google.protobuf.Struct
	1, // 4: pipeline_data.BatchProcessingEvent.commands:type_name -> pipeline_data.CommandList
	2, // 5: pipeline_data.BatchProcessingEvent.events:type_name -> pipeline_data.Event
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pipeline_data_proto_init() }
func file_pipeline_data_proto_init() {
	if File_pipeline_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipeline_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandStep); i {
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
		file_pipeline_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandList); i {
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
		file_pipeline_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_pipeline_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchProcessingEvent); i {
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
	file_pipeline_data_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pipeline_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pipeline_data_proto_goTypes,
		DependencyIndexes: file_pipeline_data_proto_depIdxs,
		MessageInfos:      file_pipeline_data_proto_msgTypes,
	}.Build()
	File_pipeline_data_proto = out.File
	file_pipeline_data_proto_rawDesc = nil
	file_pipeline_data_proto_goTypes = nil
	file_pipeline_data_proto_depIdxs = nil
}
