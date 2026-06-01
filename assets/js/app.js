/**
 * ShadowSelf - Modular Alpine.js Components & Theme Switcher
 */

// Immediately load stored theme to avoid flashing
(function() {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark');
    } else {
        document.documentElement.classList.remove('dark');
    }
})();

document.addEventListener('alpine:init', () => {
    // Theme global store
    Alpine.store('theme', {
        isDark: document.documentElement.classList.contains('dark'),
        
        toggle() {
            this.isDark = !this.isDark;
            if (this.isDark) {
                document.documentElement.classList.add('dark');
                localStorage.setItem('theme', 'dark');
            } else {
                document.documentElement.classList.remove('dark');
                localStorage.setItem('theme', 'light');
            }
        }
    });

    // Component for index.html questionnaire & landing page modal
    Alpine.data('assessmentForm', () => ({
        step: 0,
        nama: '',
        email: '',
        q1: '',
        q2: '',
        q3: '',
        isSubmitting: false,
        isOpen: false, // Modal state for questionnaire overlay
        
        init() {
            if (new URLSearchParams(window.location.search).get('start') === 'true') {
                setTimeout(() => this.openModal(), 150);
            }
        },
        openModal() {
            this.isOpen = true;
            this.step = 0;
            // Prevent scroll on body
            document.body.classList.add('overflow-hidden');
        },
        closeModal() {
            this.isOpen = false;
            document.body.classList.remove('overflow-hidden');
        },
        validateStep0() {
            return this.nama.trim().length >= 2 && 
                   this.email.includes('@') && 
                   this.email.includes('.');
        }
    }));

    // Component for paywall.html diagnostic simulator & countdown
    Alpine.data('paywallLoader', () => ({
        loading: true,
        progress: 0,
        textIndex: 0,
        loadingTexts: [
            'Sinkronisasi metrik kuesioner...',
            'Menghitung deviasi skor terhadap indeks rata-rata populasi...',
            'Menyusun pemetaan klaster perilaku defensif subjek...',
            'Mengonstruksi berkas pelaporan klinis final...'
        ],
        
        init() {
            let interval = setInterval(() => {
                if (this.progress < 100) {
                    this.progress += 2;
                    if (this.progress % 25 === 0 && this.textIndex < 3) {
                        this.textIndex++;
                    }
                } else {
                    clearInterval(interval);
                    this.loading = false;
                }
            }, 50);
        }
    }));
});
