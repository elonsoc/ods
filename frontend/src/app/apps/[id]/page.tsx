import React from 'react';
import styles from '@/styles/pages/application.module.css';
import Link from 'next/link';
import { UserAppInformation } from '@/app/api/applications/application.d';

interface ApplicationProps {
	params: {
		id: string;
	};
}

async function fetchApplication(id: String): Promise<UserAppInformation> {
	const res = await fetch(`http://localhost:3000/api/applications?id=${id}`, {
		cache: 'no-cache',
	});
	const [application] = await res.json();
	return application;
}

const ApplicationPage = async ({ params: { id } }: ApplicationProps) => {
	const application: UserAppInformation = await fetchApplication(id);
	if (!application) {
		return (
			<div className={styles.applicationContainer}>
				<BackLink />
				<h1>Application not found</h1>
			</div>
		);
	}
	const { name, description, owners, teamName } = application;
	return (
		<div className={styles.applicationContainer}>
			<BackLink />
			<h2>{teamName}</h2>
			<h1>{name}</h1>
			<p>{description}</p>
			<p>{owners}</p>
		</div>
	);
};

const BackLink = () => {
	return (
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
	);
};

export default ApplicationPage;
