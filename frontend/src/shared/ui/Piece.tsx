import { type Field } from "../../widgets/chessBoard/ChessBoard";

import Pawn from "../assets/pieces/Pawn.svg?react";
import Rook from "../assets/pieces/Rook.svg?react";
import Knight from "../assets/pieces/Knight.svg?react";
import Bishop from "../assets/pieces/Bishop.svg?react";
import Queen from "../assets/pieces/Queen.svg?react";
import King from "../assets/pieces/King.svg?react";

const Pieces = {
    'p': Pawn,
    'r': Rook,
    'n': Knight,
    'b': Bishop,
    'q': Queen,
    'k': King
}

interface PieceProps {
    field: Field;
    id?: string;
    className?: string;
}

export const Piece = ({ field, id, className }: PieceProps) => {
    if (field?.type) {
        const PieceIcon = Pieces[field.type];
        return <>
            {field?.type && <PieceIcon id={id} className={["chess__piece", field.color, className].join(' ')}/>}
        </>
    }
}