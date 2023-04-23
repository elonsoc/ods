import React from 'react';
import styles from '@/styles/pages/application.module.css';
import Link from 'next/link';

async function fetchApplication(id: any) {
	const res = await fetch(`http://localhost:3000/api/applications?id=${id}`);
	const application = await res.json();
	return application[0];
}

const ApplicationPage = async ({ params: { id } }: any) => {
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

// slash login, redirect user to Elon -> login blah blah blah -> redirected to some website
// push them to dashboard

// if have token/otherwise unauthorized

export default ApplicationPage;
