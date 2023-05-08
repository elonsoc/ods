interface Config {
	url: {
		API_URL: String;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
	},
};
const dev: Config = {
	url: {
		API_URL: `http://localhost:3000`,
	},
};

export const config = process.env.NODE_ENV === `development` ? dev : prod;
