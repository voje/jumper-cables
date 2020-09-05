package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/docker/docker/client"
)

// Global cli and context
var cli *client.Client
var ctx context.Context

func init() {
	// Init docker client and context.
	var err error
	cli, err = client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	ctx = context.Background()

}

func containerIsUnhealthy(cntID string) bool {
	info, err := cli.ContainerInspect(ctx, cntID)
	if err != nil {
		panic(err)
	}

	if info.ContainerJSONBase.State.Health.Status == "unhealthy" {
		return true
	}
	return false
}

func watchContainer(cntID string, interval time.Duration) {
	for {
		if containerIsUnhealthy(cntID) {
			log.Infof("Container %s is unhealthy - handling it!", cntID)
		}

		time.Sleep(interval)
	}
}

func main() {
	// Check once
	containerName := "jumper-cables_unhealthy-server_1"
	watchContainer(containerName, time.Second*3)
}

/*
TODOS:
- handle the unhealthy container
- add system time to log messages
*/
