import styles from './UserApp.module.css';

export interface AppInfo {
	id: string;
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

export interface UserAppInformation {
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

interface UserAppInfoProp {
	info: AppInfo;
}

const UserApp = ({
	info: { name, description, owners, teamName = 'Application' },
}: UserAppInfoProp) => {
	return (
		<div className={styles.appContainer}>
			<p className={styles.teamName}>{teamName}</p>
			<h3>{name}</h3>
			<p className={styles.description}>{description}</p>
			<p className={styles.ownerList}>{owners}</p>
		</div>
	);
};

export default UserApp;
