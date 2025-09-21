import { AppBar, Toolbar, Typography, Box, Container, BottomNavigation, BottomNavigationAction, Paper } from "@mui/material";
import React, { ReactNode } from "react";
import { useState, useEffect } from "react";
import { useRouter, usePathname } from "next/navigation";
import AccountCircleRoundedIcon from '@mui/icons-material/AccountCircleRounded';
import FactCheckRoundedIcon from '@mui/icons-material/FactCheckRounded';
import SettingsIcon from "@mui/icons-material/Settings";
import SavingsRoundedIcon from '@mui/icons-material/SavingsRounded';
import { usePoint } from '@/context/PointContext'; // Contextからフックをインポート
import { AlertDisplay } from './AlertDisplay';

type LayoutProps = {
    showLoginedLayout: Boolean;
    children: ReactNode;
};

const Layout = ({ showLoginedLayout, children }: LayoutProps) => {
    const [selectedMenu, setSelectedMenu] = useState(0);
    const [userName, setUserName] = useState('please login');
    const router = useRouter();
    const pathname = usePathname();
    const { points } = usePoint();

    // コンポーネントがマウントされた後（クライアントサイドでのみ）にlocalStorageからユーザー名を読み込む
    useEffect(() => {
        const storedUserName = localStorage.getItem('username');
        if (storedUserName) {
            setUserName(storedUserName);
        }
    }, []);

    // URLのパスが変更されたとき
    useEffect(() => {
        // selectedMenuを更新
        switch (pathname) {
            case '/':
                setSelectedMenu(0);
                break;
            case '/habit-manage':
                setSelectedMenu(1);
                break;
            default:
                // 該当するパスがない場合は、どのメニューも選択しない
                setSelectedMenu(null);
                break;
        }

        // localStorageからユーザー名を再度読み込む
        const storedUserName = localStorage.getItem('username');
        if (storedUserName) {
            setUserName(storedUserName);
        }
    }, [pathname]);

    return (
        <Box sx={{ display: "flex", flexDirection: "column", minHeight: "100vh" }}>
            {/* タイトル */}
            <AppBar position="static">
                <Toolbar sx={{ minHeight: '8vh', alignItems: 'center' }}>
                    <Typography variant="h6" sx={{ flexGrow: 1 }}>
                        <AccountCircleRoundedIcon /> {userName}
                    </Typography>
                    {showLoginedLayout &&
                        <Typography variant="h6">
                            <SavingsRoundedIcon /> {points}pt
                        </Typography>
                    }

                </Toolbar>
            </AppBar>

            {/* メインコンテンツ */}
            <Container sx={{ flex: 1, py: 3 }}>{children}</Container>

            {/* アラート */}
            <AlertDisplay />

            {/* フッターメニュー */}
            {showLoginedLayout &&
                <Paper sx={{ position: "fixed", bottom: 0, left: 0, right: 0, height: "70px" }} elevation={3}>
                    <BottomNavigation
                        sx={{ pt: "15px" }}
                        showLabels
                        value={selectedMenu}
                    >
                        <BottomNavigationAction label="Todays" icon={<FactCheckRoundedIcon />} onClick={() => router.push("/")} />
                        <BottomNavigationAction label="Settings" icon={<SettingsIcon />} onClick={() => router.push("/habit-manage")} />
                    </BottomNavigation>
                </Paper>
            }
        </Box >
    );
};

export default Layout;
