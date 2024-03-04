'use client';

import { useEffect, useState } from 'react';
import Apps from '@/components/ApplicationPreviewGallery/UserApps';
import AddAppModal from '@/components/AddApplicationModel/AddAppModal';
import styles from '@/styles/pages/applicationGallery.module.css';
import Loader from '@/components/ui/Loader/Loader';
import {
	ApplicationSimple,
	UserAppInformation,
} from '../api/applications/application.d';

export default function App() {
	const [applications, setApplications] = useState<ApplicationSimple[]>([]);
	const [modalActive, setModalActive] = useState<boolean>(false);
	const [loading, setLoading] = useState(true);
	const [hasApplications, setHasApplications] = useState(false);
	const [error, setError] = useState<any>(null);

	async function fetchApplications() {
		try {
			const res = await fetch(`/api/applications`, {
				cache: 'no-store',
			});
			const applications = await res.json();
			setApplications(applications);
			setHasApplications(applications.length > 0);
			setLoading(false);
		} catch (error) {
			setError(error);
		}
	}

	useEffect(() => {
		fetchApplications();
	}, []);

	async function handleAppSubmit(appInfo: UserAppInformation) {
		const result = await fetch(`/api/applications`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			cache: 'no-store',
			body: JSON.stringify(appInfo),
		});
		setModalActive(false);
		fetchApplications();
	}

	if (error) {
		throw new Error(error);
	}

	if (loading) {
		return <Loader />;
	}

	return (
		<>
			{!hasApplications ? (
				<NoAppsPage setModalActive={setModalActive} />
			) : (
				<div className={styles.appContainer}>
					{hasApplications && <Apps applications={applications} />}
					<button
						type='button'
						onClick={() => setModalActive(true)}
						className={`${styles.button} ${styles.topRight}`}
					>
						Add App
					</button>
				</div>
			)}

			{modalActive && (
				<AddAppModal onAdd={handleAppSubmit} onClose={setModalActive} />
			)}
		</>
	);
}

interface NoAppsPageProps {
	setModalActive: (active: boolean) => void;
}

function NoAppsPage({ setModalActive }: NoAppsPageProps) {
	return (
		<div className={styles.noAppContainer}>
			<header className={styles.statusContainer}>
				<h1 className={styles.statusTitle}>No Apps?</h1>{' '}
				<p className={styles.statusDescription}>
					You currently have no registered applications.
				</p>
			</header>
			<button
				type='button'
				onClick={() => setModalActive(true)}
				className={styles.button}
			>
				Register an Application
			</button>
		</div>
	);
}
