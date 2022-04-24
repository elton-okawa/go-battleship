# Architecture Thoughts

## 24-04-2022 Alternatives summary

After all those thoughts from yesterday, it seems that the proposed interaction between `view` and `presenter` is that the `presenter` would trigger `view` updates by using an observer pattern.

In my API the `router` fulfills the `view` role by receiving and responding to user request, so the `router` needs to be notified that the `presenter` data changed.

The problem is that the **current router handler** needs to be notified because it's the one who has the `*http.ResponseWriter` to respond the user request, but the request flow always pass through the `controller` (`router` -> `controller` -> `use case` -> `presenter`), so every bind to the current request will go through the `controller`.

For now, I'll keep the way it's now [mutating presenter values](#22-04-2022-router-accessing-presenter-struct-values) and [controller knowing about the presenter interface](#22-04-2022-controller-knowing-about-presenter-interface-expected-by-use-case) but I'm considering the [Straight forward approach](#22-04-2022-straight-forward-approach)

## 23-04-2022 Studying alternatives

Use case pull and push approaches
https://softwareengineering.stackexchange.com/a/420360
* Pull approach
  * Receive `Input` and returns `Output`
  * Receive some data, perform logic and return the output
```go
func UseCase(input Input) Output {}

// controller
res := UseCase(input)
out := Presenter(res) // same or not function
```

* Push approach
  * You want to invoke the use case and them it initiate the view  update
```go
func UseCase(input Input, presenter Presenter) {
  // stuff
  presenter(stuff)
}

// controller
UseCase(input, Presenter)
```

Discussion about use case returning data or calling presenter
https://softwareengineering.stackexchange.com/questions/357052/clean-architecture-use-case-containing-the-presenter-or-returning-data
* interactor calling the presenter
  * controller does not know about the response model
* interactor returning data
  * does not fainthfully follows the Clean Architecture
```go
repository := Repository{};
useCase := UseCase{ repository };
data := useCase.GetData();
presenter := Presenter{};
presenter.present(data);
```

[Robert C. Martin talking about Clean Architecture](https://www.youtube.com/watch?v=Nsjsiz2A9mg)


## 22-04-2022 Straight forward approach

An approach that might solve both previous problems would be changing a bit the control flow and having some data lift:
* `router` receive user request
* `router` calls the `controller`
  * `controller` calls the `use case`
    * `use case` performs and returns a DTO
  * `controller` returns the same DTO
* `router` pass the received data to the `presenter`
  * `presenter` maps DTO to the actual response
* `router` answer user request

This would change the control flow:
* before: `router` -> `controller` -> `use case` -> `presenter` -> (`controller`) -> `router`
* after: `router` -> `controller` -> `use case` -> (`controller`) -> `router` -> `presenter` -> `router`

## 22-04-2022 `controller` knowing about `presenter` interface expected by `use case`
This behavior does not break the dependency rule but it's a bit weird, we do this mainly because of the previous problem.
At the `router` layer we need the `presenter` data, we achieve this by instantiating the `presenter` at `router` layer and pass it down through the `controller` to the `use case` where it's actually needed.

One option might the `use case` struct already having a `presenter` like we do with `database` and we pass a callback in place. I think that it works but it's not an elegant solution because intermediate layers still know things that they don't need use

```go
// router
func HandleCreateAccount(ctx echo.Context) {
  cb := func (code int, body interface{}) {
    ctx.JSON(code, body)
  }

  controller(cb, ctx)
}

// controller func
useCase(cb, login, password)

// use case func
presenter(cb, dto)
```

## 22-04-2022 `router` accessing `presenter` struct values
It's a bit strange because I'm expecting that the `presenter` struct argument will be mutated by some underlying call. 

Basically I didn't find a better way to solve the flow problem. From what I understand, in the proposed `Clean Architecture`, the flow of control goes like `controller -> use case -> presenter` and `controller` should not know about `presenter` despite living in the same layer.

So the first alternative about `controller` getting `use case` response and them calling the `presenter` with it, would not fit in the architecture

```go
// controller file
out := useCase()
return presenter(out)
```

I considered `use case` returning the presenter result but it'd mean declaring the format that it's expecting the `presenter` to return. It's not good because the `use case` are enforcing something that it's not its responsability

```go
type Presenter interface {
  CreateAccountResponse(AccountDTO) (AccountOutput)
}

func UseCase() AccountOutput {
  ...
  return presenter()
}
```

Another option would be the `presenter` actually answering the user request removing the need of the `router` fetching data from it. I think that this approach also do not fit the architecture because the `interface adapter` layer are meant to map data to a convenient format and not actually performing these kind of actions.

```go
// presenter
func (p *Presenter) Response(dto UseCaseDTO) {
  code, body := mapUseCaseDTO(dto)
  p.ctx.JSON(code, body)
}
```
