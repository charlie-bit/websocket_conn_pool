# websocket_pool

Connection pool is a thread safe list of websocket conn

## How to implement connection pool?

> base on sync.Map

## How to ensure connection pool atomic?

> based on sync map,meanwhile use runtime.mutex to ensure atomic.

## In user ws connection pool situation, how to kick off user?

## How to implement user login based on several platforms?

## Support Functions

> store user conn
> update user conn
> get user ws conn pool
> delete user conn by remote url
> delete all user conn
> delete user specific conn

## Others

> About mutes，achieve origin mutex of go.
