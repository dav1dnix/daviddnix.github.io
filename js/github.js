const url = "https://api.github.com/users/dps910"
const options = {
	headers: {
		"User-Agent": "dps910"
	},
};

const request = fetch(url, options)

// Fulfil request
request.then((data) => {
	if (data.status == 200 || data.status == 301) {
		return data.json()
	} else {
		throw new Error("Couldn't reach GitHub API")
	}
});
