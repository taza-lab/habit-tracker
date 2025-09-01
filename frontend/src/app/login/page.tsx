'use client';

import React, { useState } from 'react';
import { useRouter } from "next/navigation";
import {
    Container,
    Box,
    TextField,
    Button,
    Typography,
    Alert,
    CircularProgress,
    Divider
} from '@mui/material';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { login } from '@/features/user/api';
import { usePoint } from '@/context/PointContext';

export default function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();
    const { setPoints } = usePoint();

    const handleLogin = async () => {
        setLoading(true);
        setError(null);

        try {
            // API実行
            const data = await login(username, password);

            // JWTトークンをローカルストレージに保存
            localStorage.setItem('jwt_token', data.token);
            localStorage.setItem('username', data.user.username);
            setPoints(data.user.points);

            // TOP画面リダイレクト
            router.push("/")

        } catch (err) {
            setError('Login faild');
        } finally {
            setLoading(false);
        }
    };

    return (
        <Container component="main" maxWidth="xs">
            <Box
                sx={{
                    marginTop: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <LockOutlinedIcon sx={{ fontSize: 40, color: 'primary.main' }} />
                <Typography component="h1" variant="h5" mt={1}>
                    Login
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
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!username || !password || loading}
                        onClick={() => handleLogin()}
                    >
                        {loading ? <CircularProgress size={24} /> : 'Login'}
                    </Button>

                    <Divider sx={{ my: 2 }}>or</Divider>

                    <Button
                        fullWidth
                        variant="outlined"
                        sx={{ mt: 3 }}
                        onClick={() => { router.push("/signup") }}
                    >
                        Create Account
                    </Button>
                </Box>
            </Box>
        </Container>
    );
}