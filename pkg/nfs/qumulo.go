package nfs

import (
	"context"
	"encoding/json"
	"net/http"
	"log"
)

type Cloud interface {
	CreateSnapshot(ctx context.Context, volumeName string) (snapshotInfo *SnapshotInfo, err error)
}

type cloud struct {
	BearerToken string
}

type BearerTokenResponse struct {
    BearerToken string
}

// NewCloud returns a new instance of Qumulo API
// It panics if session is invalid
func NewCloud() (Cloud, error) {
	//TODO read API keys and instantiate new HTTP client here
	url := "https://10.116.100.111:28615/api"

	authUrl := url + "/v1/session/login"

	bearerTokenResponse := new(BearerTokenResponse)
	postJsonAuth(authUrl, bearerTokenResponse)

	return &cloud{
		BearerToken: bearerTokenResponse.BearerToken,
	}, nil
}

type SnapshotInfo struct {
	SnapshotName      string
}

func (c *cloud) CreateSnapshot(ctx context.Context, volumeName string) (snapshotInfo *SnapshotInfo, err error) {
	// TODO Perform HTTP call here

	url := "https://10.116.100.111:28615/api/v1/snapshots"

    var bearer = "Bearer " + c.BearerToken

    // Create a new request using http
    req, err := http.NewRequest("POST", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

	//TODO read response into struct

	return &SnapshotInfo{
		SnapshotName: "get info from HTTP call here",
	}, nil
}

func postJsonAuth(url string, target interface{}) error {
	req, err := http.NewRequest("POST", url, nil)
	client := &http.Client{}
    r, err := client.Do(req)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}