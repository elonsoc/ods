interface Config {
	url: {
		API_URL: String;
		BACKEND_API_URL: String | undefined;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
		BACKEND_API_URL: process.env.PROD_BACKEND_API_URL,
	},
};
const dev: Config = {
	url: {
		API_URL: `http://127.0.0.1:3001`,
		BACKEND_API_URL: process.env.DEV_BACKEND_API_URL,
	},
};

console.log(process.env);

export const config = process.env.NODE_ENV == `development` ? dev : prod;
