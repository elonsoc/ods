import styles from './Apps.module.css';
import { useState } from 'react';
import App from './App/App';
import AddAppModal from './AppModal/AddAppModal';

interface AppInformation {
	name: string;
	description: string;
	owners: string;
}

const Apps = () => {
	const [apps, setApps] = useState<AppInformation[]>([]);
	const [hasApps, setHasApps] = useState<boolean>(false);
	const [modalActive, setModalActive] = useState<boolean>(false);

	function handleAddApp(name: string, description: string, owners: string) {
		const result: AppInformation[] = [
			...apps,
			{ name: name, description: description, owners: owners },
		];
		setApps(result);
		setHasApps(true);
		setModalActive(false);
	}

	return (
		<>
			<div className={styles.appContainer}>
				{hasApps &&
					apps.map((app, index) => (
						<App
							title={app.name}
							description={app.description}
							owners={app.owners}
							key={index}
						/>
					))}
				<button
					type='button'
					onClick={() => setModalActive(true)}
					className={hasApps ? styles.topRight : styles.centered}
				>
					{hasApps ? 'Add App' : 'Create an Application'}
				</button>
			</div>
			{modalActive && (
				<AddAppModal onAdd={handleAddApp} onClose={setModalActive} />
			)}
		</>
	);
};

export default Apps;
