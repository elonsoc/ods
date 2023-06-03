interface Config {
	url: {
		API_URL: String;
		BACKEND_API_URL: String;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
		BACKEND_API_URL: `https://api.ods.elon.edu/`,
	},
};
const dev: Config = {
	url: {
		API_URL: `http://localhost:3001`,
		BACKEND_API_URL: `http://localhost:3000`,
	},
};

export const config = process.env.NODE_ENV === `development` ? dev : prod;
