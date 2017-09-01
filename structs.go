package main

type Config struct {
	BindHost string `json:"bind"`
	Items    []Item `json:"items"`
}

type Item struct {
	Id        string `json:id`
	Script    string `json:script`
	ScriptDir string `json:scriptDir`
	Secret    string `json:secret`
}

type Payload struct {
	Secret string `json:"secret"` // "refs/heads/develop"
}
