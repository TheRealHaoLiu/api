// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/open-cluster-management/api/client/work/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AppliedManifestWorks returns a AppliedManifestWorkInformer.
	AppliedManifestWorks() AppliedManifestWorkInformer
	// ManifestWorks returns a ManifestWorkInformer.
	ManifestWorks() ManifestWorkInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AppliedManifestWorks returns a AppliedManifestWorkInformer.
func (v *version) AppliedManifestWorks() AppliedManifestWorkInformer {
	return &appliedManifestWorkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ManifestWorks returns a ManifestWorkInformer.
func (v *version) ManifestWorks() ManifestWorkInformer {
	return &manifestWorkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}