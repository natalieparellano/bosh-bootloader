package application

import "github.com/cloudfoundry/bosh-bootloader/storage"

type GlobalConfiguration struct {
	EndpointOverride string
	StateDir         string
	Debug            bool
}

type StringSlice []string

func (s StringSlice) ContainsAny(targets ...string) bool {
	for _, target := range targets {
		for _, element := range s {
			if element == target {
				return true
			}
		}
	}
	return false
}

type Configuration struct {
	Global          GlobalConfiguration
	Command         string
	SubcommandFlags StringSlice
	State           storage.State
}
