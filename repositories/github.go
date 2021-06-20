package repositories

import (
	"api-storage-github/models"
	"api-storage-github/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

func GetCodeInRepositoryByScriptName(scriptName string) string {
	scriptName = "https://raw.githubusercontent.com/" + os.Getenv("GITHUB_USERNAME") + "/" + os.Getenv("GITHUB_REPOSITORY") + "/main/" + scriptName
	request := utils.Request{
		Method: "GET",
		Url:    scriptName,
		Header: utils.Header{},
		Body:   []byte{},
	}
	return utils.NewHttpClient().Get(request)
}

func SendCodeInRepository(codeBody models.CodeBody) models.BodyGithub {
	code := base64.StdEncoding.EncodeToString([]byte(codeBody.Code))
	data := map[string]string{"message": "push code", "content": code}
	values, _ := json.Marshal(data)

	request := utils.Request{
		Method: "PUT",
		Url:    "https://api.github.com/repos/" + os.Getenv("GITHUB_USERNAME") + "/" + os.Getenv("GITHUB_REPOSITORY") + "/contents/" + codeBody.ScriptName,
		Header: utils.Header{
			ContentType:   "application/vnd.github.v3+json",
			Authorization: "Basic " + os.Getenv("GITHUB_TOKEN"),
		},
		Body: values,
	}

	body, err := utils.NewHttpClient().Post(request)
	if err != nil {
		fmt.Println(err)
		return models.BodyGithub{}
	}

	contentReturn := models.BodyGithub{}
	json.Unmarshal([]byte(body), &contentReturn)
	return contentReturn
}
