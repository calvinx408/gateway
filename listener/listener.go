package listener

import (
	"github.com/signalfuse/signalfxproxy/config"
	"github.com/signalfuse/signalfxproxy/core"
)

// A DatapointListener is an object that listens for input datapoints
type DatapointListener interface {
	core.StatKeeper
}

// A ListenerLoader loads a DatapointListener from a configuration definition
type ListenerLoader func(core.DatapointStreamingAPI, *config.ListenFrom) (DatapointListener, error)

// AllListenerLoaders is a map of all loaders from config, for each listener we support
var AllListenerLoaders = map[string]ListenerLoader{
	"signalfx": SignalFxListenerLoader,
	"carbon":   CarbonListenerLoader,
}