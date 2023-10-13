package main

// This should import instead, but they have incompatibilities
// with latest kustomize at the moment
// "github.com/argoproj/argo-cd/v2/cmpserver/plugin"
// "github.com/argoproj/argo-cd/v2/reposerver/apiclient"

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConfigManagementPluginKind is the string to use as Kind: in the CMP
	// plugin.yaml
	ConfigManagementPluginKind string = "ConfigManagementPlugin"
)

// PluginConfig is the structure representing a plugin.yaml configuration object
type PluginConfig struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   metav1.ObjectMeta `json:"metadata"`
	Spec       PluginConfigSpec  `json:"spec"`
}

// PluginConfigSpec is the spec from a PluginConfig
type PluginConfigSpec struct {
	Version          string     `json:"version"`
	Init             Command    `json:"init,omitempty"`
	Generate         Command    `json:"generate"`
	Parameters       Parameters `yaml:"parameters"`
	PreserveFileMode bool       `json:"preserveFileMode,omitempty"`
}

// Command holds binary path and arguments list
type Command struct {
	Command []string `json:"command,omitempty"`
	Args    []string `json:"args,omitempty"`
}

// Parameters holds static and dynamic configurations
type Parameters struct {
	Static  []*ParameterAnnouncement `yaml:"static"`
	Dynamic Command                  `yaml:"dynamic"`
}

// ParameterAnnouncement is a description for a single parameter
type ParameterAnnouncement struct {
	// name is the name identifying a parameter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// title is a human-readable text of the parameter name.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// tooltip is a human-readable description of the parameter.
	Tooltip string `protobuf:"bytes,3,opt,name=tooltip,proto3" json:"tooltip,omitempty"`
	// required defines if this given parameter is mandatory.
	Required bool `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	// itemType determines the primitive data type represented by the parameter. Parameters are always encoded as
	// strings, but this field lets them be interpreted as other primitive types.
	ItemType string `protobuf:"bytes,5,opt,name=itemType,proto3" json:"itemType,omitempty"`
	// collectionType is the type of value this parameter holds - either a single value (a string) or a collection
	// (array or map). If collectionType is set, only the field with that type will be used. If collectionType is not
	// set, `string` is the default. If collectionType is set to an invalid value, a validation error is thrown.
	CollectionType string `protobuf:"bytes,6,opt,name=collectionType,proto3" json:"collectionType,omitempty"`
	// string is the default value of the parameter if the parameter is a string.
	StringDefault string `protobuf:"bytes,7,opt,name=string,proto3" json:"string,omitempty"`
	// array is the default value of the parameter if the parameter is an array.
	Array []string `protobuf:"bytes,8,rep,name=array,proto3" json:"array,omitempty"`
	// map is the default value of the parameter if the parameter is a map.
	Map map[string]string `protobuf:"bytes,9,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
