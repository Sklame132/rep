export type User = {
    id?: string;
    username: string;
    password: string;
    first_name: string;
    last_name: string;
    address?: string;
    email?: string;
    phone_number?: string;
    created_at?: Date;
    updated_at?: Date;
    rating?: number;
    image_url?: string;
}