import { AppBar, Toolbar, Typography, Box, Container, BottomNavigation, BottomNavigationAction, Paper } from "@mui/material";
import React, { ReactNode } from "react";
import { useState } from "react";
import { useRouter } from "next/navigation";
import HomeIcon from "@mui/icons-material/Home";
import SearchIcon from "@mui/icons-material/Search";
import NotificationsIcon from "@mui/icons-material/Notifications";
import SettingsIcon from "@mui/icons-material/Settings";
import SavingsRoundedIcon from '@mui/icons-material/SavingsRounded';
import { usePoint } from '../context/PointContext'; // Contextからフックをインポート
import { AlertDisplay } from './AlertDisplay';

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    const [selectedMenu, setSelectedMenu] = useState(0);
    const router = useRouter();
    const { points } = usePoint();

    return (
        <Box sx={{ display: "flex", flexDirection: "column", minHeight: "100vh" }}>
            {/* タイトル */}
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" sx={{ flexGrow: 1 }}>
                        My App
                    </Typography>
                    <Typography variant="h6">
                        <SavingsRoundedIcon /> {points}pt
                    </Typography>
                </Toolbar>
            </AppBar>

            {/* メインコンテンツ */}
            <Container sx={{ flex: 1, py: 3 }}>{children}</Container>

            {/* アラート */}
            <AlertDisplay />

            {/* フッターメニュー */}
            <Paper sx={{ position: "fixed", bottom: 0, left: 0, right: 0, height: "65px" }} elevation={3}>
                <BottomNavigation
                    showLabels
                    value={selectedMenu}
                    onChange={(event, newValue) => {
                        setSelectedMenu(newValue);
                    }}
                >
                    <BottomNavigationAction label="ホーム" icon={<HomeIcon />} onClick={() => router.push("/")} />
                    <BottomNavigationAction label="検索" icon={<SearchIcon />} onClick={() => router.push("/")} />
                    <BottomNavigationAction label="通知" icon={<NotificationsIcon />} onClick={() => router.push("/")} />
                    <BottomNavigationAction label="設定" icon={<SettingsIcon />} onClick={() => router.push("/habit-manage")} />
                </BottomNavigation>
            </Paper>
        </Box >
    );
};

export default Layout;
