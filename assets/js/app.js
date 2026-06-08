/**
 * ShadowSelf - App JavaScript
 * 9Router-inspired premium UI interactions
 */

document.addEventListener('DOMContentLoaded', function() {
    // --- Smooth scroll for anchor links ---
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            const href = this.getAttribute('href');
            if (href === '#') return;
            const target = document.querySelector(href);
            if (target) {
                e.preventDefault();
                target.scrollIntoView({ behavior: 'smooth', block: 'start' });
            }
        });
    });

    // --- Animated progress bars on scroll ---
    const progressBars = document.querySelectorAll('.progress-fill');
    if (progressBars.length > 0) {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    const bar = entry.target;
                    const width = bar.getAttribute('data-width') || bar.style.width;
                    bar.style.width = '0%';
                    setTimeout(() => {
                        bar.style.width = width;
                    }, 100);
                    observer.unobserve(bar);
                }
            });
        }, { threshold: 0.3 });

        progressBars.forEach(bar => {
            const width = bar.style.width;
            bar.style.width = '0%';
            bar.setAttribute('data-width', width);
            observer.observe(bar);
        });
    }

    // --- Card hover tilt effect (subtle) ---
    const tiltCards = document.querySelectorAll('.card-hover');
    tiltCards.forEach(card => {
        card.addEventListener('mousemove', function(e) {
            const rect = this.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;
            const centerX = rect.width / 2;
            const centerY = rect.height / 2;
            const rotateX = (y - centerY) / 20;
            const rotateY = (centerX - x) / 20;
            this.style.transform = `perspective(1000px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) translateY(-2px)`;
        });
        card.addEventListener('mouseleave', function() {
            this.style.transform = '';
        });
    });

    // --- Copy to clipboard for code blocks ---
    document.querySelectorAll('code[data-copy]').forEach(code => {
        code.addEventListener('click', function() {
            const text = this.textContent;
            navigator.clipboard.writeText(text).then(() => {
                const original = this.innerHTML;
                this.innerHTML = '<span class="material-symbols-outlined text-sm">check</span> Copied!';
                setTimeout(() => {
                    this.innerHTML = original;
                }, 2000);
            });
        });
    });

    // --- Auto-dismiss alerts ---
    document.querySelectorAll('.alert-auto-dismiss').forEach(alert => {
        setTimeout(() => {
            alert.style.transition = 'opacity 0.3s ease';
            alert.style.opacity = '0';
            setTimeout(() => alert.remove(), 300);
        }, 5000);
    });

    // --- Print button enhancement ---
    const printBtn = document.querySelector('[onclick="window.print()"]');
    if (printBtn) {
        printBtn.addEventListener('click', function(e) {
            e.preventDefault();
            window.print();
        });
    }

    console.log('%c ShadowSelf %c Premium UI Active ',
        'background:#ef4444;color:white;padding:4px 8px;border-radius:4px 0 0 4px;font-weight:bold;',
        'background:#f3f4f6;color:#111827;padding:4px 8px;border-radius:0 4px 4px 0;font-size:11px;'
    );
});
