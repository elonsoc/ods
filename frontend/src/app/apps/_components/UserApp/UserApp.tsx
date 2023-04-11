import styles from './UserApp.module.css';

interface UserAppInformation {
	title: string;
	description: string;
	owners: string;
}

interface UserAppInfoProp {
	info: UserAppInformation;
}

const UserApp = ({ info: { title, description, owners } }: UserAppInfoProp) => {
	return (
		<div className={styles.appContainer}>
			<h3>{title}</h3>
			<p>{description}</p>
			<p>{owners}</p>
		</div>
	);
};

export default UserApp;