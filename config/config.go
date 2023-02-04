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

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env
	e := os.Getenv("APP_ENV")
	fmt.Println(e)
	envFile := ".env"
	if e != "" {
		envFile = e
	}
	godotenv.Load(envFile)

	Config = &config{}

	if err := env.Parse(&Config.Keycloak); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.Application); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.HTTPServer); err != nil {
		panic(err)
	}
}
