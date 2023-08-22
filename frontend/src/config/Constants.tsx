interface Config {
	url: {
		API_URL: string;
		BACKEND_API_URL: string | undefined;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
		BACKEND_API_URL: process.env.BACKEND_API_URL,
	},
};
const dev: Config = {
	url: {
		API_URL: `http://127.0.0.1:3001`,
		BACKEND_API_URL: process.env.BACKEND_API_URL,
	},
};

export const config = process.env.NODE_ENV == `development` ? dev : prod;
