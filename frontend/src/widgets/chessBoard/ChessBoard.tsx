import type { Chess, Square, PieceSymbol, Color, Move } from "chess.js";
import { useState, useEffect, useRef } from "react";
import { MOVE } from "../../pages/chessGame/ChessGame";

import { Piece } from "../../shared/ui/Piece"

export type Field = {
    square: Square;
    type: PieceSymbol;
    color: Color;
} | null

export const ChessBoard = ({ chess, board, setBoard, history, socket }: {
    chess: Chess;
    board: Field[][];
    setBoard: any;
    history: Move[] | null;
    socket: WebSocket;
}) => {
    const [from, setFrom] = useState<Square | null>(null);
    const availableMoves = useRef<string[]>([]);
    const isDragging = useRef<boolean>(false);
    const isClicked = useRef<boolean>(false);

    useEffect(() => {
        if (availableMoves.current) {
            for (const square of availableMoves.current) {
                document.getElementById(square)?.classList.remove("available-move", "available-eat");
            }
        }
        if (from) {
            availableMoves.current = chess.moves({square: from, verbose: true}).map((square) => square.to);
            for (const square of availableMoves.current) {
                if (!chess.get(square as Square)?.type) {
                    document.getElementById(square)?.classList.add("available-move");
                    continue;
                }
                document.getElementById(square)?.classList.add("available-eat");
            }
        } else {
            isClicked.current = false;
            availableMoves.current = [];
        }
    }, [from])

    function mousedownHandler(squareRepresentation: Square) {
        isDragging.current = true;
        if (!from) {
            setFrom(squareRepresentation);
        } else if (isClicked.current) {
            if (availableMoves.current?.includes(squareRepresentation)) {
                socket.send(JSON.stringify({
                    type: MOVE,
                    payload: {
                        move: {
                            from,
                            to: squareRepresentation
                        }
                    }
                }))
                setFrom(null)
                chess.move({
                    from,
                    to: squareRepresentation
                });
                history = updateLastMoveHighlights(chess);
                setBoard(chess.board());
            } else {
                setFrom(squareRepresentation);
                isClicked.current = false;
            }
        }
    }

    function mousemoveHandler(event: React.MouseEvent) {
        if (isDragging.current && from) {
            const piece: SVGSVGElement = document.getElementById(from)?.children[0] as SVGSVGElement;
            piece.style.left = `${event.pageX - piece.clientWidth / 2}`;
            piece.style.top = `${event.pageY - piece.clientHeight / 2}`;
        }
    }

    function mouseupHandler(squareRepresentation: Square) {
        if (from) {
            if (isDragging.current && availableMoves.current?.includes(squareRepresentation)) {
                socket.send(JSON.stringify({
                type: MOVE,
                payload: {
                    move: {
                        from,
                        to: squareRepresentation
                    }
                }
                }))
                setFrom(null);
                chess.move({
                    from,
                    to: squareRepresentation
                });
                history = updateLastMoveHighlights(chess)
                setBoard(chess.board());
            } else {
                isClicked.current ? setFrom(null) : isClicked.current = true;
            }
                isDragging.current = false;
                returnPositionPiece(from);
        }
    }
    
    function returnPositionPiece(from: Square) {
        const field = document.getElementById(from) as HTMLDivElement;
        const piece = field.children[0] as SVGSVGElement;
        const pieceCopy = field.children[1] as SVGSVGElement;
        piece.style.left = pieceCopy.style.left;
        piece.style.top = pieceCopy.style.top;
    }

    return <div className="chess__board" id="board">
        {board.map((row, i) => {
            return <div key={i} className="flex">
                {row.map((field, j) => {
                    const squareRepresentation = String.fromCharCode(97 + (j % 8)) + "" + (8 - i) as Square;
                    return <div
                        onMouseDown={() => mousedownHandler(squareRepresentation)}
                        onMouseMove={(event) => mousemoveHandler(event)}
                        onMouseUp={() => mouseupHandler(squareRepresentation)}
                        key={j}
                        id={squareRepresentation}
                        className={["chess__field", 
                            (i + j) % 2 === 0 ? "b" : "w",
                        ].join(' ')}
                    >
                        <Piece field={field} className={(isDragging.current && from === squareRepresentation) ? "dragging" : ""}/>
                        {(isDragging.current && squareRepresentation === from) && <Piece field={field} className={"copy"} />}
                    </div>
                })}
            </div>
        })}
    </div>
}

export function updateLastMoveHighlights(chess: Chess): Move[] {
    const history = chess.history({ verbose: true })
    if (history) {
        if (history.length > 1) {
            document.getElementById(history[history.length - 2].from)?.classList.remove("undo");
            document.getElementById(history[history.length - 2].to)?.classList.remove("undo");
        }
            document.getElementById(history[history.length - 1].from)?.classList.add("undo");
            document.getElementById(history[history.length - 1].to)?.classList.add("undo");
    }
    return history
}