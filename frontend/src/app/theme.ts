import { createTheme } from '@mui/material/styles';

const theme = createTheme({
    palette: {
        primary: { main: '#4caf50' }, // 緑系
        secondary: { main: '#ff9800' }, // オレンジ
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
