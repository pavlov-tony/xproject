package main

// TODO: steps to do
// 0. project name
// 1. list of zones
// 2. list of instances in the zones
// 3. instance config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type Instance struct {
	Name        string
	Zone        string // TODO: iota?
	MachineType string // TODO: iota?
	ProjectId   string
}

type InstanceConfig struct {
	Description                  string
	GuestCpus                    int64
	MemoryMb                     int64
	ImageSpaceGb                 int64
	MaximumPersistentDisks       int64
	MaximumPersistentDisksSizeGb int64
	IsSharedCpu                  bool
}

func main() {
	ctx := context.Background()

	cli, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal("1\n\n", err)
	}

	computeService, err := compute.New(cli)
	if err != nil {
		log.Fatal("2\n\n", err)
	}

	projectId := os.Getenv("APP_PROJECT_ID")
	avalibleProjZones := getAvalibleZonesForProject(projectId, ctx, computeService)

	var projInstances []Instance
	for _, zone := range avalibleProjZones {
		projInstances = append(projInstances,
			getInstancesListForProjectAndZone(projectId, zone, ctx, computeService)...)
	}

	fmt.Println(projInstances)
	fmt.Println(getInstanceConfig(projInstances[0], ctx, computeService))

}

func getInstancesListForProjectAndZone(projectId string, zone string,
	ctx context.Context, service *compute.Service) (instances []Instance) {

	if resp, err := service.Instances.List(projectId, zone).Context(ctx).Do(); err != nil {
		log.Fatal("Instance List\n\n", err)
	} else {
		if len(resp.Items) > 0 {
			for _, i := range resp.Items {
				slash := strings.LastIndex(i.MachineType, "/")
				machineType := i.MachineType[slash+1:]
				instances = append(instances, Instance{i.Name, zone, machineType, projectId})
			}
		}
	}

	return instances
}

func getAvalibleZonesForProject(projectId string,
	ctx context.Context, service *compute.Service) (avalibleZones []string) {

	resp, err := service.Zones.List(projectId).Context(ctx).Do()
	if err != nil {
		log.Fatal("Avalible Zones\n\n", err)
	}

	for _, z := range resp.Items {
		avalibleZones = append(avalibleZones, z.Name)
	}

	return avalibleZones
}

func getInstanceConfig(instance Instance,
	ctx context.Context, service *compute.Service) (cfg InstanceConfig) {

	resp, err := service.MachineTypes.Get(
		instance.ProjectId, instance.Zone, instance.MachineType).Context(ctx).Do()
	if err != nil {
		log.Fatal("InstanceConfig\n\n", err)
	}

	cfg = InstanceConfig{resp.Description, resp.GuestCpus, resp.MemoryMb,
		resp.ImageSpaceGb, resp.MaximumPersistentDisks,
		resp.MaximumPersistentDisksSizeGb, resp.IsSharedCpu}

	return cfg
}
