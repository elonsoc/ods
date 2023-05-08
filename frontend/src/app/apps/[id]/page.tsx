'use client';

import React, { useEffect, useState } from 'react';
import styles from '@/styles/pages/application.module.css';
import { UserAppInformation } from '@/app/api/applications/application.d';
import { config } from '@/config/Constants';
import ApplicationInformation from './_components/ApplicationInformation/ApplicationInformation';
import SkeletonLoader from './_components/SkeletonLoader/SkeletonLoader';
import BackLink from './_components/BackLink/BackLink';
const URL = config.url.API_URL;

interface ApplicationProps {
	params: {
		id: string;
	};
}

const ApplicationPage = ({ params: { id } }: ApplicationProps) => {
	const [application, setApplication] = useState<UserAppInformation>({
		name: '',
		description: '',
		owners: '',
		teamName: '',
	});
	const [loading, setLoading] = useState(true);

	async function fetchApplication(id: String): Promise<UserAppInformation> {
		const res = await fetch(`${URL}/api/applications/${id}`, {
			cache: 'no-cache',
		});
		const [application] = await res.json();
		setApplication(application);
		setLoading(false);
		return application;
	}

	async function handleAppSubmit(appInfo: UserAppInformation) {
		setLoading(true);
		const result = await fetch(`${URL}/api/applications/${id}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(appInfo),
		});
		fetchApplication(id);
	}

	useEffect(() => {
		fetchApplication(id);
	}, []);

	if (loading) {
		return (
			<div className={styles.applicationContainer}>
				<div className={styles.backLinkWrapper}>
					<BackLink />
				</div>
				<SkeletonLoader />
			</div>
		);
	}

	if (!application) {
		return (
			<div className={styles.applicationContainer}>
				<div className={styles.backLinkWrapper}>
					<BackLink />
				</div>
				<h1>Application not found</h1>
			</div>
		);
	}
	return (
		<div className={styles.applicationContainer}>
			<ApplicationInformation
				application={application}
				handleAppSubmit={handleAppSubmit}
			/>
		</div>
	);
};

export default ApplicationPage;
