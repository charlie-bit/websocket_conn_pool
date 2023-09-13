# websocket_pool

Connection pool is a thread safe list of websocket conn

## How to implement connection pool?

> base on sync.Map

## How to ensure connection pool automic?

> based on sync map,meanwhile use runtime.mutex to ensure automic.
> reference:

## How to implement wait-mechanism?

## In user ws connection pool situation, how to kick off user?

## How to implement user login based on several platforms?

## Support Functions

> store user conn
> update user conn
> get specific user conn pool
> get user ws conn pool
> delete user conn by remote url
> delete all user conn
> delete user specific conn
