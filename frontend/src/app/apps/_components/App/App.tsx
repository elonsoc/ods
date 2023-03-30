import styles from './App.module.css';

interface InformationDetails {
	title: string;
	description: string;
	owners: string;
}

interface AppInformation {
	appInformation: InformationDetails;
}

const App = ({
	appInformation: { title, description, owners },
}: AppInformation) => {
	return (
		<div className={styles.appContainer}>
			<h3>{title}</h3>
			<p>{description}</p>
			<p>{owners}</p>
		</div>
	);
};

export default App;
