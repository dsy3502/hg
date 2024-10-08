// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/service_subscription_list.proto

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using ServiceSubscriptionList within kubernetes types, where deepcopy-gen is used.
func (in *ServiceSubscriptionList) DeepCopyInto(out *ServiceSubscriptionList) {
	p := proto.Clone(in).(*ServiceSubscriptionList)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSubscriptionList. Required by controller-gen.
func (in *ServiceSubscriptionList) DeepCopy() *ServiceSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(ServiceSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSubscriptionList. Required by controller-gen.
func (in *ServiceSubscriptionList) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceSubscription within kubernetes types, where deepcopy-gen is used.
func (in *ServiceSubscription) DeepCopyInto(out *ServiceSubscription) {
	p := proto.Clone(in).(*ServiceSubscription)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSubscription. Required by controller-gen.
func (in *ServiceSubscription) DeepCopy() *ServiceSubscription {
	if in == nil {
		return nil
	}
	out := new(ServiceSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSubscription. Required by controller-gen.
func (in *ServiceSubscription) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
