import { Habit } from '@/types/habit';
import { getAuthHeaders } from '@/lib/api';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function fetchHabits(): Promise<Habit[]> {
    const headers = getAuthHeaders();

    const res = await fetch(`${BASE_URL}/auth/habit/list`, { headers: headers });
    if (!res.ok) throw new Error('Failed to fetch habits');
    return res.json();
}

export async function createHabit(newHabit: Habit): Promise<{ id: string; message: string }> {
    const headers = getAuthHeaders();

    const res = await fetch(
        `${BASE_URL}/auth/habit/register`,
        {
            method: 'POST',
            body: JSON.stringify(newHabit),
            headers: headers
        }
    );
    if (!res.ok) throw new Error('Failed to create habits');
    return res.json();
}

export async function deleteHabit(id: string): Promise<Habit[]> {
    const headers = getAuthHeaders();

    const res = await fetch(`${BASE_URL}/auth/habit/${id}/delete`,
        {
            method: 'DELETE',
            headers: headers
        }
    );
    if (!res.ok) throw new Error('Failed to delete habits');
    return res.json();
}
