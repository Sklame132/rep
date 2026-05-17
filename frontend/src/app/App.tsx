import { BrowserRouter, Outlet, Route, Routes } from "react-router"
import { Header } from "../widgets/header/Header";
import { ChessGame } from "../pages/chessGame/ChessGame";
import "./App.scss";
import { UserProfile } from "../pages/userProfile/UserProfile";
import { Login } from "../pages/login/Login";
import { Register } from "../pages/register/Register";

function App() {
    return (
        <>
            <AuthApp/>
        </>
    )
}

const MainLayout = () => {
    return <>
        <Header />
        <Outlet />
    </>
}

function AuthApp() {
    return (
        <BrowserRouter>
            <Routes>
                <Route element={<MainLayout />}>
                    <Route path="/profile" element={<UserProfile />}/>
                    <Route path="/play" element={<ChessGame />} />
                </Route>

                <Route path="/register" element={<Register />}/>
                <Route path="/login" element={<Login />}/>
            </Routes>
        </BrowserRouter>
    )
}

export default App