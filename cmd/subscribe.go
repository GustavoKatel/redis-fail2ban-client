// Copyright © 2017 Gustavo Sampaio <gbritosampaio@gmail.com>
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
	// "time"

	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe [channel]",
	Short: "Subscribe to a pub/sub redis channel",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		pubsub := RedisClient.Subscribe(args[0])
		defer pubsub.Close()

		// subscr, err := pubsub.ReceiveTimeout(time.Second)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(subscr)

		for msg, err := pubsub.ReceiveMessage(); err == nil; msg, err = pubsub.ReceiveMessage() {
			fmt.Println(msg)
			// TODO: fail2go
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
