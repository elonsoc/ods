import styles from './App.module.css';

interface AppInformation {
	title: string;
	description: string;
	owners: string;
}

const App = ({ title, description, owners }: AppInformation) => {
	return (
		<div className={styles.appContainer}>
			<h3>{title}</h3>
			<p>{description}</p>
			<p>{owners}</p>
		</div>
	);
};

export default App;
