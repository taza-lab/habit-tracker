"use client";

import { useRouter } from 'next/navigation';
import { useCallback } from 'react';
import { AuthenticationError } from '@/lib/api';

/**
 * 認証が必要なAPI呼び出しと認証エラーハンドリングを共通化するカスタムフック
 */
export function useAuthApi() {
  const router = useRouter();

  // Promiseを返す任意の関数を型引数として受け取る
  const handleAuthApiCall = useCallback(async <T, A extends any[]>(
    apiFunction: (...args: A) => Promise<T>,
    ...args: A
  ): Promise<T | null> => {
    try {
      return await apiFunction(...args);
    } catch (error) {
      if (error instanceof AuthenticationError) {
        // 認証エラーの場合はログイン画面にリダイレクト
        localStorage.removeItem('jwt_token');
        localStorage.removeItem('username'); 
        router.push('/login');
        return null;
      }
      throw error;
    }
  }, [router]);

  return { handleAuthApiCall };
}
