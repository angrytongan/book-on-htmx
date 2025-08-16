function saveScrollTop(e) {
  const url = new URL(window.location);
  const params = url.searchParams;
  params.set('main-scroll', e.target.scrollTop)

  window.history.replaceState({}, '', url.toString());
}

function restoreScrollTop() {
  const url = new URL(window.location);
  const params = url.searchParams;
  const scrollTop = params.get('main-scroll');
  if (scrollTop) {
    document.getElementById('main').scrollTop = scrollTop;
  }
}
