import { useEffect, useState, useRef } from "react"
import { useSocket } from "../../shared/lib/useSocket";
import { ChessBoard, updateLastMoveHighlights } from "../../widgets/chessBoard/ChessBoard";
import { Button } from "../../shared/ui/Button";
import { Chess, type Move } from "chess.js";
import { useNavigate } from "react-router";
import { useFetch, type FetchState } from "../../shared/lib/useFetch";
import type { User } from "../../shared/models/User";

export const INIT_GAME = "init_game";
export const MOVE = "move";
export const GAME_OVER = "game_over";

type Metadata = {
    blackPlayer: string;
    whitePlayer: string;
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

    const [chess, _setChess] = useState<Chess>(new Chess());
    const [board, setBoard] = useState(chess?.board());
    const history = useRef<Move[] | null>(null);
    const [started, setStarted] = useState<boolean>(false);
    const [gameMetaData, setGameMetaData] = useState<Metadata | null>(null)
    
    useEffect(() => {
        if (!socket) {
            return;
        }
        socket.onmessage = (event) => {
            const message = JSON.parse(event.data);

            switch (message.type) {
                case INIT_GAME:
                    console.log("Game initialized");
                    setBoard(chess.board());
                    setStarted(true);
                    //navigate(`/game/${message.payload.gameId}`)
                    setGameMetaData({
                        blackPlayer: message.payload.blackPlayer,
                        whitePlayer: message.payload.whitePlayer
                    })
                    console.log(setGameMetaData)
                    break;
                case MOVE:
                    console.log('asd')
                    const move = message.payload;
                    chess.move(move);
                    history.current = updateLastMoveHighlights(chess);
                    setBoard(chess.board());
                    console.log("Move made");  
                    break;
                
                case GAME_OVER:
                    console.log("Game over");
                    break;
            }
        }
    }, [chess, socket]);
    
    if (!socket) return <div>Подключение...</div>
    
    return <div className="chess">
        <div>{gameMetaData?.whitePlayer} vs {gameMetaData?.blackPlayer}</div>
        <div className="chess__container">
            <ChessBoard chess={chess} board={board} setBoard={setBoard} history={history.current} socket={socket}/>
        </div>
        {!started && <Button onClick={() => {
            socket.send(JSON.stringify({
                type: INIT_GAME,
            }))
        }}>
            Play
        </Button>}
    </div>
}