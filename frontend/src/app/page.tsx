'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from "next/navigation";
import { List, ListItem, ListItemIcon, ListItemText, Button } from '@mui/material';
import PageTitle from './components/PageTitle';
import CheckBox from './components/DailyTrack/CheckBox';
import { DailyTrack } from './types/daily-track';
import { fetchTodaysTrack, todaysHabitDone } from './features/daily-track/api';

export default function HabitManage() {
    // 定数定義
    const [todaysTrack, setTodaysTrack] = useState<DailyTrack>({ date: '', habits: [] }); //初期値をundefinedにしないために空のtypeをセット
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();

    // 初期表示
    useEffect(() => {
        const loadTodaysTrack = async () => {
            try {
                const data = await fetchTodaysTrack();
                setTodaysTrack(data);
            } catch (err) {
                setError('読み込みに失敗しました');
            } finally {
                setLoading(false);
            }
        };
        loadTodaysTrack();
    }, []);

    if (loading) return <p>読み込み中...</p>;
    if (error) return <p>{error}</p>;

    // チェック
    const handleHabitDoneCheck = async (targetId: string) => {
        try {
            // API実行
            await todaysHabitDone(targetId);

            // データ更新
            const updatedHabits = todaysTrack.habits.map((item) => {
                if (item.habit.id === targetId) {
                    return { ...item, isDone: true };
                }
                return item;
            });

            setTodaysTrack({ ...todaysTrack, habits: updatedHabits })

            // TODO: ポイント更新

            // TODO: 全部チェックしたらアニメーション表示


        } catch (err) {
            setError('更新に失敗しました');
        }
    }

    return (
        <div style={{ padding: '2rem' }}>
            <PageTitle title={todaysTrack.date ?? ''} />
            <List>
                {todaysTrack?.habits.length === 0 ? (
                    <>
                        <ListItem>
                            <ListItemText primary="習慣が登録されていません。" />
                        </ListItem>
                        <ListItem>
                            {/* TODO: 画面遷移にメニューバーのハイライトを連動させる */}
                            <Button variant="contained" color="primary" onClick={() => router.push("/habit-manage")}>習慣を登録する</Button>
                        </ListItem>
                    </>
                ) : (
                    todaysTrack?.habits.map(track => (
                        <ListItem key={track.habit.id}>
                            <ListItemIcon onClick={() => handleHabitDoneCheck(track.habit.id)}>
                                <CheckBox isChecked={track.isDone} />
                            </ListItemIcon>
                            <ListItemText primary={track.habit.name} />
                        </ListItem>
                    ))
                )}
            </List>
        </div>

    );
}