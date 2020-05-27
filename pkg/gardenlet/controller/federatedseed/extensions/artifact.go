// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extensions

import (
	"fmt"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	extensionsinformers "github.com/gardener/gardener/pkg/client/extensions/informers/externalversions"

	dnsv1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	dnsinformers "github.com/gardener/external-dns-management/pkg/client/dns/informers/externalversions"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// controllerArtifacts bundles a list of artifacts for extension kinds.
type controllerArtifacts struct {
	artifacts                           map[string]*artifact
	controllerInstallationRequiredQueue workqueue.RateLimitingInterface
	hasSyncedFuncs                      []cache.InformerSynced
}

// artifact is specified for extension kinds.
type artifact struct {
	newObjFunc     func() runtime.Object
	newListObjFunc func() runtime.Object
	informer       cache.SharedIndexInformer

	controllerInstallationExtensionQueue       workqueue.RateLimitingInterface
	controllerInstallationUpdateEventPredicate func(newObj, oldObj interface{}) bool

	shootStateQueue                workqueue.RateLimitingInterface
	shootStateUpdateEventPredicate func(newObj, oldObj interface{}) bool
}

func (c *controllerArtifacts) initialize(dnsInformers dnsinformers.SharedInformerFactory, extensionsInformers extensionsinformers.SharedInformerFactory) {
	var (
		dnsProviderInformer = dnsInformers.Dns().V1alpha1().DNSProviders()

		backupBucketInformer          = extensionsInformers.Extensions().V1alpha1().BackupBuckets()
		backupEntryInformer           = extensionsInformers.Extensions().V1alpha1().BackupEntries()
		containerRuntimeInformer      = extensionsInformers.Extensions().V1alpha1().ContainerRuntimes()
		controlPlaneInformer          = extensionsInformers.Extensions().V1alpha1().ControlPlanes()
		extensionInformer             = extensionsInformers.Extensions().V1alpha1().Extensions()
		infrastructureInformer        = extensionsInformers.Extensions().V1alpha1().Infrastructures()
		networkInformer               = extensionsInformers.Extensions().V1alpha1().Networks()
		operatingSystemConfigInformer = extensionsInformers.Extensions().V1alpha1().OperatingSystemConfigs()
		workerInformer                = extensionsInformers.Extensions().V1alpha1().Workers()
	)

	c.registerExtensionControllerArtifacts(
		dnsv1alpha1.DNSProviderKind,
		&artifact{
			newObjFunc:     func() runtime.Object { return &dnsv1alpha1.DNSProvider{} },
			newListObjFunc: func() runtime.Object { return &dnsv1alpha1.DNSProviderList{} },
			informer:       dnsProviderInformer.Informer(),
			controllerInstallationUpdateEventPredicate: func(newObj, oldObj interface{}) bool {
				var (
					newExtensionObj, ok1 = newObj.(*dnsv1alpha1.DNSProvider)
					oldExtensionObj, ok2 = oldObj.(*dnsv1alpha1.DNSProvider)
				)
				return ok1 && ok2 && oldExtensionObj.Spec.Type != newExtensionObj.Spec.Type
			},
		},
		false, true,
	)

	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.BackupBucketResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.BackupBucket{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.BackupBucketList{} },
			informer:       backupBucketInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
		},
		false, true,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.BackupEntryResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.BackupEntry{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.BackupEntryList{} },
			informer:       backupEntryInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.ContainerRuntimeResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.ContainerRuntime{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.ContainerRuntimeList{} },
			informer:       containerRuntimeInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.ControlPlaneResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.ControlPlane{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.ControlPlaneList{} },
			informer:       controlPlaneInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.ExtensionResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.Extension{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.ExtensionList{} },
			informer:       extensionInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.InfrastructureResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.Infrastructure{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.InfrastructureList{} },
			informer:       infrastructureInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.NetworkResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.Network{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.NetworkList{} },
			informer:       networkInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.OperatingSystemConfigResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.OperatingSystemConfig{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.OperatingSystemConfigList{} },
			informer:       operatingSystemConfigInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
	c.registerExtensionControllerArtifacts(
		extensionsv1alpha1.WorkerResource,
		&artifact{
			newObjFunc:     func() runtime.Object { return &extensionsv1alpha1.Worker{} },
			newListObjFunc: func() runtime.Object { return &extensionsv1alpha1.WorkerList{} },
			informer:       workerInformer.Informer(),
			controllerInstallationUpdateEventPredicate: extensionTypeChanged,
			shootStateUpdateEventPredicate:             extensionStateOrResourcesChanged,
		},
		false, false,
	)
}

func (c *controllerArtifacts) registerExtensionControllerArtifacts(kind string, artifact *artifact, disableControllerInstallationControl, disableShootStateSyncControl bool) {
	c.hasSyncedFuncs = append(c.hasSyncedFuncs, artifact.informer.HasSynced)

	if !disableControllerInstallationControl {
		artifact.controllerInstallationExtensionQueue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), fmt.Sprintf("controllerinstallation-extension-%s", kind))
	}
	if !disableShootStateSyncControl {
		artifact.shootStateQueue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), fmt.Sprintf("shootstate-%s", kind))
	}

	c.artifacts[kind] = artifact
}

func (c *controllerArtifacts) addControllerInstallationEventHandlers() {
	for _, artifact := range c.artifacts {
		if artifact.controllerInstallationExtensionQueue == nil {
			continue
		}

		artifact.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc:    createEnqueueFunc(artifact.controllerInstallationExtensionQueue),
			UpdateFunc: createEnqueueOnUpdateFunc(artifact.controllerInstallationExtensionQueue, artifact.controllerInstallationUpdateEventPredicate),
			DeleteFunc: createEnqueueFunc(artifact.controllerInstallationExtensionQueue),
		})
	}
}

func (c *controllerArtifacts) addShootStateEventHandlers() {
	for _, artifact := range c.artifacts {
		if artifact.shootStateQueue == nil {
			continue
		}

		artifact.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc:    createEnqueueFunc(artifact.shootStateQueue),
			UpdateFunc: createEnqueueOnUpdateFunc(artifact.shootStateQueue, artifact.shootStateUpdateEventPredicate),
		})
	}
}

func (c *controllerArtifacts) shutdownQueues() {
	if c.controllerInstallationRequiredQueue != nil {
		c.controllerInstallationRequiredQueue.ShutDown()
	}

	for _, artifact := range c.artifacts {
		if artifact.controllerInstallationExtensionQueue != nil {
			artifact.controllerInstallationExtensionQueue.ShutDown()
		}
		if artifact.shootStateQueue != nil {
			artifact.shootStateQueue.ShutDown()
		}
	}
}
