import type { InputState } from "../../lib/useInput"
import "./InputItem.scss"

export const InputItem = ({className, inputState, type,  ...options}: {
    className: string,
    inputState: InputState,
    type?: string,
}) => {
    return <li className={`input__item ${className}`}>
        <input className={`input__field ${inputState.error ? "input__error" : ""}`} type={type} autoComplete="new-password" {...options} />
        {inputState.error && <span>{inputState.error}</span>}
    </li>
}