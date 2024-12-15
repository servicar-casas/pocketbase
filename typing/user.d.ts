export interface User {
    id: string;
    cpf: string;
    name: string;
    phone: string;
    created_at: Date;
    updated_at: Date;
}

export interface AuthUser {
    record: User;
    token: string;
}

export interface Feedback {
    id: number;
    user: User;
    message: string;
    file_urls: string[];
    created_at: Date;
    updated_at: Date;
}
