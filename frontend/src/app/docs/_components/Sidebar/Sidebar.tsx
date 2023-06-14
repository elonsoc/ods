'use client';

import Link from 'next/link';
import React, { useState } from 'react';
import styles from './Sidebar.module.css';
import { usePathname } from 'next/navigation';
import { config } from '@/config/Constants';

const URL = config.url.API_URL;

const Sidebar = () => {
	return (
		<aside>
			<nav className={styles.sidebarContainer} aria-label='Related Topics'>
				<header>
					<p className={styles.sidebarHeader}>
						<strong>
							<Link href='docs'>Docs</Link>
						</strong>
					</p>
				</header>
				<ul className={styles.sidebarLinks}>
					<li className={styles.sectionHeader}>
						<p>
							<strong className={styles.sublistHeader}>
								<Link href='docs/getting-started'>Getting Started</Link>
							</strong>
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
							<strong className={styles.sublistHeader}>
								<Link href='docs/usage-guides'>Usage Guides</Link>
							</strong>
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
							<strong className={styles.sublistHeader}>
								<Link href='docs/reference'>Reference</Link>
							</strong>
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
							<strong className={styles.sublistHeader}>
								<Link href='docs/resources'>Resources</Link>
							</strong>
						</p>
						<ul className={styles.sublist}>
							<NavLink title={'FAQ'} link={'/faq'} />
							<NavLink title={'Contact Us'} link={'/contact_us'} />
							<NavLink title={'Terms of Use'} link={'/terms_of_use'} />
						</ul>
					</li>
				</ul>
			</nav>
		</aside>
	);
};

interface NavLinkProps {
	title: string;
	link: string;
}

const NavLink = ({ title, link }: NavLinkProps) => {
	const pathName = usePathname();
	const path = pathName.replace(URL, '');
	return (
		<li
			className={`${styles.docLink} ${path == '/' + link ? styles.active : ''}`}
		>
			<Link href={link}>{title}</Link>
		</li>
	);
};

interface NavLinkDropdownProps {
	title: string;
	link: string;
	sublinks: NavLinkProps[];
}

const NavLinkDropdown = ({ title, link, sublinks }: NavLinkDropdownProps) => {
	const [dropdownActive, setDropdownActive] = useState(false);
	const pathName = usePathname();
	const path = pathName.replace(URL, '');
	return (
		<li
			className={`${styles.docLink} ${path == '/' + link ? styles.active : ''}`}
		>
			<Link href={link}>{title}</Link>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
				<title>Dropdown</title>
				<path d='M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z' />
			</svg>
			<ul>
				{sublinks.map((sublink) => (
					<li>
						<Link href={sublink.title}>{sublink.title}</Link>
					</li>
				))}
			</ul>
		</li>
	);
};

export default Sidebar;
