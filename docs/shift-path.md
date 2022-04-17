# Router

Routing based on Axel Wagner's [ShiftPath](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html) approach

## Approach

Shift path basically returns next partial path that ends with `/` each call, e.g:
```
/games/<id>/actions/shoot

Calls:
1. games
2. <id>
3. actions
4. shoot
5. ""
```

For each shift path result we can decide to delegate to a `router` or a `methodHandler`.

### POST `/games/<id>`
Simple example:

appRouter
- ShiftPath returns `games`, `/<id>`
- Delegates to `gamesRouter` with path `/<id>`

gamesRouter
- ShiftPath returns `<id>`, `/`
- We need to call ShiftPath again to know if there is a `resource` or not
- ShiftPath returns `""`, `/`
- There is no `resource`, so we delegate handling to a `methodHandler`

### POST `/games/<gameId>/events`
Example with sub-resources:

appRouter
- ShiftPath returns `games`, `/<gameId>/events`
- Delegates to `gamesRouter` with path `/<gameId>/events`

gamesRouter
- ShiftPath returns `<gameId>`, `/events`
- We need to call ShiftPath again to know if there is a `resource` or not
- ShiftPath returns `events`, `/`
- There is an `events` resource, so we delegate it to `eventsRouter` passing `<gameId>` and remaining path `/`

eventsRouter
- ShiftPath returns `""`, `/`
- There is no `eventId`, so we delegate to a `methodHandler` with `<gameId>`
