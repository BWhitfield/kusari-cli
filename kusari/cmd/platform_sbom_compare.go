// Copyright (c) Kusari <https://www.kusari.dev/>
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"strconv"

	"github.com/kusaridev/kusari-cli/pkg/platform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	comparecmd.Flags().StringVarP(&tenantUrl, "tenant-url", "", "https://kusari.api.us.kusari.cloud/", "tenant url")
	comparecmd.Flags().StringVarP(&outputFormat, "output-format", "", "markdown", "output format (markdown or sarif)")

	// Bind flags to viper
	mustBindPFlag("tenant-url", comparecmd.Flags().Lookup("tenant-url"))
	mustBindPFlag("output-format", comparecmd.Flags().Lookup("output-format"))
}

func sbomcompare() *cobra.Command {
	comparecmd.RunE = func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		// First pass we want json. Later we'll want to respond with something better...
		// if outputFormat != "markdown" && outputFormat != "sarif" {
		// 	return fmt.Errorf("invalid output format: %s (must be 'markdown' or 'sarif')", outputFormat)
		// }

		softwareId, e := convertArgs(args, 0)
		if e != nil {
			return e
		}

		sbomIdA, e := convertArgs(args, 1)
		if e != nil {
			return e
		}
		sbomIdB, e := convertArgs(args, 2)
		if e != nil {
			return e
		}

		return platform.SbomCompare(tenantUrl, outputFormat, softwareId, sbomIdA, sbomIdB)
	}

	return comparecmd
}

var comparecmd = &cobra.Command{
	Use:   "sbom-compare <sbom_id_0> <sbom_id_1>",
	Short: "Compares the contents of 2 sboms",
	Long: `Compares the contents of 2 sboms. This is useful for build system & AI integration + policy enforcement. 
    <software_id>  software id from the Kusari platform
    <sbom_id_0>  sbom id from the Kusari platform
    <sbom_id_1>  sbom id from the Kusari platform`,
	Args: cobra.ExactArgs(3),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Update from viper (this gets env vars + config + flags)
		tenantUrl = viper.GetString("tenant-url")
		outputFormat = viper.GetString("output-format")
	},
}

func convertArgs(args []string, index int) (int, error) {
	arg := args[index]

	if len(args)-1 < index {
		return 0, fmt.Errorf("Argument at index %d must have value.", index)
	}

	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("Argument at index 0 must be a positive integer ('%s'): %w", arg, err)
	}

	if id < 0 {
		return 0, fmt.Errorf("Argument at index %d must be a positive integer", index)
	}

	return id, nil
}
