import { useContext } from "react";
import { Navigate, useRoutes } from "react-router-dom";
import { authContext } from "./context/auth";
import { DashboardLayout } from "./layouts/DashboardLayout";
import DashboardPage from "./pages/DashboardPage";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";

const Router = () => {
  const { authState } = useContext(authContext);

  const routes = useRoutes([
    {
      path: "/",
      element: authState.isLoggedIn ? (
        <Navigate to="/dashboard" />
      ) : (
        <LoginPage />
      ),
    },
    {
      path: "/dashboard",
      element: (
        <DashboardLayout
          isAdmin={authState.user ? authState.user.is_admin : false}
        />
      ),
      children: [
        {
          element: authState.isLoggedIn ? (
            <DashboardPage />
          ) : (
            <Navigate to="/" />
          ),
          index: true,
        },
        { path: "register", element: <RegisterPage /> },
      ],
    },
  ]);

  return routes;
};

export default Router;
