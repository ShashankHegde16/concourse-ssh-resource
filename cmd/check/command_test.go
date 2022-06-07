package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"stash.tools.deloitteinnovation.us/projects/CRE/repos/concourse-ssh-resource/internal/models"
)

func TestCheckCommand(t *testing.T) {
	var response []models.Version

	in := bytes.NewBufferString(`{"source":{},"version":{}}`)
	out := bytes.NewBuffer([]byte{})
	err := checkCommand(in, out)
	if !assert.NoError(t, err) {
		return
	}

	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 0, len(response))
}

func TestCheckCommandWithVersion(t *testing.T) {
	var response checkResponse

	now := time.Now().Format(time.RFC3339)

	request := checkRequest{Version: models.Version{Timestamp: now}}
	requestJSON, err := json.Marshal(&request)
	if !assert.NoError(t, err) {
		return
	}

	in := bytes.NewBuffer(requestJSON)
	out := bytes.NewBuffer([]byte{})
	err = checkCommand(in, out)
	if !assert.NoError(t, err) {
		return
	}

	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	if assert.Equal(t, 1, len(response)) {
		assert.Equal(t, now, response[0].Timestamp)
	}
}

func TestCheckCommandWithMalformedJSON(t *testing.T) {
	in := bytes.NewBufferString(`{`)
	out := bytes.NewBuffer([]byte{})

	err := checkCommand(in, out)
	assert.Contains(t, err.Error(), "unable to parse JSON from standard input")
}
