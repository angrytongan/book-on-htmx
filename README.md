Random notes that might make a good online book.

## Rendering pages

-   Each page has an associated `GET` handler.
-   The handler will render the page with whatever block data is necessary.
-   The template used for rendering the page will be either:
    - the name of the block for htmx requests; or
    - the name of the block with "-page" appended for non-htmx requests.
-   The block template only includes what will be shown in the main area
    of the page, and a navigation oob swap template.
-   The block+"-page" template includes everything required to do a full page
    render, including `<html>`, `<head>`, `<body>` tags, etc. The navigation
    oob template is _not_ to be included.

### How this works

-   `app.render()` checks if the request for this render was an htmx request.
    If it wasn't (ie. first page load or hard reload) "-page" is appended to
    the name of the block.
-   `app.render()` also reconstructs the navigation links. These links are
    oob-swapped if the block includes the nav swap
-   Other block data is set.
-   Page is rendered and sent to the client.

### Improvements

-   `app.render()`: only reconstruct navigation if necessary, not every render.

### Example:

-   Handler for the home page:

    ```go
	app.mux.Get("/", app.home)
    ...
    func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	    app.render(w, r, "home", nil, http.StatusOK)
    }
    ```

-   Templates

    ```go
    {{/* chosen on a full page load, or reload */}}
    {{block "home-page" .}}
      {{template "head" .}}
      {{template "nav-head" .}}

      {{template "home-content" .}}

      {{template "nav-foot" .}}
      {{template "foot" .}}
    {{end}}

    {{/* chosen on an htmx request */}}
    {{block "home" .}}
      {{template "home-content" .}}
      {{template "nav-oob" .}}
    {{end}}

    {{/* just the content of <main> */}}
    {{block "home-content" .}}
      <h1 class="text-xl">Home</h1>
    {{end}}
    ```
