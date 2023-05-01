import { FormEvent, useState } from 'react';
import styles from './ApplicationSettingsForm.module.css';

function handleSubmit(event: any) {
	event.preventDefault();
}

const ApplicationSettingsForm = ({
	application,
	setSettingsActive,
	handleAppSubmit,
}: any) => {
	const { name, description, owners, teamName } = application;
	const [state, setState] = useState({
		name: name,
		description: description,
		owners: owners,
		teamName: teamName,
	});

	const handleInputChange = (
		event: FormEvent<HTMLInputElement> | FormEvent<HTMLTextAreaElement>
	): void => {
		const { name, value } = event.currentTarget;
		setState((prevInfo) => ({
			...prevInfo,
			[name]: value,
		}));
	};

	const handleFormSubmit = (event: any) => {
		event.preventDefault();
		handleAppSubmit(state);
		setSettingsActive(false);
	};

	return (
		<form onSubmit={handleFormSubmit} className={styles.settingsFormContainer}>
			<header className={styles.formHeader}>
				<h1>Application Settings</h1>
				<div className={styles.formButtons}>
					<button
						type='button'
						className={`${styles.button} ${styles.cancelButton}`}
						onClick={() => setSettingsActive(false)}
					>
						Cancel
					</button>
					<button
						type='submit'
						className={`${styles.button} ${styles.saveButton}`}
					>
						Save
					</button>
				</div>
			</header>
			<div className={styles.inputWrapper}>
				<label htmlFor='name'>Name</label>
				<input
					type='text'
					id='name'
					name='name'
					value={state.name}
					onChange={handleInputChange}
					required={true}
				></input>
			</div>
			<div className={styles.inputWrapper}>
				<label htmlFor='description'>Description</label>
				<textarea
					id='description'
					name='description'
					value={state.description}
					onChange={handleInputChange}
					required={true}
				></textarea>
			</div>
			<div className={styles.inputWrapper}>
				<label htmlFor='owners'>Owners</label>
				<input
					type='text'
					id='owners'
					name='owners'
					value={state.owners}
					onChange={handleInputChange}
					required={true}
				></input>
			</div>
			<div className={styles.inputWrapper}>
				<label htmlFor='teamName'>Team Name</label>
				<input
					type='text'
					id='teamName'
					name='teamName'
					placeholder={`e.g. User's Team`}
					value={state.teamName}
					onChange={handleInputChange}
				></input>
			</div>
		</form>
	);
};

export default ApplicationSettingsForm;
