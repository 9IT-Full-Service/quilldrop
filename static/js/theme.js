(function() {
    var html = document.documentElement;
    var toggle = document.getElementById('theme-toggle');
    var stored = localStorage.getItem('theme');

    if (stored) {
        html.setAttribute('data-theme', stored);
    }

    function updateIcon() {
        if (toggle) {
            toggle.textContent = html.getAttribute('data-theme') === 'dark' ? '\u263D' : '\u2600\uFE0F';
        }
    }

    updateIcon();

    if (toggle) {
        toggle.addEventListener('click', function() {
            var next = html.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
            html.setAttribute('data-theme', next);
            localStorage.setItem('theme', next);
            updateIcon();
        });
    }

    // Hamburger menu
    var hamburger = document.getElementById('hamburger');
    var navLinks = document.getElementById('nav-links');
    if (hamburger && navLinks) {
        hamburger.addEventListener('click', function() {
            hamburger.classList.toggle('active');
            navLinks.classList.toggle('open');
        });
        // Close mobile menu when clicking a link
        navLinks.querySelectorAll('a').forEach(function(link) {
            link.addEventListener('click', function() {
                hamburger.classList.remove('active');
                navLinks.classList.remove('open');
            });
        });
    }

    // Dropdown toggles (for mobile tap support)
    document.querySelectorAll('.nav-dropdown-toggle').forEach(function(btn) {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            e.stopPropagation();
            var dropdown = btn.closest('.nav-dropdown');
            // Close other dropdowns
            document.querySelectorAll('.nav-dropdown.open').forEach(function(other) {
                if (other !== dropdown) other.classList.remove('open');
            });
            dropdown.classList.toggle('open');
        });
    });

    // Close dropdown when clicking outside
    document.addEventListener('click', function(e) {
        if (!e.target.closest('.nav-dropdown')) {
            document.querySelectorAll('.nav-dropdown.open').forEach(function(d) {
                d.classList.remove('open');
            });
        }
    });

    // TOC generator for posts with toc: true
    var toc = document.getElementById('toc');
    if (toc) {
        var headings = document.querySelectorAll('.post-content h1, .post-content h2, .post-content h3');
        if (headings.length > 0) {
            // Determine the highest (smallest number) heading level used
            var minLevel = 6;
            headings.forEach(function(h) {
                var level = parseInt(h.tagName.charAt(1));
                if (level < minLevel) minLevel = level;
            });

            var ul = document.createElement('ul');
            headings.forEach(function(h) {
                var level = parseInt(h.tagName.charAt(1));
                var li = document.createElement('li');
                // Indent relative to the highest heading level found
                li.className = 'toc-level-' + (level - minLevel);
                var a = document.createElement('a');
                a.href = '#' + h.id;
                a.textContent = h.textContent;
                li.appendChild(a);
                ul.appendChild(li);
            });
            toc.appendChild(ul);
        }
    }
})();
