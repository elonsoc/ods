'use client';

import {
	createContext,
	useState,
	useContext,
	useEffect,
	ReactNode,
} from 'react';
import { configuration } from '@/config/Constants';
import { refreshTokens } from '@/actions/token';

interface AuthProviderProps {
	children: ReactNode;
}

interface AuthContextType {
	isAuthenticated: boolean;
	loading: boolean;
}

const API_URL = configuration.url.API_URL;
const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: AuthProviderProps) {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
	const [loading, setLoading] = useState<boolean>(true);

	async function checkSession() {
		const options: RequestInit = {
			method: 'GET',
			credentials: 'include',
			cache: 'no-store',
		};

		const res = await fetch(`${API_URL}/api/login/status`, options);
		const { isAuthenticated } = await res.json();
		setIsAuthenticated(isAuthenticated);
		setLoading(false);
	}

	useEffect(() => {
		checkSession();

		const intervalId = setInterval(() => {
			if (isAuthenticated) {
				refreshTokens();
			}
		}, 4 * 60 * 1000 + 30 * 1000);

		return () => clearInterval(intervalId);
	}, [isAuthenticated]);

	return (
		<AuthContext.Provider
			value={{
				isAuthenticated: isAuthenticated,
				loading: loading,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth() {
	const context = useContext(AuthContext);
	if (context === undefined) {
		throw new Error('useAuth must be used within an AuthProvider');
	}
	return context;
}
