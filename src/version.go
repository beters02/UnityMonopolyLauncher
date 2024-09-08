package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func VerifyVersion() {

}

func GetMostRecentLauncherVersion() (string, error) {
	didDownload, err := downloadLinkAsFile("temp_version.json", "https://github.com/beters02/UnityMonopolyLauncher/blob/main/src/version.go")
	if !didDownload {
		return "", err
	}

	table, err := jsonFileToTable("temp_version.json")
	if err != nil {
		return "", err
	}

	os.Remove("temp_version.json")
	a := table["Version"].(string)
	return a, nil
}

func GetMostRecentGameVersion() (string, error) {
	didDownload, err := downloadLinkAsFile("temp_version.json", "https://github.com/beters02/UnityMonopolyLauncher/blob/main/src/version.go")
	if !didDownload {
		return "", err
	}

	table, err := jsonFileToTable("temp_version.json")
	if err != nil {
		return "", err
	}

	os.Remove("temp_version.json")
	a := table["Version"].(string)
	return a, nil
}

func GetLocalLauncherVersion() (string, error) {
	table, err := jsonFileToTable("version.json")
	if err != nil {
		return "", err
	}
	a := table["Launcher_Version"].(string)
	return a, nil
}

func GetLocalGameVersion() (string, error) {
	table, err := jsonFileToTable("version.json")
	if err != nil {
		return "", err
	}
	a := table["Game_Version"].(string)
	return a, nil
}

func jsonFileToTable(where string) (map[string]interface{}, error) {
	table := map[string]interface{}{}

	jsonFile, err := os.Open(where)
	if err != nil {
		return table, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return table, err
	}

	err = json.Unmarshal(byteValue, &table)
	if err != nil {
		return table, err
	}

	return table, nil
}

func downloadLinkAsFile(where string, link string) (bool, error) {
	out, err := os.Create(where)
	if err != nil {
		return false, err
	}
	defer out.Close()

	resp, err := http.Get(link)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return false, err
	}
	return true, nil
}
