package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/cav"
	"github.com/orange-cloudavenue/common-go/print"
	"github.com/urfave/cli/v3"
	"github.com/zalando/go-keyring"
)

type clientInfo struct {
	Context contextInfo
	c       cav.Client
}

var client *clientInfo

var (
	cmdContext = &cli.Command{
		Name:        "context",
		Usage:       "Manage contexts organizations",
		Description: "Login and switch between multiple organizations",
		Category:    "general",
		Flags:       []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List available contexts",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					contexts, err := Contexts()
					if err != nil {
						if err == keyring.ErrNotFound {
							fmt.Fprintf(cmd.Root().Writer, "No context found. Please login first.\n")
							return nil
						}
						return err
					}

					x := print.New(print.WithOutput(cmd.Root().Writer))
					x.SetHeader("Name", "Organization", "Active", "Last Used", "Created At")
					for _, ctx := range contexts.Items {
						active := ""
						if contexts.Current == ctx.Name {
							active = "*"
						}
						x.AddFields(ctx.Name, ctx.Organization, active, ctx.LastUsedAt.String(), ctx.CreatedAt.String())
					}

					x.PrintTable()
					return nil
				},
			},
			{
				Name:  "use",
				Usage: "Switch to a different context",
				ShellComplete: cli.ShellCompleteFunc(func(ctx context.Context, cmd *cli.Command) {
					contexts, err := Contexts()
					if err != nil {
						return
					}

					var orgs []string
					for _, ctx := range contexts.Items {
						orgs = append(orgs, ctx.Organization)
					}

					// Suggest organizations that are not already used in the current context
					slices.Sort(orgs)
					for _, org := range orgs {
						fmt.Fprintf(cmd.Root().Writer, "%s\n", org)
					}
				}),
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "name", UsageText: "Context name"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					contexts, err := Contexts()
					if err != nil {
						if err == keyring.ErrNotFound {
							fmt.Fprintf(cmd.Root().Writer, "No context found. Please login first.\n")
							return nil
						}
						return err
					}

					if len(contexts.Items) == 0 {
						fmt.Fprintf(cmd.Root().Writer, "No context found. Please login first.\n")
						return nil
					}
					name := cmd.StringArg("name")
					contexts.SetCurrent(name)
					err = contexts.Save()
					if err != nil {
						return err
					}

					if err := loadCavClient(); err != nil {
						fmt.Fprintf(cmd.Root().Writer, "Failed to load client: %v\n", err)
						return err
					}

					fmt.Fprintf(cmd.Root().Writer, "Switched to context: %s\n", name)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a context. Logout from an organization",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "name", UsageText: "Context name"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := cmd.StringArg("name")
					if name == "" {
						return fmt.Errorf("context name is required")
					}

					contexts, err := Contexts()
					if err != nil {
						if err == keyring.ErrNotFound {
							fmt.Fprintf(cmd.Root().Writer, "No context found. Please login first.\n")
							return nil
						}
						return err
					}
					if contexts == nil || len(contexts.Items) == 0 {
						fmt.Fprintf(cmd.Root().Writer, "No context found. Please login first.\n")
						return nil
					}

					contexts.Remove(name)
					err = contexts.Save()
					if err != nil {
						return err
					}

					fmt.Fprintf(cmd.Root().Writer, "Context '%s' deleted\n", name)
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create a new context. Login to an organization",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "Context name", Required: true},
					&cli.StringFlag{Name: "org", Aliases: []string{"o"}, Usage: "Organization name", Required: true},
					&cli.StringFlag{Name: "user", Aliases: []string{"u"}, Usage: "Username", Required: true},
					&cli.StringFlag{Name: "password", Aliases: []string{"p"}, Usage: "Password", Required: true},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := cmd.String("name")
					org := cmd.String("org")
					user := cmd.String("user")
					password := cmd.String("password")

					ctxs, err := Contexts()
					if err != nil && err != keyring.ErrNotFound {
						return err
					}
					if ctxs == nil {
						ctxs = new(contexts)
					}

					ctxs.Add(&contextInfo{
						ID:           uuid.New().String(),
						Name:         name,
						Organization: org,
						Username:     user,
						Password:     password,
						CreatedAt:    time.Now(),
						LastUsedAt:   time.Now(),
					})
					ctxs.SetCurrent(name)

					err = ctxs.Save()
					if err != nil {
						return err
					}

					fmt.Fprintf(cmd.Root().Writer, "Creating and switching context '%s' (%s)\n", name, org)

					return nil
				},
			},
		},
	}
)

var logLevel = "disabled"

func loadCavClient() error {

	contexts, err := Contexts()
	if err != nil || contexts == nil || len(contexts.Items) == 0 {
		err = fmt.Errorf("No context found. Please login first.\n") //nolint:stylecheck
		return err
	}

	cred := contexts.GetCurrent()

	// * Logger
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	if logLevel != "disabled" {
		handler := log.NewWithOptions(os.Stdout, log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.TimeOnly,
		})
		ll := log.InfoLevel

		switch logLevel {
		case "debug":
			ll = log.DebugLevel
		case "info":
			ll = log.InfoLevel
		case "warn":
			ll = log.WarnLevel
		case "error":
			ll = log.ErrorLevel
		default:
			ll = log.InfoLevel
		}

		handler.SetLevel(ll)
		logger = slog.New(handler)
		logger.Info("Logger initialized", "level", ll.String())
	}

	cavClient, err := cav.NewClient(
		cred.Organization,
		cav.WithCloudAvenueCredential(cred.Username, cred.Password),
		cav.WithLogger(logger),
	)
	if err != nil {
		return err
	}

	client = &clientInfo{
		Context: *cred,
		c:       cavClient,
	}
	return nil
}
