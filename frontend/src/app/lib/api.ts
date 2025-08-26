
export function getAuthHeaders(): Record<string, string> {
    // localStorageからトークンを取得
    const token = localStorage.getItem('jwt_token');

    // トークンが存在する場合のみ、ヘッダーを準備
    const headers: HeadersInit = {
        'Content-Type': 'application/json',
    };

    if (token) {
        // BearerトークンをAuthorizationヘッダーに追加
        headers['Authorization'] = `Bearer ${token}`;
    }

    return headers;
}