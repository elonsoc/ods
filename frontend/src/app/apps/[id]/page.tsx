import React from 'react';
import styles from '@/styles/application.module.css';

async function fetchApplication(id: any) {
	const res = await fetch(`http://localhost:3000/api/applications?id=${id}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
		cache: 'no-store',
	});
	const application = await res.json();
	return application[0];
}

const ApplicationPage = async ({ params: { id } }: any) => {
	const { name, description, owners, teamName } = await fetchApplication(id);
	return (
		<div className={styles.applicationContainer}>
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
