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

package config

type Config struct {
	LogDir      string         `json:"logDir"`
	Smtp        Email          `json:"smtp"`
	Destination string         `json:"destination"`
	DstPerm     FilePermission `json:"dstPerm"`
	Jobs        []Job          `json:"jobs"`
}

type Email struct {
	Address   string   `json:"address"`
	Port      uint16   `json:"port"`
	UseTls    bool     `json:"useTls"`
	User      string   `json:"user"`
	Password  string   `json:"password"`
	Subject   string   `json:"subject"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
}

type FilePermission struct {
	User  string `json:"user"`
	Group string `json:"group"`
}

type Job struct {
	Name     string   `json:"name"`
	Database Database `json:"database"`
}

type Database struct {
	Hostname string `json:"hostname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Filter   string `json:"filter"`
}
