import Head from 'next/head';

interface Props {
	title?: string;
	keywords?: string;
	description?: string;
}

const Meta: React.FunctionComponent<Props> = ({
	title = 'Elon ODS',
	keywords = "'data access, api provider, Elon University, open source, open data service, ods, elon'",
	description = 'Elon ODS is an Open Data Service at Elon University that provides API keys for students to access data about Elon University. Our API provider service offers data about buildings, courses, and more. Register an application today to get started!',
}: Props) => {
	return (
		<Head>
			<meta name='viewport' content='width=device-width, initial-scale=1.0' />
			<meta name='keywords' content={keywords} />
			<meta name='description' content={description} />
			<meta charSet='utf-8' />
			<link rel='icon' href='/favicon.ico' />
			<title>{title}</title>
		</Head>
	);
};

export default Meta;
