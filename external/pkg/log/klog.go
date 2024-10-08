// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	goflag "flag"
	"sync"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

var (
	KlogScope     = RegisterScope("klog", "", 0)
	configureKlog = sync.Once{}
)

// EnableKlogWithCobra enables klog to work with cobra / pflags.
// k8s libraries like client-go use klog.
func EnableKlogWithCobra() {
	gf := klogVerboseFlag()
	pflag.CommandLine.AddFlag(pflag.PFlagFromGoFlag(
		&goflag.Flag{
			Name:     "vklog",
			Value:    gf.Value,
			DefValue: gf.DefValue,
			Usage:    gf.Usage + ". Like -v flag. ex: --vklog=9",
		}))
}

// isKlogVerbose returns true if klog verbosity is non-zero.
func klogVerbose() bool {
	gf := klogVerboseFlag()
	return gf.Value.String() != "0"
}

// KlogVerboseFlag returns verbose flag from the klog library.
// After parsing it contains the parsed verbosity value.
func klogVerboseFlag() *goflag.Flag {
	fs := &goflag.FlagSet{}
	klog.InitFlags(fs)
	// --v= flag of klog.
	return fs.Lookup("v")
}
