package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tmc/nlm/internal/notebooklm/rpc"
)

// GetAudioArtifactMetadata reads the stable artifact ID without fetching audio
// bytes. Backfill can use this to skip RSS entries before any media download.
func (c *Client) GetAudioArtifactMetadata(projectID string) (*AudioOverviewResult, error) {
	resp, err := c.rpc.Do(rpc.Call{
		ID: rpc.RPCListArtifacts,
		Args: []interface{}{
			[]interface{}{2},
			projectID,
			"NOT artifact.status = \"ARTIFACT_STATUS_SUGGESTED\"",
		},
		NotebookID: projectID,
	})
	if err != nil {
		return nil, fmt.Errorf("query audio artifacts: %w", err)
	}
	return parseAudioArtifactMetadata(resp, projectID)
}

func parseAudioArtifactMetadata(raw []byte, projectID string) (*AudioOverviewResult, error) {
	var response []interface{}
	if err := json.Unmarshal(raw, &response); err != nil {
		return nil, fmt.Errorf("parse artifacts response: %w", err)
	}
	if len(response) == 0 {
		return nil, fmt.Errorf("no audio overview found for this notebook")
	}
	artifacts, ok := response[0].([]interface{})
	if !ok || len(artifacts) == 0 {
		return nil, fmt.Errorf("no audio artifacts found")
	}
	artifact, ok := artifacts[0].([]interface{})
	if !ok || len(artifact) < 2 {
		return nil, fmt.Errorf("invalid audio artifact metadata")
	}
	id, _ := artifact[0].(string)
	title, _ := artifact[1].(string)
	return &AudioOverviewResult{
		ProjectID: projectID,
		AudioID: strings.TrimSpace(id),
		Title: strings.TrimSpace(title),
	}, nil
}
