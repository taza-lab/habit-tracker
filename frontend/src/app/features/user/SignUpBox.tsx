'use client';

import React, { useState, useEffect } from 'react';
import {
    Box,
    TextField,
    Button,
    Typography,
    Alert,
    CircularProgress,
    Divider
} from '@mui/material';
import SentimentSatisfiedAltRoundedIcon from '@mui/icons-material/SentimentSatisfiedAltRounded';
import { signup } from '@/features/user/api';

interface LoginBoxProps {
    setSuccessMessage: React.Dispatch<React.SetStateAction<string | null>>,
    switchToLogin: () => void;
}

export default function AuthPage({ setSuccessMessage, switchToLogin }: LoginBoxProps) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [passwordError, setPasswordError] = useState('');
    const [isPasswordValid, setIsPasswordValid] = useState(false);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    // パスワード一致チェック
    useEffect(() => {
        // 両方のフィールドが空でなければチェックを行う
        if (password && confirmPassword) {
            if (password !== confirmPassword) {
                setPasswordError('パスワードが一致しません');
                setIsPasswordValid(false);
            } else {
                setPasswordError(''); // エラーをクリア
                setIsPasswordValid(true);
            }
        } else {
            setPasswordError('');
            setIsPasswordValid(false);
        }
    }, [password, confirmPassword]);

    const handleSignUp = async () => {
        setLoading(true);
        setError(null);

        try {
            // API実行
            signup(username, password, confirmPassword);

            // サインアップ成功ダイアログ表示
            setSuccessMessage('アカウント作成が成功しました。ログインしてください。');

            // ログインモーダルに切り替え
            switchToLogin();

        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
            }
        } finally {
            setLoading(false);
        }
    };

    return (

        <Box
            sx={{
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
            }}
        >
            <SentimentSatisfiedAltRoundedIcon sx={{ fontSize: 40, color: 'primary.main' }} />
            <Typography component="h1" variant="h5" mt={1}>
                Sign Up
            </Typography>
            <Box component="form" sx={{ mt: 1, width: '100%' }}>
                {/* エラーメッセージ */}
                {error && <Alert severity="error" sx={{ my: 2 }}>{error}</Alert>}

                <TextField
                    margin="normal"
                    required
                    fullWidth
                    label="username"
                    name="username"
                    autoComplete="username"
                    autoFocus
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <TextField
                    margin="normal"
                    required
                    fullWidth
                    name="password"
                    label="password"
                    type="password"
                    autoComplete="current-password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <TextField
                    margin="normal"
                    required
                    fullWidth
                    name="confirm_password"
                    label="password (確認)"
                    type="password"
                    autoComplete="current-password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    error={!!passwordError}
                    helperText={passwordError}
                />
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                    disabled={!username || !password || !confirmPassword || !isPasswordValid || loading}
                    onClick={() => handleSignUp()}
                >
                    {loading ? <CircularProgress size={24} /> : 'Sign Up'}
                </Button>

                <Divider sx={{ my: 2 }}>or</Divider>

                <Button
                    fullWidth
                    variant="outlined"
                    sx={{ mt: 3 }}
                    onClick={() => switchToLogin()}
                >
                    Login
                </Button>
            </Box>
        </Box>
    );
}
