import { Url } from 'next/dist/shared/lib/router/router';
import Link from 'next/link';
import React from 'react';
import styles from './Sidebar.module.css';

const Sidebar = () => {
	return (
		<nav className={styles.sidebarContainer}>
			<p className={styles.sidebarHeader}>
				<strong>Docs</strong>
			</p>
			<ul className={styles.sidebarLinks}>
				<li className={styles.sectionHeader}>
					<p>
						<strong className={styles.sublistHeader}>Getting Started</strong>
					</p>
					<ul className={styles.sublist}>
						<NavLink
							title={'Overview'}
							link={'docs/getting-started/overview'}
						/>
						<NavLink
							title={'Registering an App'}
							link={'docs/getting-started/registering-an-app'}
						/>
						<NavLink
							title={'Making API Calls'}
							link={'docs/getting-started/making-api-calls'}
						/>
					</ul>
				</li>
				<li className={styles.sectionHeader}>
					<p>
						<strong className={styles.sublistHeader}>Usage Guides</strong>
					</p>
					<ul className={styles.sublist}>
						<NavLink
							title={'Authentication'}
							link={'docs/usage-guides/authentication'}
						/>
						<NavLink
							title={'Rate Limits'}
							link={'docs/usage-guides/rate-limits'}
						/>
						<NavLink
							title={'Error Handling'}
							link={'docs/usage-guides/error-handling'}
						/>
					</ul>
				</li>
				<li className={styles.sectionHeader}>
					<p>
						<strong className={styles.sublistHeader}>Reference</strong>
					</p>
					<ul className={styles.sublist}>
						<NavLink
							title={'Data Formats'}
							link={'docs/reference/data-formats'}
						/>
					</ul>
				</li>
				<li className={styles.sectionHeader}>
					<p>
						<strong className={styles.sublistHeader}>Resources</strong>
					</p>
					<ul className={styles.sublist}>
						<NavLink title={'FAQ'} link={'/faq'} />
						<NavLink title={'Contact Us'} link={'/contact_us'} />
						<NavLink title={'Terms of Use'} link={'/terms_of_use'} />
					</ul>
				</li>
			</ul>
		</nav>
	);
};

interface NavLinkProps {
	title: String;
	link: Url;
}

const NavLink = ({ title, link }: NavLinkProps) => {
	return (
		<li className={styles.docLink}>
			<Link href={link}>{title}</Link>
		</li>
	);
};

export default Sidebar;
