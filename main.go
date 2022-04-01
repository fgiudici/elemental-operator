/*
Copyright © 2022 SUSE LLC

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

package main

import (
	"flag"
	"os"
	"time"

	"github.com/rancher-sandbox/rancheros-operator/pkg/operator"
	"github.com/rancher-sandbox/rancheros-operator/pkg/services"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/sirupsen/logrus"
)

var (
	namespace = flag.String("namespace", "cattle-rancheros-operator-system", "Namespace of the pod")
)

func main() {
	flag.Parse()
	logrus.Info("Starting controller")
	ctx := signals.SetupSignalContext()

	if os.Getenv("NAMESPACE") != "" {
		*namespace = os.Getenv("NAMESPACE")
	}

	//TODO check the proper namespace configuration, should it be the same namespace as the rancheros-operator?
	if err := operator.Run(ctx,
		operator.WithNamespace(*namespace),
		operator.WithServices(services.UpgradeChannelSync(10*time.Second, "fleet-default")),
	); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	<-ctx.Done()
}
