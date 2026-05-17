import { useEffect, useState } from "react"
import { useSocket } from "../../shared/lib/useSocket";
import { ChessBoard, updateLastMoveHighlights } from "../../widgets/chessBoard/ChessBoard";
import { Button } from "../../shared/ui/Button";
import { Chess, Move } from "chess.js";
import { useNavigate } from "react-router";
import { useFetch, type FetchState } from "../../shared/lib/useFetch";
import type { User } from "../../shared/models/User";
import ReactConfetti from "react-confetti";
import { ChessTimer } from "../../shared/ui/chessTimer/ChessTimer"; 

export const INIT_GAME = "init_game";
export const MOVE = "move";
export const GAME_OVER = "game_over";

type Metadata = {
    playerWhite: string;
    playerBlack: string;
}

export const ChessGame = () => {
    const navigate = useNavigate();
    const userState: FetchState<User> = useFetch<User>(`/user`, {
        credentials: "include",
    });
    if (userState.error) {
        navigate("/login")
    }

    const socket = useSocket();
    const [chess, setChess] = useState<Chess>(new Chess());
    const [board, setBoard] = useState(chess?.board());
    const [currentColor, setCurrentColor] = useState<"w" | "b" | null>(null);
    const [moves, setMoves] = useState<Move[]>([]);
    
    const [isStarted, setIsStarted] = useState<Boolean | "waiting">(false);
    const [isGameOver, setIsGameOver] = useState<boolean>(false);

    const [timerWhite, setTimerWhite] = useState<number>(0);
    const [timerBlack, setTimerBlack] = useState<number>(0);
    const [stopTimerWhite, setStopTimerWhite] = useState<Boolean>(true);
    const [stopTimerBlack, setStopTimerBlack] = useState<Boolean>(true);

    const [gameMetaData, setGameMetaData] = useState<Metadata | null>(null);

    const [playerWon, setPlayerWon] = useState<string>("")
    const [playerWonColor, setPlayerWonColor] = useState("")

    const initGame = (timer: number) => {
        setIsStarted("waiting");
        socket?.send(JSON.stringify({
            type: INIT_GAME,
            name: userState.data?.username,
            timer: timer,
        }))
    }

    const playerWhiteSeconds = Math.floor(timerWhite/1000);
    const playerBlackSeconds = Math.floor(timerBlack/1000);

    const playerWhiteMinutes = Math.floor(playerWhiteSeconds/60);
    const playerBlackMinutes = Math.floor(playerBlackSeconds/60);

    const playerWhiteRemainingSeconds = playerWhiteSeconds%60;
    const playerBlackRemainingSeconds = playerBlackSeconds%60;
    
    useEffect(() => {
        if (!socket) {
            return;
        }

        let remainingTimeWhite = timerWhite;
        let remainingTimeBlack = timerBlack;

        if (chess.isGameOver() && !chess.isDraw()) {
            setPlayerWonColor(chess.turn() === "w" ? "b" : "w");
            console.log(playerWon)
        }

        const interval: any = setInterval(() => {
            if (!stopTimerWhite) {
                if(timerWhite === 0){
                socket?.send(JSON.stringify({
                    type:GAME_OVER,
                    payload:{
                        result: chess.isDraw() ? "draw" : chess.turn() === "w" ? "win_b" : "win_w"
                    }
                }))  
                return clearInterval(interval)
            }
            remainingTimeWhite -= 1000;
            setTimerWhite(remainingTimeWhite);
            }

            if (!stopTimerBlack) {
                if(timerBlack === 0){
                    console.log("game end?")
                    socket.send(JSON.stringify({
                        type:GAME_OVER,
                        payload:{
                            result: chess.isDraw() ? "draw" : chess.turn() === "w" ? "win_b" : "win_w"
                        }
                    }))  
                    return clearInterval(interval)
                }
                remainingTimeBlack -= 1000;
                setTimerBlack(remainingTimeBlack);
            }
        }, 1000);

        socket.onmessage = (event) => {
        const message = JSON.parse(event.data);

        switch (message.type) {
            
            case INIT_GAME:
                setChess(new Chess());
                setBoard(chess.board());
                setIsGameOver(false);
                setIsStarted(true);
                setCurrentColor(message.payload);
                setTimerWhite(message.timer);
                setTimerBlack(message.timer);
                setGameMetaData({
                    playerWhite: message.white,
                    playerBlack: message.black
                })
                const gameStart = new Audio("../../shared/assets/audio/GameStartAudio.mp4");
                gameStart.play();
                break;

            case MOVE:
                const move = message.payload;
                chess.move(move);

                setBoard(chess.board());
                setMoves(updateLastMoveHighlights(chess));

                const gameMove = new Audio('../../shared/assets/audio/Move.mp4')
                gameMove.play();

                console.log(message)
                if (message.color === "w") {
                    setStopTimerWhite(false);
                    setStopTimerBlack(true);
                }
                if (message.color === "b") {
                    setStopTimerWhite(true);
                    setStopTimerBlack(false);
                }
                break;
            
            case GAME_OVER:
                const gameStatus = message.payload
                setIsStarted(false);
                setIsGameOver(true);
                console.log(message);

                setPlayerWon(gameStatus.result)
                setPlayerWonColor(gameStatus.color);

                setStopTimerWhite(true);
                setStopTimerBlack(true);
                break;
            }
        }
    return () => clearInterval(interval);

    }, [socket, stopTimerWhite, stopTimerBlack, timerWhite, timerBlack]);
    
    if (!socket) return <div>Подключение...</div>
    
    return <div className="chess">
        {
            (playerWonColor === currentColor) ? <ReactConfetti/> : null
        }


        <div className="chess__container">
            <ChessBoard chess={chess} board={board} setBoard={setBoard} setMoves={setMoves} setStopTimerWhite={setStopTimerWhite} setStopTimerBlack={setStopTimerBlack} setPlayerWonColor={setPlayerWonColor} socket={socket} currentColor={currentColor}/>
        </div>


        <div className={["chess__side-bar__container", currentColor === "b" && "rotate"].join(" ")}>
            <ChessTimer username={gameMetaData?.playerBlack} minutes={playerBlackMinutes} seconds={playerBlackRemainingSeconds} isMustRotate={currentColor === "b"}/>
            <div className="chess__side-bar">
                {
                isStarted === "waiting" ? <div className="chess__search">
                    Поиск...
                </div> :
                isStarted ? <div className="chess__moves__container">
                    {moves.map((move, index) => {
                        return <div key={index} className={["chess__moves__move", currentColor === "b" && "rotate"].join(" ")}>{`${index + 1}. ${move.from} ${move.to}`}</div>
                    } )}
                </div>  : 
                isGameOver ? <div className={["chess__winner", playerWonColor, currentColor === "b" && "rotate"].join(" ")}>{
                    playerWon === "win_w" ? "Победа белых" :
                    playerWon === "win_b" ? "Победа черных" : "Ничья"
                    }</div> :
                <div className="chess__modes__container">
                    <Button className="chess__modes__mode" onClick={() => initGame(5)}>
                        5 мин
                    </Button>
                    <Button className="chess__modes__mode" onClick={() => initGame(10)}>
                        10 мин
                    </Button>
                </div>
               }

            </div>
             <ChessTimer username={gameMetaData?.playerWhite} minutes={playerWhiteMinutes} seconds={playerWhiteRemainingSeconds} isMustRotate={currentColor === "b"}/>
        </div>
    </div>
}