import styles from './UserApp.module.css';

export interface UserAppInformation {
	title: string;
	description: string;
	owners: string;
	teamName: string;
}

interface UserAppInfoProp {
	info: UserAppInformation;
}

const UserApp = ({
	info: { title, description, owners, teamName },
}: UserAppInfoProp) => {
	return (
		<div className={styles.appContainer}>
			<p className={styles.teamName}>{teamName}</p>
			<h3>{title}</h3>
			<p className={styles.description}>{description}</p>
			<p className={styles.ownerList}>{owners}</p>
		</div>
	);
};

export default UserApp;
