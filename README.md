# WEB BATCH
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/cinqfive/webbatch/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/cinqfive/webbatch/tree/main)

## Description
Error handling is a very important aspect of programming. In Golang, your programming style is affected by the fact that in order to indicate that something went wrong you need to return an `error` object. This in turn forces you to, should you handle that error, use an if block that tests if the error is not equal to `nil` and do something in that event. This can yield functions of considerable sizes and even when size is not a problem separation of concern becomes one.

To this end, Web Batch was created. Web batch was created within the scope of web development where the use of streams is frequent and thus error handling code blocks.

## Usage
`go mod get github.com/cinq-five/weborder` will download this module for you and add a reference to it in your `go.mod` file. Using this module is quite simple.

### Via Execute()
Using the function exec from this package will require you to pass to it the same arguments that are passed to your route handler, namely `w http.ResponseWriter, r *http.Request`. The following arguments are a variable list of `Step` arguments. These are of type `func (w http.ResponseWriter, r *http.Request) bool`. They return false when something during their execution went wrong. This will cause all subsequent function executions to be canceled.

#### Example
```golang
func handler(w http.ResponseWriter, r *http.Request) {
   // Suppose checkPermission, validateData, storeData and sendResult are all defined somewhere
   webbatch.Execute(w, r, checkPermission, validateData, storeData, sendResult)
}
```

### Using the WebBatch type
Using this style yields to code blocks like these
#### Example
```golang
func handler(w http.ResponseWriter, r *http.Request) {
   // Suppose checkPermission, validateData, storeData and sendResult are all defined somewhere
   batch := webbatch.Batch
   batch.AddStep(checkPermission)
       .AddStep(validateData)
       .AddStep(storeData)
       .AddStep(sendResult)
       .Execute(w, r)
}
```

This style is particularly adapted to situations where you have many steps to invoke.
