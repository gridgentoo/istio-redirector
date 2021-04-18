package istio

import (
	"istio-redirector/domain"
)

type Config struct {
	Istio Istio
}

type Istio struct {
	DestinationRule                               string
	VirtualServiceNamespace                       string
	VirtualServiceDefaultDestinationHost          string
	VirtualServiceDefaultMatchingRegexDestination string
	VirtualServiceHosts                           []string
	VirtualServiceGateways                        []string
}

type Redirections struct {
	Name                            string
	Namespace                       string
	DestinationRuleName             string
	DefaultDestinationHost          string
	DefaultMatchingRegexDestination string
	Hosts                           []string
	Gateways                        []string
	Rules                           []domain.Rule
}