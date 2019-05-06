# Response

Go `API-Response package` provides easy and simple solution for your API response and makes responses more consistent.

# Usage

```go
func HttpHandler(w http.ResponseWriter, r *http.Request) {
    inputData := []string{"First user", "Second user"}
    // Responds with correspoding data
    response.CreateResponse(w, response.ResponseOK, inputData, "Custom message")
}
/*
Output JSON:
{
    "status":"OK",
    "message":"Custom message"
    "data": ["First user", "Second user"]
}
*/
```

# Installation

`go get github.com/wutchzone/api-response`
