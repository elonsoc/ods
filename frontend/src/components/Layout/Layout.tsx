import styles from './Layout.module.css';
import Navbar from '../Navbar/Navbar';
import Footer from '../Footer/Footer';

interface Props {
	children: React.ReactNode;
}

const Layout: React.FunctionComponent<Props> = ({ children }: Props) => {
	return (
		<div className={styles.container}>
			<Navbar />
			{children}
			<Footer />
		</div>
	);
};

export default Layout;
