import DefaultAvatar from "../../assets/userAvatars/Default.svg?react";
import "./UserAvatar.scss"

export const UserAvatar = () => {
    return <div className="profile__avatar__wrapper">
        <div className="profile__avatar">
            <DefaultAvatar className="profile__avatar__image"/>
        </div>
    </div>
}