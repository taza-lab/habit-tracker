'use client';

import { Alert, Box, Zoom } from '@mui/material';
import { useAlert } from '@/context/AlertContext';
import AutoAwesomeRoundedIcon from '@mui/icons-material/AutoAwesomeRounded';
import InfoRoundedIcon from '@mui/icons-material/InfoRounded';
import WarningRoundedIcon from '@mui/icons-material/WarningRounded';
import NewReleasesRoundedIcon from '@mui/icons-material/NewReleasesRounded';
import SentimentNeutralRoundedIcon from '@mui/icons-material/SentimentNeutralRounded';

export function AlertDisplay() {
    const { alert } = useAlert();

    if (!alert.message) return null;

    // severityに応じて表示するアイコンを定義
    let iconComponent = null;
    switch (alert.severity) {
        case 'success':
            iconComponent = <AutoAwesomeRoundedIcon />;
            break;
        case 'info':
            iconComponent = <InfoRoundedIcon />;
            break;
        case 'warning':
            iconComponent = <WarningRoundedIcon />;
            break;
        case 'error':
            iconComponent = <NewReleasesRoundedIcon />;
            break;
        default:
            iconComponent = <SentimentNeutralRoundedIcon />;
            break;
    }

    return (
        <Box
            sx={{
                position: 'fixed',
                bottom: 80, // フッターの上に表示
                zIndex: 1300,
                width: '100%',
                maxWidth: 600,
                px: 2,
            }}
        >
            <Zoom in={alert.show} timeout={300}>
                <Alert
                    icon={iconComponent}
                    severity={alert.severity || 'success'}
                    sx={{
                        fontSize: '1.2rem'
                    }}
                >
                    {alert.message}
                </Alert>
            </Zoom>
        </Box>
    );
}