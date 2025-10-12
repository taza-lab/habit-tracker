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

    // 入力ルールの定数
    const ALLOWED_CHARS_REGEX = /^[a-zA-Z0-9!@#$%^&*()_+=\-{}[\]|\\:;"'<>,.?/~`]*$/;
    const USERNAME_MAX_LENGTH = 10;
    const PASSWORD_MIN_LENGTH = 8;

    // ユーザー名チェック
    useEffect(() => {
        let isValid = true;
        let usernameRuleError = '';

        if (username.length > 0) {
            if (!ALLOWED_CHARS_REGEX.test(username)) {
                usernameRuleError = '使用できない文字が含まれています';
                isValid = false;
            } else if (username.length > USERNAME_MAX_LENGTH) {
                usernameRuleError = `${USERNAME_MAX_LENGTH}文字まで入力できます`;
                isValid = false;
            }

            setUsernameError(usernameRuleError);
            setIsUsernameValid(isValid)
            return;
        }

        setUsernameError('');
        setIsUsernameValid(false);

    }, [username]);


    // パスワードチェック
    useEffect(() => {
        let isValid = true;
        let passwordRuleError = '';
        let matchError = '';

        if (password.length > 0) {
            if (!ALLOWED_CHARS_REGEX.test(password)) {
                passwordRuleError = '使用できない文字が含まれています';
                isValid = false;
            } else if (password.length < PASSWORD_MIN_LENGTH) {
                passwordRuleError = `${PASSWORD_MIN_LENGTH}文字以上入力してください`;
                isValid = false;
            }

            // パスワード一致チェック (confirmPasswordが入力されている場合のみ)
            if (confirmPassword.length == 0) {
                isValid = false;
            } else if (confirmPassword !== password) {
                matchError = 'パスワードが一致しません';
                isValid = false;
            }

            setPasswordError(passwordRuleError);
            setConfirmPasswordError(matchError);
            setIsPasswordValid(isValid);
            return;
        }
        
        setPasswordError('');
        setConfirmPasswordError('');
        setIsPasswordValid(false);

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
                    inputProps={{
                        maxLength: 10,
                        pattern: '[a-zA-Z0-9!@#$%^&*()_+=\-{}[\]|\\:;"\'<>,.?/~`]*',
                    }}
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
                    inputProps={{
                        pattern: '[a-zA-Z0-9!@#$%^&*()_+=\-{}[\]|\\:;"\'<>,.?/~`]*',
                    }}
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
                    inputProps={{
                        pattern: '[a-zA-Z0-9!@#$%^&*()_+=\-{}[\]|\\:;"\'<>,.?/~`]*',
                    }}
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
