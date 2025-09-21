'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from "next/navigation";
import { List, ListItem, ListItemIcon, ListItemText, Button } from '@mui/material';
import PageTitle from '@/components/PageTitle';
import CheckBox from '@/components/DailyTrack/CheckBox';
import { DailyTrack } from '@/types/daily-track';
import { useAuthApi } from './hooks/useAuthApi';
import { fetchTodaysTrack, todaysHabitDone } from '@/features/daily-track/api';
import { usePoint } from '@/context/PointContext';
import { useAlert } from '@/context/AlertContext';

export default function Home() {
    // 定数定義
    const [todaysTrack, setTodaysTrack] = useState<DailyTrack>({ id: '', userId: '', date: '', habitStatuses: [] }); //初期値をundefinedにしないために空のtypeをセット
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();
    const { handleAuthApiCall } = useAuthApi();
    const { addPoints } = usePoint();
    const { showAlert } = useAlert();

    // 初期表示
    useEffect(() => {
        const loadTodaysTrack = async () => {
            try {
                const data = await handleAuthApiCall(fetchTodaysTrack);
                if (data) {
                    setTodaysTrack(data);
                }
            } catch (err) {
                console.log(err);
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
    const handleHabitDoneCheck = async (targetHabitId: string) => {
        try {
            // API実行
            await handleAuthApiCall(() => todaysHabitDone(targetHabitId));

            // データ更新
            const updatedHabitStatuses = todaysTrack.habitStatuses.map((item) => {
                if (item.habitId === targetHabitId) {
                    return { ...item, isDone: true };
                }
                return item;
            });
            setTodaysTrack({ ...todaysTrack, habitStatuses: updatedHabitStatuses })

            // ポイント付与
            addPoints(Number(process.env.NEXT_PUBLIC_HABIT_DONE_POINT));

            // 全部チェックしたらメッセージ表示
            const allDone = todaysTrack.habitStatuses.filter((item) => {
                return !item.isDone;
            }).length === 1;

            if (allDone) {
                showAlert('all Done!', 'success');
            }

        } catch (err) {
            setError('更新に失敗しました');
        }
    }

    return (
        <div style={{ padding: '2rem' }}>
            <PageTitle title={formatDate(todaysTrack.date)} />
            <List>
                {todaysTrack && todaysTrack.habitStatuses && todaysTrack.habitStatuses.length > 0 ? (
                    todaysTrack.habitStatuses.map(track => (
                        <ListItem key={track.habitId}>
                            <ListItemIcon onClick={() => handleHabitDoneCheck(track.habitId)}>
                                <CheckBox isChecked={track.isDone} />
                            </ListItemIcon>
                            <ListItemText primary={track.habitName} />
                        </ListItem>
                    ))
                ) : (
                    <>
                        <ListItem>
                            <ListItemText primary="習慣が登録されていません。" />
                        </ListItem>
                        <ListItem>
                            {/* TODO: 画面遷移にメニューバーのハイライトを連動させる */}
                            <Button variant="contained" color="primary" onClick={() => router.push("/habit-manage")}>習慣を登録する</Button>
                        </ListItem>
                    </>
                )}
            </List>
        </div>

    );
}

// YYYY年m月d日 の形式に変換
const formatDate = (dateString: string) => {
  if (!dateString) {
    return '';
  }
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }).format(date);
};
