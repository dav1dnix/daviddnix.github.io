const icon = document.querySelector("link[rel=icon");

// All icons
const red = "./assets/icons/rainbow/red.png";
const orange = "./assets/icons/rainbow/orange.png";
const yellow = "./assets/icons/rainbow/yellow.png";
const green = "./assets/icons/rainbow/green.png";
const blue = "./assets/icons/rainbow/blue.png";
const indigo = "./assets/icons/rainbow/indigo.png";
const violet = "./assets/icons/rainbow/violet.png";


// Every 0.1 seconds, change icon in order
let current = 0;
const icons = [red, orange, yellow, green, blue, indigo, violet];
window.setInterval(function() {
    const i = (++current % icons.length);
    const url = icons[i];

    icon.href = url
}, 1000);