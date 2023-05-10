import { useContext } from "react";
import { useRoutes } from "react-router-dom";
import { authContext } from "./context/auth";
import { DashboardLayout } from "./layouts/dashboard/DashboardLayout";
import DashboardPage from "./pages/DashboardPage";
import LoginPage from "./pages/LoginPage";
import MyPage from "./pages/MyPage";
import RegisterPage from "./pages/RegisterPage";

const Router = () => {
  const { authState } = useContext(authContext);

  const routes = useRoutes([
    {
      path: "/",
      element: authState.isLoggedIn ? <MyPage /> : <LoginPage />,
    },
    {
      path: "/dashboard",
      element: <DashboardLayout />,
      children: [
        { element: <DashboardPage />, index: true },
        { path: "register", element: <RegisterPage /> },
      ],
    },
  ]);

  return routes;
};

export default Router;
