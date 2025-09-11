# TODO: Add Random Number Button to Toasts Page

## Plan
Add a new button to the toasts page that makes a request to the server and retrieves a random number between 0 and 100, inclusive.

## Tasks

- [x] Create a new handler function for random number generation
- [x] Add route for the random number endpoint in main.go  
- [x] Add new button to toast.tmpl template
- [x] Create template block for random number toast display
- [x] Write tests for the new random number handler

## Implementation Details

1. **Handler Function**: Add `toastRandomNumber` handler in `cmd/web/toast.go` that generates random number 0-100 using `math/rand`

2. **Routing**: Register handler at `/toast/random-number` in HTMX group with delay middleware

3. **Template**: Add button after existing "Get server time" button with HTMX attributes for random number request

4. **Toast Display**: Create `toast-random-number` template block displaying random number with auto-dismiss after 5 seconds

5. **Testing**: Create `cmd/web/toast_test.go` to test random number generation, response format, and HTTP status codes

The implementation follows the existing server time pattern using HTMX for requests and hyperscript for toast behavior.

## Summary

Successfully implemented a random number button for the toasts page. All tasks completed:

### Changes Made
- **Handler Function**: Added `toastRandomNumber` in `cmd/web/toast.go:23-31` that generates random numbers 0-100 using `rand.Intn(101)`
- **Routing**: Registered `/toast/random-number` endpoint in `cmd/web/main.go:56` within the HTMX group with delay middleware
- **Template Button**: Added "Get random number" button in `templates/toast.tmpl:12-18` with `btn-secondary` styling and HTMX attributes
- **Toast Display**: Created `toast-random-number` template block in `templates/toast.tmpl:37-49` with green `alert-success` styling and auto-dismiss behavior
- **Tests**: Created comprehensive test suite in `cmd/web/toast_test.go` with range validation and multiple request testing

### Features Implemented
- Random number generation between 0-100 inclusive using Go's `math/rand` package
- HTMX-powered button that makes async requests without page refresh
- Toast notifications with 5-second auto-dismiss using hyperscript
- Manual dismiss capability with X button
- Proper HTTP status code handling and template rendering
- Full test coverage with edge case validation

The implementation maintains consistency with existing toast patterns while providing reliable random number functionality with proper testing.