import { Url } from 'next/dist/shared/lib/router/router';
import Link from 'next/link';
import React from 'react';
import styles from './Sidebar.module.css';

const Sidebar = () => {
	return (
		<nav className={styles.sidebarContainer}>
			<p>
				<strong className={styles.sidebarHeader}>Docs</strong>
			</p>
			<ul className={styles.sidebarLinks}>
				<ul className={styles.sublist}>
					<NavLink title={'Introduction'} link={'docs/introduction'} />
				</ul>
				<li className={styles.sectionHeader}>
					<strong>Getting Started</strong>
				</li>
				<ul className={styles.sublist}>
					<NavLink title={'Overview'} link={'docs/getting_started'} />
					<NavLink
						title={'Registering an App'}
						link={'docs/registering_applications'}
					/>
					<NavLink title={'Making API Calls'} link={'docs/making_api_calls'} />
				</ul>
				<li className={styles.sectionHeader}>
					<strong>Usage Guides</strong>
				</li>
				<ul className={styles.sublist}>
					<NavLink title={'Authentication'} link={'docs/authentication'} />
					<NavLink title={'Rate Limits'} link={'docs/rate_limits'} />
					<NavLink title={'Error Handling'} link={'docs/error_handling'} />
				</ul>
				<li className={styles.sectionHeader}>
					<strong>Reference</strong>
				</li>
				<ul className={styles.sublist}>
					<NavLink title={'Data Formats'} link={'docs/data_formats'} />
				</ul>
				<li className={styles.sectionHeader}>
					<strong>Resources</strong>
				</li>
				<ul className={styles.sublist}>
					<NavLink title={'FAQ'} link={'/faq'} />
					<NavLink title={'Contact Us'} link={'/contact_us'} />
					<NavLink title={'Terms of Use'} link={'/terms_of_use'} />
				</ul>
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
		<li>
			<Link href={link}>
				<strong>{title}</strong>
			</Link>
		</li>
	);
};

export default Sidebar;
