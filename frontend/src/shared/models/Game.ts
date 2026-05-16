export type Game = {
    id: string;
    fen_start: string;
    fen_end: string;
    player_w: string;
    player_b: string;
    game_type: string;
    mode: string;
    result: string;
    history: string;
    created_at: Date;
}