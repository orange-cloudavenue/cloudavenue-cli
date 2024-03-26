package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
	"github.com/spf13/cobra"
)

const (
	get         string = "get"
	add         string = "add"
	delete      string = "delete"
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
	// ? Test configuration
	// Bad configuration
	t.Run("Bad Configuration", func(t *testing.T) {
		os.Setenv("ENDPOINT", "")
		if err := cmd.Execute(); err != nil {
			check := err.Error()
			if !strings.Contains(check, "Error in CloudAvenue parameter") {
				t.Errorf("Fail %v", err)
			}
		}
	})

	// Good configuration
	t.Run("Configuration", func(t *testing.T) {
		os.Setenv("CLOUDAVENUE_URL", "https://console1.cloudavenue.orange-business.com")
		if err := cmd.Execute(); err != nil {
			t.Errorf("Fail %v", err)
		}
	})

	// ? Test all commands
	allCmd := cmd.NewRootCmd().Commands()
	if len(allCmd) == 0 {
		panic("No command found")
	}

	for _, oneCmd := range allCmd {

		globalTests := tts{}
		switch oneCmd.Use {
		// ? Test list argument
		case get:
			// ? Test all subcommands
			subCmd := oneCmd.Commands()
			for _, cmdSubCmd := range subCmd {
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
					// TODO: this test seems to broke all the tests
					// {
					// 	name: oneCmd.Use + "_" + cmdSubCmd.Use + " with output flag and a wtf args",
					// 	args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "wtf"},
					// 	fail: true, // Should fail because args for flag is unknown
					// },
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
					},
				}
				// startTest(tests, t)
				globalTests = append(globalTests, tests...)
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
						},
					}
					globalTests = append(globalTests, tests...)
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
					globalTests = append(globalTests, tests...)
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
					globalTests = append(globalTests, tests...)
				default:
					fmt.Println("No test for this subcommand")
				}
			}

		// ? Test delete argument
		case delete:
			tests := tts{
				{
					name: delete + " without args",
					args: []string{delete},
					fail: true, // Should fail because no args
				},
				{
					name: delete + " with a whatever flag",
					args: []string{delete, "--whatever"},
					fail: true, // Should fail because flag no exist
				},
				{
					name: delete + " a bucket",
					args: []string{delete, "s3", "whatever", "whatever-time"},
					fail: false,
				},
				{
					name: delete + " a publicip",
					args: []string{delete, "publicip", "whatever"},
					fail: true,
				},
				{
					name: delete + " an edgegateway",
					args: []string{delete, "edgegateway", "tn01e02ocb0006205spt104"},
					fail: false,
				},
				{
					name: delete + " a vdc",
					args: []string{delete, "vdc", "whatever", "whatever-time"},
					fail: false,
				},
			}
			globalTests = append(globalTests, tests...)

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
			fmt.Println("No test for this command")
		}

		// ? Test all commands
		startTest(globalTests, t)

	}
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

			// print the command generated
			fmt.Printf("***Generated Test command: %v %v \n", x.CommandPath(), test.args)

			// execute the command
			err := x.Execute()
			if err != nil && !test.fail {
				t.Errorf("Fail %v", err)
			}

		})
	}
}

// func sort command with vdc in first
func sortCmd(cmds []*cobra.Command) []*cobra.Command {
	cmdsSorted := []*cobra.Command{}
	for _, cmd := range cmds {
		fmt.Println("===")
		fmt.Println(cmd.Use)
		switch cmd.Use {
		case vdc:
			cmdsSorted = append([]*cobra.Command{cmd}, cmdsSorted...)
		default:
			cmdsSorted = append(cmdsSorted, cmd)
		}
		fmt.Println(cmdsSorted)
	}
	return cmdsSorted
}
