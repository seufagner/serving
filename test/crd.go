/*
Copyright 2018 The Knative Authors

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

package test

// crd contains functions that construct boilerplate CRD definitions.

import (
	"net/url"

	"knative.dev/pkg/test/helpers"
)

// ResourceNames holds names of various resources.
type ResourceNames struct {
	Config        string
	Route         string
	Revision      string
	Service       string
	TrafficTarget string
	URL           *url.URL
	Image         string
}

// AppendRandomString will generate a random string that begins with prefix. This is useful
// if you want to make sure that your tests can run at the same time against the same
// environment without conflicting. This method will seed rand with the current time when
// called for the first time.
var AppendRandomString = helpers.AppendRandomString

// ObjectNameForTest generates a random object name based on the test name.
var ObjectNameForTest = helpers.ObjectNameForTest
