import { Habit } from './habit';

export type DailyTrack = {
    date: string,
    habits: {
        habit: Habit,
        isDone: boolean
    }[]
}