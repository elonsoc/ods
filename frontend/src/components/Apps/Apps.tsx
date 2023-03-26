import styles from './Apps.module.css';
import { useState } from 'react';
import App from './App/App';
import AddAppModal from './AppModal/AddAppModal';

const Apps: React.FunctionComponent = () => {
	const [apps, setApps] = useState<string[][]>([]);
	const [hasApps, setHasApps] = useState(false);
	const [modalActive, setModalActive] = useState(false);

	function handleAddApp(name: string, description: string, owners: string) {
		const result: string[][] = [...apps, [name, description, owners]];
		setApps(result);
		setHasApps(true);
		setModalActive(false);
		console.log('here');
	}
	console.log(modalActive);
	return (
		<>
			<div className={styles.appContainer}>
				{hasApps &&
					apps.map((app, index) => <App information={app} key={index} />)}
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
