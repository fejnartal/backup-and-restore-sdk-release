// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package integration_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestDatabaseBackupAndRestore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DatabaseBackupAndRestore Suite")
}

var compiledSDKPath string
var envVars map[string]string

var _ = BeforeSuite(func() {
	var err error
	compiledSDKPath, err = gexec.Build(
		"github.com/cloudfoundry-incubator/database-backup-and-restore/cmd/database-backup-restore")
	Expect(err).NotTo(HaveOccurred())
})

var _ = BeforeEach(func() {
	envVars = map[string]string{
		"PG_CLIENT_PATH":      "non-existent",
		"PG_DUMP_9_6_PATH":    "non-existent",
		"PG_DUMP_9_4_PATH":    "non-existent",
		"PG_RESTORE_9_4_PATH": "non-existent",
		"MYSQL_CLIENT_PATH":   "non-existent",
		"MYSQL_DUMP_PATH":     "non-existent",
	}
})
