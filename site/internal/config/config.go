package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache" // Added for consistency if other services use it
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf // For the Site RPC server itself

	// Database and Cache configurations, assuming they are needed by the RPC service directly
	// or by other components it might interact with.
	// These would typically be populated from a YAML file.
	DB struct {
		DataSource string `json:",optional"` // Optional if not all deployments need direct DB
	} `json:",optional"`
	Cache cache.CacheConf `json:",optional"` // Optional

	// Configuration for clients this service might use.
	// If SiteService needs to call itself or other RPC services, it would be configured here.
	// For the specific task of SiteService needing to call SiteService (e.g. if it were an API gateway calling the RPC)
	// or if a sub-component of SiteService needed to make RPC calls.
	// Given the context, this seems to imply configuration for an *external* client that would call *this* SiteService.
	// However, if SiteService itself needs to act as a client to another RPC (or even itself through an LB),
	// this is where its client config would go.
	// The subtask description "add a new field for the RPC client configuration for SiteService"
	// is slightly ambiguous. It could mean:
	// A) Configuration for *this* service to act as a client *to* another "SiteService" (e.g. a different instance or version).
	// B) Configuration that *other* services would use to connect *to this* "SiteService".
	// C) This service (SiteService) itself is also an API gateway and needs to connect to its own RPC backend.
	//
	// Based on the example provided in the prompt which shows `SiteRpc zrpc.RpcClientConf`
	// and the typical use case where an API gateway (like our future public/admin handlers)
	// would need to call an RPC service (like this SiteService), it's most likely that
	// this config entry is intended for such API gateways or other clients.
	// However, the file path `site/internal/config/config.go` is for the SiteService RPC itself.
	// If SiteService *is* the RPC, then it doesn't need a client config *to itself* in its own primary config.
	//
	// Let's assume the intent is that this SiteService might, for some reason, need to make RPC calls
	// (perhaps to another service, or a decomposed part of itself).
	// Or, this config is a shared base, and different mains (RPC server, API gateways) would use relevant parts.
	//
	// Given the prompt's direct instruction "add a new field for the RPC client configuration for SiteService",
	// I will add `SiteRpc zrpc.RpcClientConf`. This implies that the service defined by this config
	// might also act as a client to a service named "SiteRpc".
	SiteRpc zrpc.RpcClientConf `json:",optional"`
}
