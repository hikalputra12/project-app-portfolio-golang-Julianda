// 1. Mobile Navigation
const navSlide = () => {
    const burger = document.querySelector('.burger');
    const nav = document.querySelector('.nav-links');
    const navLinks = document.querySelectorAll('.nav-links li');

    if(burger) {
        burger.addEventListener('click', () => {
            nav.classList.toggle('nav-active');
            burger.classList.toggle('toggle');
        });
        
        // Tutup menu saat link diklik
        navLinks.forEach(link => {
            link.addEventListener('click', () => {
                nav.classList.remove('nav-active');
                burger.classList.remove('toggle');
            });
        });
    }
}

// 2. Scroll Reveal Animation
const scrollReveal = () => {
    const reveals = document.querySelectorAll('.reveal');

    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
            if (entry.isIntersecting) {
                entry.target.classList.add('active');
            }
        });
    }, { threshold: 0.1 });

    reveals.forEach((element) => {
        observer.observe(element);
    });
}

// 3. Typing Effect
const typeWriter = (text, elementId, speed) => {
    let i = 0;
    const element = document.getElementById(elementId);
    if (!element) return;
    
    element.innerHTML = "";
    const type = () => {
        if (i < text.length) {
            element.innerHTML += text.charAt(i);
            i++;
            setTimeout(type, speed);
        }
    };
    type();
}

// Init
document.addEventListener('DOMContentLoaded', () => {
    navSlide();
    scrollReveal();
    // Efek ketik teks (Hanya akan jalan jika elemennya ada)
    typeWriter("I build things for the web.", "typing-text", 100);
});