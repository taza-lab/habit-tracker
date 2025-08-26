import { User } from '@/types/user';
import { getAuthHeaders } from '@/lib/api';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function login(username: string, password: string): Promise<{ token: string, user: User }> {
    const res = await fetch(
        `${BASE_URL}/login`,
        {
            method: 'POST',
            body: JSON.stringify({ username, password })
        }
    );
    if (!res.ok) throw new Error('Failed to login');
    return res.json();
}

export async function fetchUser(): Promise<User> {
    const headers = getAuthHeaders();
    const res = await fetch(`${BASE_URL}/auth/user`, { headers: headers });
    if (!res.ok) throw new Error('Failed to fetch todays track');
    return res.json();
}