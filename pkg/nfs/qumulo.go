package nfs

import (
	"context"
)

type Cloud interface {
	CreateSnapshot(ctx context.Context, volumeName string) (snapshotInfo *SnapshotInfo, err error)
}

type cloud struct {
}

// NewCloud returns a new instance of Qumulo API
// It panics if session is invalid
func NewCloud() (Cloud, error) {
	//TODO read API keys and instantiate new HTTP client here
	return &cloud{
		//client: httpClient
	}, nil
}

type SnapshotInfo struct {
	SnapshotName      string
}

func (c *cloud) CreateSnapshot(ctx context.Context, volumeName string) (snapshotInfo *SnapshotInfo, err error) {
	// TODO Perform HTTP call here
	return &SnapshotInfo{
		SnapshotName: "get info from HTTP call here",
	}, nil
}