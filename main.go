package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

const (
	serviceName        = "Test - Service"
	serviceDescription = "Simple service example"
)

type program struct{}

func (p program) Start(s service.Service) error {
	log.Println(s.String() + " started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	log.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {
	for {
		log.Println("Service is running")
		time.Sleep(1 * time.Second)
	}
}

func main() {

	logFile := "C:\\services\\log\\teste.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetLevel(log.TraceLevel)

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		log.Println("Cannot create the service: " + err.Error())
	}
	fmt.Println(s.Platform())

	status, errS := s.Status()
	log.Info("status=", status)
	log.Info("err=", errS)

	var install bool
	flag.BoolVar(&install, "instalar", false, "Instalar o serviço?")
	flag.Parse()
	log.Info("instalar=", install)

	if install {
		err = s.Install()
		if err != nil {
			log.Println("Cannot install the service: " + err.Error())
		}
	} else {
		err = s.Run()
		if err != nil {
			log.Println("Cannot start the service: " + err.Error())
		}
	}
}
