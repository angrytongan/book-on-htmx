# TODO: Unify Mobile and Desktop Navigation

## Objective
Combine the duplicate navigation markup (lines 46-66 and 70-89) into a single navigation component that uses responsive Tailwind classes to handle both mobile and desktop layouts.

## Tasks
- [x] Replace duplicate mobile and desktop navigation with unified responsive structure
- [x] Test the unified navigation on both mobile and desktop layouts  
- [x] Run linter to ensure code quality

## Summary
Successfully unified the mobile and desktop navigation into a single responsive component. The implementation:

### Changes Made
- **Reduced code duplication**: Combined ~40 lines of duplicate navigation markup into ~20 lines
- **Single navigation block**: Replaced separate mobile/desktop `<aside>` sections with one unified structure
- **Responsive classes**: Used Tailwind responsive classes to handle layout differences
- **Maintained functionality**: Preserved all existing HTMX behavior and navigation features

### Mobile Layout (< md breakpoint)
- Horizontal navigation bar fixed at top (`top-0 left-0 right-0 h-16`)
- Icons only, labels hidden (`hidden md:inline`)
- Scrollable horizontal menu when needed (`overflow-x-auto menu-horizontal`)

### Desktop Layout (>= md breakpoint)  
- Vertical sidebar fixed on left (`md:h-screen md:w-32`)
- Icons and labels both visible (`md:inline`)
- Standard vertical menu layout (`md:menu-vertical`)

The unified navigation provides better maintainability while preserving the same user experience across all screen sizes.

## Implementation Steps

### 1. Replace Duplicate Navigation with Unified Structure
- Remove both existing `<aside>` sections for mobile and desktop navigation
- Create a single `<aside>` element that adapts to screen size using Tailwind responsive classes

### 2. Responsive Layout Classes
- **Mobile**: `md:hidden fixed top-0 left-0 right-0 h-16` (horizontal bar at top)
- **Desktop**: `hidden md:block fixed bg-base-200 h-screen w-32` (vertical sidebar)

### 3. Menu Orientation & Styling
- **Mobile menu**: `menu-horizontal gap-2 min-w-max px-2 h-16`
- **Desktop menu**: `menu gap-2 w-32`
- Combine: `menu gap-2 md:menu-horizontal md:min-w-max md:px-2 md:h-16 lg:menu-vertical lg:w-32`

### 4. Label Visibility
- **Mobile**: Hide labels (`hidden`)
- **Desktop**: Show labels (`inline`)
- Combine: `hidden md:inline`

### 5. Link Styling Adjustments
- Unify the anchor tag classes to work for both layouts
- Use responsive classes for layout-specific styling (flex-col, padding, etc.)

### 6. Update Page Layout
- Ensure the main content area adapts properly with the unified navigation
- Maintain existing `pt-16 md:pt-0` and `ml-0 md:ml-32` classes on main content

## Expected Benefits
- **Reduced code duplication**: ~40 lines reduced to ~20 lines
- **Easier maintenance**: Single navigation structure to update
- **Consistent behavior**: Same HTMX functionality across both layouts
- **Cleaner template**: More maintainable responsive design pattern