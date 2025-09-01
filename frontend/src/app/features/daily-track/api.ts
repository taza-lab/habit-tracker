import { DailyTrack } from '@/types/daily-track';
import { getAuthHeaders } from '@/lib/api';

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function fetchTodaysTrack(): Promise<DailyTrack> {
    const today = createTodayString();
    const headers = getAuthHeaders();

    const res = await fetch(`${BASE_URL}/auth/daily_track/${today}`, { headers: headers });
    if (!res.ok) throw new Error('Failed to fetch todays track');
    return res.json();
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