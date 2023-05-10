import {
  Button,
  Card,
  Container,
  Grid,
  Stack,
  Typography,
} from "@mui/material";
import { Helmet } from "react-helmet-async";
import {
  EmployeesList,
  EmployeesListToolbar,
} from "../components/admin/employees";
import useListEmployees from "../hooks/api/useListEmployees";
import usePrivateAxios from "../hooks/usePrivateAxios";

const EmployeesPage = () => {
  const privateAxios = usePrivateAxios();
  const { employees } = useListEmployees(privateAxios);

  return (
    <>
      <Helmet>
        <title>Employees</title>
      </Helmet>
      <Container maxWidth="xl">
        <Stack
          direction="row"
          alignItems="center"
          justifyContent="space-between"
          sx={{ mb: 5 }}
        >
          <Typography variant="h3">Employees</Typography>
          <Button variant="contained">Register Employee</Button>
        </Stack>
        <Grid justifyContent="space-evenly" container spacing={3}>
          <Grid item xs={12}>
            <Card>
              <EmployeesListToolbar /> 
              <EmployeesList data={employees} />
            </Card>
          </Grid>
        </Grid>
      </Container>
    </>
  );
};

export default EmployeesPage;
