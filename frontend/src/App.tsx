import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import { authContext } from "./context/auth";
import LoginPage from "./pages/LoginPage";
import MyPage from "./pages/MyPage";
import RegisterPage from "./pages/RegisterPage";

const App = () => {
  const { authState } = useContext(authContext);

  return (
    <div>
      <Routes>
        <Route
          path="/"
          element={authState.isLoggedIn ? <MyPage /> : <LoginPage />}
        />
        <Route
          path="/register"
          element={
            authState.isLoggedIn && authState.user?.is_admin ? (
              <RegisterPage />
            ) : (
              <Navigate to="/" />
            )
          }
        />
      </Routes>
    </div>
  );
};

export default App;
