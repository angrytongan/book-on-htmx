# Plan: Vertical Button Layout at Half Page Width

## Tasks

- [x] Wrap buttons in a container div with half-width and vertical layout classes
- [x] Add vertical spacing between buttons  
- [x] Center the container on the page
- [x] Make buttons full-width within container
- [x] Test the layout and HTMX functionality

## Implementation Details

Modifying `templates/toast.tmpl` to:
1. Add container div around the four buttons (lines 4-38)
2. Apply TailwindCSS classes: `w-1/2 flex flex-col gap-4 mx-auto`
3. Add `w-full` class to each button for full-width within container
4. Maintain existing HTMX functionality and DaisyUI button styling

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

## Summary of Changes

Successfully modified `templates/toast.tmpl` to display four buttons in a vertical layout at half page width:

### Changes Made
- **Container Layout**: Wrapped all four buttons in a container div with classes `w-1/2 flex flex-col gap-4 mx-auto`
  - `w-1/2`: Sets container width to 50% of page width  
  - `flex flex-col`: Creates vertical flexbox layout
  - `gap-4`: Adds consistent spacing between buttons
  - `mx-auto`: Centers the container horizontally
- **Button Styling**: Added `w-full` class to each button to make them stretch full width within the container
- **Preserved Functionality**: Maintained all existing HTMX attributes and DaisyUI button styling

### Layout Achieved
- Four buttons now display vertically stacked
- Container is centered and takes up half the page width
- Consistent spacing between buttons using TailwindCSS gap utility
- Buttons stretch to full width of the container for better visual consistency
- All HTMX toast functionality preserved