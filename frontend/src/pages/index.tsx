import Image from 'next/image';
import { Inter } from 'next/font/google';
import styles from '../styles/page.module.css';
import Head from 'next/head';

const inter = Inter({ subsets: ['latin'] });

export default function Home() {
	return (
		<>
			<Head>
				<title>Elon ODS</title>
				<meta
					name='keywords'
					content='data access, api provider, Elon University, open source, open data service, ods, elon'
				></meta>
				<meta
					name='description'
					content='Elon ODS is an Open Data Service at Elon University that provides API keys for students to access data about Elon University. Our API provider service offers data about buildings, courses, and more. Register an application today to get started!'
				></meta>
			</Head>

			<div>
				<h1>Welcome to Open Date Service at Elon University</h1>
			</div>
		</>
	);
}
