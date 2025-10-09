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
    const [usernameError, setUsernameError] = useState('');
    const [isUsernameValid, setIsUsernameValid] = useState(false);
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [passwordError, setPasswordError] = useState('');
    const [confirmPasswordError, setConfirmPasswordError] = useState('');
    const [isPasswordValid, setIsPasswordValid] = useState(false);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    // ユーザー名チェック
    useEffect(() => {
        if (username) {
            if (username.length < 5) {
                setUsernameError('5文字以上入力してください');
                setIsUsernameValid(false);
            } else {
                setUsernameError('');
                setIsUsernameValid(true)
            }
        }
    }, [username]);


    // パスワードチェック
    useEffect(() => {
        // if (password) {
        //     // 文字数チェック
        //     if (password && password.length < 8) {
        //         setPasswordError('8文字以上入力してください');
        //         setIsPasswordValid(false);
        //         return;
        //     } else {
        //         setConfirmPasswordError('');
        //     }

        //     if (confirmPassword) {
        //         // パスワード一致チェック
        //         if (password !== confirmPassword) {
        //             setConfirmPasswordError('パスワードが一致しません');
        //             setIsPasswordValid(false);
        //         } else {
        //             setConfirmPasswordError('');
        //             setIsPasswordValid(true);
        //         }
        //     }
            
        // } else {
        //     setConfirmPasswordError('');
        //     setIsPasswordValid(false);
        // }

        let isValid = true;
        let lengthError = '';
        let matchError = '';

        if (password) {
            // 文字数チェック
            if (password.length < 8) {
                lengthError = '8文字以上入力してください';
                isValid = false;
            }

            // パスワード一致チェック (confirmPasswordが入力されている場合のみ)
            if (confirmPassword) {
                if (password !== confirmPassword) {
                    matchError = 'パスワードが一致しません';
                    isValid = false;
                }
            } else if (password.length >= 8) {
                // passwordが8文字以上でも、確認用が空なら無効（ただしエラーメッセージは出さない）
                isValid = false;
            }

            // state更新
            setPasswordError(lengthError);
            setConfirmPasswordError(matchError);

            // passwordが入力されていて、かつ全チェックを通過した場合のみtrue
            setIsPasswordValid(isValid);

        } else {
            setPasswordError('');
            setConfirmPasswordError('');
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
                    error={!!usernameError}
                    helperText={usernameError}
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
                    error={!!passwordError}
                    helperText={passwordError}
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
                    error={!!confirmPasswordError}
                    helperText={confirmPasswordError}
                />
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                    disabled={!isUsernameValid || !isPasswordValid || loading}
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
