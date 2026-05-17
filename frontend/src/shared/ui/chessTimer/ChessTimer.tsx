import Timer from "../../assets/Timer.svg?react"
import "./ChessTimer.scss"

export const ChessTimer = ({username, minutes, seconds, isMustRotate}: {
    username: string | undefined,
    minutes: number,
    seconds: number,
    isMustRotate?: boolean,
}) => {
    return <div className={["chess__timer__container", isMustRotate && "rotate"].join(" ")}>
        <div className="chess__timer">
            <div className="chess__timer__image__wrapper">
                <Timer/>
            </div>
            <p className="chess__timer__text">{minutes} {":"} {String(seconds).padStart(2,'0')}</p>
        </div>
        <div className="chess__timer__username">
            {username}
        </div>
    </div>
}