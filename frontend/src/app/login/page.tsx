'use client';

import React, { useState, useEffect } from 'react';
import { Container } from '@mui/material';
import LoginBox from "@/features/user/LoginBox"
import SignUpBox from "@/features/user/SignUpBox"

export default function AuthPage() {
    const [success, setSuccess] = useState<string | null>(null);
    const [showLogin, setShowLogin] = useState(true);
    const [showSignup, setShowSignup] = useState(false);

    // 初期表示
    useEffect(() => {
        localStorage.removeItem('jwt_token'); 
        localStorage.removeItem('username'); 
    }, []);

    // ログインモーダルからサインアップモーダルに切り替える関数
    const switchToSignup = () => {
        setSuccess('');
        setShowLogin(false);
        setShowSignup(true);
    };

    // サインアップモーダルからログインモーダルに切り替える関数
    const switchToLogin = () => {
        setShowSignup(false);
        setShowLogin(true);
    };

    return (
        <Container component="main" maxWidth="xs">
            {showLogin && !showSignup
                ? <LoginBox success={success} setSuccessMessage={setSuccess} switchToSignup={switchToSignup} />
                : <SignUpBox setSuccessMessage={setSuccess} switchToLogin={switchToLogin} />
            }
        </Container>
    );
}
