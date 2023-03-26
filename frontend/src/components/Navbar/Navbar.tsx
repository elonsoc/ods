import Link from 'next/link';
import styles from './Navbar.module.css';

const Navbar: React.FunctionComponent = () => {
	return (
		<nav className={styles.container}>
			<Link href='/'>Open Data @ Elon</Link>
		</nav>
	);
};

export default Navbar;
