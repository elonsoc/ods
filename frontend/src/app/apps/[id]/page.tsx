'use client';

import React, { useEffect, useState } from 'react';
import styles from '@/styles/pages/application.module.css';
import {
	ApplicationExtended,
	UserAppInformation,
} from '@/app/api/applications/application.d';
import { config } from '@/config/Constants';
import ApplicationInformation from './_components/ApplicationInformation/ApplicationInformation';
import SkeletonLoader from './_components/SkeletonLoader/SkeletonLoader';
import BackLink from './_components/BackLink/BackLink';
import { redirect, useRouter } from 'next/navigation';
const URL = config.url.API_URL;
const BACKEND_URL = config.url.BACKEND_API_URL;

interface ApplicationProps {
	params: {
		id: string;
	};
}

const ApplicationPage = ({ params: { id } }: ApplicationProps) => {
	const router = useRouter();
	const [application, setApplication] = useState<ApplicationExtended>({
		id: '',
		name: '',
		description: '',
		owners: '',
		teamName: '',
		apiKey: '',
		isValid: false,
	});
	const [loading, setLoading] = useState(true);

	async function fetchApplication(id: String): Promise<UserAppInformation> {
		const res = await fetch(`${BACKEND_URL}/applications/${id}`, {
			cache: 'no-cache',
		});
		const application = await res.json();
		setApplication(application);
		setLoading(false);
		return application;
	}

	async function handleAppSubmit(appInfo: UserAppInformation) {
		setLoading(true);
		const result = await fetch(`${BACKEND_URL}/applications/${id}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(appInfo),
		});
		fetchApplication(id);
	}

	async function handleAppDelete(id: string) {
		setLoading(true);
		const result = await fetch(`${BACKEND_URL}/applications/${id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
			},
		});
		router.replace(`/apps`);
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
				handleAppDelete={handleAppDelete}
			/>
		</div>
	);
};

export default ApplicationPage;
