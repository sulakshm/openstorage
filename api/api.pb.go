// Code generated by protoc-gen-go.
// source: api/api.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// CHANGE: no NONE, was 1 << iota
type Status int32

const (
	Status_STATUS_NONE    Status = 0
	Status_STATUS_INIT    Status = 1
	Status_STATUS_OK      Status = 2
	Status_STATUS_OFFLINE Status = 3
	Status_STATUS_ERROR   Status = 4
)

var Status_name = map[int32]string{
	0: "STATUS_NONE",
	1: "STATUS_INIT",
	2: "STATUS_OK",
	3: "STATUS_OFFLINE",
	4: "STATUS_ERROR",
}
var Status_value = map[string]int32{
	"STATUS_NONE":    0,
	"STATUS_INIT":    1,
	"STATUS_OK":      2,
	"STATUS_OFFLINE": 3,
	"STATUS_ERROR":   4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CHANGE: no NONE, was 1 << iota
type DriverType int32

const (
	DriverType_DRIVER_TYPE_NONE      DriverType = 0
	DriverType_DRIVER_TYPE_FILE      DriverType = 1
	DriverType_DRIVER_TYPE_BLOCK     DriverType = 2
	DriverType_DRIVER_TYPE_OBJECT    DriverType = 3
	DriverType_DRIVER_TYPE_CLUSTERED DriverType = 4
	DriverType_DRIVER_TYPE_GRAPH     DriverType = 5
)

var DriverType_name = map[int32]string{
	0: "DRIVER_TYPE_NONE",
	1: "DRIVER_TYPE_FILE",
	2: "DRIVER_TYPE_BLOCK",
	3: "DRIVER_TYPE_OBJECT",
	4: "DRIVER_TYPE_CLUSTERED",
	5: "DRIVER_TYPE_GRAPH",
}
var DriverType_value = map[string]int32{
	"DRIVER_TYPE_NONE":      0,
	"DRIVER_TYPE_FILE":      1,
	"DRIVER_TYPE_BLOCK":     2,
	"DRIVER_TYPE_OBJECT":    3,
	"DRIVER_TYPE_CLUSTERED": 4,
	"DRIVER_TYPE_GRAPH":     5,
}

func (x DriverType) String() string {
	return proto.EnumName(DriverType_name, int32(x))
}
func (DriverType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CHANGE: was Filesystem, no NONE, was a string
type FSType int32

const (
	FSType_FS_TYPE_NONE  FSType = 0
	FSType_FS_TYPE_BTRFS FSType = 1
	FSType_FS_TYPE_EXT4  FSType = 2
	FSType_FS_TYPE_FUSE  FSType = 3
	FSType_FS_TYPE_NFS   FSType = 4
	FSType_FS_TYPE_VFS   FSType = 5
	FSType_FS_TYPE_XFS   FSType = 6
	FSType_FS_TYPE_ZFS   FSType = 7
)

var FSType_name = map[int32]string{
	0: "FS_TYPE_NONE",
	1: "FS_TYPE_BTRFS",
	2: "FS_TYPE_EXT4",
	3: "FS_TYPE_FUSE",
	4: "FS_TYPE_NFS",
	5: "FS_TYPE_VFS",
	6: "FS_TYPE_XFS",
	7: "FS_TYPE_ZFS",
}
var FSType_value = map[string]int32{
	"FS_TYPE_NONE":  0,
	"FS_TYPE_BTRFS": 1,
	"FS_TYPE_EXT4":  2,
	"FS_TYPE_FUSE":  3,
	"FS_TYPE_NFS":   4,
	"FS_TYPE_VFS":   5,
	"FS_TYPE_XFS":   6,
	"FS_TYPE_ZFS":   7,
}

func (x FSType) String() string {
	return proto.EnumName(FSType_name, int32(x))
}
func (FSType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// CHANGE: was an int, 0 was modified, 1 was added, 2 was deleted
type GraphDriverChangeType int32

const (
	GraphDriverChangeType_GRAPH_DRIVER_CHANGE_TYPE_NONE     GraphDriverChangeType = 0
	GraphDriverChangeType_GRAPH_DRIVER_CHANGE_TYPE_MODIFIED GraphDriverChangeType = 1
	GraphDriverChangeType_GRAPH_DRIVER_CHANGE_TYPE_ADDED    GraphDriverChangeType = 2
	GraphDriverChangeType_GRAPH_DRIVER_CHANGE_TYPE_DELETED  GraphDriverChangeType = 3
)

var GraphDriverChangeType_name = map[int32]string{
	0: "GRAPH_DRIVER_CHANGE_TYPE_NONE",
	1: "GRAPH_DRIVER_CHANGE_TYPE_MODIFIED",
	2: "GRAPH_DRIVER_CHANGE_TYPE_ADDED",
	3: "GRAPH_DRIVER_CHANGE_TYPE_DELETED",
}
var GraphDriverChangeType_value = map[string]int32{
	"GRAPH_DRIVER_CHANGE_TYPE_NONE":     0,
	"GRAPH_DRIVER_CHANGE_TYPE_MODIFIED": 1,
	"GRAPH_DRIVER_CHANGE_TYPE_ADDED":    2,
	"GRAPH_DRIVER_CHANGE_TYPE_DELETED":  3,
}

func (x GraphDriverChangeType) String() string {
	return proto.EnumName(GraphDriverChangeType_name, int32(x))
}
func (GraphDriverChangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// CHANGE: no NONE
type VolumeActionParam int32

const (
	VolumeActionParam_VOLUME_ACTION_PARAM_NONE VolumeActionParam = 0
	// Maps to the boolean value false
	VolumeActionParam_VOLUME_ACTION_PARAM_OFF VolumeActionParam = 1
	// Maps to the boolean value true.
	VolumeActionParam_VOLUME_ACTION_PARAM_ON VolumeActionParam = 2
)

var VolumeActionParam_name = map[int32]string{
	0: "VOLUME_ACTION_PARAM_NONE",
	1: "VOLUME_ACTION_PARAM_OFF",
	2: "VOLUME_ACTION_PARAM_ON",
}
var VolumeActionParam_value = map[string]int32{
	"VOLUME_ACTION_PARAM_NONE": 0,
	"VOLUME_ACTION_PARAM_OFF":  1,
	"VOLUME_ACTION_PARAM_ON":   2,
}

func (x VolumeActionParam) String() string {
	return proto.EnumName(VolumeActionParam_name, int32(x))
}
func (VolumeActionParam) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// VolumeState represents the state of a volume.
// CHANGE: no NONE, was 1 << iota (and was bit or'ed/and'ed)
type VolumeState int32

const (
	VolumeState_VOLUME_STATE_NONE VolumeState = 0
	// Volume is transitioning to new state
	VolumeState_VOLUME_STATE_PENDING VolumeState = 1
	// Volume is ready to be assigned to a container
	VolumeState_VOLUME_STATE_AVAILABLE VolumeState = 2
	// Volume is attached to container
	VolumeState_VOLUME_STATE_ATTACHED VolumeState = 3
	// Volume is detached but associated with a container
	VolumeState_VOLUME_STATE_DETACHED VolumeState = 4
	// Volume detach is in progress
	VolumeState_VOLUME_STATE_DETATCHING VolumeState = 5
	// Volume is in error state
	VolumeState_VOLUME_STATE_ERROR VolumeState = 6
	// Volume is deleted, it will remain in this state
	// while resources are asynchronously reclaimed
	VolumeState_VOLUME_STATE_DELETED VolumeState = 7
)

var VolumeState_name = map[int32]string{
	0: "VOLUME_STATE_NONE",
	1: "VOLUME_STATE_PENDING",
	2: "VOLUME_STATE_AVAILABLE",
	3: "VOLUME_STATE_ATTACHED",
	4: "VOLUME_STATE_DETACHED",
	5: "VOLUME_STATE_DETATCHING",
	6: "VOLUME_STATE_ERROR",
	7: "VOLUME_STATE_DELETED",
}
var VolumeState_value = map[string]int32{
	"VOLUME_STATE_NONE":       0,
	"VOLUME_STATE_PENDING":    1,
	"VOLUME_STATE_AVAILABLE":  2,
	"VOLUME_STATE_ATTACHED":   3,
	"VOLUME_STATE_DETACHED":   4,
	"VOLUME_STATE_DETATCHING": 5,
	"VOLUME_STATE_ERROR":      6,
	"VOLUME_STATE_DELETED":    7,
}

func (x VolumeState) String() string {
	return proto.EnumName(VolumeState_name, int32(x))
}
func (VolumeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// VolumeStatus represents a health status for a volume.
// CHANGE: no NONE, was a string
type VolumeStatus int32

const (
	VolumeStatus_VOLUME_STATUS_NONE VolumeStatus = 0
	// Volume is not present
	VolumeStatus_VOLUME_STATUS_NOT_PRESENT VolumeStatus = 1
	// Volume is healthy
	VolumeStatus_VOLUME_STATUS_UP VolumeStatus = 2
	// Volume is in fail mode
	VolumeStatus_VOLUME_STATUS_DOWN VolumeStatus = 3
	// Volume is up but with degraded performance
	// In a RAID group, this may indicate a problem with one or more drives
	VolumeStatus_VOLUME_STATUS_DEGRADED VolumeStatus = 4
)

var VolumeStatus_name = map[int32]string{
	0: "VOLUME_STATUS_NONE",
	1: "VOLUME_STATUS_NOT_PRESENT",
	2: "VOLUME_STATUS_UP",
	3: "VOLUME_STATUS_DOWN",
	4: "VOLUME_STATUS_DEGRADED",
}
var VolumeStatus_value = map[string]int32{
	"VOLUME_STATUS_NONE":        0,
	"VOLUME_STATUS_NOT_PRESENT": 1,
	"VOLUME_STATUS_UP":          2,
	"VOLUME_STATUS_DOWN":        3,
	"VOLUME_STATUS_DEGRADED":    4,
}

func (x VolumeStatus) String() string {
	return proto.EnumName(VolumeStatus_name, int32(x))
}
func (VolumeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// VolumeLocator is a structure that is attached to a volume
// and is used to carry opaque metadata.
type VolumeLocator struct {
	// User friendly identifier
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A set of name-value pairs that acts as search filters
	VolumeLabels map[string]string `protobuf:"bytes,2,rep,name=volume_labels" json:"volume_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *VolumeLocator) Reset()                    { *m = VolumeLocator{} }
func (m *VolumeLocator) String() string            { return proto.CompactTextString(m) }
func (*VolumeLocator) ProtoMessage()               {}
func (*VolumeLocator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VolumeLocator) GetVolumeLabels() map[string]string {
	if m != nil {
		return m.VolumeLabels
	}
	return nil
}

type Source struct {
	// A volume id, if specified will create a clone of the parent.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Seed will seed the volume from the specified URI
	// Any additional config for the source comes from the labels in the spec
	Seed string `protobuf:"bytes,2,opt,name=seed" json:"seed,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// VolumeSpec has the properties needed to create a volume.
type VolumeSpec struct {
	// Ephemeral storage
	Ephemeral bool `protobuf:"varint,1,opt,name=ephemeral" json:"ephemeral,omitempty"`
	// Thin provisioned volume size in bytes
	Size uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	// Format disk with this FSType
	Format FSType `protobuf:"varint,3,opt,name=format,enum=openstorage.api.FSType" json:"format,omitempty"`
	// Block size for filesystem
	BlockSize int64 `protobuf:"varint,4,opt,name=block_size" json:"block_size,omitempty"`
	// Specifies the number of nodes that are
	// allowed to fail, and yet data is available
	// A value of 0 implies that data is not erasure coded,
	// a failure of a node will lead to data loss
	HaLevel int64 `protobuf:"varint,5,opt,name=ha_level" json:"ha_level,omitempty"`
	// The COS, 1 to 9
	Cos uint32 `protobuf:"varint,6,opt,name=cos" json:"cos,omitempty"`
	// Perform dedupe on this disk
	Dedupe bool `protobuf:"varint,7,opt,name=dedupe" json:"dedupe,omitempty"`
	// SnapshotInterval in minutes, set to 0 to disable snapshots
	SnapshotInterval uint32 `protobuf:"varint,8,opt,name=snapshot_interval" json:"snapshot_interval,omitempty"`
	// Volume configuration labels
	ConfigLabels map[string]string `protobuf:"bytes,9,rep,name=config_labels" json:"config_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *VolumeSpec) Reset()                    { *m = VolumeSpec{} }
func (m *VolumeSpec) String() string            { return proto.CompactTextString(m) }
func (*VolumeSpec) ProtoMessage()               {}
func (*VolumeSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VolumeSpec) GetConfigLabels() map[string]string {
	if m != nil {
		return m.ConfigLabels
	}
	return nil
}

// Volume represents a live, created volume.
type Volume struct {
	// Self referential volume ID
	Id       string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Source   *Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Readonly bool    `protobuf:"varint,3,opt,name=readonly" json:"readonly,omitempty"`
	// User specified locator
	Locator *VolumeLocator `protobuf:"bytes,4,opt,name=locator" json:"locator,omitempty"`
	// Volume creation time
	Ctime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=ctime" json:"ctime,omitempty"`
	// User specified VolumeSpec
	Spec *VolumeSpec `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	// Volume usage
	Usage uint64 `protobuf:"varint,7,opt,name=usage" json:"usage,omitempty"`
	// Time when an integrity check for run
	LastScan *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=last_scan" json:"last_scan,omitempty"`
	// Format FSType type if any
	Format FSType       `protobuf:"varint,9,opt,name=format,enum=openstorage.api.FSType" json:"format,omitempty"`
	Status VolumeStatus `protobuf:"varint,10,opt,name=status,enum=openstorage.api.VolumeStatus" json:"status,omitempty"`
	State  VolumeState  `protobuf:"varint,11,opt,name=state,enum=openstorage.api.VolumeState" json:"state,omitempty"`
	// Machine ID (node) on which this volume is attached
	// Machine ID is a node instance identifier for clustered systems.
	AttachedOn string `protobuf:"bytes,12,opt,name=attached_on" json:"attached_on,omitempty"`
	DevicePath string `protobuf:"bytes,14,opt,name=device_path" json:"device_path,omitempty"`
	AttachPath string `protobuf:"bytes,15,opt,name=attach_path" json:"attach_path,omitempty"`
	// Set of machine IDs (nodes) to which this volume is erasure coded - for clustered storage arrays
	ReplicaSet []string `protobuf:"bytes,16,rep,name=replica_set" json:"replica_set,omitempty"`
	// Last recorded error
	Error string `protobuf:"bytes,17,opt,name=error" json:"error,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Volume) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Volume) GetLocator() *VolumeLocator {
	if m != nil {
		return m.Locator
	}
	return nil
}

func (m *Volume) GetCtime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Volume) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Volume) GetLastScan() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastScan
	}
	return nil
}

type Stats struct {
	// Reads completed successfully
	Reads int64 `protobuf:"varint,1,opt,name=reads" json:"reads,omitempty"`
	// Time spent in reads in ms
	ReadMs    int64 `protobuf:"varint,2,opt,name=read_ms" json:"read_ms,omitempty"`
	ReadBytes int64 `protobuf:"varint,3,opt,name=read_bytes" json:"read_bytes,omitempty"`
	// Writes completed successfully
	Writes int64 `protobuf:"varint,4,opt,name=writes" json:"writes,omitempty"`
	// Time spent in writes in ms
	WriteMs    int64 `protobuf:"varint,5,opt,name=write_ms" json:"write_ms,omitempty"`
	WriteBytes int64 `protobuf:"varint,6,opt,name=write_bytes" json:"write_bytes,omitempty"`
	// IOs curently in progress
	IoProgress int64 `protobuf:"varint,7,opt,name=io_progress" json:"io_progress,omitempty"`
	// Time spent doing IOs ms
	IoMs int64 `protobuf:"varint,8,opt,name=io_ms" json:"io_ms,omitempty"`
	// BytesUsed
	BytesUsed uint64 `protobuf:"varint,9,opt,name=bytes_used" json:"bytes_used,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// TODO: what?
type Alerts struct {
}

func (m *Alerts) Reset()                    { *m = Alerts{} }
func (m *Alerts) String() string            { return proto.CompactTextString(m) }
func (*Alerts) ProtoMessage()               {}
func (*Alerts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type VolumeCreateRequest struct {
	// User specified volume name and labels
	Locator *VolumeLocator `protobuf:"bytes,1,opt,name=locator" json:"locator,omitempty"`
	// Source to create volume
	Source *Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// The storage spec for the volume
	Spec *VolumeSpec `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *VolumeCreateRequest) Reset()                    { *m = VolumeCreateRequest{} }
func (m *VolumeCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateRequest) ProtoMessage()               {}
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VolumeCreateRequest) GetLocator() *VolumeLocator {
	if m != nil {
		return m.Locator
	}
	return nil
}

func (m *VolumeCreateRequest) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *VolumeCreateRequest) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type VolumeResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *VolumeResponse) Reset()                    { *m = VolumeResponse{} }
func (m *VolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeResponse) ProtoMessage()               {}
func (*VolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// CHANGE: error was embedded VolumeResponse
type VolumeCreateResponse struct {
	// ID of the newly created volume
	Id             string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	VolumeResponse *VolumeResponse `protobuf:"bytes,2,opt,name=volume_response" json:"volume_response,omitempty"`
}

func (m *VolumeCreateResponse) Reset()                    { *m = VolumeCreateResponse{} }
func (m *VolumeCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateResponse) ProtoMessage()               {}
func (*VolumeCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *VolumeCreateResponse) GetVolumeResponse() *VolumeResponse {
	if m != nil {
		return m.VolumeResponse
	}
	return nil
}

// VolumeStateAction specifies desired actions.
type VolumeStateAction struct {
	// Attach or Detach volume
	Attach VolumeActionParam `protobuf:"varint,1,opt,name=attach,enum=openstorage.api.VolumeActionParam" json:"attach,omitempty"`
	// Mount or unmount volume
	Mount     VolumeActionParam `protobuf:"varint,2,opt,name=mount,enum=openstorage.api.VolumeActionParam" json:"mount,omitempty"`
	MountPath string            `protobuf:"bytes,3,opt,name=mount_path" json:"mount_path,omitempty"`
	// Device path returned in attach
	DevicePath string `protobuf:"bytes,4,opt,name=device_path" json:"device_path,omitempty"`
}

func (m *VolumeStateAction) Reset()                    { *m = VolumeStateAction{} }
func (m *VolumeStateAction) String() string            { return proto.CompactTextString(m) }
func (*VolumeStateAction) ProtoMessage()               {}
func (*VolumeStateAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type VolumeSetRequest struct {
	// User specified volume name and labels
	Locator *VolumeLocator `protobuf:"bytes,1,opt,name=locator" json:"locator,omitempty"`
	// The storage spec for the volume
	Spec *VolumeSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// State modification on this volume.
	Action *VolumeStateAction `protobuf:"bytes,3,opt,name=action" json:"action,omitempty"`
}

func (m *VolumeSetRequest) Reset()                    { *m = VolumeSetRequest{} }
func (m *VolumeSetRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeSetRequest) ProtoMessage()               {}
func (*VolumeSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *VolumeSetRequest) GetLocator() *VolumeLocator {
	if m != nil {
		return m.Locator
	}
	return nil
}

func (m *VolumeSetRequest) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *VolumeSetRequest) GetAction() *VolumeStateAction {
	if m != nil {
		return m.Action
	}
	return nil
}

// CHANGE: error was embedded VolumeResponse
type VolumeSetResponse struct {
	Volume         *Volume         `protobuf:"bytes,1,opt,name=volume" json:"volume,omitempty"`
	VolumeResponse *VolumeResponse `protobuf:"bytes,2,opt,name=volume_response" json:"volume_response,omitempty"`
}

func (m *VolumeSetResponse) Reset()                    { *m = VolumeSetResponse{} }
func (m *VolumeSetResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeSetResponse) ProtoMessage()               {}
func (*VolumeSetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *VolumeSetResponse) GetVolume() *Volume {
	if m != nil {
		return m.Volume
	}
	return nil
}

func (m *VolumeSetResponse) GetVolumeResponse() *VolumeResponse {
	if m != nil {
		return m.VolumeResponse
	}
	return nil
}

type SnapCreateRequest struct {
	// volume id
	Id       string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Locator  *VolumeLocator `protobuf:"bytes,2,opt,name=locator" json:"locator,omitempty"`
	Readonly bool           `protobuf:"varint,3,opt,name=readonly" json:"readonly,omitempty"`
}

func (m *SnapCreateRequest) Reset()                    { *m = SnapCreateRequest{} }
func (m *SnapCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapCreateRequest) ProtoMessage()               {}
func (*SnapCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SnapCreateRequest) GetLocator() *VolumeLocator {
	if m != nil {
		return m.Locator
	}
	return nil
}

type SnapCreateResponse struct {
	VolumeCreateResponse *VolumeCreateResponse `protobuf:"bytes,1,opt,name=volume_create_response" json:"volume_create_response,omitempty"`
}

func (m *SnapCreateResponse) Reset()                    { *m = SnapCreateResponse{} }
func (m *SnapCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapCreateResponse) ProtoMessage()               {}
func (*SnapCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SnapCreateResponse) GetVolumeCreateResponse() *VolumeCreateResponse {
	if m != nil {
		return m.VolumeCreateResponse
	}
	return nil
}

type VolumeInfo struct {
	VolumeId string      `protobuf:"bytes,1,opt,name=volume_id" json:"volume_id,omitempty"`
	Path     string      `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Storage  *VolumeSpec `protobuf:"bytes,3,opt,name=storage" json:"storage,omitempty"`
}

func (m *VolumeInfo) Reset()                    { *m = VolumeInfo{} }
func (m *VolumeInfo) String() string            { return proto.CompactTextString(m) }
func (*VolumeInfo) ProtoMessage()               {}
func (*VolumeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *VolumeInfo) GetStorage() *VolumeSpec {
	if m != nil {
		return m.Storage
	}
	return nil
}

// GraphDriverChanges represent a list of changes between the filesystem layers
// specified by the ID and Parent.  // Parent may be an empty string, in which
// case there is no parent.
// Where the Path is the filesystem path within the layered filesystem
// CHANGE: kind was an int
type GraphDriverChanges struct {
	Path string                `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Kind GraphDriverChangeType `protobuf:"varint,2,opt,name=kind,enum=openstorage.api.GraphDriverChangeType" json:"kind,omitempty"`
}

func (m *GraphDriverChanges) Reset()                    { *m = GraphDriverChanges{} }
func (m *GraphDriverChanges) String() string            { return proto.CompactTextString(m) }
func (*GraphDriverChanges) ProtoMessage()               {}
func (*GraphDriverChanges) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*VolumeLocator)(nil), "openstorage.api.VolumeLocator")
	proto.RegisterType((*Source)(nil), "openstorage.api.Source")
	proto.RegisterType((*VolumeSpec)(nil), "openstorage.api.VolumeSpec")
	proto.RegisterType((*Volume)(nil), "openstorage.api.Volume")
	proto.RegisterType((*Stats)(nil), "openstorage.api.Stats")
	proto.RegisterType((*Alerts)(nil), "openstorage.api.Alerts")
	proto.RegisterType((*VolumeCreateRequest)(nil), "openstorage.api.VolumeCreateRequest")
	proto.RegisterType((*VolumeResponse)(nil), "openstorage.api.VolumeResponse")
	proto.RegisterType((*VolumeCreateResponse)(nil), "openstorage.api.VolumeCreateResponse")
	proto.RegisterType((*VolumeStateAction)(nil), "openstorage.api.VolumeStateAction")
	proto.RegisterType((*VolumeSetRequest)(nil), "openstorage.api.VolumeSetRequest")
	proto.RegisterType((*VolumeSetResponse)(nil), "openstorage.api.VolumeSetResponse")
	proto.RegisterType((*SnapCreateRequest)(nil), "openstorage.api.SnapCreateRequest")
	proto.RegisterType((*SnapCreateResponse)(nil), "openstorage.api.SnapCreateResponse")
	proto.RegisterType((*VolumeInfo)(nil), "openstorage.api.VolumeInfo")
	proto.RegisterType((*GraphDriverChanges)(nil), "openstorage.api.GraphDriverChanges")
	proto.RegisterEnum("openstorage.api.Status", Status_name, Status_value)
	proto.RegisterEnum("openstorage.api.DriverType", DriverType_name, DriverType_value)
	proto.RegisterEnum("openstorage.api.FSType", FSType_name, FSType_value)
	proto.RegisterEnum("openstorage.api.GraphDriverChangeType", GraphDriverChangeType_name, GraphDriverChangeType_value)
	proto.RegisterEnum("openstorage.api.VolumeActionParam", VolumeActionParam_name, VolumeActionParam_value)
	proto.RegisterEnum("openstorage.api.VolumeState", VolumeState_name, VolumeState_value)
	proto.RegisterEnum("openstorage.api.VolumeStatus", VolumeStatus_name, VolumeStatus_value)
}

var fileDescriptor0 = []byte{
	// 1379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x73, 0xdb, 0x54,
	0x14, 0x46, 0x96, 0xec, 0x24, 0xc7, 0x79, 0xd8, 0xea, 0x4b, 0x4d, 0xdf, 0x1a, 0x5a, 0x8a, 0xa1,
	0x0e, 0xb8, 0x5d, 0x74, 0xd8, 0x39, 0xb6, 0x9c, 0x18, 0x5c, 0xdb, 0x63, 0x2b, 0xa1, 0x14, 0x66,
	0x34, 0x8a, 0x7c, 0x13, 0x9b, 0xca, 0x92, 0xd0, 0x95, 0xc3, 0x84, 0x3f, 0xc1, 0x8a, 0x05, 0xb0,
	0x83, 0x1d, 0x2c, 0xf8, 0x49, 0xfc, 0x11, 0x16, 0x9c, 0x7b, 0xaf, 0x94, 0x48, 0x71, 0x9c, 0x66,
	0x60, 0x91, 0xc9, 0xdc, 0xef, 0x9c, 0x7b, 0x1e, 0xdf, 0xf9, 0xce, 0x95, 0x61, 0xcd, 0x0e, 0x26,
	0x5b, 0xf8, 0x57, 0x0d, 0x42, 0x3f, 0xf2, 0xd5, 0x0d, 0x3f, 0x20, 0x1e, 0x8d, 0xfc, 0xd0, 0x3e,
	0x22, 0x55, 0x84, 0x37, 0x1f, 0x1c, 0xf9, 0xfe, 0x91, 0x4b, 0xb6, 0xb8, 0xf9, 0x60, 0x76, 0xb8,
	0x15, 0x4d, 0xa6, 0x84, 0x46, 0xf6, 0x34, 0x10, 0x37, 0xf4, 0x5f, 0x25, 0x58, 0xdb, 0xf7, 0xdd,
	0xd9, 0x94, 0x74, 0x7c, 0xc7, 0xc6, 0x9b, 0xea, 0x2a, 0x28, 0x9e, 0x3d, 0x25, 0x9a, 0xf4, 0x50,
	0x7a, 0xba, 0xa2, 0xee, 0xc0, 0xda, 0x31, 0x37, 0x5b, 0xae, 0x7d, 0x40, 0x5c, 0xaa, 0xe5, 0x1e,
	0xca, 0x4f, 0x8b, 0xb5, 0x4f, 0xaa, 0xe7, 0x32, 0x55, 0x33, 0x41, 0x92, 0x13, 0xbf, 0x62, 0x78,
	0x51, 0x78, 0xb2, 0xf9, 0x1c, 0xca, 0x73, 0xa0, 0x5a, 0x04, 0xf9, 0x2d, 0x39, 0x89, 0x53, 0xad,
	0x41, 0xfe, 0xd8, 0x76, 0x67, 0x04, 0x53, 0xe0, 0xf1, 0xb3, 0xdc, 0x4b, 0x49, 0x7f, 0x02, 0x85,
	0xa1, 0x3f, 0x0b, 0x1d, 0xa2, 0xae, 0x43, 0x21, 0xb0, 0x43, 0xe2, 0x45, 0xb1, 0x33, 0x56, 0x49,
	0x09, 0x19, 0x09, 0x5f, 0xfd, 0xaf, 0x1c, 0x80, 0x88, 0x3e, 0x0c, 0x88, 0xa3, 0x96, 0x61, 0x85,
	0x04, 0x63, 0x32, 0x25, 0xa1, 0xed, 0x72, 0xff, 0x65, 0xee, 0x3f, 0xf9, 0x41, 0xc4, 0x56, 0xd4,
	0x0f, 0xa0, 0x70, 0xe8, 0x87, 0x53, 0x3b, 0xd2, 0x64, 0x3c, 0xaf, 0xd7, 0x6e, 0xcd, 0xb5, 0xd3,
	0x1a, 0x9a, 0x27, 0x01, 0x51, 0x55, 0x80, 0x03, 0xd7, 0x77, 0xde, 0x5a, 0xfc, 0xb2, 0x82, 0xce,
	0xb2, 0x5a, 0x82, 0xe5, 0xb1, 0x6d, 0xb9, 0xe4, 0x98, 0xb8, 0x5a, 0x9e, 0x23, 0xd8, 0x86, 0xe3,
	0x53, 0xad, 0x80, 0x87, 0x35, 0x56, 0xe9, 0x88, 0x8c, 0x66, 0x01, 0xd1, 0x96, 0x78, 0xe6, 0xdb,
	0x50, 0xa6, 0x9e, 0x1d, 0xd0, 0xb1, 0x1f, 0x59, 0x13, 0x2f, 0x22, 0x21, 0x36, 0xa9, 0x2d, 0x73,
	0xd7, 0x26, 0xac, 0x39, 0xbe, 0x77, 0x38, 0x39, 0x4a, 0xc8, 0x5d, 0xe1, 0xe4, 0x3e, 0x5b, 0x40,
	0x2e, 0xeb, 0xad, 0xda, 0xe0, 0x17, 0xce, 0x31, 0x3b, 0x07, 0xbe, 0x93, 0xd9, 0x7f, 0x64, 0x28,
	0x88, 0xa8, 0xd8, 0x62, 0x6e, 0x32, 0x8a, 0x3d, 0x91, 0x18, 0xca, 0x09, 0xe7, 0xae, 0xc5, 0x0b,
	0x88, 0x89, 0xe7, 0x81, 0x24, 0x84, 0xc4, 0x1e, 0xf9, 0x9e, 0x7b, 0xc2, 0x39, 0x5c, 0x56, 0xb7,
	0x60, 0xc9, 0x15, 0xd3, 0xe7, 0x3c, 0x15, 0x6b, 0xf7, 0x2f, 0xd7, 0x88, 0xfa, 0x21, 0xe4, 0x1d,
	0x26, 0x47, 0x4e, 0x62, 0xb1, 0xb6, 0x59, 0x15, 0x5a, 0xad, 0x26, 0x5a, 0xad, 0x9a, 0x89, 0x56,
	0xd1, 0x55, 0xa1, 0xd8, 0x3c, 0x67, 0xb8, 0x58, 0xbb, 0x73, 0x09, 0x3f, 0xac, 0xd7, 0x19, 0x45,
	0x9c, 0xb3, 0xaf, 0xa8, 0xcf, 0x60, 0xc5, 0xb5, 0x69, 0x64, 0x51, 0xc7, 0xf6, 0x38, 0xeb, 0x97,
	0x27, 0x3a, 0x13, 0xc6, 0xca, 0xe5, 0xc2, 0x78, 0x86, 0x44, 0x45, 0x76, 0x34, 0xa3, 0x1a, 0x70,
	0xc7, 0x7b, 0x8b, 0x6a, 0xe2, 0x4e, 0xea, 0x47, 0x90, 0x67, 0xee, 0x44, 0x2b, 0x72, 0xef, 0xbb,
	0x97, 0x78, 0x13, 0xf5, 0x1a, 0x14, 0xed, 0x28, 0xb2, 0x9d, 0x31, 0x19, 0x59, 0xbe, 0xa7, 0xad,
	0xf2, 0xc9, 0x20, 0x38, 0x22, 0xc7, 0x13, 0x87, 0x58, 0x81, 0x1d, 0x8d, 0xb5, 0xf5, 0x04, 0x14,
	0x9e, 0x02, 0xdc, 0x48, 0xc0, 0x90, 0x04, 0xee, 0xc4, 0xb1, 0x2d, 0x4a, 0x22, 0xad, 0x84, 0x9a,
	0xe2, 0x12, 0x20, 0x61, 0x88, 0xb3, 0x29, 0xf3, 0x85, 0xf9, 0x43, 0x82, 0x3c, 0x4b, 0x46, 0x99,
	0x81, 0x0d, 0x92, 0x72, 0x01, 0xc8, 0xea, 0x06, 0x2c, 0xb1, 0xa3, 0x35, 0xa5, 0x5c, 0x01, 0x32,
	0xdb, 0x00, 0x0e, 0x1c, 0x9c, 0x44, 0x84, 0xf2, 0x51, 0xcb, 0x4c, 0xe2, 0xdf, 0x87, 0x13, 0x76,
	0x3e, 0xdd, 0x08, 0x7e, 0x66, 0xb7, 0xc4, 0x46, 0x60, 0x0d, 0x02, 0x11, 0xd7, 0x0a, 0x09, 0x38,
	0xf1, 0x2d, 0x64, 0xfd, 0x28, 0x24, 0x94, 0xf2, 0x01, 0xc9, 0x2c, 0x3f, 0x82, 0x78, 0x71, 0x39,
	0x49, 0xc7, 0xaf, 0x58, 0x33, 0x8a, 0xdb, 0xcd, 0x86, 0xa0, 0xe8, 0xcb, 0x50, 0xa8, 0xbb, 0x24,
	0x8c, 0xa8, 0xfe, 0x8b, 0x04, 0xd7, 0x04, 0x53, 0x0d, 0x2c, 0x2a, 0x22, 0x03, 0xf2, 0xdd, 0x0c,
	0x27, 0x97, 0xd6, 0x9e, 0x74, 0x25, 0xed, 0x5d, 0x59, 0xe7, 0x89, 0xf2, 0xe4, 0x77, 0x2a, 0x4f,
	0x7f, 0x00, 0xeb, 0xe2, 0x34, 0x20, 0x34, 0xf0, 0x3d, 0x4a, 0xce, 0x48, 0xe7, 0xcb, 0xa5, 0x7f,
	0x03, 0xd7, 0xb3, 0xc5, 0xc7, 0x6e, 0xe9, 0x05, 0x7c, 0x09, 0x1b, 0xf1, 0x7b, 0x1b, 0xc6, 0xe6,
	0xb8, 0xc2, 0x07, 0x0b, 0x52, 0x27, 0x51, 0xf4, 0xdf, 0xa5, 0xe4, 0x85, 0xe5, 0x2a, 0xaa, 0xe3,
	0x6e, 0xf9, 0x9e, 0x5a, 0x83, 0x82, 0x50, 0x08, 0x8f, 0xbf, 0x5e, 0xd3, 0x17, 0x84, 0x11, 0xee,
	0x7d, 0x3b, 0xb4, 0xa7, 0xea, 0xa7, 0x90, 0x9f, 0xfa, 0x33, 0x7c, 0x6a, 0x73, 0x57, 0xbe, 0x82,
	0x3d, 0xf0, 0x2b, 0x42, 0x87, 0xf2, 0x45, 0x8a, 0x55, 0x38, 0x07, 0xbf, 0x49, 0x50, 0x8a, 0xab,
	0x24, 0xd1, 0x7f, 0x1e, 0x5f, 0x32, 0x95, 0xdc, 0xbb, 0xdf, 0x03, 0x46, 0x00, 0x2f, 0x34, 0x1e,
	0xa1, 0x7e, 0xd9, 0xea, 0x89, 0x96, 0xf4, 0xe3, 0x53, 0x26, 0x59, 0x8d, 0xf1, 0x94, 0x50, 0x32,
	0x62, 0x32, 0x71, 0x8d, 0xb7, 0x16, 0x04, 0xfa, 0x1f, 0x23, 0x3c, 0x80, 0xf2, 0x10, 0x3f, 0x15,
	0x59, 0x6d, 0xa7, 0xd5, 0x91, 0x22, 0x2a, 0x77, 0x25, 0xa2, 0xe6, 0x9e, 0x69, 0xfd, 0x6b, 0x50,
	0xd3, 0x39, 0xe2, 0xe6, 0x0c, 0xb8, 0x19, 0xd7, 0xec, 0x70, 0xc3, 0x59, 0xe9, 0xa2, 0xd9, 0xc7,
	0x0b, 0xf2, 0x64, 0xc3, 0x60, 0xf0, 0xf8, 0x33, 0xdc, 0xf6, 0x0e, 0x7d, 0xf6, 0x19, 0x8e, 0x83,
	0x9e, 0x36, 0x80, 0x9f, 0x61, 0x2e, 0x06, 0xfe, 0x21, 0x52, 0x3f, 0x86, 0xa5, 0x38, 0xe4, 0x55,
	0xf6, 0xeb, 0x35, 0xa8, 0x3b, 0xa1, 0x1d, 0x8c, 0x9b, 0xe1, 0xe4, 0x98, 0x84, 0x8d, 0xb1, 0xed,
	0x1d, 0x11, 0x7a, 0x1a, 0x51, 0xc4, 0x7f, 0x01, 0xca, 0xdb, 0x89, 0x37, 0x8a, 0x95, 0xfb, 0x64,
	0x2e, 0xdc, 0x5c, 0x00, 0xf6, 0x98, 0x57, 0x2c, 0xfc, 0x99, 0x21, 0xde, 0xe9, 0x0d, 0x28, 0x0e,
	0xcd, 0xba, 0xb9, 0x37, 0xb4, 0xba, 0xbd, 0xae, 0x51, 0x7a, 0x2f, 0x05, 0xb4, 0xbb, 0x6d, 0xb3,
	0x24, 0xe1, 0x4e, 0xaf, 0xc4, 0x40, 0xef, 0x8b, 0x52, 0x0e, 0x85, 0xbf, 0x9e, 0x1c, 0x5b, 0xad,
	0x4e, 0x1b, 0xef, 0xb0, 0xe7, 0x70, 0x35, 0xc6, 0x8c, 0xc1, 0xa0, 0x37, 0x28, 0x29, 0x95, 0x9f,
	0x25, 0x00, 0x91, 0x95, 0x7f, 0x3c, 0xae, 0x43, 0xa9, 0x39, 0x68, 0xef, 0x1b, 0x03, 0xcb, 0xfc,
	0xaa, 0x6f, 0x24, 0xa9, 0xce, 0xa1, 0xad, 0x76, 0xc7, 0xc0, 0x7c, 0x37, 0xa0, 0x9c, 0x46, 0xb7,
	0x3b, 0xbd, 0x06, 0xcb, 0x7b, 0x13, 0xd4, 0x34, 0xdc, 0xdb, 0xfe, 0xdc, 0x68, 0x98, 0x98, 0xfb,
	0x36, 0xdc, 0x48, 0xe3, 0x8d, 0xce, 0xde, 0xd0, 0x34, 0x06, 0x46, 0xb3, 0xa4, 0x9c, 0x8f, 0xb4,
	0x33, 0xa8, 0xf7, 0x77, 0x4b, 0xf9, 0xca, 0x4f, 0x12, 0x14, 0xe2, 0x8f, 0x1a, 0x16, 0xde, 0x1a,
	0x66, 0x6a, 0x2a, 0xc3, 0x5a, 0x82, 0x6c, 0x9b, 0x83, 0xd6, 0x10, 0x0b, 0x4a, 0x39, 0x19, 0xaf,
	0xcd, 0x17, 0x58, 0x4b, 0x0a, 0x69, 0xed, 0x0d, 0x19, 0x03, 0xc8, 0xda, 0x69, 0x20, 0xbc, 0xa4,
	0xa4, 0x81, 0x7d, 0x04, 0xf2, 0x69, 0xe0, 0x35, 0x02, 0x85, 0x34, 0xf0, 0x06, 0x81, 0xa5, 0xca,
	0x9f, 0x12, 0xdc, 0xb8, 0x70, 0x5c, 0xea, 0x23, 0xb8, 0xc7, 0x8b, 0xb7, 0xe2, 0x76, 0x1a, 0xbb,
	0xf5, 0xee, 0x8e, 0x91, 0xa9, 0xfb, 0x31, 0x3c, 0x5a, 0xe8, 0xf2, 0xaa, 0xd7, 0x6c, 0xb7, 0xda,
	0x48, 0x89, 0xa4, 0xea, 0x70, 0x7f, 0xa1, 0x5b, 0xbd, 0xd9, 0x44, 0x9f, 0x9c, 0xfa, 0x3e, 0x3c,
	0x5c, 0xe8, 0xd3, 0x34, 0x3a, 0x86, 0x89, 0x5e, 0x72, 0xe5, 0xdb, 0xe4, 0xc9, 0x48, 0xbf, 0x8a,
	0x77, 0x41, 0xdb, 0xef, 0x75, 0xf6, 0x5e, 0x61, 0xb0, 0x86, 0xd9, 0xee, 0x75, 0xad, 0x7e, 0x7d,
	0x50, 0x7f, 0x95, 0xd4, 0x78, 0x07, 0x6e, 0x5d, 0x64, 0x45, 0x1d, 0x61, 0x65, 0x9b, 0x70, 0xf3,
	0x42, 0x63, 0xb7, 0x94, 0xab, 0xfc, 0x2d, 0x41, 0x31, 0xfd, 0x7b, 0x01, 0x07, 0x1b, 0xfb, 0x32,
	0xd9, 0x9d, 0x72, 0xa0, 0xe1, 0xe7, 0x26, 0x0d, 0xf7, 0x8d, 0x6e, 0xb3, 0xdd, 0xdd, 0xc9, 0x04,
	0x17, 0x96, 0xfa, 0x7e, 0xbd, 0xdd, 0xa9, 0x6f, 0xa3, 0xde, 0x72, 0x4c, 0x40, 0x59, 0x9b, 0x69,
	0xd6, 0x1b, 0xbb, 0xac, 0xc7, 0x39, 0x53, 0xd3, 0x88, 0x4d, 0x4a, 0xaa, 0x97, 0x33, 0x93, 0xd9,
	0xd8, 0x65, 0xe9, 0xf2, 0x4c, 0xab, 0x19, 0xa3, 0xd8, 0x8a, 0xc2, 0x5c, 0x81, 0x09, 0x9b, 0x4b,
	0x95, 0x1f, 0x25, 0x58, 0xcd, 0xfc, 0x7e, 0xca, 0x86, 0x38, 0x5b, 0xcf, 0x7b, 0x70, 0xfb, 0x3c,
	0x6e, 0x5a, 0xfd, 0x81, 0x31, 0x34, 0xba, 0x6c, 0x59, 0x71, 0xa5, 0xb2, 0xe6, 0xbd, 0xbe, 0xd8,
	0x9d, 0x2c, 0xda, 0xec, 0x7d, 0xd9, 0xc5, 0xfe, 0xb2, 0xb4, 0x30, 0xdc, 0xc0, 0xc9, 0x33, 0x15,
	0x28, 0xdb, 0x77, 0xe1, 0x9a, 0xe3, 0x4f, 0xcf, 0xbf, 0x27, 0x7d, 0xe9, 0x8d, 0x8c, 0xff, 0x0e,
	0x0a, 0xfc, 0xa7, 0xe4, 0xf3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x66, 0xdb, 0x13, 0x61, 0x8f,
	0x0d, 0x00, 0x00,
}