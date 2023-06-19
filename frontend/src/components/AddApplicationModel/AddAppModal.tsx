import styles from './AddAppModal.module.css';
import { FormEvent, useState } from 'react';
import {
	ApplicationSimple,
	UserAppInformation,
} from '@/app/api/applications/application.d';

interface ModalProps {
	onAdd: (state: UserAppInformation) => void;
	onClose: (isClosed: boolean) => void;
}

const AddAppModal = ({ onAdd, onClose }: ModalProps) => {
	const [state, setState] = useState({
		name: '',
		description: '',
		owners: '',
		teamName: ``,
	});

	const handleInputChange = (event: FormEvent<HTMLInputElement>): void => {
		const { name, value } = event.currentTarget;
		setState((prevInfo) => ({
			...prevInfo,
			[name]: value,
		}));
	};

	const handleSubmit = async (
		event: FormEvent<HTMLFormElement>
	): Promise<void> => {
		event.preventDefault();
		onAdd(state);
	};

	return (
		<div className={styles.fullScreenContainer}>
			<div className={styles.modalWindow}>
				<header className={styles.formHeader}>
					<h1 className={styles.modalTitle}>Application Registration</h1>
					<p className={styles.requirementText}>
						Required input fields are marked with{' '}
						<span className={styles.requiredRed}>*</span>
					</p>
				</header>
				<form onSubmit={(e) => handleSubmit(e)} method='POST'>
					<div className={styles.inputWrapper}>
						<label htmlFor='name'>
							Name <span className={styles.requiredRed}>*</span>
						</label>
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
						<label htmlFor='description'>
							Description <span className={styles.requiredRed}>*</span>
						</label>
						<input
							type='text'
							id='description'
							name='description'
							value={state.description}
							onChange={handleInputChange}
							required={true}
						></input>
					</div>
					<div className={styles.inputWrapper}>
						<label htmlFor='owners'>
							Owners <span className={styles.requiredRed}>*</span>
						</label>
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
					<div className={styles.submissionButtons}>
						<button
							className={styles.closeTextButton}
							type='button'
							onClick={() => onClose(false)}
						>
							Close
						</button>
						<button className={styles.addButton} type='submit'>
							Add
						</button>
					</div>
				</form>
				<button
					className={styles.closeButton}
					type='button'
					onClick={() => onClose(false)}
				>
					<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
						<title>close</title>
						<path d='M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z' />
					</svg>
				</button>
			</div>
		</div>
	);
};

export default AddAppModal;
