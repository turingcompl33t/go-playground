## Learning Go

Programs written while working through [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) by Jon Bodner.

### Program Contents

- `json-marshaling`: TODO
- `net-client`: Using standard library functionality from `net/http` to make HTTP requests
- `net-server0`: Using standard library functionality from `net/http` to implement the "hello world" HTTP server
- `net-server1`: Using standard library functionality from `net/http` to multiplex requests to different routes
- `net-server2`: Using standard library functionality from `net/http` to inject middleware into an HTTP server

### Notes

**Chapter 12: The Context**

- Go introduces the idea of context in order to address the need to pass metadata through call chains that are handled by "unknown" goroutines
- Other languages solve this issue with the use of thread-local variables; Go cannot do this because it introduces its own concurrency primitive (it does not use OS threads directly)
- It is a Go convention that the context is passed through a program as the first parameter to function calls
- One use for the context is to provide manual cancellation of requests
- Another use for the context is to provide automatic request cancellation with builtin support for cancellable contexts based on timeouts and deadlines
- Idiomatic Go favors the explicit over the implicit