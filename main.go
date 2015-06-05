/*
 * Copyright 2015 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/skarllot/mydump/config"
)

const (
	CONFIG_FILE_NAME = "config.json"
)

func main() {
	content, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		fmt.Println("Could not load configuration file:", err)
		return
	}

	cfg := &config.Config{}
	if err := json.Unmarshal(content, cfg); err != nil {
		fmt.Println("Invalid configuration file:", err)
		return
	}

	if b, err := pathExists(cfg.Destination); !b {
		if err != nil {
			fmt.Println("Could not stat destination directory:", err)
			return
		}
		if err := os.MkdirAll(cfg.Destination, 0750); err != nil {
			fmt.Println("Could not create destination directory:", err)
			return
		}
	}

	for _, job := range cfg.Jobs {
		dstPath := path.Join(cfg.Destination, job.Database.Hostname)
		if b, err := pathExists(dstPath); !b {
			if err != nil {
				fmt.Println("Could not stat destination directory:", err)
				return
			}
			if err := os.Mkdir(dstPath, 0750); err != nil {
				fmt.Println("Could not create destination directory:", err)
				return
			}
		}
	}
	
	fmt.Println("End")
}

func pathExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
