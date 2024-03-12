'use client';

import styles from '@/styles/pages/error.module.css';
import { usePathname, useRouter } from 'next/navigation';

export default function Error({ message, reset }: any) {
	const router = useRouter();
	const pathName = usePathname();

	return (
		<div className={styles.errorContainer}>
			<h2>Something went wrong!</h2>
			<p>We've contacted the team on this.</p>
			<button
				className={styles.tryAgainButton}
				onClick={
					// Attempt to recover by trying to re-render the segment
					() => router.replace(pathName)
				}
			>
				Try again
			</button>
		</div>
	);
}
