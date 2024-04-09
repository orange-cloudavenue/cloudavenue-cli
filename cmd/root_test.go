package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	get         string = "get"
	add         string = "add"
	del         string = "del"
	help        string = "help"
	version     string = "version"
	update      string = "update"
	vdc         string = "vdc"
	s3          string = "s3"
	t0          string = "t0"
	edgeGateway string = "edgegateway"
	publicip    string = "publicip"
)

type tt struct {
	name string
	args []string
	fail bool
}
type tts []tt

func TestRootCmd(t *testing.T) {
	// TODO - Need to fix the export variables
	// ? Test configuration
	// Bad configuration
	t.Run("Bad Configuration", func(t *testing.T) {
		os.Setenv("CLOUDAVENUE_USERNAME", "TOTO")
		if err := cmd.Execute(); err != nil {
			check := err.Error()
			if !strings.Contains(check, "Please check your configuration") {
				t.Errorf("Fail %v", err)
			}
		}
	})

	// Good configuration
	fmt.Println("=Good Configuration")
	t.Run("Configuration", func(t *testing.T) {
		os.Setenv("CLOUDAVENUE_USERNAME", "gaetan.ars")
		if err := cmd.Execute(); err != nil {
			t.Errorf("Fail %v", err)
		}
	})

	// ? Test all commands
	allCmd := cmd.NewRootCmd().Commands()
	if len(allCmd) == 0 {
		panic("No command found")
	}

	var globalTests, addTests, getTests, delTests tts
	for _, oneCmd := range allCmd {
		// ? Test all subcommands
		switch oneCmd.Use {
		// ? Test list argument
		case get:
			// ? Test all subcommands
			subCmd := oneCmd.Commands()
			for _, cmdSubCmd := range subCmd {
				// ? Test all subcommands for Get and some use case
				tests := tts{
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
						args: []string{oneCmd.Use, cmdSubCmd.Use},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
						fail: true, // Should fail because flag no exist
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with ouput flag without args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output"},
						fail: true, // Should fail because args for flag is empty
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with ouput flag wide args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "wide"},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with ouput flag json args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "json"},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with ouput flag yaml args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "yaml"},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with resource name flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--name", "whatever"},
						fail: true, // Should fail because resource doesn't exist
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with output flag and a wtf args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "wtf"},
						fail: true, // Should fail because args for flag is unknown
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
					},
				}
				getTests = append(getTests, tests...)
			}
		// ? Test add argument
		case add:
			// ? Test all subcommands
			subCmd := oneCmd.Commands()
			subCmdSorted := sortCmd(subCmd)

			for _, cmdSubCmd := range subCmdSorted {
				switch cmdSubCmd.Use {
				case vdc, s3:
					tests := tts{
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
							args: []string{oneCmd.Use, cmdSubCmd.Use},
							fail: true, // Should fail because no args
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
							fail: true, // Should fail because flag no exist
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name", "whatever"},
							fail: false,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag and an empty args",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name"},
							fail: true, // Should fail because args for flag is empty
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--time", "--name", "whatever-time"},
							fail: false,
						},
					}
					addTests = append(addTests, tests...)
				case publicip:
					tests := tts{
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
							args: []string{oneCmd.Use, cmdSubCmd.Use},
							fail: true, // Should fail because no args
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
							fail: true, // Should fail because flag no exist
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag and time flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name", "whatever", "--time"},
							fail: true, // Should fail because to preserve IP
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag and an empty args",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name"},
							fail: true, // Should fail because args for flag is empty
						},
					}
					addTests = append(addTests, tests...)
				case edgeGateway:
					tests := tts{
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
							args: []string{oneCmd.Use, cmdSubCmd.Use},
							fail: true, // Should fail because no args
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
							fail: true, // Should fail because flag no exist
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with vdc flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc", "whatever"},
							fail: false,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with vdc flag and an empty args",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc"},
							fail: true, // Should fail because args for flag is empty
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag and vdc flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--time", "--vdc", "whatever-time"},
							fail: true, // Should fail because too much egw created
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
							fail: true, // Should fail because no vdc
						},
					}
					addTests = append(addTests, tests...)
				default:
					if cmdSubCmd.Use == "" {
						fmt.Printf("No test for this subcommand: %v", cmdSubCmd.Use)
					}
				}
			}

		// ? Test delete argument
		case del:
			tests := tts{
				{
					name: del + " without args",
					args: []string{del},
					fail: true, // Should fail because no args
				},
				{
					name: del + " with a whatever flag",
					args: []string{del, "--whatever"},
					fail: true, // Should fail because flag no exist
				},
				{
					name: del + " a bucket",
					args: []string{del, "s3", "whatever", "whatever-time"},
					fail: false,
				},
				{
					name: del + " a publicip",
					args: []string{del, "publicip", "whatever"},
					fail: true,
				},
				{
					name: del + " an edgegateway",
					args: []string{del, "edgegateway", "tn01e02ocb0006205spt104"},
					fail: false,
				},
				{
					name: del + " a vdc",
					args: []string{del, "vdc", "whatever", "whatever-time"},
					fail: false,
				},
			}
			delTests = append(delTests, tests...)

		// ? Test help argument
		case help:
			tests := tts{
				{
					name: oneCmd.Use,
					args: []string{oneCmd.Use},
					fail: false,
				},
			}
			globalTests = append(globalTests, tests...)

		// ? Test version argument
		case version:
			tests := tts{
				{
					name: oneCmd.Use,
					args: []string{oneCmd.Use},
					fail: false,
				},
			}
			globalTests = append(globalTests, tests...)

		// ? Test update argument
		case update:
			tests := tts{
				{
					name: oneCmd.Use,
					args: []string{oneCmd.Use},
					fail: true,
				},
			}
			globalTests = append(globalTests, tests...)

		default:
			if oneCmd.Use == "" {
				fmt.Printf("No test for this subcommand: %v", oneCmd.Use)
			}
		}
	}
	// ? Test all commands
	startTest(addTests, t)
	startTest(getTests, t)
	startTest(delTests, t)
	startTest(globalTests, t)

}

func startTest(tests tts, t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// create the root command
			x := cmd.NewRootCmd()

			// redirect the output to a buffer
			var stdout bytes.Buffer
			x.SetOut(&stdout)

			// set the args
			x.SetArgs(test.args)

			// execute the command
			err := x.Execute()
			if err != nil && test.fail == false {
				t.Fail()
			}

			// reset the flags
			resetFlags(x)
			x.SetOut(nil)
			x.SetErr(nil)

		})

	}
}

// func sort command with vdc in first
func sortCmd(cmds []*cobra.Command) []*cobra.Command {
	cmdsSorted := []*cobra.Command{}
	for _, cmd := range cmds {
		switch cmd.Use {
		case vdc:
			cmdsSorted = append([]*cobra.Command{cmd}, cmdsSorted...)
		default:
			cmdsSorted = append(cmdsSorted, cmd)
		}
	}
	return cmdsSorted
}

// func to reset all flags recursuvely
func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Changed {
			err := f.Value.Set(f.DefValue)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			f.Changed = false
		}
	})

	for _, c := range cmd.Commands() {
		resetFlags(c)
	}
}
