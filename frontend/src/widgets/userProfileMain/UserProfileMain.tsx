import "./UserProfileMain.scss";
import { HistoryGame } from "../../shared/ui/historyGame/HistoryGame";
import type { Game } from "../../shared/models/Game";
import { useEffect, useState } from "react";
import { Button } from "../../shared/ui/Button";
import type { User } from "../../shared/models/User";

export const UserProfileMain = ({ activeEl, user, games, updateGameOffset }: {
    activeEl: string,
    user: User | undefined,
    games?: Game[] | undefined,
    updateGameOffset: () => void,
}) => {
    const [gameItems, setGameItems] = useState<Game[]>([]);

    useEffect(() => {
        if (games) {
            if (gameItems[0]?.id !== games[0]?.id) {
                setGameItems([...gameItems, ...games]);
            }
        }
    }, [games])

    switch (activeEl) {
        case "История":
            return <div className="profile__main">
                <div className="profile__history-legend"><p>История партий</p></div>

                <div className="profile__history">
                    <ul className="profile__history-column__list">
                        <li className="profile__history-column__item"><p>Соперник</p></li>
                        <li className="profile__history-column__item"><p>Ваш цвет</p></li>
                        <li className="profile__history-column__item"><p>Ходы</p></li>
                        <li className="profile__history-column__item"><p>Дата</p></li>
                        <li className="profile__history-column__item"><p>Результат</p></li>
                    </ul>
                    <ul className="profile__history-game__list">
                        {gameItems?.map((game, index) => {
                            const isWhite = (game.player_w === user?.username) ? true : false;
                            let gameResult
                            if ((game.result === "win_w" && isWhite) || (game.result === "win_b" && !isWhite)) {
                                gameResult = "Победа"
                            } else if (game.result === "draw") {
                                gameResult = "Ничья"
                            } else {
                                gameResult = "Поражение"
                            }

                            const history = JSON.parse(game.history);

                            return <HistoryGame key={index}
                                mode={game.mode}
                                username={isWhite ? game.player_b : game.player_w}
                                color={isWhite ? "Белый" : "Черный"}
                                moves={history?.length}
                                created_at={game.created_at}
                                result={gameResult}
                            />
                        })}
                    </ul>
                </div>

                <Button 
                className="profile__history-game__button"
                onClick={updateGameOffset}>
                    Показать еще
                </Button>
            </div>
        case "Информация":
            return <div className="profile__main">
                <div className="profile__history-legend"><p>Информация</p></div>

                <div className="profile__full-info">
                    <div className="profile__full-name">
                        {user?.first_name && <p className="profile__first-name">Имя: {user.first_name}</p>}
                        {user?.last_name &&<p className="profile__last-name">Фамилия: {user.last_name}</p>}
                    </div>
                    {user?.address && <div className="profile__full-location">Адрес: {user.address}</div>}
                    <div className="profile__contacts">
                        {user?.email && <div className="profile__email">Почта: {user.email}</div>}
                        {user?.phone_number && <div className="profile__phone-number">Номер телефона: {user.phone_number}</div>}
                    </div>
                </div>
            </div>

    }
}