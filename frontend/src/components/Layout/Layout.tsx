import styles from './Layout.module.css';
import Navbar from '../Navbar/Navbar';
import Footer from '../Footer/Footer';
import { Inter } from 'next/font/google';
import Meta from '../Meta';

const inter = Inter({ subsets: ['latin'] });

interface Props {
	children: React.ReactNode;
}

const Layout: React.FunctionComponent<Props> = ({ children }: Props) => {
	return (
		<>
			<Meta />
			<div className={styles.container}>
				<Navbar />
				<main className={inter.className}>{children}</main>
				<Footer />
			</div>
		</>
	);
};

export default Layout;
