function Clock() {
    const date = new Date();

    // Get hours, minutes and seconds
    const hour = date.getHours();
    let mins = date.getMinutes();
    let secs = date.getSeconds();
    const ampm = hour < 12 ? "AM" : "PM";

    // Check if mins < 10
    mins < 10 ? mins = `0${mins}` : null;
    secs < 10 ? secs = `0${secs}` : null;

    setTimeout(Clock, 500)

    // Show in html
    document.getElementById("clock").innerHTML = `${hour}:${mins}:${secs} ${ampm}`
};