const url = "https://api.github.com/users/dps910";
const options = {
	headers: {
		"User-Agent": "dps910"
	},
};

const request = fetch(url, options);

// Fulfil request
request.then((data) => {
	if (data.status == 200 || data.status == 301) {
		data.json().then((data) => {

			const id = document.getElementById("git")

			id.setAttribute("style", "white-space: pre-line;");

			id.textContent = `
				username: ${data.login}
				bio: ${data.bio}
				repos: ${data.public_repos}
				gists: ${data.public_gists}
				location: ${data.location}
			`
			console.log(data)
		});
		// document.getElementById("git").innerHTML = data.json()
	} else {
		throw new Error("Couldn't reach GitHub API");
	}
});
