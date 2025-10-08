import { DailyTrack } from '@/types/daily-track';
import { apiRequest } from '@/lib/api';

interface DailyTrackApiData {
  id: string;
  user_id: string;
  date: string;
  habit_statuses: {
    habit_id: string;
    habit_name: string;
    is_done: boolean;
  }[];
}

export async function fetchTodaysTrack(): Promise<DailyTrack> {
    const today = createTodayString();

    const apiData = await apiRequest<DailyTrackApiData>(`auth/daily_track/${today}`);
    
    return {
        id: apiData.id,
        userId: apiData.user_id,
        date: apiData.date,
        habitStatuses: apiData.habit_statuses != null ? apiData.habit_statuses.map(status => ({
            habitId: status.habit_id,
            habitName: status.habit_name,
            isDone: status.is_done,
        })) : [],
    };
}

export async function todaysHabitDone(habitId: string): Promise<void> {
    const today = createTodayString();

    await apiRequest(
        `auth/daily_track/done`, 
        {
            method: 'POST',
            body: JSON.stringify({ date: today, habit_id: habitId })
        }
    )
}


function createTodayString(): string {
    return new Date().toLocaleDateString("ja-JP", {
        year: "numeric", month: "2-digit",
        day: "2-digit"
    }).replaceAll('/', '-')
}
