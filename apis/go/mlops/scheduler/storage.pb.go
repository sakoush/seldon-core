/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: mlops/scheduler/storage.proto

package scheduler

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

type PipelineSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastVersion uint32               `protobuf:"varint,2,opt,name=lastVersion,proto3" json:"lastVersion,omitempty"`
	Versions    []*PipelineWithState `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
	Deleted     bool                 `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *PipelineSnapshot) Reset() {
	*x = PipelineSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_scheduler_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineSnapshot) ProtoMessage() {}

func (x *PipelineSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_scheduler_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineSnapshot.ProtoReflect.Descriptor instead.
func (*PipelineSnapshot) Descriptor() ([]byte, []int) {
	return file_mlops_scheduler_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineSnapshot) GetLastVersion() uint32 {
	if x != nil {
		return x.LastVersion
	}
	return 0
}

func (x *PipelineSnapshot) GetVersions() []*PipelineWithState {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *PipelineSnapshot) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type ExperimentSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
	// to mark the experiment as deleted, this is currently required as we persist all
	// experiments in the local scheduler state (badgerdb) so that events can be replayed
	// on restart, which would guard against lost events in communication.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *ExperimentSnapshot) Reset() {
	*x = ExperimentSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlops_scheduler_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentSnapshot) ProtoMessage() {}

func (x *ExperimentSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_mlops_scheduler_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentSnapshot.ProtoReflect.Descriptor instead.
func (*ExperimentSnapshot) Descriptor() ([]byte, []int) {
	return file_mlops_scheduler_storage_proto_rawDescGZIP(), []int{1}
}

func (x *ExperimentSnapshot) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

func (x *ExperimentSnapshot) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_mlops_scheduler_storage_proto protoreflect.FileDescriptor

var file_mlops_scheduler_storage_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x10, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d,
	0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0x72, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2e, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x69, 0x6f, 0x2f,
	0x73, 0x65, 0x6c, 0x64, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mlops_scheduler_storage_proto_rawDescOnce sync.Once
	file_mlops_scheduler_storage_proto_rawDescData = file_mlops_scheduler_storage_proto_rawDesc
)

func file_mlops_scheduler_storage_proto_rawDescGZIP() []byte {
	file_mlops_scheduler_storage_proto_rawDescOnce.Do(func() {
		file_mlops_scheduler_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_mlops_scheduler_storage_proto_rawDescData)
	})
	return file_mlops_scheduler_storage_proto_rawDescData
}

var file_mlops_scheduler_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mlops_scheduler_storage_proto_goTypes = []interface{}{
	(*PipelineSnapshot)(nil),   // 0: seldon.mlops.scheduler.PipelineSnapshot
	(*ExperimentSnapshot)(nil), // 1: seldon.mlops.scheduler.ExperimentSnapshot
	(*PipelineWithState)(nil),  // 2: seldon.mlops.scheduler.PipelineWithState
	(*Experiment)(nil),         // 3: seldon.mlops.scheduler.Experiment
}
var file_mlops_scheduler_storage_proto_depIdxs = []int32{
	2, // 0: seldon.mlops.scheduler.PipelineSnapshot.versions:type_name -> seldon.mlops.scheduler.PipelineWithState
	3, // 1: seldon.mlops.scheduler.ExperimentSnapshot.experiment:type_name -> seldon.mlops.scheduler.Experiment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mlops_scheduler_storage_proto_init() }
func file_mlops_scheduler_storage_proto_init() {
	if File_mlops_scheduler_storage_proto != nil {
		return
	}
	file_mlops_scheduler_scheduler_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mlops_scheduler_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineSnapshot); i {
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
		file_mlops_scheduler_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentSnapshot); i {
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
			RawDescriptor: file_mlops_scheduler_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mlops_scheduler_storage_proto_goTypes,
		DependencyIndexes: file_mlops_scheduler_storage_proto_depIdxs,
		MessageInfos:      file_mlops_scheduler_storage_proto_msgTypes,
	}.Build()
	File_mlops_scheduler_storage_proto = out.File
	file_mlops_scheduler_storage_proto_rawDesc = nil
	file_mlops_scheduler_storage_proto_goTypes = nil
	file_mlops_scheduler_storage_proto_depIdxs = nil
}
