import { createContext, useState, type ReactNode } from "react";

export const UserContext = createContext<{username: string | null, setUsername: any}>({username: null, setUsername: null});

export const AppProvider = ({ children }: { children: ReactNode}) => {
    const [username, setUsername] = useState<string | null>(null);
    const value = {username, setUsername};
    return <UserContext.Provider value={value}>
        {children}
    </UserContext.Provider>
}