'use client';

import Apps from '@/app/apps/_components/UserApps/UserApps';
import { useEffect, useState } from 'react';
import AddAppModal from './_components/UserAppModal/AddAppModal';
import styles from '@/styles/pages/applicationGallery.module.css';
import Loader from '@/ui/Loader/Loader';
import {
	ApplicationSimple,
	UserAppInformation,
} from '../api/applications/application.d';
import { config } from '@/config/Constants';

const URL = config.url.API_URL;
const BACKEND_URL = config.url.BACKEND_API_URL;

export default function App() {
	const [applications, setApplications] = useState<ApplicationSimple[]>([]);
	const [modalActive, setModalActive] = useState<boolean>(false);
	const [loading, setLoading] = useState(true);
	const [hasApplications, setHasApplications] = useState(false);

	async function fetchApplications() {
		const res = await fetch(`${BACKEND_URL}/applications`, {
			cache: 'no-cache',
		});
		const applications = await res.json();
		setApplications(applications);
		setHasApplications(applications.length > 0);
		setLoading(false);
	}

	useEffect(() => {
		fetchApplications();
	}, []);

	async function handleAppSubmit(appInfo: UserAppInformation) {
		const result = await fetch(`${BACKEND_URL}/applications`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(appInfo),
		});
		setModalActive(false);
		fetchApplications();
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
