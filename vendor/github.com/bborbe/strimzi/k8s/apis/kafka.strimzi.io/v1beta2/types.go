// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1beta2

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KafkaTopics []KafkaTopic

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KafkaTopic struct {
	// The specification of the topic.
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec *KafkaTopicSpec `json:"spec,omitempty"`

	// The status of the topic.
	Status *KafkaTopicStatus `json:"status,omitempty"`
}

func (t KafkaTopic) TopicName() string {
	if t.Spec != nil && t.Spec.TopicName != nil && *t.Spec.TopicName != "" {
		return *t.Spec.TopicName
	}
	return t.Name
}

func (t KafkaTopic) Equal(kafkaTopic KafkaTopic) bool {
	if t.TopicName() != kafkaTopic.TopicName() {
		return false
	}
	if t.Spec == nil && kafkaTopic.Spec == nil {
		return true
	}
	if t.Spec == nil || kafkaTopic.Spec == nil {
		return false
	}
	if t.Spec.Partitions != kafkaTopic.Spec.Partitions {
		return false
	}
	if t.Spec.Replicas != kafkaTopic.Spec.Replicas {
		return false
	}
	if !reflect.DeepEqual(t.Spec.Config, kafkaTopic.Spec.Config) {
		return false
	}
	return true
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of Kafka objects.
	Items []KafkaTopic `json:"items,omitempty"`
}

type KafkaTopicSpecs []KafkaTopicSpec

type KafkaTopicSpec struct {
	// The topic configuration.
	Config map[string]string `json:"config,omitempty"`

	// The number of partitions the topic should have. This cannot be decreased after
	// topic creation. It can be increased after topic creation, but it is important
	// to understand the consequences that has, especially for topics with semantic
	// partitioning. When absent this will default to the broker configuration for
	// `num.partitions`.
	Partitions *int32 `json:"partitions,omitempty"`

	// The number of replicas the topic should have. When absent this will default to
	// the broker configuration for `default.replication.factor`.
	Replicas *int32 `json:"replicas,omitempty"`

	// The name of the topic. When absent this will default to the metadata.name of
	// the topic. It is recommended to not set this unless the topic name is not a
	// valid Kubernetes resource name.
	TopicName *string `json:"topicName,omitempty"`
}

// The topic configuration.
//type KafkaTopicSpecConfig map[string]interface{}

// The status of the topic.
type KafkaTopicStatus struct {
	// List of status conditions.
	Conditions []KafkaTopicStatusConditionsElem `json:"conditions,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// Topic name.
	TopicName *string `json:"topicName,omitempty"`
}

type KafkaTopicStatusConditionsElem struct {
	// Last time the condition of a type changed from one status to another. The
	// required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about the condition's last
	// transition.
	Message *string `json:"message,omitempty"`

	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `json:"reason,omitempty"`

	// The status of the condition, either True, False or Unknown.
	Status *string `json:"status,omitempty"`

	// The unique identifier of a condition, used to distinguish between other
	// conditions in the resource.
	Type *string `json:"type,omitempty"`
}
