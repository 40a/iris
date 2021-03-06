package cmd

// Copyright © 2019 oleg2807@gmail.com
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

import (
	"fmt"
	"os"

	"github.com/olegsu/iris/pkg/logger"
)

func dieOnError(err error) {
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
}

func buildLogger(cmd string) logger.Logger {
	return logger.New(&logger.Options{
		Command: cmd,
		Verbose: true,
	})
}
