'use client';

import styles from '@/styles/pages/error.module.css';

export default function Error({ message, reset }: any) {
	return (
		<div className={styles.errorContainer}>
			<h2>Something went wrong!</h2>
			<p>We've contacted the team on this.</p>
			<button
				className={styles.tryAgainButton}
				onClick={
					// Attempt to recover by trying to re-render the segment
					() => reset()
				}
			>
				Try again
			</button>
		</div>
	);
}
