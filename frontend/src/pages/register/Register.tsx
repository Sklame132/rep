import { useEffect, useState, type SyntheticEvent } from "react";
import { useInput, type InputState } from "../../shared/lib/useInput"
import "./Register.scss"
import { Button } from "../../shared/ui/Button";
import { InputItem } from "../../shared/ui/inputItem/InputItem";
import type { User } from "../../shared/models/User";
import { useNavigate } from "react-router";

export const Register = () => {
    const navigate = useNavigate();

    const username = useInput("Имя пользователя", true);
    const password = useInput("Пароль", true);
    const passwordRepeat = useInput("Повтор пароля", true);
    const firstName = useInput("Имя", true);
    const lastName = useInput("Фамилия", true);
    const city = useInput("Город", false);
    const address = useInput("Адрес", false);
    const email = useInput("Почта", false);
    const phoneNumber = useInput("Номер телефона", false);

    const requiredStates: InputState[] = [username, password, passwordRepeat, firstName, lastName];

    const [responseStatus, setResponseStatus] = useState<string | null>(null);

    useEffect(() => {
        const timer = setTimeout(() => {
            if (password.value !== passwordRepeat.value) {
                password.setError("Пароли не совпадают");
                passwordRepeat.setError("Пароли не совпадают");
            } else {
                password.setError(null);
                passwordRepeat.setError(null);
            }
        }, 1000)

        return () => clearTimeout(timer)
    }, [password.value, passwordRepeat.value])

    useEffect(() => {
        const timer = setTimeout(() => {
            if (address.value && !city.value) {
                city.setError("Введите город");
            } else {
                city.setError(null)
            }
        }, 1000)
        
        return () => clearTimeout(timer)
    }, [address.value])

    useEffect(() => {
        email.setError(null);
        phoneNumber.setError(null);
    }, [email.value, phoneNumber.value])

    const submit = async (event: SyntheticEvent) => {
        event.preventDefault();
        setResponseStatus("Отправка...")
        let isError = false;

        requiredStates.map((state) => {
            if (!state.value) {
                state.setError("Обязательное поле")
                isError = true
            }
        })

        if (password.value != passwordRepeat.value) {
            password.setError("Пароли не совпадают");
            passwordRepeat.setError("Пароли не совпадают");
            isError = true;
        } else if (password.value.length < 8) {
            password.setError("Пароль должен содержать 8+ символов");
            passwordRepeat.setError("Пароль должен содержать 8+ символов");
            isError = true;
        }
        if (address.value && !city.value) {
            city.setError("Введите город");
            isError = true;
        }
        if (!email.value && !phoneNumber.value) {
            email.setError("Одно из полей должно быть заполнено") 
            phoneNumber.setError("Одно из полей должно быть заполнено")
            isError = true;
        }
        if (isError) {
            setResponseStatus(null)
            return 
        }

        const user: User = {
            username: username.value,
            password: password.value,
            first_name: firstName.value,
            last_name: lastName.value,
            address: `${city.value + ' ' + address.value}`
        }

        if (email.value) {user.email = email.value}
        if (phoneNumber.value) {user.phone_number = phoneNumber.value}
        
        const response = await fetch("http://localhost:5050/api/v1/register", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(user)
        })
        if (response.status === 500) {
            username.setError("Имя пользователя уже сущетсвует")
            setResponseStatus(null)
            return
        } else {
            setResponseStatus("Успешно")
            setTimeout(() => {
                navigate("/login")
            }, 2000)
        }
    }

    return <div className="register__container">
        <div className="register">
            <div className="register__title__container">
                <h1 className="register__title">Регистрация</h1>
            </div>
            <form className="register__form" autoComplete="off" onSubmit={submit}>
                <ul className="register__input__list">
                    <div className="register__input__group">
                        <InputItem className="register__input" inputState={username} {...username}/>
                    </div>
                    <div className="register__input__group">
                        <InputItem className="register__input" inputState={firstName} {...firstName}/>
                        <InputItem className="register__input" inputState={lastName} {...lastName}/>
                    </div>
                    <div className="register__input__group">
                        <InputItem className="register__input" inputState={password} type="password" {...password}/>
                        <InputItem className="register__input" inputState={passwordRepeat} type="password" {...passwordRepeat}/>
                    </div>
                    <div className="register__input__group">
                        <InputItem className="register__input" inputState={city} {...city}/>
                        <InputItem className="register__input" inputState={address} {...address}/>
                    </div>
                    <div className="register__input__group">
                        <InputItem className="register__input" inputState={email} {...email}/>
                        <InputItem className="register__input" inputState={phoneNumber} {...phoneNumber}/>
                    </div>
                </ul>
                <div className="register__actions">
                    <a className="register__login auth"onClick={() => navigate("/login")}>Вход</a>
                    <Button className="register__submit" type="submit" disabled={!!responseStatus}>
                        {(responseStatus === "Успешно") ? "Успешно" : "Отправить"}
                    </Button>
                </div>
            </form>
        </div>
    </div>
}