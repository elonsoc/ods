import styles from '@/styles/home.module.css';
import Link from 'next/link';
import { useState } from 'react';

export default function Home() {
	const [validUser, setValidUser] = useState<boolean>(false);
	return (
		<div className={styles.container}>
			<h1 className={styles.introTitle}>
				Welcome to Open Date Service at Elon University
			</h1>
			<Link
				href={validUser ? '/apps' : '/denied'}
				className={styles.loginButton}
			>
				Login with Elon
			</Link>
			<div className={styles.validUserInput}>
				<input
					type='checkbox'
					id='validUser'
					onChange={() => setValidUser(!validUser)}
				></input>
				<label htmlFor='validUser'>Valid User</label>
			</div>
		</div>
	);
}
