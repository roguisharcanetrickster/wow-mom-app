const CACHE_NAME = 'wowmom-v1';
const ASSETS = [
  '/',
  '/static/style.css',
  '/static/script.js',
  'https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css'
];

self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => cache.addAll(ASSETS))
  );
});

self.addEventListener('fetch', event => {
  event.respondWith(
    caches.match(event.request).then(response => response || fetch(event.request))
  );
});