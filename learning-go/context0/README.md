## `context0`

This is the first program of sufficient complexity that it seems to warrant an explanation here. There are two cases to consider.

The first is the following invocation:

```bash
$ context0 false
```

In this case, the string `false` is supplied to the program at invocation time as the error value. In `callBoth` we create a `WaitGroup` to synchronize execution of two goroutines. We spawn two goroutines, each of which invokes the `callServer` function to handle the actual HTTP client functionality. If an error is returned by the `callServer` function, the cancellation callback is invoked to notify cancellation on the `Context`. The error value string is passed as a query parameter on the request to the fast server. When this string is `false`, no error indicator is returned by the fast server, and both requests execute to completion.

```
fast result: ok
slow result: Slow Response
Complete
```

The second case is the following invocation:

```bash
$ context0 true
```

In this case, the string `true` is supplied to the program at invocation time as the error value. The logic of `callBoth` proceeds as before. Now, when the fast server returns an error indicator, a non-nil error is returned by the `callServer` function. In turn, the cancellation callback for the context is invoked. This invocation allows the slow request, which is still being processed at the server, to be cancelled without waiting for completion.

```
fast result: error
canelling from fast
slow response error: Get "http://127.0.0.1:35559": context canceled
Complete
```

The `http.Client` from `net/http` obviously has some internal logic to monitor the context for cancellation while a request is in progress with `client.Do()`.