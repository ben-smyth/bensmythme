document.querySelectorAll('a.nav-link').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();

        document.querySelector(this.getAttribute('href')).scrollIntoView({
            behavior: 'smooth'
        });
    });
});

window.addEventListener('scroll', function() {
    var button = document.getElementById('back-to-top');
    if (window.scrollY > 1000) {
        button.style.display = 'block';
    } else {
        button.style.display = 'none';
    }
});
