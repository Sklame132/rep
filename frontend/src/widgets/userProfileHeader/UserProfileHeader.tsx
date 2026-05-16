import "./UserProfileHeader.scss";
import { UserAvatar } from "../../shared/ui/userAvatar/UserAvatar";
import City from "../../shared/assets/ City.svg?react";
import { NavButton } from "../../shared/ui/NavButton";
import { useNavigate } from "react-router";

const ProfileNavButtonClass = "profile__nav"

export const UserProfileHeader = ({ username, city, created_at, activeEl, setActiveEl }: {
    username: string,
    city?: string,
    created_at: Date | undefined,
    activeEl: string,
    setActiveEl: (element: string) => void
}) => {
    const navigate = useNavigate();
    const date = new Date(created_at ?? "");

    return <div className="profile__header">
        <div className="profile__header__container">
            <UserAvatar />
            <div className="profile__info">
                <div className="profile__info__left-side">
                    <p>{ username }</p>
                    <div className="profile__info__location">
                        <City className="profile__info__location-image"/>
                        <p>{ city ? city : "Воронеж" }</p>
                    </div>
                </div>
                <div className="profile__info__right-side">
                    <p>Дата регистрации: { date.toLocaleDateString() }</p>
                </div>
            </div>
        </div>

        <div className="profile__nav-bar__container">
            <ul className="profile__nav-bar">
                <NavButton className={ProfileNavButtonClass} isActive={activeEl === "История" ? true : false} onClick={ () => { setActiveEl("История") } }>
                    История
                </NavButton>

                <NavButton className={ProfileNavButtonClass} isActive={activeEl === "Информация" ? true : false} onClick={() => { setActiveEl("Информация") }}>
                    Информация
                </NavButton>
            </ul>
            <NavButton className={[ProfileNavButtonClass, "-exit"].join('')} onClick={async () => {
                await fetch("http://localhost:5050/api/v1/logout", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    credentials: "include",
                });
                await navigate("/login");
            }}>
                Выход
            </NavButton>
        </div>
    </div>
}