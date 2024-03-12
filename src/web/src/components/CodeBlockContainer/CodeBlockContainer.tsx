import React from 'react';
import styles from './CodeBlockContainer.module.css';

const CodeBlockContainer = ({
	text,
	codeType,
}: {
	text: string;
	codeType: string;
}) => {
	return (
		<div className={styles.jsonContainer}>
			<span className={styles.jsonHeader}>
				<strong>{codeType}</strong>
			</span>
			<pre className={styles.jsonCode}>
				<code>{text}</code>
			</pre>
		</div>
	);
};

export default CodeBlockContainer;
