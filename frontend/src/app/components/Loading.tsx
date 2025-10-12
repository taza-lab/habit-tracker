import { Box } from "@mui/material";
import CircularProgress from '@mui/material/CircularProgress';

// Layout.tsx に定義されているヘッダーとフッターの高さ
const HEADER_HEIGHT = '75px';
const FOOTER_HEIGHT = '70px';

const Loading = () => (
    <Box
      sx={{
        height: `calc(100vh - ${HEADER_HEIGHT} - ${FOOTER_HEIGHT})`,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
      }}
    >
      <CircularProgress size={60} />
    </Box>
);

export default Loading;
