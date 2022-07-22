//go:build e2e
// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

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

package e2e

import (
	"context"
	"os"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/cluster-api/test/framework"
)

// GCPGPUSpecInput is the input for GCPGPUSpec.
type GCPGPUSpecInput struct {
	BootstrapClusterProxy framework.ClusterProxy
	Namespace             *corev1.Namespace
	ClusterName           string
	SkipCleanup           bool
}

func GCPGPUSpec(ctx context.Context, inputGetter func() GCPGPUSpecInput)  {
	var (
		specName    = "gcp-gpu"
		input       GCPGPUSpecInput
		machineType = os.Getenv("GCP_GPU_NODE_MACHINE_TYPE")
	)

	input = inputGetter()
	Expect(input.Namespace).NotTo(BeNil(), "Invalid argument. input.Namespace can't be nil when calling %s spec", specName)
	Expect(input.ClusterName).NotTo(BeEmpty(), "Invalid argument. input.ClusterName can't be empty when calling %s spec", specName)
	if machineType != "" {
		Expect(machineType).To(HavePrefix("Standard_N"), "AZURE_GPU_NODE_MACHINE_TYPE is \"%s\" which isn't a GPU SKU in %s spec", machineType, specName)
	}

	By("creating a Kubernetes client to the workload cluster")
	clusterProxy := input.BootstrapClusterProxy.GetWorkloadCluster(ctx, input.Namespace.Name, input.ClusterName)
}
