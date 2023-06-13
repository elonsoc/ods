'use client';

import React, {
	useState,
	useRef,
	useEffect,
	SetStateAction,
	Dispatch,
} from 'react';

import styles from './TableOfContents.module.css';
import { usePathname } from 'next/navigation';

const useIntersectionObserver = (
	setActiveId: Dispatch<SetStateAction<string | undefined>>,
	pathname: string
) => {
	const headingElementsRef = useRef<{
		[key: string]: IntersectionObserverEntry;
	}>({});
	useEffect(() => {
		const callback: IntersectionObserverCallback = (
			headings: IntersectionObserverEntry[]
		) => {
			headingElementsRef.current = headings.reduce((map, headingElement) => {
				map[headingElement.target.id] = headingElement;
				return map;
			}, headingElementsRef.current);

			// Get all headings that are currently visible on the page
			const visibleHeadings: IntersectionObserverEntry[] = [];
			Object.keys(headingElementsRef.current).forEach((key) => {
				const headingElement = headingElementsRef.current[key];
				if (headingElement.isIntersecting) visibleHeadings.push(headingElement);
			});

			const getIndexFromId = (id: any) =>
				headingElements.findIndex((heading) => heading.id === id);

			if (visibleHeadings.length === 1) {
				setActiveId(visibleHeadings[0].target.id);
			} else if (visibleHeadings.length > 1) {
				const sortedVisibleHeadings = visibleHeadings.sort(
					(a, b) => getIndexFromId(a.target.id) - getIndexFromId(b.target.id)
				);

				setActiveId(sortedVisibleHeadings[0].target.id);
			}
		};

		const observer = new IntersectionObserver(callback, {
			root: document.querySelector('iframe'),
		});

		const headingElements = Array.from(document.querySelectorAll('h2'));

		headingElements.forEach((element) => observer.observe(element));

		return () => observer.disconnect();
	}, [setActiveId, pathname]);
};

const SkeletonLoader = () => {
	return (
		<aside className={styles.tocContainer}>
			<p>
				<strong>On This Page</strong>
			</p>
			<nav className={styles.toc} aria-label='Table of Contents'>
				<ul>
					<li className={styles.loadingItem}>Loading...</li>
					<li className={styles.loadingItem}>Loading...</li>
					<li className={styles.loadingItem}>Loading...</li>
					<li className={styles.loadingItem}>Loading...</li>
					<li className={styles.loadingItem}>Loading...</li>
				</ul>
			</nav>
		</aside>
	);
};

const TableOfContents = () => {
	const pathName = usePathname();
	const [activeId, setActiveId] = useState<string>();
	const [headings, setHeadings] = useState<HTMLElement[]>([]);
	const [loading, setLoading] = useState(true);
	useEffect(() => {
		const headingElements: HTMLElement[] = Array.from(
			document.querySelectorAll('h2')
		);
		setHeadings(headingElements);
		setLoading(false);
	}, [pathName]);

	useIntersectionObserver(setActiveId, pathName);

	if (loading) {
		return <SkeletonLoader />;
	}

	if (!headings.length) {
		return <></>;
	}

	return (
		<aside className={styles.tocContainer}>
			<p>
				<strong>On This Page</strong>
			</p>
			<nav className={styles.toc} aria-label='Table of Contents'>
				<ul>
					{headings.map((heading: any) => (
						<li
							key={heading.id}
							className={heading.id === activeId ? `${styles.active}` : ''}
						>
							<a href={`#${heading.id}`}>{heading.innerText}</a>
						</li>
					))}
				</ul>
			</nav>
		</aside>
	);
};

export default TableOfContents;
