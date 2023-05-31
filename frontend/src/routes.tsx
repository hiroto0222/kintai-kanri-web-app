import { useContext } from "react";
import { Navigate, useRoutes } from "react-router-dom";
import { authContext } from "./context/auth";
import { DashboardLayout } from "./layouts/DashboardLayout";
import DashboardPage from "./pages/DashboardPage";
import EmployeePage from "./pages/EmployeePage";
import EmployeesPage from "./pages/EmployeesPage";
import LoginPage from "./pages/LoginPage";
import RegisterEmployeePage from "./pages/RegisterEmployeePage";

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
        {
          path: "employees",
          element:
            authState.isLoggedIn && authState.user?.is_admin ? (
              <EmployeesPage />
            ) : (
              <Navigate to="/" />
            ),
        },
        {
          path: "employees/register",
          element:
            authState.isLoggedIn && authState.user?.is_admin ? (
              <RegisterEmployeePage />
            ) : (
              <Navigate to="/" />
            ),
        },
        {
          path: "employees/:employeeId",
          element:
            authState.isLoggedIn && authState.user?.is_admin ? (
              <EmployeePage />
            ) : (
              <Navigate to="/" />
            ),
        },
      ],
    },
  ]);

  return routes;
};

export default Router;
