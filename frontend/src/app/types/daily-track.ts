export type DailyTrack = {
    id: string,
    userId: string,
    date: string,
    habitStatuses: {
        habitId: string,
        habitName: string,
        isDone: boolean
    }[]
}
