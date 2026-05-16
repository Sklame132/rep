import { NavBar } from "../../shared/ui/navBar/NavBar"
import "./Header.scss";
import { useLocation } from "react-router"

export const Header = () => {
        const path = useLocation().pathname;
        const isProfile = path.includes("/profile");

        return  <header className={["header", isProfile && "wrapper"].join(" ")}>
                <div className="header__top">
                        <div className="header__date"></div>

                        {isProfile ? <NavBar isWrapper /> : <NavBar />}

                        <div className="header__location"></div>
                </div>
                {isProfile && <div className="header__bottom">
                        <h1 className="header__title">Личный кабинет</h1>
                </div>}
        </header>
}