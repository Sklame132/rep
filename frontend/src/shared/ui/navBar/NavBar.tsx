import Ivents from "../../assets/navigation/Ivents.svg?react";
import Settings from "../../assets/navigation/Settings.svg?react";
import Contacts from "../../assets/navigation/Contacts.svg?react";
import Notifications from "../../assets/navigation/Notifications.svg?react";
import Profile from "../../assets/navigation/Profile.svg?react";
import Play from "../../assets/Flame.svg?react";
import "./NavBar.scss";
import { useNavigate } from "react-router";
import { NavButton } from "../NavButton";

interface NavBarComponent {
    component: any;
    onClick: () => void;
}

export const NavBar = ({ isWrapper }: { isWrapper?: boolean}) => {
    const navigate = useNavigate();
    const order: NavBarComponent[] = [
        { component: Ivents, onClick: () => { navigate("/ivents") }},
        { component: Settings, onClick: () => { navigate("/settings") }},
        { component: Contacts, onClick: () => { navigate("/contacts") }},
        { component: Notifications, onClick: () => {}}, 
        { component: Profile, onClick: () => { navigate("/profile") }},
        { component: Play, onClick: () => { navigate("/game") }}
    ]

    return <div className={["nav-bar", isWrapper && "wrapper"].join(" ")}>
            <nav className="nav-bar__container">
            <div className="nav-bar__logo" onClick={() => { navigate("/") }}>REP</div>
            <ul onClick={ () => {
            } }className="nav-bar__list">
                {order.map((Item, index) => {
                    return <NavButton key={index} className="nav-bar" onClick={Item.onClick}>
                        <Item.component />
                    </NavButton>
                })}
            </ul>
        </nav>
    </div>
}