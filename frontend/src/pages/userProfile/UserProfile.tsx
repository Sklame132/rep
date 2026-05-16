import { type FetchState, useFetch } from "../../shared/lib/useFetch";
import type { User } from "../../shared/models/User";
import type { Game } from "../../shared/models/Game";
import { UserProfileHeader } from "../../widgets/userProfileHeader/UserProfileHeader";
import { useState } from "react";
import { UserProfileMain } from "../../widgets/userProfileMain/UserProfileMain";
import { useNavigate } from "react-router";

const GAME_LIMIT = 3;
export const UserProfile = () => {
    const [activeEl, setActiveEl] = useState("История");
    const navigate = useNavigate();

    const userState: FetchState<User> = useFetch<User>(`/user`, {
        credentials: "include",
    });
    if (userState.error) {
        navigate("/login")
    }

    const [gameOffset, setGameOffset] = useState<number>(0);
    const updateGameOffset = () => {
        setGameOffset(gameOffset + GAME_LIMIT)
    }
    
    const gamesState: FetchState<Game[]> = useFetch<Game[]>(`/games?limit=${GAME_LIMIT}&offset=${gameOffset}&username=${userState.data?.username}`)
    
    return (
        <div className="profile">
            <div className="container">
                {userState.data && 
                <UserProfileHeader 
                    username={userState.data.username} 
                    city={userState.data.address}
                    created_at={userState.data.created_at}
                    activeEl={activeEl}
                    setActiveEl={setActiveEl}
                />}
                <UserProfileMain 
                    activeEl={activeEl}
                    user={userState.data}
                    games={gamesState.data}
                    updateGameOffset={updateGameOffset}
                />
            </div>
        </div>
    )
}