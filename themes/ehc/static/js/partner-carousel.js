/* ============================================================
   Partner-Carousel — Quilldrop Widget
   ------------------------------------------------------------
   Erwartet im DOM:
     <div id="partner-carousel"></div>
     <div id="partner-carousel-dots"></div>      (optional)

   Daten: /static/data/partners.json
   Konfig per data-Attribut auf #partner-carousel:
     data-visible="3"     Anzahl gleichzeitig sichtbarer Logos
     data-interval="3500" Auto-Rotation in ms
   ============================================================ */
(function() {
  'use strict';

  const DATA_URL = '/static/data/partners.json';

  function esc(s) {
    return String(s).replace(/[&<>"]/g, c =>
      ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' }[c])
    );
  }

  function build(host, dotsBox, partners, opts) {
    const VISIBLE  = opts.visible;
    const INTERVAL = opts.interval;

    host.classList.add('carousel');
    host.innerHTML = '';

    const track = document.createElement('div');
    track.className = 'carousel-track';
    host.appendChild(track);

    // Slides + Klone für Endlos-Loop
    const all = partners.concat(partners.slice(0, VISIBLE));
    all.forEach(p => {
      const slide = document.createElement('div');
      slide.className = 'carousel-slide';
      const card = `<div class="partner-card" title="${esc(p.name)}">
        <img src="${esc(p.src)}" alt="${esc(p.name)}" loading="lazy" />
      </div>`;
      slide.innerHTML = p.url
        ? `<a href="${esc(p.url)}" target="_blank" rel="noopener" class="partner-link" aria-label="${esc(p.name)}">${card}</a>`
        : card;
      track.appendChild(slide);
    });

    // Dots
    if (dotsBox) {
      dotsBox.innerHTML = '';
      partners.forEach((_, i) => {
        const b = document.createElement('button');
        b.className = 'carousel-dot' + (i === 0 ? ' active' : '');
        b.setAttribute('aria-label', `Partner ${i + 1}`);
        b.addEventListener('click', () => goTo(i));
        dotsBox.appendChild(b);
      });
    }

    let index = 0;
    function update() {
      track.style.transition = 'transform 0.7s cubic-bezier(0.4, 0, 0.2, 1)';
      track.style.transform  = `translateX(-${(100 / VISIBLE) * index}%)`;
      if (dotsBox) {
        const active = index % partners.length;
        [...dotsBox.children].forEach((d, i) => d.classList.toggle('active', i === active));
      }
    }
    function goTo(i) { index = i; update(); }
    function next() {
      index++;
      update();
      if (index >= partners.length) {
        setTimeout(() => {
          track.style.transition = 'none';
          index = 0;
          track.style.transform = 'translateX(0)';
          if (dotsBox) [...dotsBox.children].forEach((d, i) => d.classList.toggle('active', i === 0));
        }, 720);
      }
    }

    update();
    setInterval(next, INTERVAL);
  }

  function init() {
    const host = document.getElementById('partner-carousel');
    if (!host) return;
    const dotsBox = document.getElementById('partner-carousel-dots');

    const opts = {
      visible:  parseInt(host.dataset.visible,  10) || 3,
      interval: parseInt(host.dataset.interval, 10) || 3500
    };

    fetch(DATA_URL)
      .then(r => r.json())
      .then(partners => {
        if (!Array.isArray(partners) || partners.length === 0) return;
        build(host, dotsBox, partners, opts);
      })
      .catch(err => console.error('partner-carousel:', err));
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
