import { Habit } from '@/types/habit';
import { apiRequest } from '@/lib/api';

interface HabitApiData {
    id: string;
    user_id: string;
    name: string;
}

export async function fetchHabits(): Promise<Habit[]> {
    const apiData = await apiRequest<HabitApiData[]>(`auth/habit/list`);

    return apiData.map(habit => ({
        id: habit.id,
        userId: habit.user_id,
        name: habit.name,
    }));
}

export async function createHabit(newHabit: Habit): Promise<{ id: string; message: string }> {
    const res = await apiRequest<{ id: string; message: string }>(
        `auth/habit/register`,
        {
            method: 'POST',
            body: JSON.stringify(newHabit),
        }
    );

    return res;
}

export async function deleteHabit(id: string): Promise<void> {
    await apiRequest(
        `auth/habit/${id}/delete`,
        {
            method: 'DELETE',
        }
    );

    return;
}
