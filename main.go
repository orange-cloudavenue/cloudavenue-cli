/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package main

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/urfave/cli/v3"
)

func main() {
	subCommands := make([]*cli.Command, 0)
	subCommands = append(subCommands, cmdContext)
	subCommands = append(subCommands, builder()...)

	cmd := &cli.Command{
		Name:      "cloudavenue",
		Version:   "v1.0.0",
		Copyright: "(c) 2025 Orange Business",
		Usage:     "A unified CLI for managing Orange Cloudavenue resources and services.",
		UsageText: "cloudavenue [global options] <command> [command options] [arguments...]",
		ArgsUsage: "[command arguments]",
		Commands:  subCommands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "logger",
				Aliases: []string{"l"},
				Validator: func(level string) error {
					levels := []string{"info", "debug", "warn", "error", "fatal"}
					if !slices.Contains(levels, level) {
						return fmt.Errorf("must be one of %v", levels)
					}
					logLevel = level
					return nil
				},
				Local: true,
			},
		},
		EnableShellCompletion: true,
		HideHelp:              false,
		HideVersion:           false,
	}

	cmd.Run(context.Background(), os.Args)
}
