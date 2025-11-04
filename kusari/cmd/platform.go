// Copyright (c) Kusari <https://www.kusari.dev/>
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/spf13/cobra"
)

func Platform() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platform",
		Short: "Platform actions",
		Long:  "Handle Kusari platform actions",
	}

	cmd.AddCommand(sbomcompare())

	return cmd
}
