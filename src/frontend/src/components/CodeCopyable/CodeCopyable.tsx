'use client';

import React, { useState } from 'react';
import style from './CodeCopyable.module.css';

const CodeCopyable = ({ code }: { code: string }) => {
	const [copySuccess, setCopySuccess] = useState(false);

	function copyToClipboard() {
		setCopySuccess(true);
		navigator.clipboard.writeText(code);
		setTimeout(() => {
			setCopySuccess(false);
		}, 1000);
	}

	return (
		<div className={style.codeCopyableContainer}>
			<div className={style.copyButtonContainer}>
				{copySuccess && <CopySuccessAlert />}
				<button onClick={copyToClipboard} className={style.copyButton}>
					<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
						<title>Copy To Clipboard</title>
						<path d='M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z' />
					</svg>
				</button>
			</div>
			<p className={style.codeBlock}>
				<code className='code'>{code}</code>
			</p>
		</div>
	);
};

export default CodeCopyable;

const CopySuccessAlert = () => {
	return <span className={style.successAlert}>Copied!</span>;
};
