package main

import (
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/stooj/pulumi-talos-go-component/pkg/talos"
)

func main() {
	err := infer.NewProviderBuilder().
		WithName("talos-go-component").
		WithNamespace("stooj").
		WithComponents(
			infer.Component(talos.NewTalosCluster),
		).
		BuildAndRun()
	if err != nil {
		panic(err)
	}
}
