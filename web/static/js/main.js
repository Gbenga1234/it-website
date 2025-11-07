// Tiny progressive enhancement
document.addEventListener('DOMContentLoaded', () => {
  const links = document.querySelectorAll('a');
  links.forEach(a => a.addEventListener('focus', () => a.classList.add('focus')));
});
