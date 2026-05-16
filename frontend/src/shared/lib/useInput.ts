import { useState, type ChangeEvent, type FocusEvent } from "react";

export type InputState = {
    value: string,
    placeholder: string,
    error: string | null,
    setError: React.Dispatch<React.SetStateAction<string | null>>,
    onBlur: (event: FocusEvent<HTMLInputElement>) => void,
    onChange: (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void,
}

export function useInput(placeholder: string, required: boolean): InputState {
    const [value, setValue] = useState<string>("");
    const [error, setError] = useState<string | null>(null);

    return {
        value,
        placeholder,
        error,
        setError,
        onBlur: (event) => {
            if (!error) {
                if (!event.target.value && required) {
                    setError("Обязательное поле");
                } else {
                    setError(null)
                }
            }
        },
        onChange: (event) => {
            setValue(event.target.value)
            if (event.target.value) {
                setError(null)
            }
        }
    };
};