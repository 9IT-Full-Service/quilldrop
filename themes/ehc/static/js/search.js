(function () {
    'use strict';

    let searchIndex = null;
    let debounceTimer = null;

    const toggle = document.getElementById('search-toggle');
    const container = document.getElementById('search-container');
    const box = document.getElementById('search-box');
    const input = document.getElementById('search-input');
    const results = document.getElementById('search-results');

    if (!toggle || !input || !results) return;

    // Toggle search box
    toggle.addEventListener('click', function (e) {
        e.stopPropagation();
        const isOpen = container.classList.toggle('active');
        if (isOpen) {
            input.focus();
            loadIndex();
        } else {
            closeSearch();
        }
    });

    // Load search index (lazy, only once)
    function loadIndex() {
        if (searchIndex !== null) return;
        fetch('/search-index.json')
            .then(function (res) { return res.json(); })
            .then(function (data) { searchIndex = data; })
            .catch(function (err) {
                console.error('Search index load failed:', err);
            });
    }

    // Search on input
    input.addEventListener('input', function () {
        clearTimeout(debounceTimer);
        debounceTimer = setTimeout(doSearch, 200);
    });

    function doSearch() {
        var query = input.value.trim().toLowerCase();
        if (!query || !searchIndex) {
            results.innerHTML = '';
            results.classList.remove('visible');
            return;
        }

        var terms = query.split(/\s+/);
        var matches = searchIndex.filter(function (entry) {
            var haystack = [
                entry.title,
                entry.preview,
                (entry.tags || []).join(' '),
                (entry.categories || []).join(' ')
            ].join(' ').toLowerCase();

            return terms.every(function (term) {
                return haystack.indexOf(term) !== -1;
            });
        });

        renderResults(matches.slice(0, 8));
    }

    function renderResults(items) {
        if (items.length === 0) {
            results.innerHTML = '<div class="search-no-results">Keine Ergebnisse gefunden</div>';
            results.classList.add('visible');
            return;
        }

        var html = '';
        items.forEach(function (item) {
            var tags = (item.tags || []).slice(0, 3).map(function (t) {
                return '<span class="search-result-tag">' + escapeHtml(t) + '</span>';
            }).join('');

            html += '<a href="' + escapeHtml(item.url) + '" class="search-result-item">' +
                '<div class="search-result-title">' + highlightMatch(escapeHtml(item.title), input.value.trim()) + '</div>' +
                '<div class="search-result-meta">' +
                    '<span class="search-result-date">' + escapeHtml(item.date) + '</span>' +
                    (tags ? '<span class="search-result-tags">' + tags + '</span>' : '') +
                '</div>' +
            '</a>';
        });

        results.innerHTML = html;
        results.classList.add('visible');
    }

    function highlightMatch(text, query) {
        if (!query) return text;
        var terms = query.toLowerCase().split(/\s+/);
        terms.forEach(function (term) {
            if (!term) return;
            var regex = new RegExp('(' + term.replace(/[.*+?^${}()|[\]\\]/g, '\\$&') + ')', 'gi');
            text = text.replace(regex, '<mark>$1</mark>');
        });
        return text;
    }

    function escapeHtml(str) {
        if (!str) return '';
        return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
    }

    function closeSearch() {
        container.classList.remove('active');
        input.value = '';
        results.innerHTML = '';
        results.classList.remove('visible');
    }

    // Keyboard shortcuts
    input.addEventListener('keydown', function (e) {
        if (e.key === 'Escape') {
            closeSearch();
        }
    });

    // Close on click outside
    document.addEventListener('click', function (e) {
        if (!container.contains(e.target)) {
            closeSearch();
        }
    });

    // Prevent search box clicks from closing
    box.addEventListener('click', function (e) {
        e.stopPropagation();
    });

    // Global shortcut: Ctrl+K or Cmd+K to open search
    document.addEventListener('keydown', function (e) {
        if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
            e.preventDefault();
            container.classList.add('active');
            input.focus();
            loadIndex();
        }
    });
})();
