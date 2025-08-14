'use client';

import React, { useEffect, useState } from 'react';
import { Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, List, ListItem, ListItemIcon, ListItemText, IconButton } from '@mui/material';
import AddCircleRoundedIcon from '@mui/icons-material/AddCircleRounded';
import SentimentSatisfiedRoundedIcon from '@mui/icons-material/SentimentSatisfiedRounded';
import DeleteIcon from '@mui/icons-material/Delete';
import CloseRoundedIcon from '@mui/icons-material/CloseRounded';
import { fetchHabits, createHabit, deleteHabit } from './features/habit/api';
import { Habit } from './types/habit';

type Mode = 'create' | 'edit' | 'delete';

export default function Home() {

    // 定数定義
    const [habits, setHabits] = useState<Habit[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [modalMode, setModalMode] = useState<Mode | null>(null);
    const [selectedItem, setSelectedItem] = useState<Habit | null>(null);

    // 初期表示
    useEffect(() => {
        const loadHabits = async () => {
            try {
                const data = await fetchHabits();
                setHabits(data);
            } catch (err) {
                setError('読み込みに失敗しました');
            } finally {
                setLoading(false);
            }
        };
        loadHabits();
    }, []);

    if (loading) return <p>読み込み中...</p>;
    if (error) return <p>{error}</p>;

    // モーダル操作
    const openModal = (mode: Mode, habit?: Habit) => {
        setModalMode(mode);
        setSelectedItem(habit ?? null);
    };

    const closeModal = () => {
        setModalMode(null);
        setSelectedItem(null);
    };

    // 新規登録送信
    const handleCreateSubmit = async (name: string) => {
        const newHabit: Habit = { id: '0', name };

        try {
            const result = await createHabit(newHabit);
            newHabit.id = result.id;
            setHabits([...habits, newHabit]);
        } catch (err) {
            setError('登録に失敗しました');
        }

        closeModal();
    };

    // 削除送信
    const handleDeleteSubmit = (id: string) => {
        setHabits(prev => prev.filter(habit => habit.id !== id));

        try {
            const result = deleteHabit(id);
        } catch (err) {
            setError('削除に失敗しました');
        }
        closeModal();
    }


    return (
        <div style={{ padding: '2rem' }}>
            <h1>習慣一覧</h1>
            {habits.length === 0 ? (
                <p>習慣が登録されていません。</p>
            ) : (
                <List>
                    {habits.map(habit => (
                        <ListItem key={habit.id}>
                            <ListItemIcon>
                                <SentimentSatisfiedRoundedIcon />
                            </ListItemIcon>
                            <ListItemText primary={habit.name} />
                            <ListItemIcon>
                                <DeleteIcon onClick={() => openModal('delete', habit)} />
                            </ListItemIcon>

                        </ListItem>
                    ))}
                </List>
            )}
            <AddCircleRoundedIcon onClick={() => openModal('create')} />

            {modalMode && (
                <Modal
                    mode={modalMode}
                    habit={selectedItem}
                    onClose={closeModal}
                    onCreateSubmit={handleCreateSubmit}
                    onDeleteSubmit={handleDeleteSubmit}
                />
            )}
        </div>
    );
}

// モーダルコンポーネント
type ModalProps = {
    mode: Mode;
    habit: Habit | null;
    onClose: () => void;
    onCreateSubmit: (name: string) => void;
    onDeleteSubmit: (id: string) => void;
};

const Modal = ({ mode, habit, onClose, onCreateSubmit, onDeleteSubmit }: ModalProps) => {
    const [name, setName] = useState(habit?.name || '');

    const title =
        mode === 'create'
            ? '新規登録'
            : '削除確認';

    return (
        <div>
            <Dialog open={true} onClose={onClose}>
                <DialogTitle>{title}</DialogTitle>
                <IconButton
                    aria-label="close"
                    onClick={onClose}
                    sx={(theme) => ({
                        position: 'absolute',
                        right: 8,
                        top: 8,
                        color: theme.palette.grey[500],
                    })}
                >
                    <CloseRoundedIcon />
                </IconButton>

                {mode === 'delete' && habit !== null ? (
                    <div>
                        <DialogContent>
                            <p>「{habit.name}」を削除してもよろしいですか？</p>
                        </DialogContent>
                        <DialogActions>
                            <Button variant="contained" color="error" onClick={() => onDeleteSubmit(habit.id)}>削除</Button>
                        </DialogActions>
                    </div>
                ) : (
                    <div>
                        <DialogContent>
                            <TextField
                                autoFocus
                                required
                                margin="dense"
                                variant="standard"
                                fullWidth
                                type="text"
                                value={name}
                                onChange={e => setName(e.target.value)}
                                placeholder="新しい習慣を入力"
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button variant="contained" color="primary" onClick={() => onCreateSubmit(name)} disabled={!name.trim()}>登録</Button>
                        </DialogActions>
                    </div>

                )}
            </Dialog>
        </div>
    );
};
