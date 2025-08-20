import { User } from '../../types/user';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function fetchUser(): Promise<User> {
    const res = await fetch(`${BASE_URL}/user`);
    if (!res.ok) throw new Error('Failed to fetch todays track');
    return res.json();
}