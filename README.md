# Certmagic TLS cluster support for Redis

This plugin is refactor based on [Gamalan's](https://github.com/gamalan/caddy-tlsredi).
This plugin utilize [go-redis/redis](https://github.com/go-redis/redis) for its client access and [redislock](https://github.com/bsm/redislock)
for it's locking mechanism. See [distlock](https://redis.io/topics/distlock) for the lock algorithm.

This plugin currently work with versions of Caddy that use https://github.com/caddyserver/certmagic
and its new storage interface (> 0.11.1)

## TODO

- Add Redis Cluster or Sentinel support (probably need to update the distlock implementation first)





