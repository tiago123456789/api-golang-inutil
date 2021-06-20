package models

type CodeBody struct {
	Code       string `json:"code"`
	Language   string `json:"language"`
	ScriptName string `json:"scriptName"`
}

type ContentBodyGithub struct {
	Sha         string `json:"sha"`
	DownloadUrl string `json:"download_url"`
}

type BodyGithub struct {
	Content ContentBodyGithub `json:"content"`
}
