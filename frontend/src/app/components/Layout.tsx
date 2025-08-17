import { AppBar, Toolbar, Typography, Box, Container, BottomNavigation, BottomNavigationAction, Paper } from "@mui/material";
import React, { ReactNode } from "react";
import { useState } from "react";
import HomeIcon from "@mui/icons-material/Home";
import SearchIcon from "@mui/icons-material/Search";
import NotificationsIcon from "@mui/icons-material/Notifications";
import SettingsIcon from "@mui/icons-material/Settings";

type LayoutProps = {
    children: ReactNode;
};

const Layout = ({ children }: LayoutProps) => {
    const [value, setValue] = useState(0);

    return (
        <Box sx={{ display: "flex", flexDirection: "column", minHeight: "100vh" }}>
            {/* タイトル */}
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" sx={{ flexGrow: 1 }}>
                        My App
                    </Typography>
                </Toolbar>
            </AppBar>

            {/* メインコンテンツ */}
            <Container sx={{ flex: 1, py: 3 }}>{children}</Container>

            {/* フッターメニュー */}
            <Paper sx={{ position: "fixed", bottom: 0, left: 0, right: 0, height: "65px" }} elevation={3}>
                <BottomNavigation
                    showLabels
                    value={value}
                    onChange={(event, newValue) => {
                        setValue(newValue);
                    }}
                >
                    <BottomNavigationAction label="ホーム" icon={<HomeIcon />} />
                    <BottomNavigationAction label="検索" icon={<SearchIcon />} />
                    <BottomNavigationAction label="通知" icon={<NotificationsIcon />} />
                    <BottomNavigationAction label="設定" icon={<SettingsIcon />} />
                </BottomNavigation>
            </Paper>
        </Box >
    );
};

export default Layout;
