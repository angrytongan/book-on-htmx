# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Architecture Overview

This is a Go web application built with:
- **chi router** for HTTP routing
- **HTMX** for hypermedia-driven interactions
- **Hyperscript** (v0.9.14) for client-side scripting alongside HTMX
- **Domain-driven design** with packages in `internal/` (dog, search, theme, nav)
- **Repository pattern** for database access with PostgreSQL/SQLite support
- **Go HTML templates** with `.tmpl` extension
- **TailwindCSS + DaisyUI** for styling
- **Air** for live reloading during development

## Development Commands

Use these Makefile targets:
- `make web` - Run development server with Air live reload
- `make install` - Install npm dependencies 
- `make lint` - Run golangci-lint on cmd/web and internal packages
- `make lint-fix` - Run linter with automatic fixes
- `make tailwind` - Watch and rebuild TailwindCSS styles

## Key Configuration

- **Module name**: `bonh`
- **Server port**: 8880
- **Environment**: Uses `.env` file for configuration
- **Build output**: Air builds to `./tmp/main`
- **Database**: Supports both SQLite (`database.db`) and PostgreSQL
- **Templates**: Located in `templates/` with `.tmpl` extension
- **Static assets**: 
  - CSS: `assets/css/` (TailwindCSS input/output)
  - JavaScript: `assets/js/` (includes hyperscript@0.9.14.min.js)

## Project Structure

- `cmd/web/` - HTTP handlers, main application entry point, and middleware
- `internal/` - Domain packages with repository pattern:
  - `dog/` - Dog-related functionality
  - `search/` - Search functionality  
  - `theme/` - Theme management
  - `nav/` - Navigation functionality
- `templates/` - Go HTML templates (layout.tmpl, pages.tmpl, etc.)
- `assets/` - Static files (CSS, JavaScript)
- `sql/` - Database schemas and migrations
- `.air.toml` - Air configuration for live reloading

## Development Notes

- The application follows a domain-driven structure with clear separation between web handlers and business logic
- Each domain package typically includes a repository interface and implementation (often both pgx and SQLite versions)
- Templates use Go's html/template package with custom functions and data structures
- HTMX is used for dynamic interactions, with hyperscript providing additional client-side scripting capabilities
- TailwindCSS classes are used throughout templates, with DaisyUI providing component styles
- assets/css/style.css is generated automatically on changes to templates. You do not modify this file. When updating styles only update tailwindcss class names in templates/*.tmpl.
