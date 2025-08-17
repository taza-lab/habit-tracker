import { Habit } from '../../types/habit';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function fetchHabits(): Promise<Habit[]> {
    const res = await fetch(`${BASE_URL}/habit/list`);
    if (!res.ok) throw new Error('Failed to fetch habits');
    return res.json();
}

export async function createHabit(newHabit: Habit): Promise<{ id: string; message: string }> {
    const res = await fetch(
        `${BASE_URL}/habit/register`,
        {
            method: 'POST',
            body: JSON.stringify(newHabit)
        }
    );
    if (!res.ok) throw new Error('Failed to create habits');
    return res.json();
}

export async function deleteHabit(id: string): Promise<Habit[]> {
    const res = await fetch(`${BASE_URL}/habit/${id}/delete`, { method: 'DELETE' });
    if (!res.ok) throw new Error('Failed to delete habits');
    return res.json();
}
