package main

import "github.com/urfave/cli/v2"

func listFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "nexus_version",
			Aliases: []string{"nv"},
			Usage:   "nexus path version",
			Value:   "v2",
		},
	}
}

func tagsFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "List tags by image name",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "nexus_version",
			Aliases: []string{"nv"},
			Usage:   "nexus path version",
			Value:   "v2",
		},
	}
}

func infoFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "show tag info by image name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "tag",
			Aliases:  []string{"t"},
			Usage:    "show tag info by tag name",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "nexus_version",
			Aliases: []string{"nv"},
			Usage:   "nexus path version",
			Value:   "v2",
		},
	}
}

func deleteFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "delete tag info by image name",
		},
		&cli.StringFlag{
			Name:    "tag",
			Aliases: []string{"t"},
			Usage:   "delete tag info by tag name",
		},
		&cli.IntFlag{
			Name:    "keep",
			Aliases: []string{"k"},
			Usage:   "keep tag nb",
			Value:   10,
		},
		&cli.IntFlag{
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "current nb",
			Value:   10,
		},
		&cli.StringFlag{
			Name:    "nexus_version",
			Aliases: []string{"nv"},
			Usage:   "nexus path version",
			Value:   "v2",
		},
	}
}
