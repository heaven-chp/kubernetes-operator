package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *Resource01) DeepCopyInto(out *Resource01) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

func (in *Resource01) DeepCopy() *Resource01 {
	if in == nil {
		return nil
	}
	out := new(Resource01)
	in.DeepCopyInto(out)
	return out
}

func (in *Resource01) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *Resource01List) DeepCopyInto(out *Resource01List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Resource01, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Resource01)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

func (in *Resource01List) DeepCopy() *Resource01List {
	if in == nil {
		return nil
	}
	out := new(Resource01List)
	in.DeepCopyInto(out)
	return out
}

func (in *Resource01List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *Resource01Spec) DeepCopyInto(out *Resource01Spec) {
	*out = *in
}

func (in *Resource01Spec) DeepCopy() *Resource01Spec {
	if in == nil {
		return nil
	}
	out := new(Resource01Spec)
	in.DeepCopyInto(out)
	return out
}
