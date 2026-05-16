import { useState, type SyntheticEvent } from "react";
import { Button } from "../../shared/ui/Button";
import "./Login.scss";
import { useInput } from "../../shared/lib/useInput";
import { InputItem } from "../../shared/ui/inputItem/InputItem";
import { useNavigate } from "react-router";

export const Login = () => {
    const navigate = useNavigate();
    const username = useInput("Имя пользователя", false);
    const password = useInput("Пароль", false);

    const [access, setAccess] = useState<boolean>(false);

    const submit = async (event: SyntheticEvent) => {
        event.preventDefault();

        let isError = false;
        if (!username.value) {
            username.setError("Обязательное поле");
            isError = true;
        }
        if (!password.value) {
            password.setError("Обязательное поле");
            isError = true;
        }
        if (isError) {
            return
        }

        const response = await fetch("http://localhost:5050/api/v1/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            credentials: "include",
            body: JSON.stringify({
                username: username.value,
                password: password.value,
            })
        })
        if (response.status === 404) {
            username.setError("Пользователь не найден");
        } else if (response.status === 500) {
            password.setError("Неверный пароль");
        } else {
            setAccess(true);
            navigate("/profile");
        }
    }

    return <div className="login__container">
        <div className="login">
            <div className="login__title__container">
                <h1 className="login__title">Авторизация</h1>
            </div>
            <form className="login__form" autoComplete="off" onSubmit={submit}>
                <ul className="login__input__list">
                    <InputItem className="login__input" inputState={username} {...username}/>
                    <InputItem className="login__input" inputState={password} type="password" {...password}/>
                </ul>
                <div className="login__actions">
                    <a className="login__register auth"onClick={() => navigate("/register")}>Регистрация</a>
                    <Button className="login__submit" type="submit" disabled={access}>
                        {access ? "Успешно" : "Войти"}
                    </Button>
                </div>
            </form>
        </div>
    </div>
}