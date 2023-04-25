import React from 'react';
import styles from '@/styles/pages/application.module.css';
import Link from 'next/link';

interface ApplicationProps {
	params: {
		id: string;
	};
}

interface Application {
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

async function fetchApplication(id: String): Promise<Application> {
	const res = await fetch(`http://localhost:3000/api/applications?id=${id}`, {
		cache: 'no-cache',
	});
	const [application] = await res.json();
	return application;
}

const ApplicationPage = async ({ params: { id } }: ApplicationProps) => {
	const { name, description, owners, teamName } = await fetchApplication(id);
	return (
		<div className={styles.applicationContainer}>
			<Link href='/apps' className={styles.backLink}>
				<svg
					className={styles.leftArrow}
					xmlns='http://www.w3.org/2000/svg'
					viewBox='0 0 24 24'
				>
					<title>Back</title>
					<path d='M10.05 16.94V12.94H18.97L19 10.93H10.05V6.94L5.05 11.94Z' />
				</svg>
				Back to Apps
			</Link>
			<h2>{teamName}</h2>
			<h1>{name}</h1>
			<p>{description}</p>
			<p>{owners}</p>
		</div>
	);
};

export default ApplicationPage;
