import { useContext } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import { authContext } from "./context/auth";
import LoginPage from "./pages/LoginPage";
import MyPage from "./pages/MyPage";

const App = () => {
  const { authState } = useContext(authContext);

  return (
    <div>
      <Routes>
        <Route
          path="/"
          element={authState.isLoggedIn ? <Navigate to="/me" /> : <LoginPage />}
        />
        <Route
          path="/me"
          element={authState.isLoggedIn ? <MyPage /> : <Navigate to="/" />}
        />
      </Routes>
    </div>
  );
};

export default App;
