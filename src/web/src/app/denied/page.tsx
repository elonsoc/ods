import styles from '@/styles/pages/denied.module.css';
import Link from 'next/link';

export default function Denied() {
	return (
		<div className={styles.container}>
			<p className={styles.status}>Denied</p>
			<p className={styles.statusMessage}>
				<strong>
					You cannot access the Open Data Service since you are not an active
					affiliate. Please contact the Office of Information Technology if this
					is incorrect.
				</strong>
			</p>
			<Link href='/' className={styles.button}>
				Log out
			</Link>
		</div>
	);
}
