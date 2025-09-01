'use client';

import { createContext, useContext, useState, ReactNode } from 'react';

type AlertContextType = {
    alert: {
        show: boolean;
        message: string | null;
        severity: 'success' | 'info' | 'warning' | 'error' | null;
    };
    showAlert: (message: string, severity: 'success' | 'info' | 'warning' | 'error') => void;
};

const AlertContext = createContext<AlertContextType | undefined>(undefined);

export function AlertProvider({ children }: { children: ReactNode }) {
    const [alert, setAlert] = useState<AlertContextType['alert']>({
        show: false,
        message: null,
        severity: null,
    });

    const showAlert = (message: string, severity: 'success' | 'info' | 'warning' | 'error') => {
        setAlert({ show: true, message, severity });

        setTimeout(() => {
            setAlert((prev) => ({ ...prev, show: false }));
        }, 3000);

        setTimeout(() => {
            setAlert({ show: false, message: null, severity: null });
        }, 3500);
    };

    return (
        <AlertContext.Provider value={{ alert, showAlert }}>
            {children}
        </AlertContext.Provider>
    );
}

export function useAlert() {
    const context = useContext(AlertContext);
    if (context === undefined) {
        throw new Error('useAlert must be used within an AlertProvider');
    }
    return context;
}