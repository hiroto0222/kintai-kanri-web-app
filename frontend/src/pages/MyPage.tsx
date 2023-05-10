import { Button, Card, CardContent, Stack, Typography } from "@mui/material";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { authContext } from "../context/auth";
import useAuthApi from "../hooks/api/useAuthApi";
import useClockApi from "../hooks/api/useClockApi";
import useEmployees from "../hooks/api/useEmployeesApi";

const MyPage = () => {
  const { getEmployeeById } = useEmployees();
  const { clockIn } = useClockApi();
  const { logout } = useAuthApi();
  const { authState } = useContext(authContext);
  const navigate = useNavigate();

  return (
    <div>
      <h1>My Page</h1>
      <Card>
        <CardContent>
          <Stack spacing={3}>
            <Typography variant="h5">
              Welcome! {authState.user?.first_name} {authState.user?.last_name}
            </Typography>
          </Stack>
        </CardContent>
      </Card>
      <Button
        variant="contained"
        onClick={() => {
          if (authState.user) getEmployeeById(authState.user.id);
        }}
      >
        Get Employee
      </Button>
      <Button
        variant="contained"
        onClick={() => {
          logout();
        }}
      >
        Logout
      </Button>
      <Button
        variant="contained"
        onClick={() => {
          navigate("/register");
        }}
      >
        register
      </Button>
      <Button
        variant="contained"
        onClick={() => {
          if (authState.user) clockIn(authState.user?.id);
        }}
      >
        Clock In
      </Button>
    </div>
  );
};

export default MyPage;
