import styles from '@/styles/home.module.css';
import Link from 'next/link';

export default function Denied() {
	return (
		<div className={styles.container}>
			<p className={styles.status}>Denied</p>
			<h1 className={styles.introTitle}>
				You cannot access the Open Data Service since you are not an active
				affiliate. Please contact the Office of Information Technology if this
				is incorrect.
			</h1>
			<Link href='/' className={styles.loginButton}>
				Log out
			</Link>
		</div>
	);
}
