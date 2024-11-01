package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func update_version(version ...uint8) (major uint8, minor uint8, patch uint8) {
	patch = version[2]
	minor = version[1]
	major = version[0]
	if patch < 10 {
		patch += 1
	} else if minor < 10 {
		patch = 0
		minor += 1
	} else {
		patch = 0
		minor = 0
		major += 1
	}
	return major, minor, patch
}

func main() {
	file := "./ts_types/package.json"
	file_bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var package_json map[string]interface{}
	if err := json.Unmarshal(file_bytes, &package_json); err != nil {
		panic(err)
	}
	version := package_json["version"].(string)
	split_version := strings.Split(version, ".")
	parsed_version := make([]uint8, len(split_version))
	for i, t := range split_version {
		v, err := strconv.ParseUint(t, 10, 8)
		if err != nil {
			panic(err)
		}
		parsed_version[i] = uint8(v)
	}
	parsed_version[0], parsed_version[1], parsed_version[2] = update_version(parsed_version...)
	for i, t := range parsed_version {
		split_version[i] = fmt.Sprint(t)
	}
	final_version := strings.Join(split_version, ".")
	package_json["version"] = final_version
	data, err := json.MarshalIndent(package_json, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile(file, data, 0644)
}
