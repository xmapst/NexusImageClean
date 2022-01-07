package main

import "github.com/urfave/cli/v2"

func Command() cli.Commands {
	return cli.Commands{
		{
			Name:   "configure",
			Usage:  "Configure Nexus Credentials",
			Action: Configure,
		},
		{
			Name:  "image",
			Usage: "Mange Docker Images",
			Subcommands: cli.Commands{
				{
					Name:    "list",
					Aliases: []string{"ls", "l"},
					Flags:   listFlag(),
					Action:  listImagesAction,
				},
				{
					Name:    "tags",
					Aliases: []string{"tag", "t"},
					Usage:   "Display all image tags",
					Flags:   tagsFlag(),
					Action:  listTagsByImage,
				},
				{
					Name:   "info",
					Usage:  "Show image details",
					Flags:  infoFlag(),
					Action: showImageInfo,
				},
				{
					Name:    "delete",
					Aliases: []string{"del"},
					Usage:   "Delete an image",
					Flags:   deleteFlag(),
					Action:  deleteAction,
				},
			},
		},
	}
}
