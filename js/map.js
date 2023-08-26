let map;
// initMap is now async
async function initMap() {
  // marker list
  const markers = {
    sis: { center: { lat: 38.8712041, lng: -77.1143355 } },
    antonios: { center: { lat: 40.0005234, lng: -75.2943201 } },
  };

  // Request libraries when needed, not in the script tag.
  const { Map } = await google.maps.importLibrary("maps");

  // Short namespaces can be used.
  map = new Map(document.getElementById("map"), {
    center: { lat: 39.2903848, lng: -76.6121893 },
    zoom: 9,
    mapTypeId: "terrain",
  });

  for (const m in markers) {
    new google.maps.Circle({
      strokeColor: "#FF0000",
      strokeOpacity: 0.8,
      strokeWeight: 2,
      fillColor: "#FF0000",
      fillOpacity: 0.35,
      map: map,
      center: markers[m].center,
      radius: 1e5, // 100km
    });

    new google.maps.Marker({
      position: markers[m].center,
      map: map,
      label: m,
    });
  }
}

initMap();
