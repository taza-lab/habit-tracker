import { DailyTrack } from '@/types/daily-track';
import { getAuthHeaders } from '@/lib/api';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function fetchTodaysTrack(): Promise<DailyTrack> {
    const today = createTodayString();
    const headers = getAuthHeaders();

    const res = await fetch(`${BASE_URL}/auth/daily_track/${today}`, { headers: headers });
    if (!res.ok) throw new Error('Failed to fetch todays track');

    const apiData = await res.json();
    
    return {
        id: apiData.id,
        userId: apiData.user_id,
        date: apiData.date,
        habitStatuses: apiData.habit_statuses.map(status => ({
            habitId: status.habit_id,
            habitName: status.habit_name,
            isDone: status.is_done,
        })),
    };
}

export async function todaysHabitDone(habitId: string): Promise<void> {
    const today = createTodayString();
    const headers = getAuthHeaders();

    const res = await fetch(
        `${BASE_URL}/auth/daily_track/done`,
        {
            method: 'POST',
            headers: headers,
            body: JSON.stringify({ date: today, habit_id: habitId })
        }
    );
    if (!res.ok) throw new Error('Failed to fetch todays track');
}


function createTodayString(): string {
    return new Date().toLocaleDateString("ja-JP", {
        year: "numeric", month: "2-digit",
        day: "2-digit"
    }).replaceAll('/', '-')
}
