const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

// 認証エラーを表すカスタムエラー
export class AuthenticationError extends Error {
  constructor(message = '認証エラーが発生しました。') {
    super(message);
    this.name = 'AuthenticationError';
  }
}

export function getAuthHeaders(): Record<string, string> {
    // localStorageからトークンを取得
    const token = localStorage.getItem('jwt_token');

    const headers: HeadersInit = {
        'Content-Type': 'application/json',
    };

    if (token) {
        // BearerトークンをAuthorizationヘッダーに追加
        headers['Authorization'] = `Bearer ${token}`;
    }

    return headers;
}

/**
 * APIリクエストを共通化する汎用関数
 */
export async function apiRequest<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  // 認証ヘッダーを自動で付加
  const headers = {
    ...getAuthHeaders(),
    ...options.headers, // 呼び出し元からのヘッダーをマージ
  };

  // 認証が必要なリクエストで認証ヘッダーが存在しない場合、APIコール前にエラーを投げる
   if (endpoint.includes('auth/') && !headers.hasOwnProperty('Authorization')) {
    throw new AuthenticationError('認証情報を取得できませんでした。');
  }

  const url = `${BASE_URL}/${endpoint}`;

  const res = await fetch(url, { ...options, headers });

  // 認証エラーを検出
  if (res.status === 401 || res.status === 400) {
    const errorText = await res.json();
    throw new AuthenticationError(errorText.error ?? 'トークンの有効期限が切れました。');
  }

  // その他のネットワークエラーやサーバーエラーを検出
  if (!res.ok) {
    console.log(`APIリクエストに失敗しました: ${res.status} ${res.statusText}`);

    const errorText = await res.json();
    throw new Error(errorText.error ?? 'APIリクエストに失敗しました');
  }

  // レスポンスのJSONデータを返す
  return (await res.json()) as T;
}
