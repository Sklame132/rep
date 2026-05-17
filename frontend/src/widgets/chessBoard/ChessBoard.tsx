import type { Chess, Square, PieceSymbol, Color, Move } from "chess.js";
import { useState, useEffect, useRef } from "react";
import { MOVE } from "../../pages/chessGame/ChessGame";
import "./ChessBoard.scss"

import { Piece } from "../../shared/ui/piece/Piece"

export type Field = {
    square: Square;
    type: PieceSymbol;
    color: Color;
} | null

export const ChessBoard = ({ chess, board, setBoard, setMoves, setStopTimerWhite, setStopTimerBlack, setPlayerWonColor, socket, currentColor }: {
    chess: Chess;
    board: Field[][];
    setBoard: any;
    setMoves: any;
    setStopTimerWhite: any;
    setStopTimerBlack: any;
    setPlayerWonColor: any;
    socket: WebSocket;
    currentColor: "w" | "b" | null;
}) => {
    const [from, setFrom] = useState<Square | null>(null);
    const availableMoves = useRef<string[]>([]);
    const isDragging = useRef<boolean>(false);
    const isClicked = useRef<boolean>(false);
    const whiteKingHTML = document.getElementById("k_w");
    const blackKingHTML = document.getElementById("k_b");
    const boardHTML = document.getElementById("board") as HTMLDivElement;

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

    useEffect(() => {
        if (chess.isCheckmate()) {
            const checkMate = new Audio("../../shared/assets/audio/Checkmate.mp4");
            checkMate.play();

            if (chess.turn() === "b") {
                blackKingHTML?.classList.add("checkmate");
            } else {
                whiteKingHTML?.classList.add("checkmate");
            }
            return
        }
        if (chess.isCheck()) {
            const check = new Audio("../../shared/assets/audio/Check.mp4");
            check.play();

            if (chess.turn() === "b") {
                blackKingHTML?.classList.add("check");
            } else {
                whiteKingHTML?.classList.add("check");
            }
            return
        }

        whiteKingHTML?.classList.remove("check", "checkmate")
        blackKingHTML?.classList.remove("check", "checkmate")
    }, [chess.isCheck(), chess.isCheckmate()])

    function mousedownHandler(squareRepresentation: Square, pieceColor: Color | undefined) {
        if (currentColor && (currentColor !== chess.turn() || currentColor !== pieceColor)) return;
        isDragging.current = true;
        if (!from) {
            setFrom(squareRepresentation);
        } else if (isClicked.current) {
            if (availableMoves.current?.includes(squareRepresentation)) {
                chess.move({
                    from,
                    to: squareRepresentation
                });

                socket.send(JSON.stringify({
                    type: MOVE,
                    payload: {
                        move: {
                            from,
                            to: squareRepresentation
                        },
                        color:chess.turn()
                    }
                }))
                if(chess.turn() === "b"){
                    setStopTimerBlack(false)        
                    setStopTimerWhite(true)
                } else {
                    setStopTimerWhite(false)        
                    setStopTimerBlack(true)
                }

                const gameMove = new Audio("../../shared/assets/audio/Move.mp4")
                gameMove.play();

                setFrom(null);
                setBoard(chess.board());
                setMoves(updateLastMoveHighlights(chess));
            } else {
                setFrom(squareRepresentation);
                isClicked.current = false;
            }
        }
    }

    function mousemoveHandler(event: React.MouseEvent) {
        if (isDragging.current && from) {
            const piece: SVGSVGElement = document.getElementById(from)?.children[0] as SVGSVGElement;

            if (currentColor === "b") {
                piece.style.left = `${boardHTML.offsetLeft + boardHTML.clientWidth - event.clientX - piece.clientWidth / 2}`;
                piece.style.top = `${boardHTML.offsetTop + boardHTML.clientHeight - event.clientY - piece.clientHeight / 2}`;
            } else {
                piece.style.left = `${event.clientX - piece.clientWidth / 2}`;
                piece.style.top = `${event.clientY - piece.clientHeight / 2}`;
            }
          
        }
    }

    function mouseupHandler(squareRepresentation: Square) {
        if (from) {
            if (isDragging.current && availableMoves.current?.includes(squareRepresentation)) {
                chess.move({
                    from,
                    to: squareRepresentation
                });

                socket.send(JSON.stringify({
                type: MOVE,
                payload: {
                    move: {
                        from,
                        to: squareRepresentation
                    },
                    color: chess.turn()
                }
                }));

                if(chess.turn() === "b"){
                    setStopTimerBlack(false)        
                    setStopTimerWhite(true)
                } else {
                    setStopTimerWhite(false)        
                    setStopTimerBlack(true)
                }

                const gameMove = new Audio("../../shared/assets/audio/Move.mp4")
                gameMove.play();

                setFrom(null);
                setBoard(chess.board());
                setMoves(updateLastMoveHighlights(chess));
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

    return <div className={["chess__board", currentColor === "b" ? "rotate" : ""].join(" ")} id="board">
        {board.map((row, i) => {
            return <div key={i} className="flex">
                {row.map((field, j) => {
                    const squareRepresentation = String.fromCharCode(97 + (j % 8)) + "" + (8 - i) as Square;
                    return <div
                        onMouseDown={() => mousedownHandler(squareRepresentation, field?.color)}
                        onMouseMove={(event) => mousemoveHandler(event)}
                        onMouseUp={() => mouseupHandler(squareRepresentation)}
                        key={j}
                        id={squareRepresentation}
                        className={["chess__field", 
                            (i + j) % 2 === 0 ? "w" : "b",
                        ].join(' ')}
                    >
                        <Piece field={field} className={[(isDragging.current && from === squareRepresentation) ? "dragging" : "", currentColor === "b" ? "rotate" : ""].join(" ")}/>
                        {(isDragging.current && squareRepresentation === from) && <Piece field={field} className={["copy", currentColor === "b" ? "rotate" : ""].join(" ")} />}
                    </div>
                })}
            </div>
        })}
    </div>
}

export function updateLastMoveHighlights(chess: Chess): Move[] {
    const moves = chess.history({ verbose: true })
    if (moves) {
        if (moves.length > 1) {
            document.getElementById(moves[moves.length - 2].from)?.classList.remove("undo");
            document.getElementById(moves[moves.length - 2].to)?.classList.remove("undo");
        }
            document.getElementById(moves[moves.length - 1].from)?.classList.add("undo");
            document.getElementById(moves[moves.length - 1].to)?.classList.add("undo");
    }
    return moves
}