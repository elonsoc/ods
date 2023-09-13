interface Config {
	url: {
		API_URL: string;
		BACKEND_API_URL: string | undefined;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
		BACKEND_API_URL: process.env.NEXT_PUBLIC_BACKEND_API_URL,
	},
};
const dev: Config = {
	url: {
		API_URL: `http://localhost:3001`,
		BACKEND_API_URL: 'http://localhost:3000',
	},
};

export const config = process.env.NODE_ENV == `development` ? dev : prod;
