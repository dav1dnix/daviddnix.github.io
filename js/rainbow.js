const icon = document.querySelector("link[rel=icon");

// All icons
const red = "./assets/icons/rainbow/red.png";
const orange = "./assets/icons/rainbow/orange.png";
const yellow = "./assets/icons/rainbow/yellow.png";
const green = "./assets/icons/rainbow/green.png";
const blue = "./assets/icons/rainbow/blue.png";
const indigo = "./assets/icons/rainbow/indigo.png";
const violet = "./assets/icons/rainbow/violet.png";


// Every second, change icon in order
let current = 0;

// Add null as first element, forcing the array to start from red instead of orange.
const icons = [null, red, orange, yellow, green, blue, indigo, violet];

window.setInterval(function() {
    // What this does is it adds to the current, then uses % to get the remainder which is 0. It will go along until it gets to the last element index, and then restart. 
    const url = icons[++current % icons.length];

    icon.href = url;
}, 1000);