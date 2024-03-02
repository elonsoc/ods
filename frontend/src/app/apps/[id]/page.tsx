'use client';

import React, { useEffect, useState } from 'react';
import styles from '@/styles/pages/application.module.css';
import {
	ApplicationExtended,
	UserAppInformation,
} from '@/app/api/applications/application.d';
import ApplicationInformation from '@/components/ApplicationInformation/ApplicationInformation';
import SkeletonLoader from '@/components/SkeletonLoader/SkeletonLoader';
import BackLink from '@/components/BackLink/BackLink';
import { useRouter } from 'next/navigation';

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
		const res = await fetch(`/api/applications/${id}`, {
			cache: 'no-store',
		});
		const application = await res.json();
		setApplication(application);
		setLoading(false);
		return application;
	}

	async function handleAppSubmit(appInfo: UserAppInformation) {
		setLoading(true);
		const result = await fetch(`/api/applications/${id}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(appInfo),
			cache: 'no-store',
		});
		fetchApplication(id);
	}

	async function handleAppDelete(id: string) {
		setLoading(true);
		const result = await fetch(`/api/applications/${id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
			},
			cache: 'no-store',
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
