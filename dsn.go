// Copyright 2019 The SQLFlow Authors. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goalisa

// Config is the deserialization of connect string
type Config struct {
	AccessID  string
	AccessKey string
	Endpoint  string
}

// ParseDSN deserialize the connect string
func ParseDSN(dsn string) (*Config, error) {
	// TODO(weiguoz)
	return &Config{}, nil
}

// FormatDSN serialize a config to connect string
func (cfg *Config) FormatDSN() string {
	// TODO(weiguoz)
	return ""
}
