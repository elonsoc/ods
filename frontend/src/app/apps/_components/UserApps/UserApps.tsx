import styles from './UserApps.module.css';
import { useState } from 'react';
import UserApp from '../UserApp/UserApp';
import AddAppModal from '../UserAppModal/AddAppModal';

interface InformationDetails {
	title: string;
	description: string;
	owners: string;
}

const UserApps = () => {
	const [apps, setApps] = useState<InformationDetails[]>([]);
	const [hasApps, setHasApps] = useState<boolean>(false);
	const [modalActive, setModalActive] = useState<boolean>(false);

	function handleAddApp(name: string, description: string, owners: string) {
		const result: InformationDetails[] = [
			...apps,
			{ title: name, description: description, owners: owners },
		];
		setApps(result);
		setHasApps(true);
		setModalActive(false);
	}

	return (
		<>
			<div className={styles.appContainer}>
				{hasApps &&
					apps.map((app, index) => <UserApp info={app} key={index} />)}
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

export default UserApps;
