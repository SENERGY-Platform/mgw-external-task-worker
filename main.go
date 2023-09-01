/*
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"log"
	"os"
	"os/signal"
	"time"

	"syscall"
)

func main() {
	configLocation := flag.String("config", "config.json", "configuration file")
	flag.Parse()

	config, err := configuration.Load(*configLocation)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("wait for cluster routing")
	time.Sleep(10 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		sig := <-shutdown
		log.Println("received shutdown signal", sig)
		cancel()
	}()

	pkg.Start(ctx, config)
	log.Println("worker stopped")
}
