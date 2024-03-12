export interface UserAppInformation {
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

export interface ApplicationSimple {
	id: string;
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

export interface ApplicationExtended {
	id: string;
	name: string;
	description: string;
	owners: string;
	teamName: string;
	apiKey: string;
	isValid: boolean;
}

export interface RequestOptions {
	method: string;
	duplex: string;
	headers: {
		'Content-Type': string;
		credentials: string;
	};
	body: ReadableStream<Uint8Array> | null;
	cache: string;
}
