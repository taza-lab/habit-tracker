import { User } from '@/types/user';
import { getAuthHeaders } from '@/lib/api';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function login(username: string, password: string): Promise<{ token: string, user: User }> {
    const res = await fetch(
        `${BASE_URL}/login`,
        {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        }
    );
    if (!res.ok) throw new Error('Failed to login');
    return res.json();
}

export async function signup(username: string, password: string, confirmPassword: string): Promise<void> {
    const res = await fetch(
        `${BASE_URL}/signup`,
        {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password, confirm_password: confirmPassword })
        }
    );

    // ステータスコードが200番台以外の場合
    if (!res.ok) {
        const errorData = await res.json();
        const errorMessage = errorData.error || 'Failed to signup';
        throw new Error(errorMessage);
    }

    return res.json();
}

export async function fetchUser(): Promise<User> {
    const headers = getAuthHeaders();
    const res = await fetch(`${BASE_URL}/auth/user`, { headers: headers });
    if (!res.ok) throw new Error('Failed to fetch todays track');
    return res.json();
}