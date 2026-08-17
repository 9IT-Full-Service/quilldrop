/* ============================================================
   Partner-Grid — Quilldrop Widget
   ------------------------------------------------------------
   Erwartet im DOM:
     <div id="partner-grid"></div>

   Daten: /static/data/partners.json
   Konfig per data-Attribut auf #partner-grid:
     data-sort="name|none"   alphabetisch sortieren (Default: name)
   ============================================================ */
(function() {
  'use strict';

  const DATA_URL = '/static/data/partners.json';

  function esc(s) {
    return String(s).replace(/[&<>"]/g, c =>
      ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' }[c])
    );
  }

  function build(host, partners) {
    host.classList.add('partner-grid');
    host.innerHTML = '';

    partners.forEach(p => {
      const tag = p.url ? 'a' : 'div';
      const el  = document.createElement(tag);
      el.className = 'partner-tile';
      if (p.url) {
        el.href   = p.url;
        el.target = '_blank';
        el.rel    = 'noopener';
      }
      el.innerHTML = `
        <div class="logo">
          <img src="${esc(p.src)}" alt="${esc(p.name)}" loading="lazy" />
        </div>
        <div class="info">
          <div class="name">${esc(p.name)}</div>
          ${p.url
            ? `<div class="visit">Website besuchen &rarr;</div>`
            : `<div class="no-url">Website folgt</div>`}
        </div>`;
      host.appendChild(el);
    });
  }

  function init() {
    const host = document.getElementById('partner-grid');
    if (!host) return;

    const sortMode = host.dataset.sort || 'name';

    fetch(DATA_URL)
      .then(r => r.json())
      .then(partners => {
        if (!Array.isArray(partners) || partners.length === 0) return;
        if (sortMode === 'name') {
          partners.sort((a, b) => a.name.localeCompare(b.name, 'de'));
        }
        build(host, partners);
      })
      .catch(err => console.error('partner-grid:', err));
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
