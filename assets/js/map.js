function createMap(id) {
  const map = L.map(id);

  map.setView([-25.6101, 134.3548], 4);
  L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
  }).addTo(map);

  return map;
}

function addMarker(map, data, oldMarker) {
	const { lat, lng } = data;
	const marker = L.marker([lat, lng]);
	
	if (oldMarker) {
		map.removeLayer(oldMarker)
	}

	marker.addTo(map);
	map.setView([lat, lng], 17);

	return marker;
}
