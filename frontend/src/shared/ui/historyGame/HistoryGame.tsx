import "./HistoryGame.scss"
import Timer from "../../assets/Timer.svg?react"

export const HistoryGame = ({mode, username, color, moves, created_at, result}: {
    mode: string,
    username: string,
    color: string,
    moves: number,
    created_at: Date,
    result: string,
}) => {
    const date = new Date(created_at)
    let resultClass = `profile__history-game__result`
    if (result === "Победа") {
        resultClass += "-win"
    } else if (result === "Поражение") {
        resultClass += "-defeat"
    } else {
        resultClass += "-draw"
    }

    return <li className="profile__history-game__item">
        <div className="profile__history-game__username">
            <div className="profile__history-game__mode">
                <div className="profile__history-game__mode-timer">
                    <Timer className="profile__history-game__mode-timer-image"/>
                </div>
                <p>{mode}</p>
            </div>
            <p>{username}</p>
        </div>
        <div className="profile__history-game__color">{color}</div>
        <div className="profile__history-game__moves">{moves}</div>
        <div className="profile__history-game__date">{ date.toLocaleString().replace(",", '') }</div>
        <div className={resultClass}>{result}</div>
    </li>
}