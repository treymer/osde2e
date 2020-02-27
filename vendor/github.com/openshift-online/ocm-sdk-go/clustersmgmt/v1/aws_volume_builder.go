/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AWSVolumeBuilder contains the data and logic needed to build 'AWS_volume' objects.
//
// Holds settings for an AWS storage volume.
type AWSVolumeBuilder struct {
	iops  *int
	size  *int
	type_ *string
}

// NewAWSVolume creates a new builder of 'AWS_volume' objects.
func NewAWSVolume() *AWSVolumeBuilder {
	return new(AWSVolumeBuilder)
}

// IOPS sets the value of the 'IOPS' attribute
// to the given value.
//
//
func (b *AWSVolumeBuilder) IOPS(value int) *AWSVolumeBuilder {
	b.iops = &value
	return b
}

// Size sets the value of the 'size' attribute
// to the given value.
//
//
func (b *AWSVolumeBuilder) Size(value int) *AWSVolumeBuilder {
	b.size = &value
	return b
}

// Type sets the value of the 'type' attribute
// to the given value.
//
//
func (b *AWSVolumeBuilder) Type(value string) *AWSVolumeBuilder {
	b.type_ = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AWSVolumeBuilder) Copy(object *AWSVolume) *AWSVolumeBuilder {
	if object == nil {
		return b
	}
	b.iops = object.iops
	b.size = object.size
	b.type_ = object.type_
	return b
}

// Build creates a 'AWS_volume' object using the configuration stored in the builder.
func (b *AWSVolumeBuilder) Build() (object *AWSVolume, err error) {
	object = new(AWSVolume)
	if b.iops != nil {
		object.iops = b.iops
	}
	if b.size != nil {
		object.size = b.size
	}
	if b.type_ != nil {
		object.type_ = b.type_
	}
	return
}
