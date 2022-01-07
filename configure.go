package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
	"github.com/xmapst/NexusImageClean/nexus"
	"html/template"
	"os"
)

const (
	CredentialsFile      = ".credentials"
	CredentialsTemplates = `# Nexus Credentials
host = "{{ .Host }}"
username = "{{ .Username }}"
password = "{{ .Password }}"
repository = "{{ .Repository }}"`
)

func Configure(*cli.Context) error {
	var hostname, repository, username, password string
	fmt.Print("Enter Nexus Host: ")
	_, _ = fmt.Scan(&hostname)
	fmt.Print("Enter Nexus Repository Name: ")
	_, _ = fmt.Scan(&repository)
	fmt.Print("Enter Nexus Username: ")
	_, _ = fmt.Scan(&username)
	fmt.Print("Enter Nexus Password: ")
	_, _ = fmt.Scan(&password)

	data := nexus.ServerInfo{
		Host:       hostname,
		Username:   username,
		Password:   password,
		Repository: repository,
	}

	tmpl, err := template.New(CredentialsFile).Parse(CredentialsTemplates)
	if err != nil {
		return err
	}

	f, err := os.Create(CredentialsFile)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func newConnNexus() (client nexus.Client, err error) {
	r := nexus.ServerInfo{}
	if _, err := os.Stat(CredentialsFile); os.IsNotExist(err) {
		return client, errors.New(fmt.Sprintf("%s file not found\n", CredentialsFile))
	} else if err != nil {
		return client, err
	}
	if _, err := toml.DecodeFile(CredentialsFile, &r); err != nil {
		return client, err
	}
	return nexus.New(r)
}
