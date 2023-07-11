package server

import (
	"github.com/google/wire"
)

var ProviderServer = wire.NewSet(NewGRPCServer, NewHTTPServer)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(ProviderServer, ProviderHttp)
