import { User } from '@/types/user';
import { apiRequest } from '@/lib/api';

interface UserApiData {
    token: string;
    user: {
        id: string;
        username: string;
        password: string;
        points: number;
    }
}

export async function login(username: string, password: string): Promise<{ token: string, user: User }> {
    const apiData = await apiRequest<UserApiData>(
        `login`,
        {
            method: 'POST',
            body: JSON.stringify({ username, password })
        }
    );

    return {
        token: apiData.token,
        user: {
            id: apiData.user.id,
            username: apiData.user.username,
            points: apiData.user.points,
        }
    };
}

export async function signup(username: string, password: string, confirmPassword: string): Promise<void> {
    await apiRequest<UserApiData>(
        `signup`,
        {
            method: 'POST',
            body: JSON.stringify({ username, password, confirm_password: confirmPassword })
        }
    );

    return;
}
