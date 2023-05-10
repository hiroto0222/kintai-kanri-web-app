import { Button, Card, CardContent, Stack, Typography } from "@mui/material";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { authContext } from "../context/auth";
import useAuthApi from "../hooks/api/useAuthApi";
import useClockInsClockOutsApi from "../hooks/api/useClockInsClockOutsApi";
import useEmployees from "../hooks/api/useEmployeesApi";
import useListClockIns from "../hooks/api/useListClockIns";
import usePrivateAxios from "../hooks/usePrivateAxios";

const MyPage = () => {
  const privateApi = usePrivateAxios();
  const { getEmployeeById } = useEmployees(privateApi);
  const { clockIn, clockOut } = useClockInsClockOutsApi(privateApi);
  const { logout } = useAuthApi(privateApi);
  const { authState } = useContext(authContext);
  const navigate = useNavigate();

  const { clockIns } = useListClockIns(privateApi);

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
      <Button
        variant="contained"
        onClick={() => {
          if (authState.user) clockOut(authState.user?.id);
        }}
      >
        Clock Out
      </Button>
      <Card>
        <CardContent>
          <div>
            {clockIns &&
              clockIns.map((clockIn) => (
                <ul key={clockIn.clock_in_id}>
                  <li>clock in id: {clockIn.clock_in_id}</li>
                  <li>
                    clock in time: {new Date(clockIn.clock_in_time).toString()}
                  </li>
                  <li>clock out id: {clockIn.clock_out_id.Int32}</li>
                  <li>
                    clock out time:{" "}
                    {new Date(clockIn.clock_out_time.Time).toString()}
                  </li>
                </ul>
              ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default MyPage;
