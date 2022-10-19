package mycontext

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/digitalocean/go-libvirt"
	"github.com/eskpil/salmon/cmd/agent/utils"
	"github.com/eskpil/salmon/pkg/definitions"
	"github.com/eskpil/salmon/pkg/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Context struct {
	Virt   *libvirt.Libvirt
	Client definitions.M2MClient
}

func NewContext() (*Context, error) {
	virtConn, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Errorf("Could not dial the libvirt socket: %v\n", err)
		return nil, err
	}

	l := libvirt.New(virtConn)
	if err := l.Connect(); err != nil {
		log.Errorf("Could not connect with libvirt: %v\n", err)
		return nil, err
	}

	conn, err := grpc.Dial("192.168.0.74:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Could not connect with the salmon api: %v\n", err)
		return nil, err
	}

	c := definitions.NewM2MClient(conn)

	return &Context{
		Virt:   l,
		Client: c,
	}, nil
}

func (c *Context) PerformRoutine() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	machineId, err := utils.GetMachineId()

	if err != nil {
		log.Errorf("Failed to get the machineId: %v\n", err)
		return nil
	}

	r, err := c.Client.Heartbeat(ctx, &definitions.HeartbeatRequest{MachineId: machineId})

	if err != nil {
		log.Errorf("Failed to perform a heartbeat: %v\n", err)
		return nil
	}

	wg := &sync.WaitGroup{}

	for _, task := range r.GetTasks() {
		wg.Add(1)
		defer wg.Done()

		log.Infof("Performing task: \"%s\"\n", task.Name)

		switch {
		case task.Name == "collectHostname":
			{
				hostname, err := os.Hostname()
				if err != nil {
					log.Errorf("Could not get system hostname: %v\n", err)
					return err
				}

				type response struct {
					Hostname string `json:"hostname"`
				}

				res := response{Hostname: hostname}
				data, err := json.Marshal(res)

				if err != nil {
					log.Errorf("Could not marshal the json response: %v\n", err)
					return err
				}

				c.Client.FinishTask(ctx, &definitions.FinishTaskRequest{Id: task.Id, Data: data})
			}
		case task.Name == "collectMachines":
			{
				type outgoing struct {
					Id       string   `json:"id"`
					Name     string   `json:"name"`
					Groups   []string `json:"groups"`
					Hostname string   `json:"hostname"`
				}

				domains, err := c.Virt.Domains()
				if err != nil {
					log.Errorf("Could not list the libvirt domains: %v\n", err)
					return err
				}

				res := []outgoing{}

				for _, d := range domains {
					id, err := uuid.FromBytes(d.UUID[:])

					if err != nil {
						log.Errorf("Could not construct a machineId: %v\n", err)
						return err
					}

					// 2 means we use the guest agent living on the domain.
					hostname, err := c.Virt.DomainGetHostname(d, 2)

					if err != nil {
						log.Errorf("Could not get the machines hostname: %v\n", err)
						hostname = "<unknown>"
					}

					data := outgoing{Id: id.String(), Name: d.Name, Groups: []string{}, Hostname: hostname}
					res = append(res, data)
				}

				data, err := json.Marshal(res)

				if err != nil {
					log.Errorf("Could not marshal the json response: %v\n", err)
					return err
				}

				c.Client.FinishTask(ctx, &definitions.FinishTaskRequest{Id: task.Id, Data: data})
			}
		case task.Name == "collectMachineInterfaces":
			{
				type outgoing struct {
					MachineId  string             `json:"machine_id"`
					Interfaces []models.Interface `json:"interfaces"`
				}

				var res []outgoing

				domains, err := c.Virt.Domains()
				if err != nil {
					log.Errorf("Could not list the libvirt domains: %v\n", err)
					return err
				}

				for _, d := range domains {
					machineId, err := uuid.FromBytes(d.UUID[:])

					if err != nil {
						log.Errorf("Could not construct a machineId: %v\n", err)
						return err
					}

					interfaces, err := c.Virt.DomainInterfaceAddresses(d, 1, 0)
					if err != nil {
						log.Errorf("Could not get interfaces for domain: %d\n", d.Name)
					}

					parsedInterfaces := []models.Interface{}

					for _, i := range interfaces {
						ipAddrs := []models.IpAddr{}

						for _, a := range i.Addrs {
							ipAddrs = append(ipAddrs, models.IpAddr{
								Type:   a.Type,
								Addr:   a.Addr,
								Prefix: a.Prefix,
							})
						}

						c := models.Interface{
							Id:      uuid.New().String(),
							Name:    i.Name,
							Mac:     i.Hwaddr[0],
							IpAddrs: ipAddrs,
						}

						parsedInterfaces = append(parsedInterfaces, c)
					}

					d := outgoing{
						MachineId:  machineId.String(),
						Interfaces: parsedInterfaces,
					}

					res = append(res, d)
				}

				data, err := json.Marshal(res)

				if err != nil {
					log.Errorf("Could not marshal the json response: %v\n", err)
					return err
				}

				c.Client.FinishTask(ctx, &definitions.FinishTaskRequest{Id: task.Id, Data: data})

			}
		}
	}

	wg.Wait()
	return nil
}
