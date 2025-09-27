package main

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/commands"
	"github.com/urfave/cli/v3"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/api/draas/v1"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/api/edgegateway/v1"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/api/organization/v1"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/api/vdc/v1"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/api/vdcgroup/v1"
)

func builder() (cliCommands []*cli.Command) {
	for _, ns := range commands.NewRegistry().GetNamespaces() {
		nsDetail := commands.NewRegistry().Get(ns, "", "")

		cliNsCmd := cli.Command{
			// High level command info
			Name:        strings.ToLower(nsDetail.Namespace),
			Usage:       nsDetail.ShortDocumentation,
			Description: nsDetail.LongDocumentation,
			Action:      nil,
			Metadata: map[string]any{
				"command": nsDetail,
			},
		}

		// Get all cav top commands (get, list, create, update, delete, ...) for this namespace
		cmdNsTop := commands.NewRegistry().GetCommandsByFilter(func(cmd commands.Command) bool {
			// List high level commands (create, update, get, delete, list, ...)
			return cmd.GetNamespace() == ns && cmd.GetResource() == "" && cmd.GetVerb() != ""
		})

		// Get all cav sub commands (vdc create, vdc delete, edgegateway list, t0 get, etc.) for this namespace
		cmdNsSub := commands.NewRegistry().GetCommandsByFilter(func(cmd commands.Command) bool {
			// List sub commands (create firewall-rule, update security-group, etc.)
			return cmd.GetNamespace() == ns && cmd.GetResource() != "" && cmd.GetVerb() != ""
		})

		// cliSubCommands holds the top level commands for this namespace
		cliSubCommands := []*cli.Command{}

		// Convert cav top commands to cli commands
		for _, cmd := range cmdNsTop {
			cliSubCommands = append(cliSubCommands, cavCmdToCliCommand(cmd))
		}

		// Group sub commands by resource (e.g., firewall-rule, security-group, etc.)
		resources := map[string][]commands.Command{}
		for _, cmd := range cmdNsSub {
			resources[cmd.GetResource()] = append(resources[cmd.GetResource()], cmd)
		}

		// For each resource, create a sub-command with its verbs as sub-commands
		for resource, cmds := range resources {
			resourceSubCommands := []*cli.Command{}
			for _, cmd := range cmds {
				resourceSubCommands = append(resourceSubCommands, cavCmdToCliCommand(cmd))
			}

			cliCmdResource := cli.Command{
				// Ex vdc storageprofiles
				Name:        strings.ToLower(resource),
				Usage:       fmt.Sprintf("Manage %s %s", nsDetail.Namespace, resource),
				Description: fmt.Sprintf("Commands to manage %s %s", nsDetail.Namespace, resource),
				// Category:    strings.ToLower(resource),
				Commands: resourceSubCommands,
				Metadata: map[string]any{
					"command": nsDetail,
				},
			}

			cliSubCommands = append(cliSubCommands, &cliCmdResource)
		}

		cliNsCmd.Commands = cliSubCommands

		cliCommands = append(cliCommands, &cliNsCmd)
	}

	return cliCommands
}

func cavCmdToCliCommand(cavcmd commands.Command) *cli.Command {
	return &cli.Command{
		Name:        strings.ToLower(cavcmd.Verb),
		Usage:       cavcmd.ShortDocumentation,
		Description: cavcmd.LongDocumentation,
		Flags:       paramsSpecsToCliFlags(cavcmd),
		Before: cli.BeforeFunc(func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			// || slices.ContainsFunc(cavcmd.ParamsSpecs, func(p commands.ParamsSpec) bool {
			// 	return p.Required
			// })

			if slices.ContainsFunc(cmd.Args().Slice(), func(a string) bool {
				return a == "--help" || a == "-h"
			}) {
				// Show help information
				return nil, cli.ShowSubcommandHelp(cmd)
			}

			// Ensure client is loaded
			if client == nil {
				if err := loadCavClient(); err != nil {
					return nil, cli.Exit(fmt.Sprintf("Error loading CloudAvenue client: %v", err), 1)
				}
			}

			// * CLI Flags to command params mapping

			// Parse command line arguments to extract flag values
			// []string{"--name","example"}
			commandParams := map[string]string{}

			for i, f := range cmd.Args().Slice() {
				if strings.HasPrefix(f, "--") {
					flagName := strings.TrimPrefix(f, "--")
					if i+1 < len(cmd.Args().Slice()) {
						flagValue := cmd.Args().Slice()[i+1]
						if strings.HasPrefix(flagValue, "--") {
							// Next argument is another flag, so this is a boolean flag
							commandParams[flagName] = "true"
						} else {
							commandParams[flagName] = flagValue
						}
					} else {
						// No next argument, so this is a boolean flag
						commandParams[flagName] = "true"
					}
				}
			}

			// Load command parameters into rVal
			// Init rVal with reflect Value nil
			rVal := reflect.ValueOf(nil)

			if cavcmd.ParamsType != nil {
				rType := reflect.TypeOf(cavcmd.ParamsType)
				// Override rVal with a new instance of the command's ParamsType
				rVal = reflect.New(rType).Elem()

				// for each param in commandParams, set the value in rVal
				for paramName, paramValue := range commandParams {
					if err := commands.StoreValueAtPath(rVal.Addr().Interface(), paramName, paramValue); err != nil {
						return nil, cli.Exit(fmt.Sprintf("Error storing parameter value: %v", err), 1)
					}
				}
			}

			cmd.Metadata["rVal"] = rVal

			return ctx, nil
		}),
		Metadata: map[string]any{
			"command": cavcmd,
		},
		Action:          runCommand,
		SkipFlagParsing: true,
	}
}

func paramsSpecsToCliFlags(cmd commands.Command) (flags []cli.Flag) {
	// Helper to determine flag category
	category := func(param commands.ParamsSpec) string {
		if param.Required {
			return "Required"
		}
		return "Optional"
	}

	// Reflect ParamType to determine flag type
	for _, param := range cmd.ParamsSpecs {
		pT, err := commands.GetParamType(reflect.TypeOf(cmd.ParamsType), param.Name)
		if err != nil {
			fmt.Printf("Error getting param type for %s: %v\n", param.Name, err)
			continue
		}

		switch pT.Kind() {
		case reflect.Bool:
			flags = append(flags, &cli.BoolFlag{
				Name:     param.Name,
				Usage:    param.Description,
				Required: param.Required,
				Category: category(param),
			})
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			flags = append(flags, &cli.IntFlag{
				Name:     param.Name,
				Usage:    param.Description,
				Required: param.Required,
				Category: category(param),
			})
		case reflect.Float32, reflect.Float64:
			flags = append(flags, &cli.Float64Flag{
				Name:     param.Name,
				Usage:    param.Description,
				Required: param.Required,
				Category: category(param),
			})
		case reflect.String:
			flags = append(flags, &cli.StringFlag{
				Name:     param.Name,
				Usage:    param.Description,
				Required: param.Required,
				Category: category(param),
			})
		case reflect.Slice:
			switch pT.Elem().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				flags = append(flags, &cli.IntSliceFlag{
					Name:     param.Name,
					Usage:    fmt.Sprintf("%s (multiple values allowed)", param.Description),
					Required: param.Required,
					Category: category(param),
				})
			case reflect.Float32, reflect.Float64:
				flags = append(flags, &cli.Float64SliceFlag{
					Name:     param.Name,
					Usage:    fmt.Sprintf("%s (multiple values allowed)", param.Description),
					Required: param.Required,
					Category: category(param),
				})
			default:
				continue
			}
		default:
			// Fallback to string flag for other types (e.g., slices, structs)
			flags = append(flags, &cli.StringFlag{
				Name:     param.Name,
				Usage:    fmt.Sprintf("%s (type: %s)", param.Description, pT.Kind().String()),
				Required: param.Required,
				Category: category(param),
			})
		}
	}

	return flags
}

func runCommand(ctx context.Context, cmd *cli.Command) error {
	// If the client is not initialized, it means we are in help mode
	// and we do nothing
	// because the command's Before has already displayed the help
	if client == nil {
		return nil
	}

	// Retrieve the cav command from metadata
	cavcmd := cmd.Metadata["command"].(commands.Command)

	// Initialize sub-client based on namespace
	var cmdClient any
	switch strings.ToLower(cavcmd.GetNamespace()) {
	case "vdc":
		cmdClient, _ = vdc.New(client.c)
	case "edgegateway", "t0":
		cmdClient, _ = edgegateway.New(client.c)
	case "vdcgroup":
		cmdClient, _ = vdcgroup.New(client.c)
	case "draas":
		cmdClient, _ = draas.New(client.c)
	case "organization":
		cmdClient, _ = organization.New(client.c)
	default:
		return cli.Exit(fmt.Sprintf("Unsupported namespace: %s", cavcmd.GetNamespace()), 1)
	}

	// Call the command's RunnerFunc if defined
	if cavcmd.RunnerFunc == nil {
		return cli.Exit(fmt.Errorf("No action defined for command %s %s %s\n", cavcmd.GetNamespace(), cavcmd.GetResource(), cavcmd.GetVerb()), 1)
	}

	var (
		result any
		err    error
	)

	// Retrieve rVal from metadata (set in Before function)
	rVal := cmd.Metadata["rVal"].(reflect.Value)

	if !rVal.IsValid() {
		result, err = cavcmd.Run(context.Background(), cmdClient, nil)
	} else {
		result, err = cavcmd.Run(context.Background(), cmdClient, rVal.Interface())
	}
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	return Print(result, cmd)
}
