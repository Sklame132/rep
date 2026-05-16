import { WebSocket } from "ws"
import { Chess, Move } from "chess.js"
import { GAME_OVER, INIT_GAME, MOVE } from "./messages";
import { randomUUID } from "node:crypto";

export type Player = {
    username: string;
    socket: WebSocket
}

export class Game {
    public gameId: string;
    public playerWhite: WebSocket;
    public playerBlack: WebSocket;
    public board: Chess;
    private startTime: Date;
    private moveCount = 0;

    constructor(playerWhite: WebSocket, playerBlack: WebSocket) {
        this.gameId = randomUUID();
        this.playerWhite = playerWhite;
        this.playerBlack = playerBlack;
        this.board = new Chess();
        this.startTime = new Date();
        this.playerWhite.send(JSON.stringify({
            type: INIT_GAME,
            payload: {
                color: "white",
                //player_w: this.playerWhite.username,
                game_id: this.gameId,
            }
        }))
        this.playerBlack.send(JSON.stringify({
            type: INIT_GAME,
            payload: {
                color: "black",
                //player_b: this.playerWhite.username,
                game_id: this.gameId,
            }
        }))
    }

    makeMove(socket: WebSocket, move: {
        from: string;
        to: string;
    }) {
         if (this.playerWhite) {
                this.playerWhite.emit(JSON.stringify({
                type: GAME_OVER,
                payload: {
                    winner: this.board.turn() === "w" ? "black" : "white"
                }
            }))
            }
            if (this.playerBlack) {
                this.playerBlack.emit(JSON.stringify({
                    type: GAME_OVER,
                    payload: {
                        winner: this.board.turn() === "w" ? "black" : "white"
                    }
                }))
            }
        if (this.moveCount % 2 === 0 && socket !== this.playerWhite) {
            return
        }
        if (this.moveCount % 2 === 1 && socket !== this.playerBlack) {
            return
        }

        try {
            this.board.move(move);
        } catch(err) {
            console.log(err);
            return;
        }

        if (this.board.isGameOver()) {
           
        }

        if (this.moveCount % 2 === 0) {
            this.playerBlack.send(JSON.stringify({
                type: MOVE,
                payload: move,
            }))
        } else {
            this.playerWhite.send(JSON.stringify({
                type: MOVE,
                payload: move,
            }))
        }
        this.moveCount++;
    }
}