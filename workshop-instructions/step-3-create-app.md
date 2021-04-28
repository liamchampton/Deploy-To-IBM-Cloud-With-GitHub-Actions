# Step 3 - Create a Go App

This workshop uses a Go application, but you can chose your preferred language as long as there is a buildpack in IBM Cloud available. More information on IBM Cloud Foundry can be found on the [product docs](https://cloud.ibm.com/docs/cloud-foundry-public).

**Before we start to code, we need to initiate `go modules`. To do this run the command `go mod init` from the root of the project. If you get an error, try and run `go mod init <path-to-project-folder>`**

Now, let's create the test case. Open up the file `main_test.go` and copy in the code below.

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	t.Run("returns message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Home(response, request)

		got := response.Body.String()
		want := "This is a test for GitHub Actions"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
```
This code will test the function `Home()`. This function is a HTTP method on the route `/`. The test will pass a mock request into the function and return a response. Once the call has completed, the test code will then test the expected response against the actual response it received from the call.

The TDD unit test cycle should follow RED -> AMBER -> GREEN steps. These steps represent the status of the test code.

RED - The test is failing.

AMBER - The test is passing but the tested code could perform better or be refactored.

GREEN - The test is passing, and the tested code has been refactored.

That being said, try and run the test code with the command `go test -v`

The output should be a FAIL. This represents the RED stage.

Now let's write enough code for the test to pass. Open the `main.go` file and replace the existing `Hello World!` code with the code below.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create the route handler listening on '/'
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Assign the 'msg' variable with a string value
	msg := "This is a test for GitHub Actions"

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(msg)))
}
```

Save the file and then run the tests again using the command `go test -v`.

This time the tests should pass!

To run this locally, use the command `go run main.go` and then navigate to `localhost:8080` in a web browser to see the output from the `/` route.

We have now written enough code for the test to pass, representing the AMBER stage.

You would now look into the code and attempt to refactor. This isn't always possible to do and, in this instance, this code is fairly concise already so there isn't much, if anything, to refactor. This would land us into the GREEN stage.

Once it is all working locally, commit these changes to your forked repository in GitHub.

Recap:
- We have written a test
- Briefly looked at the different stages for TDD
- We have written the code to pass the test

[Step 4 - Manually Deploy Into IBM Cloud Foundry](./step-4-push-to-cf.md)