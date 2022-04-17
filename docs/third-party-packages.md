# About third party packages

# Table of Contents
1. [oapi-codegen](#open-api-codegen)
2. [testify](#testify)

## Open Api Codegen
GitHub repository: [oapi-codegen](https://github.com/deepmap/oapi-codegen)

I started with the builtin `http` package with the [ShiftPath](./shift-path.md) router approach. It worked but as I started writing more features I faced three problems:

1. There were a lot of boilerplate code to create a new endpoint
2. All request validation were made manually for every endpoint
3. There wasn't any endpoint documentation with the expected body and possible responses

I knew I could solve all those problems by writing a OpenAPI Specification and generating bindings and validations based on that. (I saw this approach in NodeJS, so why not in Golang?)

I found two main approaches, one was writing the OpenAPI Specification and generating the code or writing annotations in code and generating the OpenAPI Specification. I chose the former because it seemed to require less code to be written.

Open api and router bindings generator alternatives and thoughts:
* https://github.com/deepmap/oapi-codegen (better doc)
* https://github.com/getkin/kin-openapi
* https://openapi.tools/ (reference)

## Testify
GitHub repository: [testify](https://github.com/stretchr/testify)

I started using the built-in package `testing`, it worked without any problems but I start noticing that most of assertions had the pattern:

```go
if want == got {
  t.Errorf("Expected %s to be equals %s", got, want)
}
```

I could write some test helpers to make those comparisons for me I knew that I'd end up writing the same code as a public available test package.

One helpful resource was the blog [Exploring the landscape of Go testing frameworks](https://bmuschko.com/blog/go-testing-frameworks/) which pointed out some golang testing package alternatives.

I decided to use `testify` because of the following reasons:
1. It solved my repetition problem.
2. It still follows the same test style of the builtin `testing` package.
3. It was the most popular from the alternatives (16k stars on Github) and had active development which means new features, updates and bug correction.

Testing alternatives and thoughts:
* https://github.com/stretchr/testify (same approach as `testing`)
* https://github.com/onsi/ginkgo (BDD)