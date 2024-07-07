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

function toggleText(button) {
    const content = button.previousElementSibling;
    if (content.style.display === "none" || content.style.display === "") {
      content.style.display = "block";
        button.textContent = "-";
    } else {
      content.style.display = "none";
        button.textContent = "+";
    }
}
