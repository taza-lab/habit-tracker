'use client';

import React, { useEffect, useState } from 'react';
import { Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, List, ListItem, ListItemIcon, ListItemText, IconButton, Fab } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import SentimentSatisfiedRoundedIcon from '@mui/icons-material/SentimentSatisfiedRounded';
import DeleteIcon from '@mui/icons-material/Delete';
import CloseRoundedIcon from '@mui/icons-material/CloseRounded';
import { fetchHabits, createHabit, deleteHabit } from '@/features/habit-manage/api';
import { useAuthApi } from '../hooks/useAuthApi';
import { Habit } from '@/types/habit';
import PageTitle from '@/components/PageTitle';
import Loading from '@/components/Loading';
import { useAlert } from '@/context/AlertContext';

type Mode = 'create' | 'delete';

export default function HabitManage() {

    // 定数定義
    const { handleAuthApiCall } = useAuthApi();
    const { showAlert } = useAlert();
    const [habits, setHabits] = useState<Habit[]>([]);
    const [loading, setLoading] = useState(true);
    const [modalMode, setModalMode] = useState<Mode | null>(null);
    const [selectedItem, setSelectedItem] = useState<Habit | null>(null);

    // 初期表示
    useEffect(() => {
        const loadHabits = async () => {
            try {
                const data = await handleAuthApiCall(fetchHabits);
                if (data) {
                    setHabits(data);
                }
                
            } catch (err) {
                console.log(err);
                showAlert('読み込みに失敗しました', 'error');
            } finally {
                setLoading(false);
            }
        };
        loadHabits();
    }, [handleAuthApiCall]);

    if (loading) return <Loading />;

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
        const newHabit: Habit = { id: '0', userId: '', name };

        try {
            const result = await handleAuthApiCall(() => createHabit(newHabit));
            if (result) {
                newHabit.id = result.id;
                setHabits([...habits, newHabit]);
            }
        } catch (err) {
            console.log(err);
            showAlert('登録に失敗しました', 'error');
        }

        closeModal();
    };

    // 削除送信
    const handleDeleteSubmit = (id: string) => {
        try {
            handleAuthApiCall(() => deleteHabit(id));
            setHabits(prev => prev.filter(habit => habit.id !== id));
            
        } catch (err) {
            console.log(err);
            showAlert('削除に失敗しました', 'error');
        }
        closeModal();
    }


    return (
        <div style={{ padding: '2rem' }}>
            <PageTitle title="習慣一覧" />
            <List>
                {habits.length === 0 ? (
                    <ListItem>
                        <ListItemText primary="習慣が登録されていません。" />
                    </ListItem>
                ) : (
                    habits.map(habit => (
                        <ListItem key={habit.id} sx={{ px: 0 }}>
                            <ListItemIcon sx={{minWidth: 38}}>
                                <SentimentSatisfiedRoundedIcon />
                            </ListItemIcon>
                            <ListItemText primary={habit.name} />
                            <ListItemIcon sx={{minWidth: 45, justifyContent: 'flex-end'}}>
                                <DeleteIcon onClick={() => openModal('delete', habit)} />
                            </ListItemIcon>

                        </ListItem>
                    ))
                )}
            </List>
            <Fab
                color="secondary"
                aria-label="add"
                disabled={habits.length >= 5}
                sx={{
                    position: "fixed",
                    bottom: 120,
                    right: 50,
                    width: 60,
                    height: 60,
                }}
            >
                <AddIcon
                    sx={{ fontSize: 36 }}
                    onClick={() => openModal('create')}
                />
            </Fab>
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
            <Dialog open={true} onClose={onClose} fullWidth={true}>
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
                                inputProps={{
                                    maxLength: 20, // ここに制限したい最大文字数を指定 (例: 50文字)
                                }}
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
