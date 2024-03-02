'use client';

import Link from 'next/link';
import styles from './Navbar.module.css';
import { useAuth } from '@/context/auth/auth';
import { configuration } from '@/config/Constants';

const BACKEND_URL = configuration.url.BACKEND_API_URL;

const Navbar = () => {
	const { isAuthenticated } = useAuth();

	return (
		<nav className={styles.container}>
			<Link href='/' className={styles.logo}>
				Open Data Service
			</Link>
			<Link href='/' className={`${styles.logo} ${styles.mobileLogo}`}>
				ODS
			</Link>
			<div className={styles.navOptions}>
				<ul className={styles.links}>
					<li className={styles.link}>
						<Link href='/docs'>Docs</Link>
					</li>
					<li className={styles.link}>
						<Link href='/apps'>Apps</Link>
					</li>
				</ul>
				{isAuthenticated ? (
					<Link className={styles.loginButton} href={BACKEND_URL + '/logout'}>
						Log out
					</Link>
				) : (
					<Link className={styles.loginButton} href={BACKEND_URL}>
						Log in
					</Link>
				)}
			</div>
		</nav>
	);
};

export default Navbar;
