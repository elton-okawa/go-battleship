## Feature
By priority

### High
* Better http responses
* Implement `presenter` for cli
* Document (with images) current architecture

### Medium
* Document api
* Improve tests
* Create player account
* Games by player
* Two player battleship game

### Low
* Persist player statistic over games
* Random placement bigger ships - now we only place single square
* Keep track if an entire ship was sinked (more than one square)

## Fix
* Random is always generating the same map

## Think about
* ShiftPath router has a lot of boiler plate
* Presenters and controllers are being passed down as pointers

## Take a look
Open api and router bindings generator
* https://github.com/deepmap/oapi-codegen (better doc)
* https://github.com/getkin/kin-openapi
* https://openapi.tools/ (reference)
Google Wire dependency injection package