var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}
function refresher() {
	var container = document.getElementById("controls");
	container.src=container.src;
}


function sendTemp(el){
	var temp = document.getElementById("temp").value;
	var request = "v1/"+el.id+"/settemp";

	fetch(request,{
		method: "POST",
		headers:{
			"Content-type": "application/json; charset=UTF-8",
			"temp": temp
		}
	});
	console.log("Nueva temperatura de "+el.id+" es: "+temp);
	refresher();
}