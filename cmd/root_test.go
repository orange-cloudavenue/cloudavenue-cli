package cmd_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
)

const (
	get    string = "get"
	create string = "create"
	delete string = "delete"
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

		// ? Test all subcommands
		subCmd := oneCmd.Commands()
		for _, cmdSubCmd := range subCmd {
			tests := tts{}
			switch cmdSubCmd.Use {
			// ? Test list argument
			case get:
				tests = tts{
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
						args: []string{oneCmd.Use, cmdSubCmd.Use},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
						fail: true,
					},
				}

			// ? Test delete argument
			case delete:
				tests = tts{
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
						args: []string{oneCmd.Use, cmdSubCmd.Use},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever argument",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "whatever"},
					},
				}

			// ? Test all subcommands
			case create:
				tests = tts{
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " without flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a unknow flag with argument",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--unknow", "whatever"},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a good flag without argument",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--name"},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a good flag with argument",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--name", "whatever"},
					},
				}
				// if strings.Contains(oneCmd.Use, cmd.CmdEdgeGateway) {
				// 	tests[3] = tt{
				// 		name: oneCmd.Use + "_" + cmdSubCmd.Use + " with good flags with argument",
				// 		args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc", "whatever", "--t0", "1"},
				// 		fail: true,
				// 	}
				// }

			}
			startTest(tests, t)

		}

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

			// execute the command
			err := x.Execute()
			if err != nil && !test.fail {
				t.Errorf("Fail %v", err)
			}

		})
	}
}
