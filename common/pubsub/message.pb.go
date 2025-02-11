// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/proto/message.proto

package pubsub

import (
	segmenters "github.com/caraml-dev/xp/common/segmenters"
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

type MessagePublishState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Update:
	//	*MessagePublishState_ExperimentCreated
	//	*MessagePublishState_ExperimentUpdated
	//	*MessagePublishState_ProjectSettingsCreated
	//	*MessagePublishState_ProjectSettingsUpdated
	//	*MessagePublishState_ProjectSegmenterCreated
	//	*MessagePublishState_ProjectSegmenterUpdated
	//	*MessagePublishState_ProjectSegmenterDeleted
	Update isMessagePublishState_Update `protobuf_oneof:"update"`
}

func (x *MessagePublishState) Reset() {
	*x = MessagePublishState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePublishState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePublishState) ProtoMessage() {}

func (x *MessagePublishState) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePublishState.ProtoReflect.Descriptor instead.
func (*MessagePublishState) Descriptor() ([]byte, []int) {
	return file_api_proto_message_proto_rawDescGZIP(), []int{0}
}

func (m *MessagePublishState) GetUpdate() isMessagePublishState_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *MessagePublishState) GetExperimentCreated() *ExperimentCreated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ExperimentCreated); ok {
		return x.ExperimentCreated
	}
	return nil
}

func (x *MessagePublishState) GetExperimentUpdated() *ExperimentUpdated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ExperimentUpdated); ok {
		return x.ExperimentUpdated
	}
	return nil
}

func (x *MessagePublishState) GetProjectSettingsCreated() *ProjectSettingsCreated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ProjectSettingsCreated); ok {
		return x.ProjectSettingsCreated
	}
	return nil
}

func (x *MessagePublishState) GetProjectSettingsUpdated() *ProjectSettingsUpdated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ProjectSettingsUpdated); ok {
		return x.ProjectSettingsUpdated
	}
	return nil
}

func (x *MessagePublishState) GetProjectSegmenterCreated() *segmenters.ProjectSegmenterCreated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ProjectSegmenterCreated); ok {
		return x.ProjectSegmenterCreated
	}
	return nil
}

func (x *MessagePublishState) GetProjectSegmenterUpdated() *segmenters.ProjectSegmenterUpdated {
	if x, ok := x.GetUpdate().(*MessagePublishState_ProjectSegmenterUpdated); ok {
		return x.ProjectSegmenterUpdated
	}
	return nil
}

func (x *MessagePublishState) GetProjectSegmenterDeleted() *segmenters.ProjectSegmenterDeleted {
	if x, ok := x.GetUpdate().(*MessagePublishState_ProjectSegmenterDeleted); ok {
		return x.ProjectSegmenterDeleted
	}
	return nil
}

type isMessagePublishState_Update interface {
	isMessagePublishState_Update()
}

type MessagePublishState_ExperimentCreated struct {
	ExperimentCreated *ExperimentCreated `protobuf:"bytes,1,opt,name=experiment_created,json=experimentCreated,proto3,oneof"`
}

type MessagePublishState_ExperimentUpdated struct {
	ExperimentUpdated *ExperimentUpdated `protobuf:"bytes,2,opt,name=experiment_updated,json=experimentUpdated,proto3,oneof"`
}

type MessagePublishState_ProjectSettingsCreated struct {
	ProjectSettingsCreated *ProjectSettingsCreated `protobuf:"bytes,3,opt,name=project_settings_created,json=projectSettingsCreated,proto3,oneof"`
}

type MessagePublishState_ProjectSettingsUpdated struct {
	ProjectSettingsUpdated *ProjectSettingsUpdated `protobuf:"bytes,4,opt,name=project_settings_updated,json=projectSettingsUpdated,proto3,oneof"`
}

type MessagePublishState_ProjectSegmenterCreated struct {
	ProjectSegmenterCreated *segmenters.ProjectSegmenterCreated `protobuf:"bytes,5,opt,name=project_segmenter_created,json=projectSegmenterCreated,proto3,oneof"`
}

type MessagePublishState_ProjectSegmenterUpdated struct {
	ProjectSegmenterUpdated *segmenters.ProjectSegmenterUpdated `protobuf:"bytes,6,opt,name=project_segmenter_updated,json=projectSegmenterUpdated,proto3,oneof"`
}

type MessagePublishState_ProjectSegmenterDeleted struct {
	ProjectSegmenterDeleted *segmenters.ProjectSegmenterDeleted `protobuf:"bytes,7,opt,name=project_segmenter_deleted,json=projectSegmenterDeleted,proto3,oneof"`
}

func (*MessagePublishState_ExperimentCreated) isMessagePublishState_Update() {}

func (*MessagePublishState_ExperimentUpdated) isMessagePublishState_Update() {}

func (*MessagePublishState_ProjectSettingsCreated) isMessagePublishState_Update() {}

func (*MessagePublishState_ProjectSettingsUpdated) isMessagePublishState_Update() {}

func (*MessagePublishState_ProjectSegmenterCreated) isMessagePublishState_Update() {}

func (*MessagePublishState_ProjectSegmenterUpdated) isMessagePublishState_Update() {}

func (*MessagePublishState_ProjectSegmenterDeleted) isMessagePublishState_Update() {}

var File_api_proto_message_proto protoreflect.FileDescriptor

var file_api_proto_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x05, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x5a, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x5a, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x61, 0x0a, 0x19, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x61,
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x61, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x17, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_message_proto_rawDescOnce sync.Once
	file_api_proto_message_proto_rawDescData = file_api_proto_message_proto_rawDesc
)

func file_api_proto_message_proto_rawDescGZIP() []byte {
	file_api_proto_message_proto_rawDescOnce.Do(func() {
		file_api_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_message_proto_rawDescData)
	})
	return file_api_proto_message_proto_rawDescData
}

var file_api_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_message_proto_goTypes = []interface{}{
	(*MessagePublishState)(nil),                // 0: pubsub.MessagePublishState
	(*ExperimentCreated)(nil),                  // 1: pubsub.ExperimentCreated
	(*ExperimentUpdated)(nil),                  // 2: pubsub.ExperimentUpdated
	(*ProjectSettingsCreated)(nil),             // 3: pubsub.ProjectSettingsCreated
	(*ProjectSettingsUpdated)(nil),             // 4: pubsub.ProjectSettingsUpdated
	(*segmenters.ProjectSegmenterCreated)(nil), // 5: segmenters.ProjectSegmenterCreated
	(*segmenters.ProjectSegmenterUpdated)(nil), // 6: segmenters.ProjectSegmenterUpdated
	(*segmenters.ProjectSegmenterDeleted)(nil), // 7: segmenters.ProjectSegmenterDeleted
}
var file_api_proto_message_proto_depIdxs = []int32{
	1, // 0: pubsub.MessagePublishState.experiment_created:type_name -> pubsub.ExperimentCreated
	2, // 1: pubsub.MessagePublishState.experiment_updated:type_name -> pubsub.ExperimentUpdated
	3, // 2: pubsub.MessagePublishState.project_settings_created:type_name -> pubsub.ProjectSettingsCreated
	4, // 3: pubsub.MessagePublishState.project_settings_updated:type_name -> pubsub.ProjectSettingsUpdated
	5, // 4: pubsub.MessagePublishState.project_segmenter_created:type_name -> segmenters.ProjectSegmenterCreated
	6, // 5: pubsub.MessagePublishState.project_segmenter_updated:type_name -> segmenters.ProjectSegmenterUpdated
	7, // 6: pubsub.MessagePublishState.project_segmenter_deleted:type_name -> segmenters.ProjectSegmenterDeleted
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_message_proto_init() }
func file_api_proto_message_proto_init() {
	if File_api_proto_message_proto != nil {
		return
	}
	file_api_proto_experiment_proto_init()
	file_api_proto_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePublishState); i {
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
	file_api_proto_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MessagePublishState_ExperimentCreated)(nil),
		(*MessagePublishState_ExperimentUpdated)(nil),
		(*MessagePublishState_ProjectSettingsCreated)(nil),
		(*MessagePublishState_ProjectSettingsUpdated)(nil),
		(*MessagePublishState_ProjectSegmenterCreated)(nil),
		(*MessagePublishState_ProjectSegmenterUpdated)(nil),
		(*MessagePublishState_ProjectSegmenterDeleted)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_message_proto_goTypes,
		DependencyIndexes: file_api_proto_message_proto_depIdxs,
		MessageInfos:      file_api_proto_message_proto_msgTypes,
	}.Build()
	File_api_proto_message_proto = out.File
	file_api_proto_message_proto_rawDesc = nil
	file_api_proto_message_proto_goTypes = nil
	file_api_proto_message_proto_depIdxs = nil
}
