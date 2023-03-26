import styles from './AddAppModal.module.css';
import { useState } from 'react';

interface ModalProps {
	onAdd: (a: string, b: string, c: string) => void;
	onClose: (a: boolean) => void;
}

const AddAppModal: React.FunctionComponent<ModalProps> = ({
	onAdd,
	onClose,
}: ModalProps) => {
	const [name, setName] = useState<string>('');
	const [description, setDescription] = useState<string>('');
	const [owners, setOwners] = useState<string>('');

	return (
		<div className={styles.container}>
			<div className={styles.contentWrapper}>
				<form>
					<div className={styles.input}>
						<label htmlFor='name'>Name</label>
						<input
							type='text'
							id='name'
							onChange={(e) => setName(e.target.value)}
						></input>
					</div>
					<div className={styles.input}>
						<label htmlFor='description'>Description</label>
						<input
							type='text'
							id='description'
							onChange={(e) => setDescription(e.target.value)}
						></input>
					</div>
					<div className={styles.input}>
						<label htmlFor='owners'>Owners</label>
						<input
							type='text'
							id='owners'
							onChange={(e) => setOwners(e.target.value)}
						></input>
					</div>
					<button
						className={styles.addButton}
						type='submit'
						onClick={() => onAdd(name, description, owners)}
					>
						Add
					</button>
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
