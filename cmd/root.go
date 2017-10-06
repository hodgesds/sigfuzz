// Copyright © 2017 Daniel Hodges <hodges.daniel.scott@gmail.com>
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

package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sigfuzz",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		stopChan := make(chan struct{})

		sigs, err := getSignals(viper.GetStringSlice("signal"))
		if err != nil {
			log.Fatalln(err)
		}

		processes, err := getProcesses(viper.GetStringSlice("pid"))
		if err != nil {
			log.Fatalln(err)
		}

		interval := viper.GetDuration("interval")
		number := viper.GetInt("number")
		exit := viper.GetBool("exit")

		var wg sync.WaitGroup

		for _, process := range processes {
			wg.Add(1)
			go func(process *os.Process) {
				defer wg.Done()
				err := fuzzProcess(
					stopChan,
					interval,
					exit,
					number,
					process,
					sigs...,
				)
				if err != nil {
					log.Println(err)
				}
			}(process)
		}

		wg.Wait()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config", "c",
		"",
		"config file (default is $HOME/.sigfuzz.yaml)",
	)

	RootCmd.PersistentFlags().DurationP(
		"interval", "i",
		1*time.Second,
		"time duration between signals",
	)

	RootCmd.PersistentFlags().StringSliceP(
		"signal", "s",
		[]string{},
		"signal to fuzz",
	)

	RootCmd.PersistentFlags().StringSliceP(
		"pid", "p",
		[]string{},
		"pids to fuzz",
	)

	RootCmd.PersistentFlags().IntP(
		"number", "n",
		1,
		"number of times to signal",
	)

	RootCmd.PersistentFlags().BoolP(
		"exit", "x",
		false,
		"exit on signal failure",
	)

	viper.BindPFlags(RootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".sigfuzz") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("reading in config: %s\n", cfgFile)
	}
}

func fuzzProcess(
	stopChan chan struct{},
	interval time.Duration,
	exit bool,
	number int,
	process *os.Process,
	signals ...os.Signal,
) error {
	t := time.NewTicker(interval)

	for {
		select {
		case <-t.C:
			for _, sig := range signals {
				if err := process.Signal(sig); err != nil {
					if exit {
						return err
					}
					log.Println(err)
				}
			}
			number = number - 1
			if number == 0 {
				return nil
			}
		case <-stopChan:
			return nil
		}
	}
}
