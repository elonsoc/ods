import styles from './UserApps.module.css';
import { useState } from 'react';
import UserApp from '../UserApp/UserApp';
import AddAppModal from '../UserAppModal/AddAppModal';
import { Raleway } from 'next/font/google';
import Link from 'next/link';

const raleway = Raleway({ subsets: ['latin'] });

interface InformationDetails {
	title: string;
	description: string;
	owners: string;
	teamName: string;
}

const UserApps = () => {
	const [apps, setApps] = useState<InformationDetails[]>([]);
	const [hasApps, setHasApps] = useState<boolean>(false);
	const [modalActive, setModalActive] = useState<boolean>(false);

	function handleAddApp(
		name: string,
		description: string,
		owners: string,
		teamName: string
	) {
		const result: InformationDetails[] = [
			...apps,
			{
				title: name,
				description: description,
				owners: owners,
				teamName: teamName,
			},
		];
		setApps(result);
		setHasApps(true);
		setModalActive(false);
	}

	return (
		<>
			{' '}
			{!hasApps ? (
				<div className={styles.noAppContainer}>
					<header className={styles.statusContainer}>
						<h1 className={`${raleway.className} ${styles.statusTitle}`}>
							No Apps?
						</h1>{' '}
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
			) : (
				<div className={styles.appContainer}>
					{hasApps &&
						apps.map((app, index) => (
							<Link href={`apps/${index}`} key={index}>
								<UserApp info={app} key={index} />
							</Link>
						))}

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
				<AddAppModal onAdd={handleAddApp} onClose={setModalActive} />
			)}
		</>
	);
};

export default UserApps;
