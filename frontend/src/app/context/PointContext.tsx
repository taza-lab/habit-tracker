'use client';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';

type PointContextType = {
    points: number;
    setPoints: (points: number) => void;
    addPoints: (amount: number) => void;
    subPoints: (amount: number) => void;
};

const PointContext = createContext<PointContextType | undefined>(undefined);
const POINT_STORAGE_KEY = 'user_points';

export function PointProvider({ children }: { children: ReactNode }) {
    const [points, setPoints] = useState(0);

    // コンポーネントがマウントされた後（クライアントサイドでのみ）にlocalStorageからデータを読み込む
    useEffect(() => {
        const storedPoints = localStorage.getItem(POINT_STORAGE_KEY);
        if (storedPoints) {
            setPoints(parseInt(storedPoints, 10));
        }
    }, []); 
    
    // pointsが変更されるたびにlocalStorageに保存
    useEffect(() => {
        localStorage.setItem(POINT_STORAGE_KEY, points.toString());
    }, [points]);

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
