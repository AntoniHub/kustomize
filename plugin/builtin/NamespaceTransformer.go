// Code generated by pluginator on NamespaceTransformer; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
	"sigs.k8s.io/kustomize/pkg/transformers"
	"sigs.k8s.io/kustomize/pkg/transformers/config"
	"sigs.k8s.io/yaml"
)

// Change or set the namespace of non-cluster level resources.
type NamespaceTransformerPlugin struct {
	Namespace  string             `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	FieldSpecs []config.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

//noinspection GoUnusedGlobalVariable
func NewNamespaceTransformerPlugin() *NamespaceTransformerPlugin {
	return &NamespaceTransformerPlugin{}
}

func (p *NamespaceTransformerPlugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, c []byte) (err error) {
	p.Namespace = ""
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *NamespaceTransformerPlugin) Transform(m resmap.ResMap) error {
	return transformers.NewNamespaceTransformer(
		p.Namespace, p.FieldSpecs).Transform(m)
}