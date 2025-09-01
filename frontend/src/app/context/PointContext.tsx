'use client';

import { createContext, useContext, useState, ReactNode } from 'react';

type PointContextType = {
    points: number;
    setPoints: (points: number) => void;
    addPoints: (amount: number) => void;
    subPoints: (amount: number) => void;
};

const PointContext = createContext<PointContextType | undefined>(undefined);

export function PointProvider({ children }: { children: ReactNode }) {
    const [points, setPoints] = useState(0);

    const addPoints = (amount: number) => {
        setPoints((prev) => prev + amount);
    };

    const subPoints = (amount: number) => {
        setPoints((prev) => prev - amount);
    };

    return (
        <PointContext.Provider value={{ points, setPoints, addPoints, subPoints }}>
            {children}
        </PointContext.Provider>
    );
}

export function usePoint() {
    const context = useContext(PointContext);
    if (context === undefined) {
        throw new Error('usePoint must be used within a PointProvider');
    }
    return context;
}