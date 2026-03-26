package commands

import (
	"context"
	"fmt"
)

func switchCommand() Definition {
	return Definition{
		Name:        "switch",
		Description: "Switch model",
		SubCommands: []SubCommand{
			{
				Name:        "model",
				Description: "Switch to a different model",
				ArgsUsage:   "to <name>",
				Handler: func(_ context.Context, req Request, rt *Runtime) error {
					if rt == nil || rt.SwitchModel == nil {
						return req.Reply(unavailableMsg)
					}
					// Parse: /switch model to <value>
					value := nthToken(req.Text, 3) // tokens: [/switch, model, to, <value>]
					if nthToken(req.Text, 2) != "to" || value == "" {
						return req.Reply("Usage: /switch model to <name>")
					}
					oldModel, err := rt.SwitchModel(value)
					if err != nil {
						return req.Reply(err.Error())
					}
					return req.Reply(fmt.Sprintf("Switched model from %s to %s", oldModel, value))
				},
			},
			{
				Name:        "agent",
				Description: "Switch to a different agent (separate context/history)",
				ArgsUsage:   "to <id>",
				Handler: func(_ context.Context, req Request, rt *Runtime) error {
					if rt == nil || rt.SwitchAgent == nil {
						return req.Reply(unavailableMsg)
					}
					value := nthToken(req.Text, 3)
					if nthToken(req.Text, 2) != "to" || value == "" {
						if rt.ListAgentIDs != nil {
							ids := rt.ListAgentIDs()
							return req.Reply(fmt.Sprintf("Usage: /switch agent to <id>\nAvailable: %s", joinOrNone(ids)))
						}
						return req.Reply("Usage: /switch agent to <id>")
					}
					oldAgent, err := rt.SwitchAgent(value)
					if err != nil {
						return req.Reply(err.Error())
					}
					return req.Reply(fmt.Sprintf("🔄 Switched from %s → %s (separate history & context)", oldAgent, value))
				},
			},
			{
				Name:        "channel",
				Description: "Moved to /check channel",
				Handler: func(_ context.Context, req Request, _ *Runtime) error {
					return req.Reply("This command has moved. Please use: /check channel <name>")
				},
			},
		},
	}
}
