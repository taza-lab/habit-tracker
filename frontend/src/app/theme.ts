import { createTheme } from '@mui/material/styles';

const theme = createTheme({
    palette: {
        primary: { main: '#4cafafff' },
        secondary: { main: '#ff9800' },
        text: {
            primary: '#171717', // 主要な文字色
            secondary: '#757575',
        },
    },
    shape: {
        borderRadius: 12, // 角丸を全体的に
    },
    typography: {
        fontFamily: 'Roboto, Arial, sans-serif',
        button: { textTransform: 'none' }, // ボタンの大文字化OFF
    },
});

export default theme;
