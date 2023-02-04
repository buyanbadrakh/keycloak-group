// Copyright 2019 Muhammet Arslan <github.com/geass>
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

package config

import "time"

// Config stores configuration values
var Config *config

type config struct {

	// Token provides the token configurations.
	Keycloak struct {
		URL          string `env:"URL"`
		Realm        string `env:"REALM"`
		PublicKey    string `env:"PUBLIC_KEY"`
		GrantType    string `env:"GRANT_TYPE"`
		ClientID     string `env:"CLIENT_ID"`
		ClientSecret string `env:"CLIENT_SECRET"`
	}

	// Application provides the application configurations.
	Application struct {
		Name        string `env:"APPLICATION_NAME"`
		Environment string `env:"APPLICATION_ENVIRONMENT"`
		Version     string `env:"APPLICATION_VERSION"`
		APIv        string `env:"APPLICATION_API_VERSION"`
	}

	// HTTPServer provides the HTTP server configuration.
	HTTPServer struct {
		Listen               string        `env:"SERVER_LISTEN"`
		ReadTimeout          time.Duration `env:"SERVER_READ_TIMEOUT"`
		WriteTimeout         time.Duration `env:"SERVER_WRITE_TIMEOUT"`
		MaxConnsPerIP        int           `env:"SERVER_MAX_CONN_PER_IP"`
		MaxRequestsPerConn   int           `env:"SERVER_MAX_REQUESTS_PER_CONN"`
		MaxKeepaliveDuration time.Duration `env:"SERVER_MAX_KEEP_ALIVE_DURATION"`
	}
}
