/*
Copyright © 2022 - 2024 SUSE LLC

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
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/rancher/elemental-operator/pkg/hostinfo"
)

func main() {
	var output string

	flag.StringVar(&output, "o", "labels", "output type [labels, raw]")
	flag.Parse()

	hData, err := hostinfo.Host()
	if err != nil {
		log.Fatalf("Cannot retrieve host data: %s", err)
	}

	var rawData interface{}

	switch output {
	case "raw":
		rawData = hData
	case "labels":
		dataMap, err := hostinfo.ExtractLabels(hData)
		if err != nil {
			log.Fatalf("Cannot convert host data to labels: %s", err)
		}
		rawData = dataMap
	}

	jsonData, err := json.MarshalIndent(rawData, "", "  ")
	if err != nil {
		log.Fatalf("Cannot convert raw data to json: %s", err)
	}
	fmt.Printf("%+v\n", string(jsonData))
}
