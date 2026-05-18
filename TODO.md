# WOW Mom App Development TODO

## 1. Bugfixes & Stability
- [x] Fix `api/api.go` field names: (Verified correct on disk, LLM filter caused confusion).
- [x] Verify database schema: Check `database/schema.sql` against the `api/api.go` structs.
- [ ] Implement better error logging in `main.go` and `api/api.go` for easier debugging on Pi 4.
- [x] Fix potential "sensitive data" redaction issues: Confirmed as tool-result mask, not file bug.

## 2. PWA & UI Enhancements
- [x] Add `manifest.json`: Define the PWA manifest for installability.
- [x] Add Service Worker: Implement a basic service worker in `static/sw.js` for caching and PWA functionality.
- [x] Update `index.html`:
    - [x] Add `<link rel="manifest" ...>` and meta tags for mobile.
    - [x] Register the Service Worker.
    - [x] Link a common CSS library (Pico.css).
- [x] Make CRUD Interface Reactive: Refactored `static/script.js` with simple fetch/innerHtml reactivity.

## 3. Workflow Implementation
- [ ] Finalize "Mother Registration" flow.
- [ ] Implement "Leader Interview" scheduling logic in the API.

## Execution Order
1. Bugfixes (API stability)
2. PWA infrastructure (Manifest/SW)
3. UI Refactor (Reactivity & CSS)
4. Feature completion
